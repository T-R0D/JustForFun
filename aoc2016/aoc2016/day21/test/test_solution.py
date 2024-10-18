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
import aoc2016.day21.solution as day21


TEST_INSTRUCTIONS = """swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d"""

class TestSolution(unittest.TestCase):
    def test_scramble_password_correctly_scrambles_abcde_into_decab(self):
        password = "abcde"
        expected = "decab"

        instructions = day21.parse_instruction_list(TEST_INSTRUCTIONS)
        result = day21.scramble_password(password, instructions)

        self.assertEqual(result, expected)
