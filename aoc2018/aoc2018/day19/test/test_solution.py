import unittest
import aoc2018.day19.solution as solution

SAMPLE_PROGRAM = """#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5"""


class TestDay19(unittest.TestCase):
    def test_part_one_computes_register_zero_value(self):
        expected = "7"

        result = solution.part_one(SAMPLE_PROGRAM)

        self.assertEqual(result, expected)
