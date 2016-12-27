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
import re

LETTER_BASE = ord('a')
N_LETTERS = ord('z') - LETTER_BASE + 1
TARGET_ROOM_NAME = 'Northpole Object storage'.lower()


def part_one(puzzle_input):
    total = 0
    for line in puzzle_input.split('\n'):
        name, sector, checksum = decompose_line(line)
        computed_checksum = compute_checksum(name.replace('-', ''))
        if checksum == computed_checksum:
            total += int(sector)

    return str(total)


def part_two(puzzle_input):
    for line in puzzle_input.split('\n'):
        encrypted_name, sector, checksum = decompose_line(line)
        computed_checksum = compute_checksum(encrypted_name.replace('-', ''))
        if checksum == computed_checksum:
            room_name = rotate_name(encrypted_name, int(sector))
            if room_name == TARGET_ROOM_NAME:
                return sector


def decompose_line(line):
    groups =  re.match('([a-z\-]+)(\d+)\[([a-z]+)\]', line).groups()
    return groups[0][0:-1], groups[1], groups[2]


def compute_checksum(name):
    counts = {}
    for letter in name:
        counts[letter] = counts.get(letter, 0) + 1

    sorted_by_count = sorted(counts.items(), key=letter_count_key, reverse=True)
    return ''.join([item[0] for item in sorted_by_count[0:5]])


def letter_count_key(count_pair):
    letter, count = count_pair
    return str(count) + str(ord('z') + 10 - ord(letter))


def rotate_name(encrypted_name, key):
    name = [' '] * len(encrypted_name)
    for i in range(len(encrypted_name)):
        char = encrypted_name[i]
        if char == '-':
            name[i] = ' '
        else:
            name[i] = rotate_letter(char, key)
    return ''.join(name)


def rotate_letter(letter, key):
    return chr(((ord(letter) - LETTER_BASE + key) % N_LETTERS) + LETTER_BASE)
