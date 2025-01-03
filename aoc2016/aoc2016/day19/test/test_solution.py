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
import aoc2016.day19.solution as day19


class TestSolution(unittest.TestCase):
    def test_part_one_solves_correctly_with_odd_input(self):
        expected = "3"
        result = day19.part_one("5")
        self.assertEqual(expected, result)

    def test_part_one_solves_correctly_with_even_input(self):
        expected = "1"
        result = day19.part_one("8")
        self.assertEqual(expected, result)

    def test_part_two_solves_correctly_with_odd_input(self):
        expected = "2"
        result = day19.part_two("5")
        self.assertEqual(expected, result)

    def test_part_two_solves_correctly_with_even_input(self):
        expected = "7"
        result = day19.part_two("8")
        self.assertEqual(expected, result)
