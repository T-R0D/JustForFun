import unittest
import aoc2018.day13.solution as solution

TEST_TRACK = """/->-\\        
|   |  /----\\
| /-+--+-\\  |
| | |  | v  |
\\-+-/  \\-+--/
  \\------/   """


TEST_TRACK2 = """/-\  
\>+-\\
  \</"""


TEST_TRACK3 = """/>-<\\  
|   |  
| /<+-\\
| | | v
\\>+</ |
  |   ^
  \\<->/"""


class TestDay13(unittest.TestCase):
    def test_part_one_finds_first_collision(self):
        expected = "(7, 3)"

        result = solution.part_one(TEST_TRACK)

        self.assertEqual(result, expected)

    def test_part_one_finds_first_collision2(self):
        expected = "(0, 1)"

        result = solution.part_one(TEST_TRACK2)

        self.assertEqual(result, expected)

    def test_part_two_locates_last_cart(self):
        expected = "(6, 4)"

        result = solution.part_two(TEST_TRACK3)

        self.assertEqual(result, expected)
