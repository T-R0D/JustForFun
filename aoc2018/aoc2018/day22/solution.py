import heapq


def part_one(puzzle_input):
    depth, target = parse_depth_and_target(puzzle_input)

    region_types = compute_region_types(target, depth)
    risk_level = sum_risk_level(region_types)

    return str(risk_level)


def part_two(puzzle_input):
    depth, target = parse_depth_and_target(puzzle_input)

    # The 100 buffer is somewhat arbitrary. It seemed easier to precompute
    # some buffer than compute erosion levels/region types on the fly.
    region_types = compute_region_types(target, depth, 100)

    time_to_find = search_for_target(target, region_types)

    return str(time_to_find)


def parse_depth_and_target(puzzle_input):
    depth_line, target_line = puzzle_input.split("\n")

    depth = int(depth_line.replace("depth: ", ""))

    target = tuple(
        reversed([int(x) for x in target_line.replace("target: ", "").split(",")])
    )

    return depth, target


def compute_region_types(target, depth, buffer=0):
    erosion_levels = [
        [0 for _ in range(target[1] + 1 + buffer)]
        for _ in range(target[0] + 1 + buffer)
    ]

    for i in range(len(erosion_levels)):
        for j in range(len(erosion_levels[0])):
            geologic_index = 0
            if i == 0 and j == 0:
                geologic_index = 0

            elif i == target[0] and j == target[1]:
                geologic_index = 0

            elif i == 0:
                geologic_index = j * 16807

            elif j == 0:
                geologic_index = i * 48271

            else:
                geologic_index = erosion_levels[i][j - 1] * erosion_levels[i - 1][j]

            erosion_levels[i][j] = (geologic_index + depth) % 20183

    region_types = [
        [0 for _ in range(len(erosion_levels[0]))] for _ in range(len(erosion_levels))
    ]

    for i in range(len(erosion_levels)):
        for j in range(len(erosion_levels[0])):
            region_types[i][j] = erosion_levels[i][j] % 3

    return region_types


def sum_risk_level(region_types):
    risk_level = 0
    for i in range(len(region_types)):
        for j in range(len(region_types[0])):
            risk_level += region_types[i][j]

    return risk_level


def search_for_target(target, region_types):
    NOTHING_EQUIPPED = 0x00
    TORCH = 0x01
    CLIMBING_GEAR = 0x02

    EQUIPMENT_FOR_REGION_TYPE = {
        0: {CLIMBING_GEAR, TORCH},
        1: {CLIMBING_GEAR, NOTHING_EQUIPPED},
        2: {TORCH, NOTHING_EQUIPPED},
    }

    frontier = [(0, TORCH, (0, 0))]
    seen = set()

    while frontier:
        t, current_equipment, (i, j) = heapq.heappop(frontier)

        if (i, j) == target and current_equipment == TORCH:
            return t

        if ((i, j), current_equipment) in seen:
            continue

        current_region_type = region_types[i][j]
        for d_i, d_j in ((-1, 0), (1, 0), (0, -1), (0, 1)):
            r, s = i + d_i, j + d_j

            if r < 0 or len(region_types) <= r or s < 0 or len(region_types[0]) <= s:
                continue

            candidate_region_type = region_types[r][s]

            if current_region_type == candidate_region_type:

                heapq.heappush(frontier, (t + 1, current_equipment, (r, s)))
                continue

            equipment_required_to_move = list(
                EQUIPMENT_FOR_REGION_TYPE[current_region_type]
                & EQUIPMENT_FOR_REGION_TYPE[candidate_region_type]
            )[0]
            if current_equipment == equipment_required_to_move:
                heapq.heappush(frontier, (t + 1, current_equipment, (r, s)))

        next_equipment = list(
            EQUIPMENT_FOR_REGION_TYPE[current_region_type] - {current_equipment}
        )[0]
        heapq.heappush(frontier, (t + 7, next_equipment, (i, j)))

        seen.add(((i, j), current_equipment))

    return -1
