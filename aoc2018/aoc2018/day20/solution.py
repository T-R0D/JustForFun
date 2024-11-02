def part_one(puzzle_input):
    path_regex = parse_input_regex(puzzle_input)

    _, room_min_steps = build_map_from_regex(path_regex)

    furthest_room_from_origin = max(room_min_steps.values())

    return str(furthest_room_from_origin)


def part_two(puzzle_input):
    path_regex = parse_input_regex(puzzle_input)

    _, room_min_steps = build_map_from_regex(path_regex)

    n_rooms_one_thousand_steps_away = len(
        [steps for steps in room_min_steps.values() if steps >= 1000]
    )

    return str(n_rooms_one_thousand_steps_away)


def parse_input_regex(puzzle_input):
    return list(puzzle_input)


def build_map_from_regex(path_regex):
    deltas = {
        "N": (-1, 0),
        "S": (1, 0),
        "E": (0, 1),
        "W": (0, -1),
    }

    max_possible_steps = len(path_regex)

    room_min_steps = {(0, 0): 0}
    room_to_rooms = {}

    i, j = 0, 0
    current_steps = 0

    stack = []

    for r in path_regex:
        if r == "^":
            continue

        elif r in ("N", "S", "E", "W"):
            a, b = deltas[r]

            room_to_rooms[(i, j)] = room_to_rooms.get((i, j), []) + [(i + a, j + b)]

            current_steps += 1
            i += a
            j += b

            room_min_steps[(i, j)] = min(
                room_min_steps.get((i, j), max_possible_steps), current_steps
            )

        elif r == "(":
            stack.append(((i, j), current_steps))

        elif r == ")":
            (i, j), current_steps = stack.pop()

        elif r == "|":
            (i, j), current_steps = stack[-1]

        elif r == "$":
            continue

        else:
            raise ValueError(f"unrecognized rune: {r}")

    return room_to_rooms, room_min_steps
