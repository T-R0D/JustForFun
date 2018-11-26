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
    builder = Machine.Builder()
    builder.add_disc(17, 15)
    builder.add_disc(3, 2)
    builder.add_disc(19, 4)
    builder.add_disc(13, 2)
    builder.add_disc(7, 2)
    builder.add_disc(5, 0)
    machine = builder.build()

    return str(machine.find_full_drop_start_time())


def part_two(puzzle_input):
    builder = Machine.Builder()
    builder.add_disc(17, 15)
    builder.add_disc(3, 2)
    builder.add_disc(19, 4)
    builder.add_disc(13, 2)
    builder.add_disc(7, 2)
    builder.add_disc(5, 0)
    machine = builder.build()

    machine.add_disc(11, 0)

    return str(machine.find_full_drop_start_time())


class Machine(object):
    class Builder(object):
        def __init__(self):
            self.discs = []

        def add_disc(self, slots, start_pos=0):
            self.discs.append(Machine.Disc(slots, start_pos))

        def build(self):
            return Machine(self.discs)

    class Disc(object):
        def __init__(self, slots, start_pos=0):
            self.start_pos = start_pos
            self.slots = slots
            self.pos = start_pos

        def reset(self):
            self.pos = self.start_pos

        def advance(self, steps=1):
            self.pos = (self.pos + steps) % self.slots

        def is_open(self):
            return self.pos == 0

    def __init__(self, discs):
        self.discs = discs

    def find_full_drop_start_time(self):
        full_drop_achieved = False
        start_time = -1
        while not full_drop_achieved:
            start_time += 1
            full_drop_achieved = self.drop(start_time)

        return start_time

    def drop(self, t=0):
        self.reset_discs()

        level = 0
        self.advance_discs(t)
        for _ in range(len(self.discs)):
            self.advance_discs(1)

            if not self.discs[level].is_open():
                break

            level += 1

        return level == len(self.discs)

    def advance_discs(self, t):
        if t < 1:
            return

        for disc in self.discs:
            disc.advance(t)

    def reset_discs(self):
        for disc in self.discs:
            disc.reset()

    def add_disc(self, slots, start_pos=0) -> None:
        self.discs.append(Machine.Disc(slots, start_pos))
