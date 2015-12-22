import math
import collections


def main():
    presents = 0
    with open('input.txt') as input_file:
        presents = int(input_file.read().strip())

    # part_1(presents)
    part_2(presents)


def part_1(presents):
    n = 1
    while get_presents_for_house(n) < presents:
        n += 1

    print("The first house to get {} presents is house {}.".format(presents, n))


def get_presents_for_house(n):
    factors = get_factors(n)
    return sum(factors) * 10


def get_factors(x):
    factors = {x}
    stop = int(math.sqrt(x)) + 1
    for i in range(1, stop):
        if x % i == 0:
            factors.add(i)
            factors.add(x // i)

    return factors


def part_2(presents):
    n = 1
    while get_presents_for_house_other(n) < presents:
        n += 1

    print("The first house to get {} presents is house {}.".format(presents, n))


def get_presents_for_house_other(n):
    factors = get_factors(n)
    factors = filter(lambda x: n // x <= 50, factors)

    return sum(factors) * 11

    # 3272728 too high

    



if __name__ == '__main__':
    main()
