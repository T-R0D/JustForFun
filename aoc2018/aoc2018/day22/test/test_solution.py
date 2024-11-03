import unittest
import aoc2018.day22.solution as solution


class TestDay22(unittest.TestCase):
    def test_part_one_finds_risk_level(self):
        expected = "114"
        puzzle_input = """depth: 510\ntarget: 10,10"""

        result = solution.part_one(puzzle_input)

        self.assertEqual(result, expected)

    def test_part_two_finds_target_in_optimal_time(self):
        expected = "45"
        puzzle_input = """depth: 510\ntarget: 10,10"""

        result = solution.part_two(puzzle_input)

        self.assertEqual(result, expected)