import unittest
import day07.solution as solution

class TestDay07(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        test_input = [
            ('C', 'A',),
            ('C', 'F',),
            ('A', 'B',),
            ('A', 'D',),
            ('B', 'E',),
            ('D', 'E',),
            ('F', 'E',),
        ]
        self.assertEqual('CABDFE', solution.part_one(test_input))

    def test_part_two(self):
        pass
