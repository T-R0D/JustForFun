import unittest
import day02.solution as solution

class TestDay02(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        test_input = [
            'abcdef',
            'bababc',
            'abbcde',
            'abcccd',
            'aabcdd',
            'abcdee',
            'ababab',
        ]
        self.assertEqual('12', solution.part_one(test_input))

    def test_part_two(self):
        test_input = [
            'abcde',
            'fghij',
            'klmno',
            'pqrst',
            'fguij',
            'axcye',
            'wvxyz',
        ]
        self.assertEqual('fgij', solution.part_two(test_input))

    def test_has_n_letters(self):
        self.assertFalse(solution.has_n_letters('abcdef', 2))
        self.assertTrue(solution.has_n_letters('bababc', 2))
        self.assertTrue(solution.has_n_letters('abbcde', 2))
        self.assertFalse(solution.has_n_letters('abcccd', 2))
        self.assertTrue(solution.has_n_letters('aabcdd', 2))
        self.assertTrue(solution.has_n_letters('abcdee', 2))
        self.assertFalse(solution.has_n_letters('ababab', 2))

        self.assertFalse(solution.has_n_letters('abcdef', 3))
        self.assertTrue(solution.has_n_letters('bababc', 3))
        self.assertFalse(solution.has_n_letters('abbcde', 3))
        self.assertTrue(solution.has_n_letters('abcccd', 3))
        self.assertFalse(solution.has_n_letters('aabcdd', 3))
        self.assertFalse(solution.has_n_letters('abcdee', 3))
        self.assertTrue(solution.has_n_letters('ababab', 3))

    def test_find_difference_in_ids(self):
        n_different, differing_letters, same_letters = solution.find_difference_in_ids('abcde', 'axcye')
        self.assertEqual(2, n_different)

        n_different, differing_letters, same_letters = solution.find_difference_in_ids('fghij', 'fguij')
        self.assertEqual(1, n_different)
