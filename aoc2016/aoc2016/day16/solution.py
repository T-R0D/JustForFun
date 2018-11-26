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
    generator = DragonCurveGenerator('11101000110010100')
    generator.generate(272)
    return str(generator.get_checksum())


def part_two(puzzle_input):
    generator = DragonCurveGenerator('11101000110010100')
    generator.generate(35651584)
    return str(generator.get_checksum())


class DragonCurveGenerator(object):
    def __init__(self, initial_data = '0'):
        self.initial_data = initial_data
        self.data = initial_data

    def generate(self, required_bits):
        data = self.initial_data
        while len(data) < required_bits:
            data = self.iterate(data)
        self.data = data[:required_bits]

    def get_data(self):
        return self.data

    def get_checksum(self):
        return DragonCurveGenerator.checksum(self.data)

    def data_len(self):
        return len(self.data)

    @staticmethod
    def iterate(a):
        b = reversed(a)
        b = map(lambda x: '0' if x == '1' else '1', b)
        return a + '0' + ''.join(b)

    @staticmethod
    def checksum(data):
        result = ''.join(['1' if a == b else '0' for a, b in zip(data[::2], data[1::2])])
        if len(result) % 2 == 0:
            return DragonCurveGenerator.checksum(result)
        else:
            return result


