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
import aoc2016.day06.solution as sut

class TestDay06Solution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        puzzle_input = 'eedadn\n' \
                       'drvtee\n' \
                       'eandsr\n' \
                       'raavrd\n' \
                       'atevrs\n' \
                       'tsrnev\n' \
                       'sdttsa\n' \
                       'rasrtv\n' \
                       'nssdts\n' \
                       'ntnada\n' \
                       'svetve\n' \
                       'tesnvt\n' \
                       'vntsnd\n' \
                       'vrdear\n' \
                       'dvrsen\n' \
                       'enarar\n'
        self.assertEqual('easter', sut.part_one(puzzle_input))

    def test_part_two(self):
        puzzle_input = 'eedadn\n' \
                       'drvtee\n' \
                       'eandsr\n' \
                       'raavrd\n' \
                       'atevrs\n' \
                       'tsrnev\n' \
                       'sdttsa\n' \
                       'rasrtv\n' \
                       'nssdts\n' \
                       'ntnada\n' \
                       'svetve\n' \
                       'tesnvt\n' \
                       'vntsnd\n' \
                       'vrdear\n' \
                       'dvrsen\n' \
                       'enarar\n'
        self.assertEqual('advent', sut.part_two(puzzle_input))
