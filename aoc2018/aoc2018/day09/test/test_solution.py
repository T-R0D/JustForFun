import unittest
import aoc2018.day09.solution as solution


class TestDay09(unittest.TestCase):
    def test_part_one_determines_best_game_score(self):
        test_cases = [
            ("9 players; last marble is worth 27 points", "32"),
            ("10 players; last marble is worth 1618 points", "8317"),
            ("13 players; last marble is worth 7999 points", "146373"),
            ("17 players; last marble is worth 1104 points", "2764"),
            ("21 players; last marble is worth 6111 points", "54718"),
            ("30 players; last marble is worth 5807 points", "37305"),
        ]

        for puzzle_input, expected in test_cases:
            with self.subTest(puzzle_input=puzzle_input, expected=expected):
                result = solution.part_one(puzzle_input)

                self.assertEqual(result, expected)
