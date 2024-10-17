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


def part_one(puzzle_input):
    block_list = parse_intervals(puzzle_input)
    block_list = merge_intervals(block_list)

    first_allowed_address = (
        block_list[0][0] - 1 if block_list[0][0] > 0 else block_list[0][1] + 1
    )

    return str(first_allowed_address)


# 13628342 too high
def part_two(puzzle_input):
    block_list = parse_intervals(puzzle_input)
    block_list = merge_intervals(block_list)

    N_IP_ADDRESSES = 1 << 32

    n_blocked_addresses = 0
    for interval in block_list:
        n_blocked_addresses += interval[1] - interval[0] + 1

    n_allowed_addresses = N_IP_ADDRESSES - n_blocked_addresses

    return str(n_allowed_addresses)


def parse_intervals(puzzle_input):
    intervals = [
        [int(x) for x in entry.split("-")] for entry in puzzle_input.split("\n")
    ]
    intervals.sort()
    return intervals


def merge_intervals(intervals):
    if not intervals:
        return intervals

    result = []
    a, b = intervals[0]
    for c, d in intervals[1:]:
        if c <= b + 1:
            b = max(b, d)
            continue

        result.append((a, b))
        a, b = c, d
    result.append((a, b))

    return result
