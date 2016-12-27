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
    n_triangles = 0
    for line in puzzle_input.split('\n'):
        a, b, c = decompose_line(line)
        if is_triangle(a, b, c):
            n_triangles += 1

    return str(n_triangles)


def part_two(puzzle_input):
    n_triangles = 0
    input_lines = puzzle_input.split('\n')
    for line_group in zip(input_lines[::3], input_lines[1::3], input_lines[2::3],):
        line_group = [decompose_line(line) for line in line_group]
        for i in range(3):
            sides = [line[i] for line in line_group]
            if is_triangle(*sides):
                n_triangles += 1

    return str(n_triangles)


def is_triangle(a, b, c):
    a, b, c = sorted((a, b, c))
    return a + b > c


def decompose_line(line):
    return list(map(int, line.strip().split()))
