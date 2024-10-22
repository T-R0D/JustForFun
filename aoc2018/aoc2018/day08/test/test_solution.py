import unittest
import aoc2018.day08.solution as solution

TEST_LICENSE_FILE = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"


class TestDay08(unittest.TestCase):
    def test_part_one_sums_metadata_entries(self):
        expected = "138"

        result = solution.part_one(TEST_LICENSE_FILE)

        self.assertEqual(result, expected)

    def test_part_two_finds_root_node_value(self):
        expected = "66"

        result = solution.part_two(TEST_LICENSE_FILE)

        self.assertEqual(result, expected)
