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
import aoc2016.day07.solution as sut


class TestDay07Solution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_contains_abba(self):
        self.assertTrue(sut.contains_abba('abba'))
        self.assertTrue(sut.contains_abba('xyyx'))
        self.assertTrue(sut.contains_abba('bddb'))
        self.assertTrue(sut.contains_abba('ioxxoj'))
        self.assertFalse(sut.contains_abba('aaaa'))

    def test_supports_tls(self):
        self.assertTrue(sut.supports_tls('abba[mnop]qrst'))
        self.assertTrue(sut.supports_tls('ioxxoj[asdfgh]zxcvbn'))
        self.assertTrue(sut.supports_tls('abcd[efgh]ijkl[mnop]qrrq'))

        self.assertFalse(sut.supports_tls('aaaa[qwer]tyui'))
        self.assertFalse(sut.supports_tls('abcd[bddb]xyyx'))
        self.assertFalse(sut.supports_tls('abcd[effe]ijkl[mnop]qrst'))

    def test_find_all_aba(self):
        self.assertEqual({('a', 'b')}, sut.find_all_aba('aba'))
        self.assertEqual(set(), sut.find_all_aba('aaa'))
        self.assertEqual({('z', 'a'), ('z', 'b')}, sut.find_all_aba('zazbz'))

    def test_contains_corresponding_bab(self):
        self.assertTrue(sut.contains_corresponding_bab('kek', {('e', 'k')}))
        self.assertTrue(sut.contains_corresponding_bab('bzb', {('z', 'a'), ('z', 'b')}))

    def test_supports_ssl(self):
        self.assertTrue(sut.supports_ssl('aba[bab]xyz'))
        self.assertTrue(sut.supports_ssl('aaa[kek]eke'))
        self.assertTrue(sut.supports_ssl('zazbz[bzb]cdb'))

        self.assertFalse(sut.supports_ssl('xyx[xyx]xyx'))

    def test_part_one(self):
        pass

    def test_part_two(self):
        pass
