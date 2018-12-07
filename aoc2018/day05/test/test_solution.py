import unittest
import day05.solution as solution

class TestDay05(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        self.assertEqual('0', solution.part_one('aA'))

        self.assertEqual('0', solution.part_one('abBA'))

        self.assertEqual('4', solution.part_one('abAB'))

        self.assertEqual(str(len('dabCBAcaDA')), solution.part_one('dabAcCaCBAcCcaDA'))

    def test_part_two(self):
        self.assertEqual('4', solution.part_two('dabAcCaCBAcCcaDA'))
