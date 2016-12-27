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

class Screen(object):
    def __init__(self, height=6, width=50):
        self.height = height
        self.width = width
        self.pixels = None
        self.clear()

    def rect(self, height, width):
        for i in range(height):
            for j in range(width):
                self.pixels[i][j] = 1

    def rotate_row(self, row, by):
        temp = [self.pixels[row][(j - by) % self.width] for j in range(self.width)]
        for j in range(self.width):
            self.pixels[row][j] = temp[j]

    def rotate_col(self, col, by):
        temp = [self.pixels[(i - by) % self.height][col] for i in range(self.height)]
        for i in range(self.height):
            self.pixels[i][col] = temp[i]

    def pixels_on(self):
        sum = 0
        for i in range(self.height):
            for j in range(self.width):
                sum += self.pixels[i][j]
        return sum

    def clear(self):
        self.pixels = [[0 for i in range(self.width)] for j in range(self.height)]

    def __str__(self):
        s = ''
        def convert(value):
            return '#' if value else ' '
        for row in self.pixels:
            s += ''.join(map(convert, row)) + '\n'
        return s
