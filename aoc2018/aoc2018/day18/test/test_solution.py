import unittest
import aoc2018.day18.solution as solution


TEST_LAYOUT = """.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|."""


class TestDay18(unittest.TestCase):
    def test_part_one_computes_resource_value_after_ten_minutes(self):
        expected = "1147"

        result = solution.part_one(TEST_LAYOUT)

        self.assertEqual(result, expected)
