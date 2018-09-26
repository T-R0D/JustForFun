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

import argparse


PATH = '../aoc2016/'

def main():
    argument_parser = argparse.ArgumentParser(description='Generate a blank module for an AoC solution')
    argument_parser.add_argument('day', type=int, choices=range(1, 26))
    args = argument_parser.parse_args()

    day = '{0:02d}'.format(args.day)

if __name__  == '__main__':
    main()
