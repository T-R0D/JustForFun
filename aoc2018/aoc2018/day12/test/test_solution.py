import unittest
import aoc2018.day12.solution as solution

TEST_NOTES = """initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #"""


class TestDay12(unittest.TestCase):
    def test_part_one_computes_score_correctly(self):
        expected = "325"

        result = solution.part_one(TEST_NOTES)

        self.assertEqual(result, expected)
