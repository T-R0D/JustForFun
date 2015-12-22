import numpy

DIM = 100

def main():
    light_grid = numpy.zeros((DIM, DIM))
    with open('input.txt') as input_file:
        i = 0
        for line in input_file:
            j = 0
            for c in line.strip():
                light_grid[i][j] = 1 if c == '#' else 0
                j += 1
            i += 1

    part_1(light_grid)
    part_2(light_grid)


def part_1(light_grid):
    next_grid = numpy.zeros((DIM, DIM))

    for t in range(0, 100):
        for i in range(0, DIM):
            for j in range(0, DIM):
                neighbors = count_neighbors(light_grid, i, j)
                
                if (light_grid[i][j] == 1 and neighbors in (2, 3)) or neighbors == 3:
                    next_grid[i][j] = 1
                else:
                    next_grid[i][j] = 0

        light_grid, next_grid = next_grid, light_grid

    lights_on = 0
    for i in range(0, DIM):
        for j in range(0, DIM):
            lights_on += light_grid[i][j]

    print("At the end of the show, there are {} lights on.".format(lights_on))


def count_neighbors(light_grid, i, j):
    neighbors = 0
    for x in range(max(0, i - 1), min(i + 2, DIM)):
        for y in range(max(0, j - 1), min(j + 2, DIM)):
            if x != i or y != j:
                neighbors += light_grid[x][y]
    return neighbors


def part_2(light_grid):
    light_grid[0 ][0 ] = 1
    light_grid[0 ][99] = 1
    light_grid[99][0 ] = 1
    light_grid[99][99] = 1

    next_grid = numpy.zeros((DIM, DIM))

    for t in range(0, 100):
        for i in range(0, DIM):
            for j in range(0, DIM):
                if (i == 0 and j == 0) or (i == 0 and j == DIM - 1) or \
                   (i == DIM - 1 and j == 0) or (i == DIM - 1 and j == DIM - 1):
                    next_grid[i][j] = 1

                else:
                    neighbors = count_neighbors(light_grid, i, j)
                    
                    if (light_grid[i][j] == 1 and neighbors in (2, 3)) or \
                       neighbors == 3:
                        next_grid[i][j] = 1
                    else:
                        next_grid[i][j] = 0

        light_grid, next_grid = next_grid, light_grid

    lights_on = 0
    for i in range(0, DIM):
        for j in range(0, DIM):
            lights_on += light_grid[i][j]

    print("At the end of the show (with a broken grid), there are {} lights on.".format(lights_on))

if __name__ == '__main__':
    main()

    # consider doing it with list of lists representation since matrix is sparse
    # sets work wonderfully
