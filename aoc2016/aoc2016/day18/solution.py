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

SAFE_TILE = '.'
TRAP_TILE = '^'

def part_one(puzzle_input):
    return count_safe_tiles_in_room(first_row_of_tiles=puzzle_input, n_rows=40)


def part_two(puzzle_input):
    return count_safe_tiles_in_room(first_row_of_tiles=puzzle_input, n_rows=400000)


def count_safe_tiles_in_room(first_row_of_tiles, n_rows):
    current_row = list(first_row_of_tiles)
    n_safe_tiles = count_safe_tiles(current_row)
    for _ in range(n_rows - 1):
        current_row = decode_next_row_of_tiles(current_row)
        n_safe_tiles += count_safe_tiles((current_row))

    return n_safe_tiles

def count_safe_tiles(row_of_tiles):
    n_traps = 0
    for tile in row_of_tiles:
        if tile == SAFE_TILE:
            n_traps += 1
    return n_traps


def decode_next_row_of_tiles(input_row):
    new_row = ['' for _ in range(len(input_row))]
    new_row[0] = determine_tile(SAFE_TILE, input_row[0], input_row[1])
    new_row[-1] = determine_tile(input_row[-2], input_row[-1], SAFE_TILE)
    for i in range(1, len(input_row) - 1):
        new_row[i] = determine_tile(*input_row[i - 1: i + 2])
    return new_row


def determine_tile(left, center, right):
    if (left == TRAP_TILE and center == SAFE_TILE and right == SAFE_TILE) or \
        (left == SAFE_TILE and center == SAFE_TILE and right == TRAP_TILE) or \
        (left == TRAP_TILE and center == TRAP_TILE and right == SAFE_TILE) or \
        (left == SAFE_TILE and center == TRAP_TILE and right == TRAP_TILE):
        return TRAP_TILE

    return SAFE_TILE
