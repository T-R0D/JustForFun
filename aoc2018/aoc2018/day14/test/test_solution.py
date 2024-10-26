import unittest
import aoc2018.day14.solution as solution

class TestDay14(unittest.TestCase):
    def test_part_one_determines_ten_scores_after_target(self):
        test_cases = [
            ("9", "5158916779"),
            ("5", "0124515891"),
            ("18", "9251071085"),
            ("2018", "5941429882"),
        ]

        for target, expected in test_cases:
            with self.subTest(f"{target} -> {expected}"):
                result = solution.part_one(target)

                self.assertEqual(result, expected)

    def test_part_two_counts_recipes_before_target_score_sequence(self):
        test_cases = [
            ("51589", "9"),
            ("01245", "5"),
            ("92510", "18"),
            ("59414", "2018"),
        ]

        for target, expected in test_cases:
            with self.subTest(f"{target} -> {expected}"):
                result = solution.part_two(target)

                self.assertEqual(result, expected)