# This file is part of aoc2016.
#
# aoc2016 is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#  
# aoc2016 is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#  
# You should have received a copy of the GNU General Public License
# along with aoc2016.  If not, see <http://www.gnu.org/licenses/>.

import unittest
import aoc2016.day12.solution as solution

class TestSolution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_parse_line(self):
        self.assertEqual(('cpy', '42', 'a'), solution.parse_input_line('cpy 42 a'))

    def test_cpu_cpy(self):
        program = [('cpy', '42' , 'a')]
        cpu = solution.CarrotProcessingUnit()
        cpu.set_program(program)
        cpu.run_program()

        result = cpu.check_register(solution.CarrotProcessingUnit.REG_A)
        self.assertEqual(42, result)

    def test_part_one(self):
        input = 'cpy 41 a\n' + \
                'inc a\n' + \
                'inc a\n' + \
                'dec a\n' + \
                'jnz a 2\n' + \
                'dec a'

        result = solution.part_one(input)
        self.assertEqual(42, result)
