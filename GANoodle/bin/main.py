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
import operator
import random
import copy


TARGET = 5
POPULATION_SIZE = 200
MAX_GENERATIONS = 1000


def main():
    generations_to_solution = 0
    population = [Chromosome() for _ in range(0, POPULATION_SIZE)]

    for _ in range(0, MAX_GENERATIONS):
        solution_found = False
        for chromosome in population:
            chromosome.fitness = compute_fitness(chromosome.bits, TARGET)
            # print(chromosome.fitness)

            if chromosome.fitness == 999:
                solution_found = True
                print('found it!')

        print(max([chromosome.fitness for chromosome in population]))

        if solution_found:
            print("generations: ", generations_to_solution)
            break


        population = build_new_population(population, len(population), 2)

        generations_to_solution += 1


def compute_fitness(chromosome_bits, target):
    result = 0

    step_size = 2 * 4

    for i in range(0, len(chromosome_bits) // step_size, step_size):
        segment = chromosome_bits[i:i + step_size]
        # print(segment)
        operator = int(segment[0:4])
        operand = int(segment[4:])

        if operator == 0:
            result += operand
        elif operator == 1:
            result -= operand
        elif operator == 2:
            result *= operand
        elif operator == 3:
            result // operand

    if result == target:
        return 999
    else:
        return 1 / abs(target - result)

def build_new_population(population, population_size, tournament_size):
    new_population = []

    while len(new_population) < population_size:
        a = tournament(population, tournament_size)
        b = tournament(population, tournament_size)

        new_population.append(a.clone())
        new_population.append(b.clone())

        a2 = a.clone()
        b2 = b.clone()

        a2.mutate(0.0001)
        b2.mutate(0.0001)

        new_population.append(a2)
        new_population.append(b2)

        new_population.extend(Chromosome.crossover(a, b))

    return new_population[:population_size]

def tournament(population, tournament_size):
    competitors = random.sample(population, tournament_size)
    fittest = competitors[0]
    for chromosome in competitors:
        if chromosome.fitness > fittest.fitness:
            fittest = chromosome

    return fittest



class Chromosome(object):
    def __init__(self, bit_length=320):
        bits_fmt = '{:0' + str(bit_length) + 'b}'
        bits = random.getrandbits(bit_length)

        self.bit_length = bit_length
        self.bits = (bits_fmt.format(bits))#.replace(' ', '0')
        self.fitness = 0


        # print(len(self.bits), self.bits)

    def clone(self):
        return copy.deepcopy(self)

    def mutate(self, mutation_rate):
        new_str = ''

        for bit in self.bits:
            if random.random() < mutation_rate:
                if bit == '0':
                    new_str += '1'
                else:
                    new_str += '0'

        self.bits = new_str

    @classmethod
    def crossover(cls, parent_1, parent_2, crossover_probability=0.7):
        if random.random() < crossover_probability:
            new_chromosome_1 = cls()
            new_chromosome_2 = cls()

            temp_1 = parent_1.bits
            temp_2 = parent_2.bits

            crossover_point = int(random.random() * len(temp_1))

            new_chromosome_1.bits = temp_1[:crossover_point] + temp_2[crossover_point:]
            new_chromosome_1.bits = temp_2[:crossover_point] + temp_1[crossover_point:]

            return [new_chromosome_1, new_chromosome_2]
        else:
            return []


if __name__ == '__main__':
    main()