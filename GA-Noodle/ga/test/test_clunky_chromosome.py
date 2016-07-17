# This file is part of GANoodle.
#
# GANoodle is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#  
# GANoodle is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#  
# You should have received a copy of the GNU General Public License
# along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
import unittest

from ga.clunky_chromosome import BasicChromosome


class TestChromosome(unittest.TestCase):
    def setup(self):
        pass

    def tearDown(self):
        pass

    def test_mutate(self):
        individual = BasicChromosome()
        individual.representation = '00000000'

        clone = individual.get_mutated_clone(points=[1, 3])

        self.assertEqual('01010000', clone.representation)


    def test_crossover(self):
        parent_1 = BasicChromosome()
        parent_1.representation = '00000000'

        parent_2 = BasicChromosome()
        parent_2.representation = '11111111'

        child = parent_1.crossover(parent_2, cuts=[4, 2])

        self.assertEqual('00110000', child.representation)
