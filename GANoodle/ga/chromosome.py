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
# along with GANoodle.  If not, see <http://www.gnu.org/licenses/>.
import copy
import random


class Chromosome(object):
    def __init__(self, genes=10):

        self.genes = [random.randint(0, 255) for _ in range(0, genes)]

        self.fitness = 0

        # print(len(self.bits), self.bits)

    def representation(self):
        ret = ''
        for gene in self.genes:
            operator, number = Chromosome.decompose_gene(gene)
            ret += '{} {:>2} '.format(operator, number)

        ret += '= {:>3}'.format(self.evaluate())

        return ret

    def evaluate(self):
        result = 0
        for gene in self.genes:
            operator, number = Chromosome.decompose_gene(gene)
            if operator == '+':
                result += number
            elif operator == '-':
                result -= number
            elif operator == '*':
                result *= number
            elif operator == '/':
                if number != 0:
                    result //= number
            else:
                raise Exception('bad operator {}'.format(operator))

        return abs(result)

    def clone(self):
        return copy.deepcopy(self)

    def mutate(self, mutation_rate):
        for g in range(len(self.genes)):
            for b in range(0, 8):
                if random.random() < mutation_rate:
                    self.genes[g] = random.randint(0, 255)
                    # print('mutate!')
                    # print(Chromosome.decompose_gene(self.genes[g]))
                    # self.genes[g] ^= 1 << b
                    # print(Chromosome.decompose_gene(self.genes[g]))

    def __str__(self):
        return "{}; fitness: {:<08}".format(self.representation(), str(round(self.fitness, 6)))

    @classmethod
    def crossover(cls, parent_1, parent_2, crossover_probability=0.7):
        if random.random() < crossover_probability:
            new_chromosome_1 = cls()
            new_chromosome_2 = cls()

            temp_1 = parent_1.genes
            temp_2 = parent_2.genes

            crossover_point = int(random.random() * len(temp_1))

            new_chromosome_1.genes = temp_1[:crossover_point] + temp_2[crossover_point:]
            new_chromosome_1.genes = temp_2[:crossover_point] + temp_1[crossover_point:]

            return [new_chromosome_1, new_chromosome_2]
        else:
            return []

    @staticmethod
    def decompose_gene(gene):
        operator = gene >> 6
        if operator == 0:
            operator = '+'
        elif operator == 1:
            operator = '-'
        elif operator == 2:
            operator = '*'
        elif operator == 3:
            operator = '/'
        else:
            raise Exception('decompose operator barf')

        number = (gene & 0b00111111)

        return operator, number
