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
import aoc2016.day02.solution as sut


class TestSolution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        puzzle_input = 'ULL\nRRDDD\nLURDL\nUUUUD'
        expected = '1985'
        self.assertEqual(expected, sut.part_one(puzzle_input))

    def test_part_two(self):
        pass

    def test_move(self):
        self.assertEqual(2, sut.move(5, 'U'))
        self.assertEqual(8, sut.move(5, 'D'))
        self.assertEqual(4, sut.move(5, 'L'))
        self.assertEqual(6, sut.move(5, 'R'))

        self.assertEqual(1, sut.move(1, 'U'))
        self.assertEqual(1, sut.move(1, 'L'))
        self.assertEqual(3, sut.move(3, 'U'))
        self.assertEqual(3, sut.move(3, 'R'))
        self.assertEqual(7, sut.move(7, 'D'))
        self.assertEqual(7, sut.move(7, 'L'))
        self.assertEqual(9, sut.move(9, 'D'))
        self.assertEqual(9, sut.move(9, 'R'))
