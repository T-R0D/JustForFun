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
from aoc2016.day01 import day01


class TestDay01(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_turn(self):
        self.assertEqual(day01.EAST, day01.turn(day01.NORTH, day01.RIGHT))
        self.assertEqual(day01.SOUTH, day01.turn(day01.EAST, day01.RIGHT))
        self.assertEqual(day01.WEST, day01.turn(day01.SOUTH, day01.RIGHT))
        self.assertEqual(day01.NORTH, day01.turn(day01.WEST, day01.RIGHT))

        self.assertEqual(day01.WEST, day01.turn(day01.NORTH, day01.LEFT))
        self.assertEqual(day01.SOUTH, day01.turn(day01.WEST, day01.LEFT))
        self.assertEqual(day01.EAST, day01.turn(day01.SOUTH, day01.LEFT))
        self.assertEqual(day01.NORTH, day01.turn(day01.EAST, day01.LEFT))

    def test_move(self):
        pass

    def test_part_one(self):
        pass

    def test_part_two(self):
        self.assertEqual('4', day01.part_two('R8, R4, R4, R8'))
