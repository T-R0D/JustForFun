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
import aoc2016.day05.solution as sut


class TestDay05Solution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_get_digest(self):
        self.assertEqual('1', sut.get_digest('abc', 3231929)[5])
        self.assertEqual('8', sut.get_digest('abc', 5017308)[5])
        self.assertEqual('f', sut.get_digest('abc', 5278568)[5])

    def test_part_one(self):
        self.assertEqual('18f47a30', sut.part_one('abc'))

    def test_part_two(self):
        self.assertEqual('05ace8e3', sut.part_two('abc'))
