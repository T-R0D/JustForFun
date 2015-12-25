import copy
import collections


class SearchState(object):
    def __init__(self, presents):
        self.presents = copy.deepcopy(presents)


def main():
    presents = set()
    with open('input.txt') as input_file:
        for line in input_file:
            presents.add(int(line))

    part_1(presents)
    part_2(presents)


def part_1(presents):
    total_weight = sum(presents)
    target = total_weight // 3

    groups = []

    # need to barf if we can't make 3 evenly weighted groups, for this challenge
    # it "works out"
    groups.append(get_combo(presents, target))
    groups.append(get_combo(presents - groups[0], target))
    groups.append(get_combo((presents - groups[0]) - groups[1], target))

    groups = list(sorted(groups, key=compute_quantum_entanglement))

    print("The quantum entanglement of the group with the fewest items is {}.".format(
        compute_quantum_entanglement(groups[0])))

    for group in groups:
        print(group, compute_quantum_entanglement(group))


def get_combo(presents, target):
    explored = set()
    queue = collections.deque([SearchState(set())])

    while queue:
        current_state = queue.pop()

        if current_state.presents in explored or \
           sum(current_state.presents) > target:
            continue

        if sum(current_state.presents) == target:
            # multiple combos are possible, need to account for this
            return current_state.presents

        for present in presents:
            if present not in current_state.presents:
                new_state = SearchState(current_state.presents)
                new_state.presents.add(present)
                queue.appendleft(new_state)

        explored.add(frozenset(current_state.presents))

    return None

def compute_quantum_entanglement(group):
    entanglement = 1
    for item in group:
        entanglement *= item
    return entanglement


def part_2(presents):
    total_weight = sum(presents)
    target = total_weight // 4

    groups = []

    # need to barf if we can't make 3 evenly weighted groups, for this challenge
    # it "works out"
    groups.append(get_combo(presents, target))
    groups.append(get_combo(presents - groups[0], target))
    groups.append(get_combo((presents - groups[0]) - groups[1], target))
    groups.append(get_combo(((presents - groups[0]) - groups[1]) - groups[2], target))

    groups = list(sorted(groups, key=compute_quantum_entanglement))

    print(
        "The quantum entanglement of the group with the fewest items in the 4" \
        " compartment scenario is {}.".format(
        compute_quantum_entanglement(groups[0])))

    for group in groups:
        print(group, compute_quantum_entanglement(group))


if __name__ == '__main__':
    main()
