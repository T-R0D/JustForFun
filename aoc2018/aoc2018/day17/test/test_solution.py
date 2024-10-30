import unittest
import aoc2018.day17.solution as solution


TEST_SCAN = """x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504"""


class TestDay17(unittest.TestCase):
    def test_part_one_counts_watered_sand_squares_correctly(self):
        expected = "57"

        result = solution.part_one(TEST_SCAN)

        self.assertEqual(result, expected)
    def test_part_two_counts_retained_water_correctly(self):
        expected = "29"

        result = solution.part_two(TEST_SCAN)

        self.assertEqual(result, expected)
