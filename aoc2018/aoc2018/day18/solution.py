def part_one(puzzle_input):
    n_simulation_minutes = 10

    current_state = parse_initial_layout(puzzle_input)
    resource_value = 0

    for _ in range(n_simulation_minutes):
        current_state, resource_value = compute_next_state(current_state)

    return str(resource_value)


def part_two(puzzle_input):
    n_simulation_minutes = 1_000_000_000

    current_state = parse_initial_layout(puzzle_input)
    resource_value = 0


    pre_loop_t = 0
    seen_states = set()
    current_state_str = state_to_str(current_state)
    while current_state_str not in seen_states:
        pre_loop_t += 1
        current_state, resource_value = compute_next_state(current_state)
        seen_states.add(current_state_str)
        current_state_str = state_to_str(current_state)

    loop_t = 0
    t_to_resource_value = {0: 0}
    marker_state_str = state_to_str(current_state)
    current_state_str = ""
    while current_state_str != marker_state_str:
        t_to_resource_value[loop_t] = resource_value 
        current_state, resource_value = compute_next_state(current_state)
        current_state_str = state_to_str(current_state)
        loop_t += 1

    effective_minutes = (n_simulation_minutes - pre_loop_t) % loop_t

    return str(t_to_resource_value[effective_minutes])


def parse_initial_layout(puzzle_input):
    return [[rune for rune in line] for line in puzzle_input.split("\n")]


def compute_next_state(current_state):
    n_wooded = 0
    n_lumberyards = 0
    next_state = [["E" for _ in row] for row in current_state]

    for i in range(len(current_state)):
        for j in range(len(current_state[0])):
            adjacent_empty = 0
            adjacent_wooded = 0
            adjacent_lumberyards = 0

            for k in (-1, 0, 1):
                for l in (-1, 0, 1):
                    u = i + k
                    v = j + l
                    if (
                        u < 0
                        or len(current_state) <= u
                        or v < 0
                        or len(current_state[0]) <= v
                    ):
                        continue

                    if u == i and v == j:
                        continue

                    current_rune = current_state[u][v]
                    if current_rune == ".":
                        adjacent_empty += 1
                    elif current_rune == "|":
                        adjacent_wooded += 1
                    elif current_rune == "#":
                        adjacent_lumberyards += 1

            current_rune = current_state[i][j]
            next_state[i][j] = current_rune
            if current_rune == "." and adjacent_wooded >= 3:
                next_state[i][j] = "|"
            elif current_rune == "|" and adjacent_lumberyards >= 3:
                next_state[i][j] = "#"
            elif current_rune == "#" and not (
                adjacent_lumberyards >= 1 and adjacent_wooded >= 1
            ):
                next_state[i][j] = "."

            if next_state[i][j] == "|":
                n_wooded += 1
            elif next_state[i][j] == "#":
                n_lumberyards += 1

    return next_state, n_wooded * n_lumberyards

def state_to_str(state):
    return "\n".join(["".join(row) for row in state])


#632425 too high
#385645 too high