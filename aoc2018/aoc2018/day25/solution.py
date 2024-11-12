import collections


def part_one(puzzle_input):
    coordinates = parse_coordinates(puzzle_input)

    adjacency_list = create_graph(coordinates)

    cliques = identify_cliques(adjacency_list)

    return str(len(cliques))


def part_two(puzzle_input):
    return "Merry Christmas!"


def parse_coordinates(puzzle_input):
    return [tuple(int(x) for x in line.split(",")) for line in puzzle_input.split("\n")]


def create_graph(coordinates):
    adjacency_list = {}
    for src in coordinates:
        adjacency_list[src] = []
        for dst in coordinates:
            if src is dst:
                continue

            if manhattan4_distance(src, dst) <= 3:
                adjacency_list[src].append(dst)

    return adjacency_list


def manhattan4_distance(a, b):
    distance = 0
    for i in range(4):
        distance += abs(a[i] - b[i])
    return distance


def identify_cliques(adjacency_list):
    cliques = []
    identified_clique_members = set()
    for x in adjacency_list.keys():
        if x in identified_clique_members:
            continue

        clique = set()
        frontier = collections.deque([x])
        while frontier:
            src = frontier.pop()

            if src in clique:
                continue

            for dst in adjacency_list[src]:
                frontier.appendleft(dst)

            clique.add(src)
            identified_clique_members.add(src)

        cliques.append(clique)

    return cliques
