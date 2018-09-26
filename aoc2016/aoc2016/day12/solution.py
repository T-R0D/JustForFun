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
import numbers


def part_one(puzzle_input):
    program = parse_input(puzzle_input)
    cpu = CarrotProcessingUnit()
    cpu.set_program(program)
    cpu.run_program()
    return cpu.check_register(CarrotProcessingUnit.REG_A)

def part_two(puzzle_input):
    program = parse_input(puzzle_input)
    cpu = CarrotProcessingUnit()
    cpu.set_program(program)
    cpu.set_register(CarrotProcessingUnit.REG_C, 1)
    cpu.run_program()
    return cpu.check_register(CarrotProcessingUnit.REG_A)

def parse_input(puzzle_input):
    lines = puzzle_input.split('\n')
    return [parse_input_line(line) for line in lines]

def parse_input_line(line):
    parts = line.split(' ')

    if len(parts[1:]) == 1:
        return tuple(parts + [None])
    else:
        return tuple(parts)


class CarrotProcessingUnit(object):
    REG_A = 'a'
    REG_B = 'b'
    REG_C = 'c'
    REG_D = 'd'

    OP_CPY = 'cpy'
    OP_INC = 'inc'
    OP_DEC = 'dec'
    OP_JNZ = 'jnz'

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
        else:
            raise Exception('{} is not a valid instruction.'.format(instruction))


    def copy(self, val, reg):
        if reg not in self.registers.keys():
            raise Exception('CPY must target a register.')
        self.registers[reg] = self.get_value_of_operand(val)
        return 1

    def increment(self, reg):
        if reg not in self.registers.keys():
            raise Exception('INC must target a register.')
        self.registers[reg] += 1
        return 1

    def decrement(self, reg):
        if reg not in self.registers.keys():
            raise Exception('DEC must target a register.')
        self.registers[reg] -= 1
        return 1

    def jump_non_zero(self, operand, count):
        return 1 if self.get_value_of_operand(operand) == 0 else self.get_value_of_operand(count)

    def check_register(self, reg):
        if reg not in self.registers.keys():
            raise Exception('{} is not a valid register.')
        return self.registers[reg]

    def set_register(self, reg, val):
        if reg not in self.registers.keys():
            raise Exception('{} is not a valid register.')
        self.registers[reg] = val

    def get_value_of_operand(self, operand):
        if operand in self.registers.keys():
            return self.registers[operand]
        else:
            return int(operand)