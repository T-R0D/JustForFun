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
import aoc2016.day24.solution as day24

TEST_MAP = """###########
#0.1.....2#
#.#######.#
#4.......3#
###########"""


class TestSolution(unittest.TestCase):
    def test_part_one_finds_shortest_path_to_cover_all_numbers(self):
        expected = "14"

        result = day24.part_one(TEST_MAP)

        self.assertEqual(result, expected)
