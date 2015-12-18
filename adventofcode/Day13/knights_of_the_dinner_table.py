import copy


class SearchState(object):
    def __init__(self):
        self.path = []
        self.distance = 0

    def __repr__(self):
        return str(self.distance) + ' ' + ' -> '.join(self.path)

    def __str__(self):
        return repr(self)


def main():
    graph = {}

    with open('input.txt') as input_file:
        for line in input_file:
            parts = line.strip().split(' ')
            
            first_person = parts[0]
            sign = 1 if parts[2] == 'gain' else -1
            value = int(parts[3])
            other_person = parts[-1].replace('.', '')

            if first_person not in graph:
                graph[first_person] = {}

            graph[first_person].update({other_person: sign * value})

    part_1(graph)
    part_2(graph)


def part_1(_input):
    people = list(_input.keys())

    best_arrangement = SearchState()
    best_arrangement.distance = -9999999999999999

    n_people = len(people)

    initial_state = SearchState()
    initial_state.path.append(people[0])
    initial_state.distance = 0

    stack = [initial_state]

    while stack:
        state = stack.pop()

        if len(state.path) == n_people:
            first_person = state.path[0]
            last_persion = state.path[-1]

            # print(first_person, last_persion)
            # print(_input[last_persion])
            # print(_input[last_persion][first_person])

            state.path.append(first_person)
            state.distance += _input[last_persion][first_person] + _input[first_person][last_persion]

            if state.distance > best_arrangement.distance:
                best_arrangement = state

        else:
            one_person = state.path[-1]
            for next_person, happiness in _input[one_person].items():
                if next_person not in state.path:
                    new_state = copy.deepcopy(state)
                    new_state.path.append(next_person)
                    happiness = _input[one_person][next_person] + _input[next_person][one_person]
                    new_state.distance += happiness
                    stack.append(new_state)

    print("The best seating arrangement is {}".format(best_arrangement))


def part_2(_input):

    _input['me'] = {}

    for person in _input.keys():
        _input['me'][person] = 0
        _input[person]['me'] = 0

    people = list(_input.keys())

    best_arrangement = SearchState()
    best_arrangement.distance = -9999999999999999

    n_people = len(people)

    initial_state = SearchState()
    initial_state.path.append('me')
    initial_state.distance = 0

    stack = [initial_state]

    while stack:
        state = stack.pop()

        if len(state.path) == n_people:
            first_person = state.path[0]
            last_persion = state.path[-1]

            # print(first_person, last_persion)
            # print(_input[last_persion])
            # print(_input[last_persion][first_person])

            state.path.append(first_person)
            state.distance += _input[last_persion][first_person] + _input[first_person][last_persion]

            if state.distance > best_arrangement.distance:
                best_arrangement = state

        else:
            one_person = state.path[-1]
            for next_person, happiness in _input[one_person].items():
                if next_person not in state.path:
                    new_state = copy.deepcopy(state)
                    new_state.path.append(next_person)
                    happiness = _input[one_person][next_person] + _input[next_person][one_person]
                    new_state.distance += happiness
                    stack.append(new_state)

    print("The best seating arrangement (including myself) is {}".format(best_arrangement))


if __name__ == '__main__':
    main()
