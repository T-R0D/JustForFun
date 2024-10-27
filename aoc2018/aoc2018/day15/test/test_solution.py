import unittest
import aoc2018.day15.solution as solution

TEST_COMBAT_0 = """#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########"""

TEST_COMBAT_1 = """#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######"""

TEST_COMBAT_2 = """#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######"""

TEST_COMBAT_3 = """#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######"""

TEST_COMBAT_4 = """#######
#E.G#.#
#.#G..#
#G.#.G# 
#G..#.#
#...E.#
#######"""

TEST_COMBAT_5 = """#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######"""

TEST_COMBAT_6 = """#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...# 
#.G...G.#
#.....G.#
#########"""

TEST_COMBAT_7 = """#########
#G#.....#
#G#.....#
#.#.....#
#.#.....#
#.#.....#
#....E..#
#..GE...#
#########"""
TEST_COMBAT_8 = """########
#...####
#..G..##
#......#
##.....#
##....##
##G..G.#
##...E.#
##E....#
########"""

TEST_COMBAT_9 = """#######
#######
#.E..G#
#.#####
#G#####
#######
#######"""

TEST_COMBAT_10 = """####
#GG#
#.E#
####"""

TEST_COMBAT_11 = """#######
#####G#
#####.#
#..E..#
#G#####
#######
#######"""

TEST_COMBAT_12 = """########
#..E..G#
#G######
########"""

TEST_COMBAT_13 = """######
#.G..#
##..##
#...E#
#E...#
######"""

TEST_COMBAT_14 = """######################
#...................E#
#.####################
#....................#
####################.#
#....................#
#.####################
#....................#
###.##################
#EG.#................#
######################"""

TEST_COMBAT_15 = """###########
#G..#....G#
###..E#####
###########"""

# This test case will simulate a unit moving into a defeated unit's position.
# We should not see that unit attack twice (it will be the goblin that starts
# at the bottom of the map).
TEST_COMBAT_16 = """##############
##########G###
#.........EG.#
#.########E###
#.############
#............#
############.#
#.......G....#
##############"""


class TestDay15(unittest.TestCase):  #
    def test_part_one_computes_battle_outcome_correctly(self):
        test_cases = [
            (TEST_COMBAT_1, "27730"),
            (TEST_COMBAT_2, "36334"),
            (TEST_COMBAT_3, "39514"),
            (TEST_COMBAT_4, "27755"),
            (TEST_COMBAT_5, "28944"),
            (TEST_COMBAT_6, "18740"),
            (TEST_COMBAT_0, "27828"),
            (TEST_COMBAT_7, "16704"),
            (TEST_COMBAT_8, "17628"),
            (TEST_COMBAT_9, "10234"),
            (TEST_COMBAT_10, "9933"),
            (TEST_COMBAT_11, "10430"),
            (TEST_COMBAT_12, "10234"),
            (TEST_COMBAT_13, "10430"),
            (TEST_COMBAT_14, "13332"),
            (TEST_COMBAT_15, "10804"),
            (TEST_COMBAT_16, "29973"),
        ]

        for i, (puzzle_input, expected) in enumerate(test_cases):
            with self.subTest(f"{i+1} -> expected"):
                result = solution.part_one(puzzle_input)

                self.assertEqual(result, expected)

    def test_part_two_finds_correct_outcome_with_minimal_attack_boost(self):
        test_cases = [
            (TEST_COMBAT_1, "4988"),
            (TEST_COMBAT_3, "31284"),
            (TEST_COMBAT_4, "3478"),
            (TEST_COMBAT_5, "6474"),
            (TEST_COMBAT_6, "1140"),
        ]

        for i, (puzzle_input, expected) in enumerate(test_cases):
            with self.subTest(f"{i+1} -> expected"):
                result = solution.part_two(puzzle_input)

                self.assertEqual(result, expected)
