import unittest
import aoc2018.day07.solution as solution

TEST_INSTRUCTIONS = """Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin."""


class TestDay07(unittest.TestCase):
    def test_part_one_determines_correct_construction_order(self):
        expected = "CABDFE"

        result = solution.part_one(TEST_INSTRUCTIONS)

        self.assertEqual(result, expected)

    def test_determine_construction_time_counts_construction_time_correctly(self):
        expected = 15
        dependencies = solution.parse_dependency_list(TEST_INSTRUCTIONS)
        n_workers = 2
        startup_cost = 0

        result = solution.determine_construction_time(
            dependencies, n_workers, startup_cost
        )

        self.assertEqual(result, expected)
