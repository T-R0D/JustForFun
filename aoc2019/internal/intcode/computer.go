package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	RESULT_ADDR = 0

	OP_NUL = 0
	OP_ADD = 1
	OP_MUL = 2
	OP_INP = 3
	OP_OUT = 4
	OP_JIT = 5
	OP_JIF = 6
	OP_LST = 7
	OP_EQU = 8
	OP_REL = 9
	OP_TER = 99

	PARAM_MODE_POSITION  = 0
	PARAM_MODE_IMMEDIATE = 1
	PARAM_MODE_RELATIVE  = 2

	STATE_RUN_COMPLETE    = 0
	STATE_RUNTIME_ERROR   = 1
	STATE_AWAITING_INPUT  = 2
	STATE_AWAITING_OUTPUT = 3
	STATE_READY           = 4

	MODE_BATCH         = 0
	MODE_INTERRUPTIBLE = 1

	kSTORE_OPERATION      = "store"
	kREAD_OPERATION       = "read"
	kADVANCE_PC_OPERATION = "advance_instruction_pointer"
)

func StateToString(state int) string {
	switch state {
	case STATE_RUN_COMPLETE:
		return "RUN COMPLETE"
	case STATE_RUNTIME_ERROR:
		return "ERROR"
	case STATE_AWAITING_INPUT:
		return "AWAITING INPUT"
	case STATE_AWAITING_OUTPUT:
		return "AWAITING OUTPUT"
	case STATE_READY:
		return "READY"
	default:
		return "WTF"
	}
}

type outOfBoundsError struct {
	operation string
	addr      int
}

func (e *outOfBoundsError) Error() string {
	return fmt.Sprintf(
		"Out of Bounds Error: illegal %q access at %d (legal access is 0 <= addr)",
		e.operation, e.addr)
}

type Computer struct {
	mem           map[int]int
	mode          int
	pc            int
	relativeBase  int
	state         int
	awaitingInput bool
	inReg         int
	outReg        int
	input         []int
	inPtr         int
	output        []int
}

type memMap map[int]int

func NewComputer() *Computer {
	return &Computer{
		mem: memMap{
			0: OP_TER,
		},
		mode:          MODE_BATCH,
		pc:            0,
		relativeBase:  0,
		state:         STATE_READY,
		awaitingInput: false,
		input:         make([]int, 0),
		inPtr:         0,
		output:        make([]int, 0),
	}
}

func (c *Computer) Clone() *Computer {
	memCpy := memMap{}
	for k, v := range c.mem {
		memCpy[k] = v
	}
	return &Computer{
		mem:           memCpy,
		mode:          c.mode,
		pc:            c.pc,
		relativeBase:  c.relativeBase,
		state:         c.state,
		awaitingInput: c.awaitingInput,
		inReg:         c.inReg,
		outReg:        c.outReg,
		input:         append(c.input[:0:0], c.input...),
		inPtr:         c.inPtr,
		output:        append(c.output[:0:0], c.output...),
	}
}

func (c *Computer) SetInterruptibleMode() {
	c.mode = MODE_INTERRUPTIBLE
}

func (c *Computer) InputProgram(p string) error {
	mem, err := programStrToMemMap(p)
	if err != nil {
		return err
	}
	c.mem = mem
	return nil
}

func programStrToMemLayout(p string) ([]int, error) {
	opsAsStrs := strings.Split(p, ",")

	if len(opsAsStrs) <= 0 {
		return nil, fmt.Errorf("Program is not a suitable length (len=%d)", len(opsAsStrs))
	}

	prog := make([]int, len(opsAsStrs))
	for i, op := range opsAsStrs {
		opAsInt, err := strconv.Atoi(op)
		if err != nil {
			return nil, err
		}
		prog[i] = opAsInt
	}

	return prog, nil
}

func programStrToMemMap(p string) (memMap, error) {
	opsAsStrs := strings.Split(p, ",")

	if len(opsAsStrs) <= 0 {
		return nil, fmt.Errorf("Program is not a suitable length (len=%d)", len(opsAsStrs))
	}

	prog := memMap{}
	for i, op := range opsAsStrs {
		opAsInt, err := strconv.Atoi(op)
		if err != nil {
			return nil, err
		}
		prog[i] = opAsInt
	}

	return prog, nil
}

