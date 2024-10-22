import unittest
import day06.solution as solution

class TestDay06(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        test_input = [
            (1, 1,),
            (1, 6,),
            (8, 3,),
            (3, 4,),
            (5, 5,),
            (8, 9,),
        ]
        self.assertEqual('17', solution.part_one(test_input))

    def test_part_two(self):
        test_neighbor_area = 32
        test_input = [
            (1, 1,),
            (1, 6,),
            (8, 3,),
            (3, 4,),
            (5, 5,),
            (8, 9,),
        ]
        self.assertEqual('16', solution.part_two(test_input, test_neighbor_area))
