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

import aoc2016.day19.solution as sut
import unittest

class TestSolution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_odd_number_last_elf_standing(self):
        expected = 3
        result = sut.brute_force_find_last_elf_standing(5)
        self.assertEqual(expected, result)

    def test_even_number_last_elf_standing(self):
        expected = 1
        result = sut.brute_force_find_last_elf_standing(8)
        self.assertEqual(expected, result)

