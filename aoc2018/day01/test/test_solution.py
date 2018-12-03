import unittest
import day01.solution as solution

class TestDay01(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        self.assertEqual('3', solution.part_one([+1, -2, +3, +1]))
        self.assertEqual('3', solution.part_one([+1, +1, +1]))
        self.assertEqual('0', solution.part_one([+1, +1, -2]))
        self.assertEqual('-6', solution.part_one([-1, -2, -3]))

    def test_part_two(self):
        self.assertEqual('2', solution.part_two([+1, -2, +3, +1]))
        self.assertEqual('0', solution.part_two([+1, -1]))
        self.assertEqual('10', solution.part_two([+3, +3, +4, -2, -4]))
        self.assertEqual('5', solution.part_two([-6, +3, +8, +5, -6]))
        self.assertEqual('14', solution.part_two([+7, +7, -2, -7, -4]))
