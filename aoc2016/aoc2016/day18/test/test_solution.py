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
import aoc2016.day18.solution as sut


class TestSolution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass


    def test_determine_tile(self):
        expected = sut.TRAP_TILE
        result = sut.determine_tile(sut.TRAP_TILE, sut.TRAP_TILE, sut.SAFE_TILE)
        self.assertEqual(expected, result)

        expected = sut.TRAP_TILE
        result = sut.determine_tile(sut.SAFE_TILE, sut.TRAP_TILE, sut.TRAP_TILE)
        self.assertEqual(expected, result)

        expected = sut.TRAP_TILE
        result = sut.determine_tile(sut.TRAP_TILE, sut.SAFE_TILE, sut.SAFE_TILE)
        self.assertEqual(expected, result)

        expected = sut.TRAP_TILE
        result = sut.determine_tile(sut.SAFE_TILE, sut.SAFE_TILE, sut.TRAP_TILE)
        self.assertEqual(expected, result)

        expected = sut.SAFE_TILE
        result = sut.determine_tile(sut.SAFE_TILE, sut.SAFE_TILE, sut.SAFE_TILE)
        self.assertEqual(expected, result)

        expected = sut.SAFE_TILE
        result = sut.determine_tile(sut.SAFE_TILE, sut.TRAP_TILE, sut.SAFE_TILE)
        self.assertEqual(expected, result)

        expected = sut.SAFE_TILE
        result = sut.determine_tile(sut.TRAP_TILE, sut.SAFE_TILE, sut.TRAP_TILE)
        self.assertEqual(expected, result)

        expected = sut.SAFE_TILE
        result = sut.determine_tile(sut.TRAP_TILE, sut.TRAP_TILE, sut.TRAP_TILE)
        self.assertEqual(expected, result)

    def test_small_case(self):
        expected = 6
        result = sut.count_safe_tiles_in_room('..^^.', 3)
        self.assertEqual(expected, result)

    def test_large_case(self):
        expected = 38
        result = sut.count_safe_tiles_in_room('.^^.^.^^^^', 10)
        self.assertEqual(expected, result)
