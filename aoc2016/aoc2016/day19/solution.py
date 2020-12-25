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
    return brute_force_find_last_elf_standing(3005290)


def part_two(puzzle_input):
    return last_elf_standing_cross_circle_steal_brute_force(3005290)


def find_last_elf_standing(n_elves):
    ret = 1
    generator = elf_circle_result()
    for i in range(n_elves):
        ret = next(generator)
        print('For {} elves, the last elf standing is {}'.format(i + 1, ret))
    return ret


def brute_force_find_last_elf_standing(n_elves):
    elves = list(range(1, n_elves + 1, 2))  # Keep only odds as a minor optimization.
    if n_elves > 1 and n_elves & 1 == 1:
        del elves[0]

    while len(elves) > 1:
        elves_remaining = len(elves)
        elves_to_delete = 0
        if elves_remaining & 1 == 0:
            elves_to_delete = reversed(range(1, elves_remaining + 1, 2))
        else:
            elves_to_delete = itertools.chain(reversed(range(1, elves_remaining - 1, 2)), [0])
        for i in elves_to_delete:
            del elves[i]
    return elves[0]

def last_elf_standing_cross_circle_steal_brute_force(n_elves):
    elves = list(range(1, n_elves + 1))

    while len(elves) > 1:
        for i in range(len(elves) // 2):
            elves_in_circle = len(elves)
            removal_index = (i + (elves_in_circle // 2)) % elves_in_circle



def screw_around():
    # brute_force_find_last_elf_standing(6)
    for i in range(1, 17 + 1):
        print('=== {} elves in the circle ==='.format(i))

        print(brute_force_find_last_elf_standing(i))

    print(brute_force_find_last_elf_standing(3005290))  # 1816277

if __name__ == '__main__':
    screw_around()