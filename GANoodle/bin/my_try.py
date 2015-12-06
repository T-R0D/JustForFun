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

from ga.chromosome import Chromosome
import random


TARGET = 150
POPULATION_SIZE = 100
MAX_GENERATIONS = 1000


def main():
    # c = Chromosome()
    # print(c)


# def m():
    generations_to_solution = 0
    population = [Chromosome() for _ in range(0, POPULATION_SIZE)]

    fittest_guy = None

    for _ in range(0, MAX_GENERATIONS):
        solution_found = False
        for chromosome in population:
            chromosome.fitness = compute_fitness(chromosome, TARGET)
            # print(chromosome)

            if chromosome.fitness == 999:
                solution_found = True
                print('found it!')
                print(chromosome.representation())

        # break

        fittest_guys = sorted([round(chromosome.fitness, 6) for chromosome in population])[-7:]
        fittest_guy = max(population, key=operator.attrgetter('fitness'))
        # print(fittest_guys)


        if solution_found:
            print("generations: ", generations_to_solution)
            break

        population = build_new_population(population, len(population), 5)

        generations_to_solution += 1

    print(fittest_guy)

def compute_fitness(chromosome, target):
    result = chromosome.evaluate()

    if result == target:
        return 999
    else:
        return 1 / (target - result)

def build_new_population(population, population_size, tournament_size):
    new_population = []

    while len(new_population) < population_size:
        a = tournament(population, tournament_size)
        b = tournament(population, tournament_size)

        new_population.append(a.clone())
        new_population.append(b.clone())

        a2 = a.clone()
        b2 = b.clone()

        a2.mutate(0.001)
        b2.mutate(0.001)

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


if __name__ == '__main__':
    main()