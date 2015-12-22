import numpy


class Ingredient(object):
    def __init__(self, line, name, capacity, durability, flavor, texture, calories):
        self.line = line
        self.name = name
        self.capacity = int(capacity)
        self.durability = int(durability)
        self.flavor = int(flavor)
        self.texture = int(texture)
        self.calories = int(calories)

        self.non_caloric_vector = \
            numpy.array([self.capacity, self.durability, self.flavor, self.texture])

    def __str__(self):
        return self.line

    def __repr__(self):
        return str(self)


def main():
    ingredients = []
    with open('input.txt') as input_file:
        for line in input_file:
            parts = line.strip().replace(':', '').replace(',', '').split(' ')
            ingredients.append(Ingredient(line, parts[0], parts[2], parts[4], parts[6], parts[8], parts[10]))


    # print(ingredients)
    # r = generate_permutations(3, 3)
    # print(r)

    part_1(ingredients)
    part_2(ingredients)

def part_1(ingredients):
    ingredient_matrix = numpy.array(
        [ingredient.non_caloric_vector for ingredient in ingredients])

    permutations = generate_permutations(100, len(ingredients))

    best = 0, None
    for permutation in permutations:
        attribute_totals = numpy.dot(permutation, ingredient_matrix)

        positives = numpy.greater(attribute_totals, [0,0,0,0])

        if numpy.all(positives):

            product = numpy.prod(attribute_totals)

            if product > best[0]:
                best = product, permutation

    print("The best combination is {1} with a score of {0}".format(*best))



def generate_permutations(resouces_remaining, buckets):
    if buckets == 1:
        generate_permutations(0, 0)
        return [[resouces_remaining]]
    elif buckets:
        ret = []
        end = resouces_remaining + 1
        for i in range(0, end):
            sub_lists = generate_permutations(resouces_remaining - i, buckets - 1)
            for sub_list in sub_lists:
                ret.append([i] + sub_list)

        return ret
    else:
        return None



def part_2(ingredients):
    ingredient_matrix = numpy.array(
        [ingredient.non_caloric_vector for ingredient in ingredients])

    calorie_vector = numpy.array(
        [ingredient.calories for ingredient in ingredients])

    permutations = generate_permutations(100, len(ingredients))

    best = 0, None
    for permutation in permutations:
        attribute_totals = numpy.dot(permutation, ingredient_matrix)

        positives = numpy.greater(attribute_totals, [0,0,0,0])

        if numpy.all(positives):

            product = numpy.prod(attribute_totals)

            calories = numpy.dot(permutation, calorie_vector)

            if calories == 500 and product > best[0]:
                best = product, permutation

    print("The best combination is {1} with a score of {0}".format(*best))



if __name__ == '__main__':
    main()
