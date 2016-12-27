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

class GridHistory(object):
    def __init__(self):
        self.history = {}

    def add_straight_path(self, x1, y1, x2, y2):

        if x1 == x2:
            self.add_line(x1, y1, y2)
        elif y1 == y2:
            for x in range(x1, x2):
                self.add_line(x, y1, y1)
            self.add_line(x2, y1, y1)
        else:
            raise Exception('Both x values or both y values must be the same.')

    def add_line(self, x, y1, y2):
        start = min(y1, y2)
        end = max(y1, y2)
        segments = self.history.get(x, [])

        for a, b in segments:
            if a < y1 <= b <= y2:
                start = a

            if y1 <= a <= y2 < b:
                end = b

        def keep(segment):
            return not (start < segment[0] and segment[1] < end)

        self.history[x] = list(filter(keep, segments)) + [(start, end)]
