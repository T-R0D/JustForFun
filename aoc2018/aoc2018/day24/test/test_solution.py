import unittest
import aoc2018.day24.solution as solution

TEST_INPUT = """Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4"""


class TestDay24(unittest.TestCase):
    def test_part_one_finds_correct_number_of_remaining_units_on_winning_side(self):
        expected = "5216"

        result = solution.part_one(TEST_INPUT)

        self.assertEqual(result, expected)

    def test_part_two_finds_minimum_effective_immune_boost(self):
        expected = "51"

        result = solution.part_two(TEST_INPUT)

        self.assertEqual(result, expected)
