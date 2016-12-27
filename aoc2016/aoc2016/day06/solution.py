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
    return decode_message(puzzle_input)


def part_two(puzzle_input):
    return decode_message(puzzle_input, -1)


def decode_message(lines, frequency_rank=0):
    counts = get_column_frequencies(lines)
    letters = [sorted(count.items(), key=use_item_value, reverse=True)[frequency_rank][0] for count in counts]
    return ''.join(letters)


def get_column_frequencies(lines):
    line_len = len(lines.split('\n', 1)[0])
    counts = [{} for i in range(line_len)]
    for line in lines.split('\n'):
        for i in range(len(line)):
            letter = line[i]
            counts[i][letter] = counts[i].get(letter, 0) + 1
    return counts


def use_item_value(item):
    return item[1]
