package intcode

import (
	"reflect"
	"testing"
)

func TestRunBatchProgram(t *testing.T) {
	testRunBatchProgramHappyPath(t, "1,9,10,3,2,3,11,0,99,30,40,50", 3500, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50})

	testRunBatchProgramHappyPath(t, "1,0,0,0,99", 2, []int{2, 0, 0, 0, 99})

	testRunBatchProgramHappyPath(t, "2,3,0,3,99", 2, []int{2, 3, 0, 6, 99})

	testRunBatchProgramHappyPath(t, "2,4,4,5,99,0", 2, []int{2, 4, 4, 5, 99, 9801})

	testRunBatchProgramHappyPath(t, "1,1,1,4,99,5,6,0,99", 30, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})

	testRunBatchProgramHappyPath(t, "1101,100,-1,4,0", 1101, []int{1101, 100, -1, 4, 99})

	testRunBatchProgramHappyPath(t, "1002,4,3,4,33", 1002, []int{1002, 4, 3, 4, 99})
}

func testRunBatchProgramHappyPath(t *testing.T, prog string, expectedResult int, expectedMem []int) {
	c := NewComputer()
	err := c.InputProgram(prog)
	if err != nil {
		t.Fatalf("Expected input program success: %v", err)
	}
	state, err := c.RunProgram()
	if err != nil {
		t.Fatalf("Expected run program success: %v", err)
	} else if state != STATE_RUN_COMPLETE {
		t.Fatalf("Expected a complete program run, program is in state: %v", state)
	}
	resultMem := c.GetMemory()
	result, err := c.GetProgramResult()
	if err != nil {
		t.Fatalf("Expected to successfully get result")
	}
	if !reflect.DeepEqual(expectedMem, resultMem) || result != expectedResult || err != nil {
		t.Fatalf("Program run failed; expected resulting program: %v, got program: %v, expecting result: %v, got result: %v err: %v", expectedMem, resultMem, expectedResult, result, err)
	}
}

func TestBatchProgramWithInputAndOutput(t *testing.T) {
	testBatchProgramWithInputAndOutput(t, "3,0,4,0,99", []int{5}, []int{5})

	testBatchProgramWithInputAndOutput(
		t,
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		[]int{7},
		[]int{999})

	testBatchProgramWithInputAndOutput(
		t,
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		[]int{8},
		[]int{1000})

	testBatchProgramWithInputAndOutput(
		t,
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		[]int{9},
		[]int{1001})
}

func testBatchProgramWithInputAndOutput(t *testing.T, prog string, input, expectedOutput []int) {
	c := NewComputer()
	err := c.InputProgram(prog)
	if err != nil {
		t.Fatalf("Expected input program success: %v", err)
	}
	for _, i := range input {
		c.AddBatchInput(i)
	}
	state, err := c.RunProgram()
	if err != nil {
		t.Fatalf("Expected run program success: %v", err)
	} else if state != STATE_RUN_COMPLETE {
		t.Fatalf("Expected complete program run; computer is in state: %v", state)
	}
	resultOutput := c.GetBatchOutput()
	if !reflect.DeepEqual(resultOutput, expectedOutput) {
		t.Fatalf("Failure in output - expected: %v, got: %v", expectedOutput, resultOutput)
	}
}

func TestProgramStrToMemLayout(t *testing.T) {
	testProgramStrToMemLayoutHappyPath(t, "1,0,0,0,99", []int{1, 0, 0, 0, 99})

	testProgramStrToMemLayoutHappyPath(t, "2,3,0,3,99", []int{2, 3, 0, 3, 99})

	testProgramStrToMemLayoutHappyPath(t, "2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 0})

	testProgramStrToMemLayoutHappyPath(t, "1,1,1,4,99,5,6,0,99", []int{1, 1, 1, 4, 99, 5, 6, 0, 99})

	in := ""
	_, err := programStrToMemLayout(in)
	if nil == err {
		t.Fatalf("Program is malformed, expected failure")
	}
}

func testProgramStrToMemLayoutHappyPath(t *testing.T, in string, expect []int) {
	result, err := programStrToMemLayout(in)
	if !reflect.DeepEqual(result, expect) || err != nil {
		t.Fatalf("Conversion failed; expected: %v, got: %v, err: %v", expect, result, err)
	}
}

func TestParseOpAndParams(t *testing.T) {
	testParseOpAndParams(t, 99, OP_TER, []int{0, 0, 0})

	testParseOpAndParams(t, 1002, OP_MUL, []int{0, 1, 0})

	testParseOpAndParams(t, 1101, OP_ADD, []int{1, 1, 0})
}

func testParseOpAndParams(t *testing.T, in int, expectedOp int, expectedParams []int) {
	op, params := parseOpAndParams(in)
	if op != expectedOp || !reflect.DeepEqual(params, expectedParams) {
		t.Fatalf("Failure: expected op: %v, got: %v; expected params: %v, got: %v", expectedOp, op, expectedParams, params)
	}
}

func TestRunQuineProgram(t *testing.T) {
	input := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	expected := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	testBatchProgramByOutput(t, input, expected)
}

func TestRunBigOuptutProgram(t *testing.T) {
	input := "1102,34915192,34915192,7,4,7,99,0"
	expected := []int{1219070632396864}
	testBatchProgramByOutput(t, input, expected)
}

func TestRunBigValueInProgram(t *testing.T) {
	input := "104,1125899906842624,99"
	expected := []int{1125899906842624}
	testBatchProgramByOutput(t, input, expected)
}

func testBatchProgramByOutput(t *testing.T, prog string, expected []int) {
	c := NewComputer()
	err := c.InputProgram(prog)
	if err != nil {
		t.Fatalf("Unexpected error in program input: %v", err)
	}

	state, err := c.RunProgram()
	if err != nil {
		t.Fatalf("Expected healthy program run: %v", err)
	} else if state != STATE_RUN_COMPLETE {
		t.Fatalf("Expected program to run to completion, not be in state: %v", state)
	}

	output := c.GetBatchOutput()
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Output does not match expectation. Expected\n%v,\ngot:\n%v", expected, output)
	}
}

