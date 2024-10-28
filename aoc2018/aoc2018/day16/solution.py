def part_one(puzzle_input):
    instruction_samples, _ = parse_sample_instructions_and_program(puzzle_input)

    samples_with_three_or_more_possible_op_codes = 0
    for before, (_, a, b, c), after in instruction_samples:
        possible_ops = 0
        for i in range(16):
            processor = Processor(before)
            processor.op_list[i](a, b, c)
            if processor.check() == after:
                possible_ops += 1

        if possible_ops >= 3:
            samples_with_three_or_more_possible_op_codes += 1

    return str(samples_with_three_or_more_possible_op_codes)


def part_two(puzzle_input):
    instruction_samples, program = parse_sample_instructions_and_program(puzzle_input)

    processor = Processor.from_samples(instruction_samples)

    processor.run_mapped_program(program)

    return str(processor.check()[0])


def parse_sample_instructions_and_program(puzzle_input):
    sections = puzzle_input.split("\n\n\n\n")

    sample_instructions = []
    for sample in sections[0].split("\n\n"):
        before, instruction, after = sample.split("\n")
        sample_instructions.append(
            (
                tuple(
                    int(x)
                    for x in before.replace("Before: [", "")
                    .replace("]", "")
                    .split(", ")
                ),
                tuple(int(x) for x in instruction.split(" ")),
                tuple(
                    int(x)
                    for x in after.replace("After:  [", "").replace("]", "").split(", ")
                ),
            )
        )

    program = []
    for line in sections[1].split("\n"):
        program.append(tuple(int(x) for x in line.split(" ")))

    return sample_instructions, program


class Processor:
    def __init__(self, initial_register_values, op_code_mappings=None):
        self.registers = list(initial_register_values)
        self.op_list = (
            self.addr,
            self.addi,
            self.mulr,
            self.muli,
            self.banr,
            self.bani,
            self.borr,
            self.bori,
            self.setr,
            self.seti,
            self.gtir,
            self.gtri,
            self.gtrr,
            self.eqir,
            self.eqri,
            self.eqrr,
        )

        if op_code_mappings is None:
            op_code_mappings = {i: i for i in range(16)}
        self.op_code_mappings = {
            src: self.op_list[target] for src, target in op_code_mappings.items()
        }

    @classmethod
    def from_samples(cls, samples):
        possible_op_code_mappings = {i: [j for j in range(16)] for i in range(16)}

        for before, (src_opcode, a, b, c), after in samples:
            possible_targets = list(possible_op_code_mappings[src_opcode])
            for target in possible_targets:
                processor = Processor(before)
                processor.op_list[target](a, b, c)
                if processor.check() != after:
                    if target in possible_op_code_mappings[src_opcode]:
                        possible_op_code_mappings[src_opcode].remove(target)

            if len(possible_op_code_mappings[src_opcode]) > 1:
                continue

            next_opcodes_to_remove = possible_op_code_mappings[src_opcode]
            if next_opcodes_to_remove:
                opcodes_to_remove = next_opcodes_to_remove
                next_opcodes_to_remove = []

                for opcode in opcodes_to_remove:
                    for targets in possible_op_code_mappings.values():
                        if len(targets) == 1 and targets[0] == opcode:
                            continue

                        if opcode in targets:
                            targets.remove(opcode)
        
                        if len(targets) == 1:
                            next_opcodes_to_remove.append(targets[0])

        if any(len(targets) != 1 for targets in possible_op_code_mappings.values()):
            raise ValueError(
                f"one-to-one mapping not determined: {possible_op_code_mappings}"
            )

        op_code_mappings = {
            src: targets[0] for src, targets in possible_op_code_mappings.items()
        }

        return cls((0, 0, 0, 0), op_code_mappings)

    def run_mapped_program(self, program):
        for instruction in program:
            self.run_mapped_instruction(instruction)

    def run_mapped_instruction(self, instruction):
        op, a, b, c = instruction
        self.op_code_mappings[op](a, b, c)

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
