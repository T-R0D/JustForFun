import unittest
import aoc2018.day11.solution as solution


class TestDay11(unittest.TestCase):
    def test_compute_power_level_returns_correct_power_level(self):
        test_cases = [
            (8, 3, 5, 4),
            (57, 122, 79, -5),
            (39, 217, 196, 0),
            (71, 101, 153, 4),
        ]
        for grid_serial_number, x, y, expected in test_cases:
            with self.subTest(f"{grid_serial_number} {x},{y} => {expected}"):
                result = solution.compute_power_level(grid_serial_number, x, y)

                self.assertEqual(result, expected)

    def test_find_coordinates_of_highest_powered_cell_block_works(self):
        test_cases = [
            ("18", "(33, 45)"),
            ("42", "(21, 61)"),
        ]
        for grid_serial_number, expected in test_cases:
            with self.subTest(f"{grid_serial_number} => {expected}"):
                result = solution.part_one(grid_serial_number)

                self.assertEqual(result, expected)
