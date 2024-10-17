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
import aoc2016.day20.solution as day20


class TestSolution(unittest.TestCase):
    def test_part_one_solves_correctly_with_example_block_list(self):
        puzzle_input = "5-8\n0-2\n4-7"
        expected = "3"

        result = day20.part_one(puzzle_input)

        self.assertEqual(result, expected)

    def test_part_two_solves_correctly_with_example_block_list(self):
        puzzle_input = "5-8\n0-2\n4-7"
        expected = str(4294967288)

        result = day20.part_two(puzzle_input)

        self.assertEqual(result, expected)
