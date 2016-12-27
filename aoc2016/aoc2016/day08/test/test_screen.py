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
from aoc2016.day08.screen import Screen


class TestScreen(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_rect(self):
        sut = Screen(height=2, width=2)
        sut.rect(2,2)
        expected = '##\n' + \
                   '##\n'
        self.assertEqual(expected, str(sut))

        sut.clear()
        sut.rect(1,1)
        expected = '# \n' + \
                   '  \n'
        self.assertEqual(expected, str(sut))

        sut = Screen(height=3, width=7)
        sut.rect(height=2, width=3)
        expected = '###    \n' + \
                   '###    \n' + \
                   '       \n'
        self.assertEqual(expected, str(sut))

    def test_rotate_row(self):
        sut = Screen(height=3, width=7)
        sut.rect(height=2, width=3)
        sut.rotate_row(row=0, by=4)
        expected = '    ###\n' + \
                   '###    \n' + \
                   '       \n'
        self.assertEqual(expected, str(sut))

    def test_rotate_col(self):
        sut = Screen(height=3, width=7)
        sut.rect(height=2, width=3)
        sut.rotate_col(col=1, by=2)
        expected = '###    \n' + \
                   '# #    \n' + \
                   ' #     \n'
        self.assertEqual(expected, str(sut))

    def test_scenario(self):
        sut = Screen(height=3, width=7)
        sut.rect(height=2, width=3)
        sut.rotate_col(col=1, by=1)
        sut.rotate_row(row=0, by=4)
        sut.rotate_col(col=1, by=1)
        expected = ' #  # #\n' +\
                   '# #    \n' +\
                   ' #     \n'
        self.assertEqual(expected, str(sut))
        self.assertEqual(6, sut.pixels_on())