func (c *Computer) UpdateProgram(updates map[int]int) error {
	for i, v := range updates {
		err := c.storeValue(i, PARAM_MODE_IMMEDIATE, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Computer) AddBatchInputs(inputs []int) {
	for _, i := range inputs {
		c.AddBatchInput(i)
	}
}

func (c *Computer) AddBatchInput(v int) {
	c.input = append(c.input, v)
}

func (c *Computer) Input(v int) {
	c.inReg = v
}

func (c *Computer) CollectOutput() int {
	return c.outReg
}

func (c *Computer) GetState() int {
	return c.state
}

func (c *Computer) RunProgram() (int, error) {
	op := OP_NUL
	var params []int
	err := error(nil)
	for c.pc < len(c.mem) && err == nil {
		opAndParams := c.mem[c.pc]
		op, params = parseOpAndParams(opAndParams)

		switch op {
		case OP_ADD:
			c.pc, err = c.opAdd(c.pc, params)

		case OP_MUL:
			c.pc, err = c.opMul(c.pc, params)

		case OP_INP:
			if !c.awaitingInput && c.mode == MODE_INTERRUPTIBLE {
				c.awaitingInput = true
				c.state = STATE_AWAITING_INPUT
				return STATE_AWAITING_INPUT, nil
			} else {
				c.awaitingInput = false
			}
			c.pc, err = c.opSto(c.pc, params)

		case OP_OUT:
			c.pc, err = c.opOut(c.pc, params)
			if c.mode == MODE_INTERRUPTIBLE {
				c.state = STATE_AWAITING_OUTPUT
				return STATE_AWAITING_OUTPUT, nil
			}

		case OP_JIT:
			c.pc, err = c.opJit(c.pc, params)

		case OP_JIF:
			c.pc, err = c.opJif(c.pc, params)

		case OP_LST:
			c.pc, err = c.opLst(c.pc, params)

		case OP_EQU:
			c.pc, err = c.opEqu(c.pc, params)

		case OP_REL:
			c.pc, err = c.opRel(c.pc, params)

		case OP_TER:
			c.pc, err = c.opTer(c.pc, params)

		default:
			err = fmt.Errorf("Unrecognized op code: %d", op)
		}

		if err != nil {
			c.state = STATE_RUNTIME_ERROR
			return STATE_RUNTIME_ERROR, err
		}
	}

	c.state = STATE_RUN_COMPLETE
	return STATE_RUN_COMPLETE, nil
}

func (c *Computer) GetProgramResult() (int, error) {
	return c.getValue(RESULT_ADDR, PARAM_MODE_IMMEDIATE)
}

func (c *Computer) GetBatchOutput() []int {
	return c.output
}

func (c *Computer) GetMemory() []int {
	if len(c.mem) == 0 {
		return []int{}
	}

	memLen := 0
	for k := range c.mem {
		if k+1 > memLen {
			memLen = k + 1
		}
	}

	memSlice := make([]int, memLen)
	for k, v := range c.mem {
		memSlice[k] = v
	}

	return memSlice
}

func parseOpAndParams(opAndParams int) (op int, params []int) {
	op = opAndParams % 100

	params = make([]int, 0, 4)
	opAndParams /= 100
	for i := 0; i < 3; i++ {
		params = append(params, opAndParams%10)
		opAndParams /= 10
	}

	return
}

func (c *Computer) opAdd(pc int, params []int) (int, error) {
	if len(c.mem) <= pc+4 {
		return 0, &outOfBoundsError{
			operation: kADVANCE_PC_OPERATION,
			addr:      pc + 4,
		}
	}
	addend1, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}
	addend2, err := c.getValue(pc+2, params[1])
	if err != nil {
		return 0, err
	}
	resultAddr, err := c.getValue(pc+3, PARAM_MODE_IMMEDIATE)
	if err != nil {
		return 0, err
	}
	err = c.storeValue(resultAddr, params[2], addend1+addend2)
	if err != nil {
		return 0, err
	}

	return pc + 4, nil
}

func (c *Computer) opMul(pc int, params []int) (int, error) {
	if len(c.mem) <= pc+4 {
		return 0, &outOfBoundsError{
			operation: kADVANCE_PC_OPERATION,
			addr:      pc + 4,
		}
	}
	multiplicand1, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}
	multiplicand2, err := c.getValue(pc+2, params[1])
	if err != nil {
		return 0, err
	}
	resultAddr, err := c.getValue(pc+3, PARAM_MODE_IMMEDIATE)
	if err != nil {
		return 0, err
	}
	err = c.storeValue(resultAddr, params[2], multiplicand1*multiplicand2)
	if err != nil {
		return 0, err
	}

	return pc + 4, nil
}

