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

UP = 'U'
DOWN = 'D'
LEFT = 'L'
RIGHT = 'R'


def part_one(puzzle_input):
    position = 5
    code = ''
    for line in puzzle_input.split('\n'):
        for instruction in line:
            position = move(position, instruction)
        code += str(position)
    return code


def part_two(puzzle_input):
    position = '5'
    code = ''
    for line in puzzle_input.split('\n'):
        for instruction in line:
            position = KEY_PAD[position].get(instruction, position)
        code += position
    return code


def move(start, instruction):
    if instruction == UP:
        if start <= 3:
            return start
        return start - 3
    elif instruction == DOWN:
        if 6 < start:
            return start
        return start + 3
    elif instruction == LEFT:
        if (start % 3) == 1:
            return start
        return start - 1
    elif instruction == RIGHT:
        if start % 3 == 0:
            return start
        return start + 1


KEY_PAD = {
    '1': {
        DOWN: '3',
    },
    '2': {
        RIGHT: '3',
        DOWN: '6',
    },
    '3': {
        UP: '1',
        DOWN: '7',
        LEFT: '2',
        RIGHT: '4',
    },
    '4': {
        LEFT: '3',
        DOWN: '8',
    },
    '5': {
        RIGHT: '6',
    },
    '6': {
        UP: '2',
        DOWN: 'A',
        LEFT: '5',
        RIGHT: '7',
    },
    '7': {
        UP: '3',
        DOWN: 'B',
        LEFT: '6',
        RIGHT: '8',
    },
    '8': {
        UP: '4',
        DOWN: 'C',
        LEFT: '7',
        RIGHT: '9',
    },
    '9': {
        LEFT: '8',
    },
    'A': {
        UP: '6',
        RIGHT: 'B',
    },
    'B': {
        UP: '7',
        DOWN: DOWN,
        LEFT: 'A',
        RIGHT: 'C',
    },
    'C': {
        UP: '8',
        LEFT: 'B',
    },
    DOWN: {
        UP: 'B',
    },
}
