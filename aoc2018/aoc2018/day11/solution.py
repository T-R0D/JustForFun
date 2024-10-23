import collections


def part_one(puzzle_input):
    grid_serial_number = parse_grid_serial_number(puzzle_input)

    grid_size = 300

    grid = generate_power_level_grid(grid_serial_number, grid_size)
    summed_power_table = compute_summed_power_table(grid)

    best_block_coordinates = (0, 0)
    best_block_power = -5 * (grid_size * grid_size)
    block_size = 3
    for i in range(len(summed_power_table) - (block_size - 1)):
        for j in range(len(summed_power_table[0]) - (block_size - 1)):
            block_power = compute_block_power(summed_power_table, i, j, block_size)

            if block_power > best_block_power:
                best_block_coordinates = (j + 1, i + 1)
                best_block_power = block_power

    return str(best_block_coordinates)


def part_two(puzzle_input):
    grid_serial_number = parse_grid_serial_number(puzzle_input)

    grid_size = 300

    grid = generate_power_level_grid(grid_serial_number, grid_size)
    summed_power_table = compute_summed_power_table(grid)

    best_block_coordinates = (0, 0)
    best_block_power = -5 * (grid_size * grid_size)
    best_block_size = 0
    for block_size in range(1, 300 + 1):
        for i in range(len(summed_power_table) - (block_size - 1)):
            for j in range(len(summed_power_table[0]) - (block_size - 1)):
                block_power = compute_block_power(summed_power_table, i, j, block_size)

                if block_power > best_block_power:
                    best_block_coordinates = (j + 1, i + 1)
                    best_block_power = block_power
                    best_block_size = block_size

    return str((best_block_coordinates, best_block_size))


def parse_grid_serial_number(puzzle_input):
    return int(puzzle_input)


def generate_power_level_grid(grid_serial_number, size):
    return [
        [compute_power_level(grid_serial_number, x, y) for x in range(1, size + 1)]
        for y in range(1, size + 1)
    ]


def compute_power_level(grid_serial_number, x, y):
    rack_id = x + 10
    power_level = rack_id * y
    power_level += grid_serial_number
    power_level *= rack_id
    power_level = (power_level // 100) % 10
    power_level -= 5
    return power_level


def compute_summed_power_table(grid):
    table = [[grid[i][j] for j in range(len(grid[0]))] for i in range(len(grid))]
    for i in range(len(table)):
        for j in range(len(table[0])):
            if i - 1 >= 0:
                table[i][j] += table[i - 1][j]

            if j - 1 >= 0:
                table[i][j] += table[i][j - 1]

            if i - 1 >= 0 and j - 1 >= 0:
                table[i][j] -= table[i - 1][j - 1]

    return table


def compute_block_power(summed_power_table, i, j, block_size):
    block_power = summed_power_table[i + (block_size) - 1][j + (block_size) - 1]

    if i - 1 >= 0:
        block_power -= summed_power_table[i - 1][j + (block_size - 1)]

    if j - 1 >= 0:
        block_power -= summed_power_table[i + (block_size - 1)][j - 1]

    if i - 1 >= 0 and j - 1 >= 0:
        block_power += summed_power_table[i - 1][j - 1]

    return block_power