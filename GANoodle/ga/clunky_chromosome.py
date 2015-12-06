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
import random


# random.seed(0)

class BasicChromosome(object):
    def __init__(self, gene_bits=4, n_genes=2, p_mutation=0.01, p_crossover=0.7):
        self.gene_bits = gene_bits
        self.n_genes = n_genes
        self.p_mutation = p_mutation
        self.p_crossover = p_crossover

        self.representation = ''
        for _ in range(0, n_genes):
            for ___ in range(0, gene_bits):
                self.representation += str(random.randint(0, 1))

        self.fitness = 0

    @classmethod
    def from_representation(cls, parent, representation):
        new_chromosome = cls(parent.gene_bits, parent.n_genes, parent.p_mutation, parent.p_crossover)
        new_chromosome.gene_bits = parent.gene_bits
        new_chromosome.n_genes = parent.n_genes
        new_chromosome.p_mutation = parent.p_mutation
        new_chromosome.p_crossover = parent.p_crossover

        new_chromosome.representation = representation

        return new_chromosome

    def genes(self):
        genes = []
        print(type(len(self.representation)), type(self.gene_bits))

        for i in range(len(self.representation) // int(self.gene_bits)):
            genes.append(self.representation[i:i + self.gene_bits])

        return genes

    def get_mutated_clone(self, points=None):
        if points is None:
            points = [None if random.random() < self.p_mutation else 1 for _ in
                      range(0, random.randint(0, len(self.representation)))]

        new_representation = list(self.representation)
        for point in points:
            if point is not None:
                new_representation[point] = str(abs(1 - int(new_representation[point])))

        return BasicChromosome.from_representation(self, ''.join(new_representation))

    def crossover(self, other, cuts=None):
        if cuts is None:
            cuts = [random.randint(0, len(self.representation)) if random.random() < self.p_crossover else -1 for _ in
                    range(0, random.randint(0, len(self.representation) - 1))]

        child_chromosome = self.representation
        i = 0
        for cut in sorted(cuts):
            if cut != -1:
                if i % 2 == 0:
                    child_chromosome = child_chromosome[:cut] + other.representation[cut:]
                else:
                    child_chromosome = child_chromosome[:cut] + self.representation[cut:]
                i += 1

        return BasicChromosome.from_representation(parent=self, representation=child_chromosome)
