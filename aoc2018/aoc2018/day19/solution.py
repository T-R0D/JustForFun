def part_one(puzzle_input):
    instruction_pointer_register, program = parse_program(puzzle_input)

    processor = Processor([0 for _ in range(6)], instruction_pointer_register)
    processor.run_program(program)

    registers = processor.check()

    return str(registers[0])

def part_two(puzzle_input):
    # This one took some manual analysis. Basically, the program boils down to
    # a convoluted summing all of the numbers that an initial constant is
    # divisible by (at least for my input).
    # 
    # That constant gets bigger based on
    # a sort of toggle based on the initial value of the `0` register, so it's
    # a small number for an initial value of 0, much larger for 1.
    # 
    # The convoluted program works by initializing the constant in register
    # `5`, using register `3` as a jump flag, register `1` as a jump flag and
    # `4` as a sort of accumulator (and `0`, of course, is the accumulator for
    # the end result).
    #
    # The program counts up in register `4` until it exceeds `5`
    # (for x in range(1, *5 + 1)). Register `1` is used as a second counter.
    # If `*1 * *4 == *5` ever, then the value in `4` is added to `0`.
    # 
    # This results in A LOT of unnecessary incrementing/counting, which is
    # why the runtime is long enough to justify editing the program and
    # writing a second, faster one.
    instruction_pointer_register, program = parse_program(puzzle_input)

    program[-1] = ("seti", 255, 0, instruction_pointer_register)

    processor = Processor(
        [1 if i == 0 else 0 for i in range(6)], instruction_pointer_register
    )
    processor.run_program(program)

    registers = processor.check()

    big_constant = registers[-1]

    total = 0
    for i in range(1, big_constant + 1):
        if big_constant % i == 0:
            total += i

    return str(total)


def parse_program(puzzle_input):
    instruction_pointer_register = 0
    program = []
    for line in puzzle_input.split("\n"):
        if line.startswith("#ip"):
            instruction_pointer_register = int(line.replace("#ip ", ""))
            continue

        op, *param_strs = line.split()
        params = (int(x) for x in param_strs)

        program.append((op, *params))

    return instruction_pointer_register, program


class Processor:
    def __init__(self, initial_registers, instruction_pointer_register):
        self.registers = [x for x in initial_registers]
        self.instruction_pointer_register = instruction_pointer_register
        self.op_code_mappings = {
            "addr": self.addr,
            "addi": self.addi,
            "mulr": self.mulr,
            "muli": self.muli,
            "banr": self.banr,
            "bani": self.bani,
            "borr": self.borr,
            "bori": self.bori,
            "setr": self.setr,
            "seti": self.seti,
            "gtir": self.gtir,
            "gtri": self.gtri,
            "gtrr": self.gtrr,
            "eqir": self.eqir,
            "eqri": self.eqri,
            "eqrr": self.eqrr,
        }

    def run_program(self, program):
        while 0 <= self.registers[self.instruction_pointer_register] < len(program):
            op, a, b, c = program[self.registers[self.instruction_pointer_register]]
            self.op_code_mappings[op](a, b, c)
            self.registers[self.instruction_pointer_register] += 1

    def check(self):
        return tuple(self.registers)

    def addr(self, a, b, c):
        self.registers[c] = self.registers[a] + self.registers[b]

    def addi(self, a, b, c):
        self.registers[c] = self.registers[a] + b

    def mulr(self, a, b, c):
        self.registers[c] = self.registers[a] * self.registers[b]

    def muli(self, a, b, c):
        self.registers[c] = self.registers[a] * b

    def banr(self, a, b, c):
        self.registers[c] = self.registers[a] & self.registers[b]

    def bani(self, a, b, c):
        self.registers[c] = self.registers[a] & b

    def borr(self, a, b, c):
        self.registers[c] = self.registers[a] | self.registers[b]

    def bori(self, a, b, c):
        self.registers[c] = self.registers[a] | b

    def setr(self, a, _b, c):
        self.registers[c] = self.registers[a]

    def seti(self, a, _b, c):
        self.registers[c] = a

    def gtir(self, a, b, c):
        self.registers[c] = 1 if a > self.registers[b] else 0

    def gtri(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] > b else 0

    def gtrr(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] > self.registers[b] else 0

    def eqir(self, a, b, c):
        self.registers[c] = 1 if a == self.registers[b] else 0

    def eqri(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] == b else 0

    def eqrr(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] == self.registers[b] else 0
