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
import time

import aoc2016.day01.day01 as day01
import aoc2016.day02.solution as day02
import aoc2016.day03.solution as day03
import aoc2016.day04.solution as day04
import aoc2016.day05.solution as day05
import aoc2016.day06.solution as day06
import aoc2016.day07.solution as day07
import aoc2016.day08.solution as day08
import aoc2016.day09.solution as day09
import aoc2016.day10.solution as day10
import aoc2016.day11.solution as day11
import aoc2016.day12.solution as day12
import aoc2016.day13.solution as day13
import aoc2016.day14.solution as day14
import aoc2016.day15.solution as day15
import aoc2016.day16.solution as day16
import aoc2016.day17.solution as day17
import aoc2016.day18.solution as day18
import aoc2016.day19.solution as day19
import aoc2016.day20.solution as day20
import aoc2016.day21.solution as day21
import aoc2016.day22.solution as day22
import aoc2016.day23.solution as day23


def main():
    args = parse_args()

    puzzle_input = args.input.read().strip()

    day = args.day
    part = args.part

    solve = {
        1: {
            1: (day01.Day01()).part_a,
            2: day01.part_two,
        },
        2: {
            1: day02.part_one,
            2: day02.part_two,
        },
        3: {
            1: day03.part_one,
            2: day03.part_two,
        },
        4: {
            1: day04.part_one,
            2: day04.part_two,
        },
        5: {
            1: day05.part_one,
            2: day05.part_two,
        },
        6: {
            1: day06.part_one,
            2: day06.part_two,
        },
        7: {
            1: day07.part_one,
            2: day07.part_two,
        },
        8: {
            1: day08.part_one,
            2: day08.part_two,
        },
        9: {
            1: day09.part_one,
            2: day09.part_two,
        },
        10: {
            1: day10.part_one,
            2: day10.part_two,
        },
        11: {
            1: day11.part_one,
            2: day11.part_two,
        },
        11: {
            1: day11.part_one,
            2: day11.part_two,
        },
        12: {
            1: day12.part_one,
            2: day12.part_two,
        },
        13: {
            1: day13.part_one,
            2: day13.part_two,
        },
        14: {
            1: day14.part_one,
            2: day14.part_two,
        },
        15: {
            1: day15.part_one,
            2: day15.part_two,
        },
        16: {
            1: day16.part_one,
            2: day16.part_two,
        },
        17: {
            1: day17.part_one,
            2: day17.part_two,
        },
        18: {
            1: day18.part_one,
            2: day18.part_two,
        },
        19: {
            1: day19.part_one,
            2: day19.part_two,
        },
        20: {
            1: day20.part_one,
            2: day20.part_two,
        },
        21: {
            1: day21.part_one,
            2: day21.part_two,
        },
        22: {
            1: day22.part_one,
            2: day22.part_two,
        },
        23: {
            1: day23.part_one,
            2: day23.part_two,
        },
        # 24: {1: day24.part_one, 2: day24.part_two,},
        # 25: {1: day25.part_one, 2: day24.part_two,},
    }[day][part]

    result = ""
    start = 0
    end = 0
    try:
        start = time.time()
        result = solve(puzzle_input)
    except Exception as e:
        result = str(e)
    finally:
        end = time.time()

    print(result)
    print(f"{end - start}")


def parse_args():
    parser = argparse.ArgumentParser(
        prog="aoc2016",
        description="solvers for the 2016 Advent of Code programming challenge.",
    )
    parser.add_argument(
        "--input", type=argparse.FileType("r"), help="File path for the puzzle's input."
    )
    parser.add_argument(
        "--day",
        type=int,
        choices=list(range(1, 26)),
        help="The day of the advent to solve.",
    )
    parser.add_argument(
        "--part",
        type=int,
        choices=(1, 2),
        help="The part of the day's problem to solve.",
    )

    return parser.parse_args()


if __name__ == "__main__":
    main()