func (c *Computer) opSto(pc int, params []int) (int, error) {
	var v int
	var err error = nil
	if c.mode == MODE_INTERRUPTIBLE {
		v = c.inReg
	} else {
		v, err = c.takeInput()
	}
	if err != nil {
		return 0, err
	}
	storeAddr, err := c.getValue(pc+1, PARAM_MODE_IMMEDIATE)
	if err != nil {
		return 0, err
	}

	err = c.storeValue(storeAddr, params[0], v)
	if err != nil {
		return 0, err
	}

	return pc + 2, nil
}

func (c *Computer) opOut(pc int, params []int) (int, error) {
	v, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}

	if c.mode == MODE_INTERRUPTIBLE {
		c.outReg = v
	} else {
		c.sendOutput(v)
	}

	return pc + 2, nil
}

func (c *Computer) opJit(pc int, params []int) (int, error) {
	v, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}

	if v != 0 {
		return c.getValue(pc+2, params[1])
	}

	return pc + 3, nil
}

func (c *Computer) opJif(pc int, params []int) (int, error) {
	v, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}

	if v == 0 {
		return c.getValue(pc+2, params[1])
	}

	return pc + 3, nil
}

func (c *Computer) opLst(pc int, params []int) (int, error) {
	v1, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}
	v2, err := c.getValue(pc+2, params[1])
	if err != nil {
		return 0, err
	}

	var r int
	if v1 < v2 {
		r = 1
	} else {
		r = 0
	}
	storeAddr, err := c.getValue(pc+3, PARAM_MODE_IMMEDIATE)
	if err != nil {
		return 0, err
	}
	err = c.storeValue(storeAddr, params[2], r)
	if err != nil {
		return 0, err
	}

	return pc + 4, nil
}

func (c *Computer) opEqu(pc int, params []int) (int, error) {
	v1, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}
	v2, err := c.getValue(pc+2, params[1])
	if err != nil {
		return 0, err
	}

	var r int
	if v1 == v2 {
		r = 1
	} else {
		r = 0
	}
	storeAddr, err := c.getValue(pc+3, PARAM_MODE_IMMEDIATE)
	if err != nil {
		return 0, err
	}
	err = c.storeValue(storeAddr, params[2], r)
	if err != nil {
		return 0, err
	}

	return pc + 4, nil
}

func (c *Computer) opRel(pc int, params []int) (int, error) {
	v, err := c.getValue(pc+1, params[0])
	if err != nil {
		return 0, err
	}

	c.relativeBase += v

	return pc + 2, nil
}

func (c *Computer) opTer(pc int, params []int) (int, error) {
	return len(c.mem), nil
}

func (c *Computer) getValue(addr, mode int) (int, error) {
	if addr < 0 {
		return 0, &outOfBoundsError{
			operation: kREAD_OPERATION,
			addr:      addr,
		}
	}
	v, ok := c.mem[addr]
	if !ok {
		c.mem[addr] = 0
		v = c.mem[addr]
	}

	if mode == PARAM_MODE_IMMEDIATE {
		return v, nil
	} else if mode == PARAM_MODE_POSITION {
		return c.getValue(v, PARAM_MODE_IMMEDIATE)
	} else { // mode == PARAM_MODE_RELATIVE
		return c.getValue(c.relativeBase+v, PARAM_MODE_IMMEDIATE)
	}
}

func (c *Computer) storeValue(addr, param, v int) error {
	if param == PARAM_MODE_RELATIVE {
		addr = c.relativeBase + addr
	}

	if addr < 0 {
		return &outOfBoundsError{
			operation: kSTORE_OPERATION,
			addr:      addr,
		}
	}

	c.mem[addr] = v
	return nil
}

func (c *Computer) takeInput() (int, error) {
	if c.inPtr >= len(c.input) {
		return 0, fmt.Errorf("Not enough input to the program")
	}

	v := c.input[c.inPtr]
	c.inPtr++
	return v, nil
}

func (c *Computer) sendOutput(v int) {
	c.output = append(c.output, v)
}
