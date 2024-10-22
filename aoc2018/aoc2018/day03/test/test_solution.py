import unittest
import day03.solution as solution

class TestDay03(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        test_input = [
            (1, 1, 3, 4, 4),
            (2, 3, 1, 4, 4),
            (3, 5, 5, 2, 2),
        ]
        self.assertEqual('4', solution.part_one(test_input))

    def test_part_two(self):
        test_input = [
            (1, 1, 3, 4, 4),
            (2, 3, 1, 4, 4),
            (3, 5, 5, 2, 2),
        ]
        self.assertEqual('3', solution.part_two(test_input))
