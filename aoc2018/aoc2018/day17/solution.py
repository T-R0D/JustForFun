def part_one(puzzle_input):
    grid, (min_i, _), _ = parse_scan_data(puzzle_input)

    n_saturated_sand_blocks, _ = simulate_water_flow(grid, (min_i, 500))

    return str(n_saturated_sand_blocks)


def part_two(puzzle_input):
    grid, (min_i, _), _ = parse_scan_data(puzzle_input)

    _, n_blocks_retained_water = simulate_water_flow(grid, (min_i, 500))

    return str(n_blocks_retained_water)


def parse_scan_data(puzzle_input):
    range_specs = []
    m, n = 0, 0
    min_i, min_j = 999_999, 999_999
    for line in puzzle_input.split("\n"):
        a, b = line.split(", ")
        a_axis, value_str = a.split("=")
        value = int(value_str)
        if a_axis == "x":
            n = max(n, value + 2)
            min_j = min(min_j, value)
        else:
            m = max(m, value + 1)
            min_i = min(min_i, value)

        b_axis, b_range_str = b.split("=")
        range_start, range_end = tuple(int(x) for x in b_range_str.split(".."))
        if b_axis == "x":
            n = max(n, range_end + 2)
            min_j = min(min_j, range_start)
        else:
            m = max(m, range_end + 1)
            min_i = min(min_i, range_start)

        range_specs.append((a_axis + b_axis, value, (range_start, range_end)))

    grid = [[" " for _ in range(n)] for _ in range(m)]
    for axis_config, value, (range_start, range_end) in range_specs:
        if axis_config == "yx":
            for j in range(range_start, range_end + 1):
                grid[value][j] = "#"
        else:
            for i in range(range_start, range_end + 1):
                grid[i][value] = "#"

    return grid, (min_i, m), (min_j, n)


def simulate_water_flow(grid, water_source):
    src_i, src_j = water_source

    n_saturated_sand_blocks = 0
    n_blocks_of_retained_water = 0

    eval_stack = [FlowEvalState(src_i, src_j, stage=FlowEvalState.INITIAL)]
    blocked_left = False
    blocked_right = False
    while eval_stack:
        state = eval_stack.pop()
        i, j = state.i, state.j

        if i < src_i or len(grid) <= i or j < 0 or len(grid[0]) <= j:
            continue

        if state.stage in FlowEvalState.INITIAL and grid[i][j] == " ":
            n_saturated_sand_blocks += 1
            grid[i][j] = "|"
            eval_stack.append(FlowEvalState(i, j, stage=FlowEvalState.INITIATE_RUNOFF))
            eval_stack.append(FlowEvalState(i + 1, j, stage=FlowEvalState.INITIAL))
            continue

        elif (
            state.stage == FlowEvalState.INITIATE_RUNOFF
            and i + 1 < len(grid)
            and grid[i + 1][j] in ("#", "~")
        ):
            eval_stack.append(
                FlowEvalState(i, j, stage=FlowEvalState.INITIATE_BLOCKED_CHECK)
            )
            eval_stack.append(FlowEvalState(i, j - 1, stage=FlowEvalState.INITIAL))
            eval_stack.append(FlowEvalState(i, j + 1, stage=FlowEvalState.INITIAL))
            continue

        elif state.stage == FlowEvalState.INITIATE_BLOCKED_CHECK:
            blocked_left = False
            blocked_right = False
            eval_stack.append(FlowEvalState(i, j, stage=FlowEvalState.MARK_BLOCKED))
            eval_stack.append(
                FlowEvalState(i, j - 1, stage=FlowEvalState.CHECK_LEFT_BLOCKED)
            )
            eval_stack.append(
                FlowEvalState(i, j + 1, stage=FlowEvalState.CHECK_RIGHT_BLOCKED)
            )
            continue

        elif state.stage == FlowEvalState.CHECK_LEFT_BLOCKED and grid[i][j] != " ":
            if grid[i][j] in ("#", "~"):
                blocked_left = True
                continue

            eval_stack.append(
                FlowEvalState(i, j - 1, stage=FlowEvalState.CHECK_LEFT_BLOCKED)
            )
            continue

        elif state.stage == FlowEvalState.CHECK_RIGHT_BLOCKED and grid[i][j] != " ":
            if grid[i][j] in ("#", "~"):
                blocked_right = True
                continue

            eval_stack.append(
                FlowEvalState(i, j + 1, stage=FlowEvalState.CHECK_RIGHT_BLOCKED)
            )
            continue

        elif (
            state.stage == FlowEvalState.MARK_BLOCKED
            and blocked_left
            and blocked_right
            and grid[i][j] == "|"
        ):
            grid[i][j] = "~"
            n_blocks_of_retained_water += 1
            eval_stack.append(
                FlowEvalState(i, j - 1, stage=FlowEvalState.MARK_LEFT_BLOCKED)
            )
            eval_stack.append(
                FlowEvalState(i, j + 1, stage=FlowEvalState.MARK_RIGHT_BLOCKED)
            )
            continue

        elif state.stage == FlowEvalState.MARK_LEFT_BLOCKED and grid[i][j] == "|":
            grid[i][j] = "~"
            n_blocks_of_retained_water += 1
            eval_stack.append(
                FlowEvalState(i, j - 1, stage=FlowEvalState.MARK_LEFT_BLOCKED)
            )
            continue

        elif state.stage == FlowEvalState.MARK_RIGHT_BLOCKED and grid[i][j] == "|":
            grid[i][j] = "~"
            n_blocks_of_retained_water += 1
            eval_stack.append(
                FlowEvalState(i, j + 1, stage=FlowEvalState.MARK_RIGHT_BLOCKED)
            )
            continue

    return n_saturated_sand_blocks, n_blocks_of_retained_water


class FlowEvalState:
    INITIAL = "INITIAL"
    INITIATE_RUNOFF = "INITIATE_RUNOFF"
    INITIATE_BLOCKED_CHECK = "INITIATE_BLOCKED_CHECK"
    CHECK_LEFT_BLOCKED = "CHECK_LEFT_BLOCKED"
    CHECK_RIGHT_BLOCKED = "CHECK_RIGHT_BLOCKED"
    MARK_BLOCKED = "MARK_BLOCKED"
    MARK_LEFT_BLOCKED = "MARK_LEFT_BLOCKED"
    MARK_RIGHT_BLOCKED = "MARK_RIGHT_BLOCKED"

    def __init__(self, i, j, /, stage):
        self.i = i
        self.j = j
        self.stage = stage
