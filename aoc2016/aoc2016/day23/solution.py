# This file is part of aoc2016.
#
# aoc2016 is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# aoc2016 is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with aoc2016.  If not, see <http://www.gnu.org/licenses/>.


def part_one(puzzle_input):
    program = parse_input(puzzle_input)

    result = run_computation(program, 7)

    return str(result)


def part_two(puzzle_input):
    # The program will run for a long time with the larger starting input.
    # Careful inspection of the program's code (with _some_ execution of it
    # by hand) reveals a few facts:
    # - The program works in essentially 3 phases, with the 2nd phase nested
    #   in the first
    # - The 1st phase is a factorial computation
    # - The 2nd phase utilizes the single `tgl` instruction to rewrite
    #   instructions in the 3rd phase to remove infinite loops and make
    #   operations become only "additive"
    # - The 3rd phase takes the result of the 1st and adds the result of 2
    #   constants multiplied together
    # I couldn't figure out a way to automate this analysis, so I
    # hardcode(-ish) a solution here.

    return str(compute_program_after_manual_analysis(12, 73, 71))


def parse_input(puzzle_input):
    lines = puzzle_input.split("\n")
    return [parse_input_line(line) for line in lines]


def parse_input_line(line):
    parts = line.split(" ")

    if len(parts[1:]) == 1:
        return tuple(parts + [None])
    else:
        return tuple(parts)


def run_computation(program, initial_val):
    cpu = CarrotProcessingUnit()
    cpu.set_register(CarrotProcessingUnit.REG_A, initial_val)
    cpu.set_program(program)
    cpu.run_program()
    return cpu.check_register(CarrotProcessingUnit.REG_A)


class CarrotProcessingUnit(object):
    REG_A = "a"
    REG_B = "b"
    REG_C = "c"
    REG_D = "d"

    OP_CPY = "cpy"
    OP_INC = "inc"
    OP_DEC = "dec"
    OP_JNZ = "jnz"
    OP_TGL = "tgl"

    def __init__(self):
        self.program_counter = 0
        self.registers = {
            CarrotProcessingUnit.REG_A: 0,
            CarrotProcessingUnit.REG_B: 0,
            CarrotProcessingUnit.REG_C: 0,
            CarrotProcessingUnit.REG_D: 0,
        }
        self.program = []

    def set_program(self, program):
        self.program = program

    def run_program(self):
        self.program_counter = 0
        program_length = len(self.program)
        while self.program_counter < program_length:
            instruction, x, y = self.program[self.program_counter]
            self.execute_instruction(instruction, x, y)

    def execute_instruction(self, instruction, x, y):
        if instruction == CarrotProcessingUnit.OP_CPY:
            self.program_counter += self.copy(x, y)
        elif instruction == CarrotProcessingUnit.OP_INC:
            self.program_counter += self.increment(x)
        elif instruction == CarrotProcessingUnit.OP_DEC:
            self.program_counter += self.decrement(x)
        elif instruction == CarrotProcessingUnit.OP_JNZ:
            self.program_counter += self.jump_non_zero(x, y)
        elif instruction == CarrotProcessingUnit.OP_TGL:
            self.program_counter += self.toggle(x)
        else:
            raise Exception("{} is not a valid instruction.".format(instruction))

    def copy(self, val, reg):
        if reg not in self.registers:
            raise Exception("CPY must target a register.")
        self.registers[reg] = self.get_value_of_operand(val)
        return 1

    def increment(self, reg):
        if reg not in self.registers:
            raise Exception("INC must target a register.")
        self.registers[reg] += 1
        return 1

    def decrement(self, reg):
        if reg not in self.registers:
            raise Exception("DEC must target a register.")
        self.registers[reg] -= 1
        return 1

    def jump_non_zero(self, operand, count):
        return (
            1
            if self.get_value_of_operand(operand) == 0
            else self.get_value_of_operand(count)
        )

    def toggle(self, reg):
        offset = self.check_register(reg)
        instruction_pointer = self.program_counter + offset

        if instruction_pointer < 0 or len(self.program) <= instruction_pointer:
            return 1

        instruction, x, y = self.program[self.program_counter + offset]

        new_instruction = CarrotProcessingUnit.OP_INC
        if instruction == CarrotProcessingUnit.OP_INC:
            new_instruction = CarrotProcessingUnit.OP_DEC
        elif instruction == CarrotProcessingUnit.OP_CPY:
            new_instruction = CarrotProcessingUnit.OP_JNZ
        elif instruction == CarrotProcessingUnit.OP_JNZ:
            new_instruction = CarrotProcessingUnit.OP_CPY

        self.program[instruction_pointer] = (new_instruction, x, y)

        return 1

    def check_register(self, reg):
        if reg not in self.registers:
            raise Exception("{} is not a valid register.")
        return self.registers[reg]

    def set_register(self, reg, val):
        if reg not in self.registers:
            raise Exception("{} is not a valid register.")
        self.registers[reg] = val

    def get_value_of_operand(self, operand):
        if operand in self.registers:
            return self.registers[operand]
        else:
            return int(operand)


def compute_program_after_manual_analysis(initial_value, constant_a, constant_b):
    return factorial(initial_value) + (constant_a * constant_b)


def factorial(x):
    result = 1
    for b in range(1, x + 1):
        result *= b
    return result
