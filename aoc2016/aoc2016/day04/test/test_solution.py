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
from aoc2016.day04 import solution as sut


class TestSolution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_decompose_line(self):
        self.assertTupleEqual(('aaaaa-bbb-z-y-x', '123', 'abxyz'), sut.decompose_line('aaaaa-bbb-z-y-x-123[abxyz]'))
        self.assertTupleEqual(('a-b-c-d-e-f-g-h', '987', 'abcde'), sut.decompose_line('a-b-c-d-e-f-g-h-987[abcde]'))
        self.assertTupleEqual(('not-a-real-room', '404', 'oarel'), sut.decompose_line('not-a-real-room-404[oarel]'))
        self.assertTupleEqual(('totally-real-room', '200', 'decoy'), sut.decompose_line('totally-real-room-200[decoy]'))

    def test_compute_checksum(self):
        self.assertEqual('abxyz', sut.compute_checksum('aaaaabbbzyx'))
        self.assertEqual('abcde', sut.compute_checksum('abcdefgh'))
        self.assertEqual('oarel', sut.compute_checksum('notarealroom'))
        self.assertNotEqual('decoy', sut.compute_checksum('totallyrealroom'))

    def test_rotate_name(self):
        self.assertEqual('very encrypted name', sut.rotate_name('qzmt-zixmtkozy-ivhz', 343))
