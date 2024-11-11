import unittest
import aoc2018.day23.solution as solution


TEST_LIST = """pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1"""

TEST_LIST_2 = """pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5"""

class TestDay23(unittest.TestCase):
    def test_part_one_counts_bots_in_rage_of_strongest_sensor(self):
        expected = "7"

        result = solution.part_one(TEST_LIST)

        self.assertEqual(result, expected)

    def test_part_two_finds_distance_to_be_in_range_of_the_most_nanobots(self):
        expected = "36"

        result = solution.part_two(TEST_LIST_2)

        self.assertEqual(result, expected)
