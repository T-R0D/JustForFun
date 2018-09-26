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
from aoc2016.day10 import solution as sut


class TestDay10Solution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_nothing(self):
        pass

    def test_create_process_graph(self):
        chips, bots = sut.create_processing_graph(TEST_INPUT)
        expected_chips = {
            5: [('bot', 2)],
            3: [('bot', 1)],
            2: [('bot', 2)],
        }
        self.assertDictEqual(expected_chips, chips)
        expected_bots = {
            2: (('bot', 1), ('bot', 0)),
            1: (('output', 1), ('bot', 0)),
            0: (('output', 2), ('output', 0)),
        }
        self.assertDictEqual(expected_bots, bots)

    def test_contruct_chip_paths(self):
        chips, bots = sut.create_processing_graph(TEST_INPUT)
        paths = sut.construct_chip_paths(chips, bots)
        expected_paths = {
            5: [('bot', 2), ('bot', 0), ('output', 0)],
            3: [('bot', 1), ('bot', 0), ('output', 2)],
            2: [('bot', 2), ('bot', 1),  ('output', 1)],
        }
        self.assertDictEqual(expected_paths, paths)


TEST_INPUT = '''value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2'''
