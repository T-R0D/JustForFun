# This one was a bit defeating. I tried to fully reverse engineer the
# code, but couldn't figure it out all of the way. If you want to know what
# the elf code does, it is a known algorithm and you can find out on Reddit.
#
# What I did get out of the reverse engineering was that the input register
# is only used once - to check against register `3`. So after this arbitrary
# process, there is an equality check against register `0`. I used this
# information to store the incoming value in that equality check, and build
# a list until there was a loop. Rather than try different starting values
# for register `0`, I just used the value 0 (a known "bad" input) to make
# the elf code loop through the possible values endlessly (stopping, of
# course, when I found a loop in the "acceptable" values).
#
# Unfortunately, this requires a runtime of almost 10 minutes, but many
# people in the solution thread did the same, so eh... The major differences
# to get the speed up seem to be just rewriting the algorithm in a more
# efficient way in a target language of your choice. I could _probably_ do
# that with more time, but I kind of don't want to spend more time on this.
#
# To truly reverse engineer this requires reversing a hashing function (hint),
# which is a little bit beyond me without a lot more time.


def part_one(puzzle_input):
    instruction_pointer_register, program = parse_program(puzzle_input)

    processor = Processor(0, instruction_pointer_register)
    halting_values = processor.run_program_to_find_halting_values(program)

    return str(halting_values[0][0])


def part_two(puzzle_input):
    instruction_pointer_register, program = parse_program(puzzle_input)

    processor = Processor(0, instruction_pointer_register)
    halting_values = processor.run_program_to_find_halting_values(program)

    return str(halting_values[-1][0])


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
    def __init__(self, initial_zero_register_value, instruction_pointer_register):
        self.registers = [
            initial_zero_register_value if i == 0 else 0 for i in range(6)
        ]
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

    def run_program_to_find_halting_values(self, program):
        halting_values = []
        instruction_count = 0
        while 0 <= self.registers[self.instruction_pointer_register] < len(program):
            op, a, b, c = program[self.registers[self.instruction_pointer_register]]

            if op == "eqrr" and (a == 0 or b == 0):
                halting_value = 0
                if a == 0:
                    halting_value = self.registers[b]
                else:
                    halting_value = self.registers[a]

                if any(halting_value == item[0] for item in halting_values):
                    break

                # should be `instruction_count + 2`, but eh...`
                halting_values.append((halting_value, instruction_count))

            self.op_code_mappings[op](a, b, c)
            instruction_count += 1
            self.registers[self.instruction_pointer_register] += 1

        return halting_values

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
