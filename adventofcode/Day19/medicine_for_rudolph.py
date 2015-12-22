import re
import collections
import random


def main():
    productions = {}
    target = ''

    with open('input.txt') as input_file:
        for line in input_file:
            if '=>' in line:
                parts = line.strip().replace(' ', '').split('=>')

                if parts[0] not in productions:
                    productions[parts[0]] = []

                productions[parts[0]].append(parts[1])

            elif len(line) > 3:
                target = line.strip()

    part_1(productions, target)
    part_2(productions, target)


def part_1(productions, target):
    produced_molecules = set()

    for lhs in productions.keys():
        locations = [m.start() for m in re.finditer(lhs, target)]
        length = len(lhs)
        for i in locations:
            for rhs in productions[lhs]:
                new_molecule = target[:i] + rhs + target[i + length: ]
                produced_molecules.add(new_molecule)

    print(
        "The calibration molecule produces {} other unique molecules.".format(
            len(produced_molecules)))


def part_2(productions, target):
    inverted_productions = {}
    for key, value in productions.items():
        for rhs in value:
            inverted_productions[rhs] = key

    sorted_rhss = list(sorted(inverted_productions.keys(), key=lambda x: -len(x)))

    solution_evades = True
    while solution_evades:
        steps = 0
        reduction = target
        while reduction != 'e':
            # print(reduction)

            reductions_made = False

            for rhs in sorted_rhss:
                while rhs in reduction:
                    reductions_made = True
                    steps += reduction.count(rhs) 
                    reduction = reduction.replace(rhs, inverted_productions[rhs])

            if not reductions_made:
                print('uh oh')
                random.shuffle(sorted_rhss)
                steps = 0
                break

        if reduction == 'e':
            solution_evades = False

    print("The medicine molecule can be produced in {} steps.".format(steps))


    # BFS takes WAAAAAAAAAAAAAAAAAAAAY too long/much! #
    # explored = {'e'}
    # queue = collections.deque([('e', 0)])
    # required_steps = -1

    # length_max = 1

    # target_len = len(target)

    # while queue:
    #     current = queue.popleft()

    #     molecule, steps = current

    #     if molecule == target:
    #         required_steps = current[1]
    #         break

    #     elif len(molecule) <= target_len:
    #         for lhs in productions.keys():
    #             locations = [m.start() for m in re.finditer(lhs, molecule)]
    #             length = len(lhs)
    #             for i in locations:
    #                 for rhs in productions[lhs]:
    #                     new_molecule = molecule[:i] + rhs + molecule[i + length:]
    #                     if new_molecule not in explored:
    #                         if len(new_molecule) >= length_max:
    #                             queue.append((new_molecule, steps + 1))
    #                             length_max = len(new_molecule)

    #     explored.add(molecule)

    # print("The medicine molecule can be made in {} steps.".format(required_steps))

if __name__ == '__main__':
    main()
