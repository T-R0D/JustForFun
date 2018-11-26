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

import unittest
from aoc2016.day16.solution import DragonCurveGenerator


class TestDay16(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_iterate(self):
        expected = '100'
        actual = DragonCurveGenerator.iterate('1')
        self.assertEqual(expected, actual)

        expected = '001'
        actual = DragonCurveGenerator.iterate('0')
        self.assertEqual(expected, actual)

        expected = '11111000000'
        actual = DragonCurveGenerator.iterate('11111')
        self.assertEqual(expected, actual)

        expected = '1111000010100101011110000'
        actual = DragonCurveGenerator.iterate('111100001010')
        self.assertEqual(expected, actual)

    def test_checksum(self):
        expected = '100'
        actual = DragonCurveGenerator.checksum('110010110100')
        self.assertEqual(expected, actual)

    def test_part_one(self):
        generator = DragonCurveGenerator('10000')
        generator.generate(20)
        self.assertEqual('10000011110010000111', generator.get_data())
        self.assertEqual('01100', generator.get_checksum())
