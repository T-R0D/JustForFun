import copy
import collections

class SearchState(object):
    def __init__(self):
        self.containers_used = set()
        self.capacity = 0

    def __repr__(self):
        return "{} -> {}".format(self.capacity, ', '.join([str(c) for c in self.containers_used]))

    def __str__(self):
        return repr(self)


def main():
    containers = []
    with open('input.txt') as input_file:
        for i, line in enumerate(input_file):
            containers.append(int(line.strip()))

    part_1(containers)
    part_2(containers)


def part_1(containers):
    solutions = set()
    explored = set()

    for i in range(0, len(containers)):
        initial_state = SearchState()
        initial_state.containers_used.add(i)
        initial_state.capacity += containers[i]

        stack = [initial_state]

        while stack:
            current = stack.pop()

            if current.containers_used in explored:
                pass

            elif current.capacity == 150:
                solutions.add(frozenset(current.containers_used))

            elif current.capacity < 150:
                for j in range(i + 1, len(containers)):
                    if j not in current.containers_used:
                        new_state = copy.deepcopy(current)
                        new_state.containers_used.add(j)
                        new_state.capacity += containers[j]

                        if new_state.containers_used not in explored:
                            stack.append(new_state)

            explored.add(frozenset(current.containers_used))

    print("There are {} ways to hold the eggnog.".format(len(solutions)))



def part_2(containers):
    solution_size = len(containers)
    solutions = set()
    explored = set()

    queue = collections.deque()
    for i in range(0, len(containers)):
        initial_state = SearchState()
        initial_state.containers_used.add(i)
        initial_state.capacity += containers[i]

        queue.append(initial_state)

    while queue:
        current = queue.popleft()

        if current.containers_used in explored:
            pass

        elif current.capacity == 150:
            if solutions:
                if len(current.containers_used) == solution_size:
                    solutions.add(frozenset(current.containers_used))
            else:
                solution_size = len(current.containers_used)
                solutions.add(frozenset(current.containers_used))

        elif current.capacity < 150:
            for j in range(0, len(containers)):
                if j not in current.containers_used:
                    new_state = copy.deepcopy(current)
                    new_state.containers_used.add(j)
                    new_state.capacity += containers[j]

                    if new_state.containers_used not in explored:
                        queue.append(new_state)

        explored.add(frozenset(current.containers_used))

    print("There are {} ways to hold the eggnog in {} containers.".format(
        len(solutions), solution_size))


if __name__ == '__main__':
    main()
