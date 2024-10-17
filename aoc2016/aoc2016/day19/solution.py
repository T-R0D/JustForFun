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

# The cool kids either solved this analytically (O(1)), or (like in part 2)
# used 2 dequeues that they shuffled elves in and out of. I was a little more
# straightforward at the cost of a little time (~2s each part).

def part_one(puzzle_input):
    elves = [Elf(i + 1) for i in range(int(puzzle_input))]

    while len(elves) > 1:
        safe_elves = []
        for i, elf in enumerate(elves):
            if elf.skip:
                continue

            elves[(i + 1) % len(elves)].skip = True
            safe_elves.append(elf)

        elves = safe_elves

    return str(elves[0].i)


def part_two(puzzle_input):
    elves = [Elf(i + 1) for i in range(int(puzzle_input))]

    while len(elves) > 1:
        n_remaining_elves = len(elves)
        n_skipped = 0
        for i, elf in enumerate(elves):
            if elf.skip:
                n_skipped -= 1
                continue

            next_elf_to_skip = (i + (n_remaining_elves // 2) + n_skipped) % len(elves)
            elves[next_elf_to_skip].skip = True
            n_skipped += 1
            n_remaining_elves -= 1

        elves = [elf for elf in elves if not elf.skip]

    return str(elves[0].i)


class Elf:
    def __init__(self, i):
        self.i = i
        self.skip = False
