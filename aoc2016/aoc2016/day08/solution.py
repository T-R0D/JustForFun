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
from aoc2016.day08.screen import Screen


def part_one(puzzle_input):
    screen = Screen(height=6, width=50)
    for line in puzzle_input.split('\n'):
        parts = line.split()
        if parts[0] == 'rect':
            width, height = parts[1].split('x')
            screen.rect(height=int(height), width=int(width))
        elif parts[1] == 'row':
            row = int(parts[2].replace('y=', ''))
            by = int(parts[4])
            screen.rotate_row(row=row, by=by)
        elif parts[1] == 'column':
            col = int(parts[2].replace('x=', ''))
            by = int(parts[4])
            screen.rotate_col(col=col, by=by)

    return str(screen.pixels_on())


def part_two(puzzle_input):
    screen = Screen(height=6, width=50)
    for line in puzzle_input.split('\n'):
        parts = line.split()
        if parts[0] == 'rect':
            width, height = parts[1].split('x')
            screen.rect(height=int(height), width=int(width))
        elif parts[1] == 'row':
            row = int(parts[2].replace('y=', ''))
            by = int(parts[4])
            screen.rotate_row(row=row, by=by)
        elif parts[1] == 'column':
            col = int(parts[2].replace('x=', ''))
            by = int(parts[4])
            screen.rotate_col(col=col, by=by)

    return str(screen)
