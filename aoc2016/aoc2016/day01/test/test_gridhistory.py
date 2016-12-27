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
from aoc2016.day01 import gridhistory

class TestGridHistory(unittest.TestCase):
    def setUp(self):
        self.sut = gridhistory.GridHistory()

    def tearDown(self):
        self.sut = None

    def test_add_point(self):
        start = (0, 0)
        end = (0, 0)
        self.sut.add_straight_path(*start, *end)
        expected = {
            0: [(0, 0)],
        }
        self.assertDictEqual(expected, self.sut.history)

    def test_add_horizontal_path(self):
        start = (0, 0)
        end = (5, 0)
        self.sut.add_straight_path(*start, *end)
        expected = {
            0: [(0, 0)],
            1: [(0, 0)],
            2: [(0, 0)],
            3: [(0, 0)],
            4: [(0, 0)],
            5: [(0, 0)],
        }
        self.assertDictEqual(expected, self.sut.history)

    def test_add_vertical_path(self):
        start = (0, -5)
        end = (0, 5)
        self.sut.add_straight_path(*start, *end)
        expected = {
            0: [(-5, 5)],
        }
        self.assertDictEqual(expected, self.sut.history)

    def test_add_disjoint_paths(self):
        start = (0, -5)
        end = (0, 5)
        self.sut.add_straight_path(*start, *end)

        start = (1, -5)
        end = (1, 5)
        self.sut.add_straight_path(*start, *end)
        expected = {
            0: [(-5, 5)],
            1: [(-5, 5)],
        }
        self.assertDictEqual(expected, self.sut.history)

    def test_add_overlapping_paths(self):
        start = (0, -5)
        end = (0, 5)
        self.sut.add_straight_path(*start, *end)

        start = (0, 0)
        end = (0, 10)
        self.sut.add_straight_path(*start, *end)
        expected = {
            0: [(-5, 10)],
        }
        self.assertDictEqual(expected, self.sut.history)
