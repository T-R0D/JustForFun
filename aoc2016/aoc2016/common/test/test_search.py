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
import aoc2016.common.search as search


class TestSearch(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_graph_search(self):
        class TestState(search.State):
            def __init__(self, val=0):
                search.State.__init__(self)
                self.val = val

            def __str__(self):
                return '{}'.format(self.val)

            def generate_next_states(self):
                return [TestState(i) for i in range(self.val + 1, self.val + 4)]

            def is_goal(self):
                return self.val == 8

        result, path = search.graph_search(TestState())
        self.assertEqual(result.val, 8)
        self.assertListEqual(path, ['0', '2', '5',])
