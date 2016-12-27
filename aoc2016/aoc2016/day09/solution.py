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


# 692043 < x
def part_one(puzzle_input):
    return str(len(decompress_message(puzzle_input)))


def part_two(puzzle_input):
    return str(get_nested_decompressed_length(puzzle_input))


def decompress_message(compressed_message):
    buffer = []
    message = []
    decompressing = False
    parsing_marker = False
    chars_to_repeat, times = 0, 0
    for char in compressed_message:
        if not parsing_marker and not decompressing:
            if char == '(':
                parsing_marker = True
            else:
                message.append(char)
        elif parsing_marker:
            if char == ')':
                chars_to_repeat, times = parse_compression_marker(''.join(buffer))
                buffer.clear()
                parsing_marker = False
                decompressing = True
            else:
                buffer.append(char)
        elif decompressing:
            buffer.append(char)
            chars_to_repeat -= 1
            if chars_to_repeat == 0:
                for j in range(times):
                    message.extend(buffer)
                buffer.clear()
                decompressing = False

    return ''.join(message)


def parse_compression_marker(marker):
    chars_to_repeat, times = marker.split('x')
    return int(chars_to_repeat), int(times)


def get_nested_decompressed_length(compressed_message):
    length = 0
    buffer = []
    parsing_marker = False
    decompressing = False
    chars_to_repeat, times = 0, 0
    for char in compressed_message:
        if not parsing_marker and not decompressing:
            if char == '(':
                parsing_marker = True
            else:
                length += 1
        elif parsing_marker:
            if char == ')':
                chars_to_repeat, times = parse_compression_marker(''.join(buffer))
                buffer.clear()
                parsing_marker = False
                decompressing = True
            else:
                buffer.append(char)
        elif decompressing:
            buffer.append(char)
            chars_to_repeat -= 1
            if chars_to_repeat == 0:
                length += times * get_nested_decompressed_length(''.join(buffer))
                buffer.clear()
                decompressing = False

    return length
