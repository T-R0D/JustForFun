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
import aoc2016.day14.solution as solution
import hashlib


class TestSolution(unittest.TestCase):

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_blah(self):
        for x in range(22728, 23729):
            h = hashlib.md5('abc{}'.format(x).encode()).digest().hex()
            char_3 = solution.OTPGenerator.find_runs(h, 5)
            if char_3 == 'c':
                print(x, h, sep=": ") # 22804

    def test_has_run(self):
        self.assertEqual('a', solution.OTPGenerator.find_runs('aaabcdefg', 3))
        self.assertEqual('e', solution.OTPGenerator.find_runs('abcdeeefg', 3))
        self.assertEqual('g', solution.OTPGenerator.find_runs('abcdefggg', 3))

        self.assertEqual('8', solution.OTPGenerator.find_runs('cc38887a5', 3))

        h = hashlib.md5('abc200'.encode()).digest().hex()
        self.assertEqual('9', solution.OTPGenerator.find_runs(h, 5))

        h = hashlib.md5('abc22728'.encode()).digest().hex()
        self.assertEqual('c', solution.OTPGenerator.find_runs(h, 3))

        h = hashlib.md5('abc200'.encode()).digest().hex()
        self.assertEqual('9', solution.OTPGenerator.find_runs(h, 5))

    def test_get_otps(self):
        generator = solution.OTPGenerator(salt='abc')
        otps = generator.get_otps(n_needed=64, proximity=1000)

        # self.assertEqual(92, otps[0], 'Could not find the earliest possible key.')
        self.assertTrue(39 in otps, 'The first triple (39) could not be found in the set of keys.')
        self.assertEqual(22728, otps[-1], 'The last OTP was not 22728.')

