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
import aoc2016.day13.solution as solution


class TestSolution(unittest.TestCase):
    DUMMY_DESIGNER_NUMBER = 0
    DUMMY_EXIT = (0,0)

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_count_bits(self):
        room = solution.DilbertLand(self.DUMMY_DESIGNER_NUMBER, self.DUMMY_EXIT)

        n_bits = room.count_bits(0)
        self.assertEqual(0, n_bits, '0 does not have {} bits'.format(n_bits))

        n_bits = room.count_bits(1)
        self.assertEqual(1, n_bits, '1 does not have {} bits'.format(n_bits))

        n_bits = room.count_bits(5)
        self.assertEqual(2, n_bits, '5 does not have {} bits'.format(n_bits))

        n_bits = room.count_bits(1742346774)
        self.assertEqual(16, n_bits, '1742346774 does not have {} bits'.format(n_bits))

    def test_find_exit_path(self):
        room = solution.DilbertLand(10, (7, 4))
        path = room.find_route_to_exit()
        self.assertEqual(11, len(path))

    def test_find_highest_coverage(self):
        room = solution.DilbertLand(10, (7, 4))
        path = room.find_highest_coverage_path(10)
        print(path)
        self.assertEqual(9, len(set(path)))
