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
import itertools

def part_one(puzzle_input):
    # A Good ol' brute force (with a minor optimization) solves this pretty
    # quickly, so I'll leave that code in place.
    # However, we can solve it _even faster_ (in CPU time, not human time)
    # by doing some careful analysis of the assembunny code. Analyzing the
    # code reveals:
    # - The 1st segment computes an initial value of the form:
    #   `(initial register 'a') + (a constant) * (another constant)`
    #   This initial value is essentially used for the initialization for
    #   a sub-loop within the program's infinite loop.
    # - The program falls into a sub-loop that loops until that initial value
    #   is reduced to zero
    # - Yet another sub-loop takes the current value of the
    #   value-that-started-with-the-initial-value and integer divides it by 2
    #   and keeps the remainder
    # - After the yet-another-sub-loop, the program does some "extra"
    #   computation to produce the value that is `out`-ed (1 or 0)
    # - If the value-that-started-with-the-initial-value is not yet 0, the
    #   outermost sub-loop is restarted
    # - Otherwise, The infinite loop is restarted and the "initial-value" is
    #   restored
    # - Repeat forever
    # There's probably a more mathy way to solve this, but reversing the
    # operations basically boils down to finding the first multiple of two
    # that is greater than `(a constant) * (another constant)` that can be
    # odd, then even, then odd... as we divide by 2 repeatedly. We can do this
    # starting with 0 (since we know the sequence will collapse to zero to
    # invoke the infinite loop) by alternating multiplying by 2 and
    # multiplying by 2 and adding 1. Now, which to do first? I'm not sure,
    # so let's try both and take the lesser. Once we have that first multiple
    # of 2 that is greater than `(a constant) * (another constant)`, we can do
    # some simple subtraction to find the correct initial register "a" value.

    # Manual inspection aided fast solution:
    #return str(find_initial_register_a_value_after_manual_inspection(7, 365))

    # "Brute force", "slow" solution:
    program = parse_input(puzzle_input)

    best_initial_input = 0
    for initial in range(999_999):
        cpu = CarrotProcessingUnit()
        cpu.set_program(program)
        cpu.set_register(CarrotProcessingUnit.REG_A, initial)
        sample = cpu.sample_clock_signal()

        if all((j & 1) == x for j, x in enumerate(sample)):
            best_initial_input = initial
            break

    return str(best_initial_input)


def part_two(puzzle_input):
    return "Merry Christmas!"


def find_initial_register_a_value_after_manual_inspection(constant_a, constant_b):
    constant_product = constant_a * constant_b

    candidate_a = 0
    for i in itertools.cycle((0, 1)):
        candidate_a *= 2
        candidate_a += i
        if candidate_a > constant_product:
            break

    candidate_b = 0
    for i in itertools.cycle((1, 0)):
        candidate_b *= 2
        candidate_b += i
        if candidate_b > constant_product:
            break

    candidate = min(candidate_a, candidate_b)
    return candidate - constant_product



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
    OP_OUT = "out"

    def __init__(self):
        self.program_counter = 0
        self.registers = {
            CarrotProcessingUnit.REG_A: 0,
            CarrotProcessingUnit.REG_B: 0,
            CarrotProcessingUnit.REG_C: 0,
            CarrotProcessingUnit.REG_D: 0,
        }
        self.program = []
        self.clock_signal = []

    def set_program(self, program):
        self.program = program

    def run_program(self):
        self.program_counter = 0
        program_length = len(self.program)
        while self.program_counter < program_length:
            instruction, x, y = self.program[self.program_counter]
            self.execute_instruction(instruction, x, y)

    def sample_clock_signal(self):
        self.program_counter = 0
        program_length = len(self.program)
        # I assume 100 is a long enough sample that we can be sure we have an
        # infinite loop.
        while self.program_counter < program_length and len(self.clock_signal) < 100:
            instruction, x, y = self.program[self.program_counter]
            self.execute_instruction(instruction, x, y)
            if instruction == CarrotProcessingUnit.OP_OUT:
                i = len(self.clock_signal) - 1
                if i & 1 != self.clock_signal[i]:
                    break

        return self.clock_signal

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
        elif instruction == CarrotProcessingUnit.OP_OUT:
            self.program_counter += self.out(x)
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

    def out(self, reg):
        self.clock_signal.append(self.check_register(reg))
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
