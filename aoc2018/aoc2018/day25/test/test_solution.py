import unittest
import aoc2018.day25.solution as solution

TEST_COORDINATES_0 = """0,0,0,0
 3,0,0,0
 0,3,0,0
 0,0,3,0
 0,0,0,3
 0,0,0,6
 9,0,0,0
12,0,0,0"""

TEST_COORDINATES_1 = """-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0"""

TEST_COORDINATES_2 = """1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2"""

TEST_COORDINATES_3 = """1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2"""

class TestDay25(unittest.TestCase):
    def test_part_one_counts_constellations(self):
        test_cases = [
            (TEST_COORDINATES_0, "2"),
            (TEST_COORDINATES_1, "4"),
            (TEST_COORDINATES_2, "3"),
            (TEST_COORDINATES_3, "8"),
        ]

        for i, (puzzle_input, expected) in enumerate(test_cases):
            with self.subTest(f"Case {i} has {expected} constellations"):
                result = solution.part_one(puzzle_input)

                self.assertEqual(result, expected)
