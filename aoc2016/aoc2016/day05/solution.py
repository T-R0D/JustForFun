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
import hashlib


PREFIX = '0' * 5


def part_one(puzzle_input):
    password = ''
    for index in range(0, 99999999):
        digest = get_digest(puzzle_input, index)
        if digest.startswith(PREFIX):
            password += digest[5]
            if len(password) == 8:
                break
    return password


def part_two(puzzle_input):
    password = ['-'] * 8
    set_digits = 0x00
    for index in range(0, 99999999):
        digest = get_digest(puzzle_input, index)
        if digest.startswith(PREFIX):
            password, set_digits = place_digit(password, digest, set_digits)
            if set_digits == 0xFF:
                break
    return ''.join(password)


def get_digest(door_id, index):
    h = hashlib.md5()
    h.update(bytes(door_id + str(index), 'utf-8'))
    return h.hexdigest()


def place_digit(password, digest, set_digits=0):
    place = int(digest[5], base=16)
    place_flag = 1 << place
    if place < 8 and (set_digits & place_flag) == 0:
        password[place] = digest[6]
        set_digits |= place_flag
    return password, set_digits
