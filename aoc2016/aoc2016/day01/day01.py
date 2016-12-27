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


class Day01(object):
    def __init__(self):
        self.heading = 0
        self.x = 0
        self.y = 0
        self.visited = []



    def part_a(self, puzzle_input):
        for instruction in puzzle_input.split(', '):
            direction = instruction[0]
            if direction == 'R':
                self.heading += 1
            else:
                self.heading -= 1
            self.heading %= 4

            distance = int(instruction[1: ])
            if self.heading == 0:
                self.y += distance
            elif self.heading == 1:
                self.x += distance
            elif self.heading == 2:
                self.y -= distance
            else:
                self.x -= distance

        return str(abs(self.x) + abs(self.y))

    def part_b(self, puzzle_input):
        while (self.x, self.y) not in self.visited:
            pass


    def move_once(self, instruction):
        pass

RIGHT = 'R'
LEFT = 'L'

N_DIRECTIONS = 4
NORTH = 0
EAST = 1
SOUTH = 2
WEST = 3


def part_two(puzzle_input):
    heading = NORTH
    location = (0, 0)
    visited = set()

    for instruction in puzzle_input.split(', '):
        direction = instruction[0]
        distance = int(instruction[1: ])
        heading = turn(heading, direction)
        location = move(location, heading, distance)

        for block in range(distance):
            location = move(location, heading, 1)

            if location in visited:
                break

            visited.add(location)

    return str(abs(location[0]) + abs(location[1]))

def turn(heading, direction):
    if direction == RIGHT:
        return (heading + 1) % N_DIRECTIONS
    else:
        return (heading - 1) % N_DIRECTIONS


def move(start, heading, distance, ):
    x, y = start
    if heading == NORTH:
        return x, y + 1
    elif heading == EAST:
        return x + 1, y
    elif heading == SOUTH:
        return x, y - 1
    elif heading == WEST:
        return x - 1, y