func TestOpAdd(t *testing.T) {
	testOp(t, "1,5,6,0,99,2,3", []int{}, 5, []int{})
	testOp(t, "101,5,6,0,99,2,3", []int{}, 8, []int{})
	testOp(t, "1001,5,6,0,99,2,3", []int{}, 8, []int{})
	testOp(t, "1101,5,6,0,99,2,3", []int{}, 11, []int{})
}

func TestOpMul(t *testing.T) {
	testOp(t, "2,5,6,0,99,2,3", []int{}, 6, []int{})
	testOp(t, "102,5,6,0,99,2,3", []int{}, 15, []int{})
	testOp(t, "1002,5,6,0,99,2,3", []int{}, 12, []int{})
	testOp(t, "1102,5,6,0,99,2,3", []int{}, 30, []int{})
}

func TestOpSto(t *testing.T) {
	testOp(t, "3,0,99", []int{3}, 3, []int{})

	testOp(t, "109,5,203,-5,99", []int{69}, 69, []int{})
}

func TestOpOut(t *testing.T) {
	testOp(t, "4,3,99,7", []int{}, 4, []int{7})
}

func TestOpJit(t *testing.T) {
	// Expected behavior:
	// Check if value at pos 9 is "true"
	// Jump to position stored at position 11 (6)
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "5,9,11,4,10,99,4,9,99,1,0,6", []int{}, 5, []int{1})

	// Expected behavior:
	// Check if first parameter is "true"
	// Jump to position stored at position 11 (6)
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "105,1,11,4,10,99,4,9,99,1,0,6", []int{}, 105, []int{1})

	// Expected behavior:
	// Check if Check if value at pos 9 is "true"
	// Jump to position 6
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "1005,9,6,4,10,99,4,9,99,1,0,6", []int{}, 1005, []int{1})

	// Expected behavior:
	// Check if first parameter is "true"
	// Jump to position 6
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "1105,1,6,4,10,99,4,9,99,1,0,6", []int{}, 1105, []int{1})

	// Expected behavior:
	// Check if value at pos 10 is "true" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "5,10,11,4,10,99,4,9,99,1,0,6", []int{}, 5, []int{0})

	// Expected behavior:
	// Check if first parameter is "true" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "105,0,11,4,10,99,4,9,99,1,0,6", []int{}, 105, []int{0})

	// Expected behavior:
	// Check if Check if value at pos 10 is "true" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "1005,10,6,4,10,99,4,9,99,1,0,6", []int{}, 1005, []int{0})

	// Expected behavior:
	// Check if first parameter is "true" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "1105,0,6,4,10,99,4,9,99,1,0,6", []int{}, 1105, []int{0})

	// Expected behavior:
	// Check if first parameter is "true"
	// Jump to position 6
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "1105,-1,6,4,10,99,4,9,99,1,0,6", []int{}, 1105, []int{1})
}

