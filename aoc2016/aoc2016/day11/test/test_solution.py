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
from aoc2016.common.search import graph_search
from aoc2016.day11 import solution

TEST_INPUT = """The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant."""


class TestDay11Solution(unittest.TestCase):
    def test_get_element_to_floor_locations_captures_things_correctly(self):
        expected = {
            "hydrogen": [1, 0],
            "lithium": [2, 0],
        }

        result = solution.get_element_to_floor_locations(TEST_INPUT)

        self.assertDictEqual(result, expected)

    def test_part_one_solves_correctly(self):
        expected = "11"

        result = solution.part_one(TEST_INPUT)

        self.assertEqual(result, expected)

    def test_part_two_solves_correctly(self):
        expected = "No solution found"

        result = solution.part_two(TEST_INPUT)

        self.assertEqual(result, expected)
    
