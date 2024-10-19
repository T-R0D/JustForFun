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
import aoc2016.day23.solution as day23


TEST_PROGRAM = """cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a"""

class TestSolution(unittest.TestCase):
    def test_part_one_runs_test_program_correctly(self):
        expected = "3"

        result = day23.part_one(TEST_PROGRAM)

        self.assertEqual(result, expected)
