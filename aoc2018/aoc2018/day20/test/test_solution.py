import unittest
import aoc2018.day20.solution as solution


class TestDay20(unittest.TestCase):
    def test_part_one_computes_furthest_room(self):
        test_cases = [
            ("^WNE$", "3"),
            ("^ENWWW(NEEE|SSE(EE|N))$", "10"),
            ("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", "18"),
        ]

        for path_regex, expected in test_cases:
            with self.subTest(f"{path_regex} => {expected}"):
                result = solution.part_one(path_regex)

                self.assertEqual(result, expected)
