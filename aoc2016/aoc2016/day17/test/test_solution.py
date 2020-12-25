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
from aoc2016.day17.solution import MazeState, part_one, part_two


class TestSolution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_hash_passcode_and_path(self):
        expected = 'ced9fc52441937264674bca3f4ba7588'
        result = MazeState.hash_passcode_and_path('hijkl', [])
        self.assertEqual(expected, result)
        print(result)

    def test_part_one(self):
        expected = 'DDRRRD'
        result = part_one('ihgpwlah')
        self.assertEqual(expected, result)

        expected = 'DDUDRLRRUDRD'
        result = part_one('kglvqrro')
        self.assertEqual(expected, result)

        expected = 'DRURDRUDDLLDLUURRDULRLDUUDDDRR'
        result = part_one('ulqzkmiv')
        self.assertEqual(expected, result)

    def test_part_two(self):
        expected = 370
        result = part_two('ihgpwlah')
        self.assertEqual(expected, result)

        expected = 492
        result = part_two('kglvqrro')
        self.assertEqual(expected, result)

        expected = 830
        result = part_two('ulqzkmiv')
        self.assertEqual(expected, result)