func TestOpJif(t *testing.T) {
	// Expected behavior:
	// Check if value at pos 10 is "false"
	// Jump to position stored at position 11 (6)
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "6,10,11,4,10,99,4,9,99,1,0,6", []int{}, 6, []int{1})

	// Expected behavior:
	// Check if first parameter is "false"
	// Jump to position stored at position 11 (6)
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "106,0,11,4,10,99,4,9,99,1,0,6", []int{}, 106, []int{1})

	// Expected behavior:
	// Check if Check if value at pos 10 is "false"
	// Jump to position 6
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "1006,10,6,4,10,99,4,9,99,1,0,6", []int{}, 1006, []int{1})

	// Expected behavior:
	// Check if first parameter is "false"
	// Jump to position 6
	// Output value at position 9 (1)
	// Terminate
	testOp(t, "1106,0,6,4,10,99,4,9,99,1,0,6", []int{}, 1106, []int{1})

	// Expected behavior:
	// Check if value at pos 9 is "false" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "6,9,11,4,10,99,4,9,99,1,0,6", []int{}, 6, []int{0})

	// Expected behavior:
	// Check if first parameter is "false" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "106,1,11,4,10,99,4,9,99,1,0,6", []int{}, 106, []int{0})

	// Expected behavior:
	// Check if Check if value at pos 9 is "false" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "1006,9,6,4,10,99,4,9,99,1,0,6", []int{}, 1006, []int{0})

	// Expected behavior:
	// Check if first parameter is "false" (it's not)
	// Move to the next instruction
	// Output value at position 10 (0)
	// Terminate
	testOp(t, "1106,-1,6,4,10,99,4,9,99,1,0,6", []int{}, 1106, []int{0})
}

func TestOpLst(t *testing.T) {
	testOp(t, "7,5,6,0,99,0,1", []int{}, 1, []int{})
	testOp(t, "7,5,6,0,99,1,0", []int{}, 0, []int{})

	testOp(t, "107,5,6,0,99,0,6", []int{}, 1, []int{})
	testOp(t, "107,7,6,0,99,0,6", []int{}, 0, []int{})

	testOp(t, "1007,5,6,0,99,-1,1", []int{}, 1, []int{})
	testOp(t, "1007,5,6,0,99,7,0", []int{}, 0, []int{})

	testOp(t, "1107,5,6,0,99,0,1", []int{}, 1, []int{})
	testOp(t, "1107,7,6,0,99,1,0", []int{}, 0, []int{})
}

func TestOpEqu(t *testing.T) {
	testOp(t, "8,5,6,0,99,0,0", []int{}, 1, []int{})
	testOp(t, "8,5,6,0,99,1,0", []int{}, 0, []int{})

	testOp(t, "108,5,6,0,99,0,5", []int{}, 1, []int{})
	testOp(t, "108,-5,6,0,99,0,5", []int{}, 0, []int{})

	testOp(t, "1008,5,-1,0,99,-1,1", []int{}, 1, []int{})
	testOp(t, "1008,5,6,0,99,-1,0", []int{}, 0, []int{})

	testOp(t, "1108,3,3,0,99,0,1", []int{}, 1, []int{})
	testOp(t, "1108,3,-3,0,99,1,0", []int{}, 0, []int{})
}

func TestOpRel(t *testing.T) {
	testOp(t, "9,5,204,0,99,5", []int{}, 9, []int{5})

	testOp(t, "109,5,204,0,99,5", []int{}, 109, []int{5})

	testOp(t, "209,5,204,0,99,5", []int{}, 209, []int{5})
}

func testOp(t *testing.T, prog string, input []int, expectedResult int, expectedOutput []int) {
	c := NewComputer()

	err := c.InputProgram(prog)
	if err != nil {
		t.Fatalf("Expected input program success: %v", err)
	}

	c.AddBatchInputs(input)

	state, err := c.RunProgram()
	if err != nil {
		t.Fatalf("Expected run program success: %v", err)
	} else if state != STATE_RUN_COMPLETE {
		t.Fatalf("Expected complete program run; computer is in state: %v", state)
	}

	result, err := c.GetProgramResult()
	if err != nil {
		t.Fatalf("Error getting result: %v", err)
	} else if result != expectedResult {
		t.Fatalf("Incorrect result - expected: %v, got: %v", expectedResult, result)
	}

	output := c.GetBatchOutput()
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf("Incorrect output - expected: %v, got: %v", expectedOutput, output)
	}
}

func TestInterruptibleProgram(t *testing.T) {
	prog := "3,5,4,5,99,0"
	v := 42069 // Nice.
	var state int
	var err error

	c := NewComputer()
	c.SetInterruptibleMode()

	err = c.InputProgram(prog)
	if err != nil {
		t.Fatalf("Expected input of program to be successful: %v", err)
	}

	state, err = c.RunProgram()
	if err != nil {
		t.Fatalf("Expected run to be error free: %v", err)
	} else if state != STATE_AWAITING_INPUT {
		t.Fatalf("Expected program to stop to await input, not be in state: %v", state)
	}
	c.Input(v)

	state, err = c.RunProgram()
	if err != nil {
		t.Fatalf("Expected run to be error free: %v", err)
	} else if state != STATE_AWAITING_OUTPUT {
		t.Fatalf("Expected program to stop to await output, not be in state: %v", state)
	}
	output := c.CollectOutput()
	if output != v {
		t.Fatalf("Output incorrect; expected: %v, got: %v", v, output)
	}

	state, err = c.RunProgram()
	if err != nil {
		t.Fatalf("Expected run to be error free: %v", err)
	} else if state != STATE_RUN_COMPLETE {
		t.Fatalf("Expected program to complete, not be in state: %v", state)
	}
}
