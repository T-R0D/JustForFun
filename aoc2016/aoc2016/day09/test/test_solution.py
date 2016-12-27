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
from aoc2016.day09 import solution as sut

class TestDay09Solution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_decompress_message(self):
        self.assertEqual('ADVENT', sut.decompress_message('ADVENT'))
        self.assertEqual('ABBBBBC', sut.decompress_message('A(1x5)BC'))
        self.assertEqual('XYZXYZXYZ', sut.decompress_message('(3x3)XYZ'))
        self.assertEqual('ABCBCDEFEFG', sut.decompress_message('A(2x2)BCD(2x2)EFG'))
        self.assertEqual('ABCBCDEFEFG', sut.decompress_message('A(2x2)BCD(2x2)EFG'))
        self.assertEqual('(1x3)A', sut.decompress_message('(6x1)(1x3)A'))
        self.assertEqual('X(3x3)ABC(3x3)ABCY', sut.decompress_message('X(8x2)(3x3)ABCY'))