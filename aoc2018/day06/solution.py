def main():
    NEIGHBOR_AREA = 10000
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input, NEIGHBOR_AREA))


def parse_input():
    destinations = []
    with open('input.txt') as f:
        destinations = [tuple(map(int, line.split(', '))) for line in f]
    return destinations


def part_one(puzzle_input):
    destinations = {}
    for i, coordinate in enumerate(puzzle_input):
        destinations[str(i)] = coordinate

    bounding_box = determine_bounding_box(destinations.values())

    grid = generate_grid(destinations, bounding_box)

    grid = mark_grid(grid, bounding_box, destinations)

    destinations_with_finite_area =\
        set(destinations.keys()) - find_infinite_area_labels(grid, bounding_box, destinations)

    candidate_area_sizes = find_sizes_of_finite_areas(grid, destinations_with_finite_area)

    destination, largest_finite_area = find_largest_area(candidate_area_sizes)

    return str(largest_finite_area)


def part_two(puzzle_input, neighbor_area):
    bounding_box = determine_bounding_box(puzzle_input)

    coordinates_in_neighbor_area = set()
    for x in range(bounding_box.min_x, bounding_box.max_x):
        for y in range(bounding_box.min_y, bounding_box.max_y):
            total_distance_to_destinations = 0
            for destination in puzzle_input:
                total_distance_to_destinations += manhattan_distance(x, y, *destination)
                # Early exit to speed things up
                if total_distance_to_destinations >= neighbor_area:
                    break

            if total_distance_to_destinations < neighbor_area:
                coordinates_in_neighbor_area.add((x, y))

    return str(len(coordinates_in_neighbor_area))

def determine_bounding_box(points):
    min_x = 999
    min_y = 999
    max_x = -1
    max_y = -1
    for x, y in points:
        if x < min_x:
            min_x = x
        if x > max_x:
            max_x = x
        if y < min_y:
            min_y = y
        if y > max_y:
            max_y = y
    return BoundingBox(min_x - 1, min_y - 1, max_x + 1, max_y + 1)


def generate_grid(destinations, bounding_box):
    return  [[' ' for _ in range(bounding_box.max_y)] for _ in range(bounding_box.max_x)]


def mark_grid(grid, bounding_box, destinations):
    for x in range(bounding_box.min_x, bounding_box.max_x):
        for y in range(bounding_box.min_y, bounding_box.max_y):
            min_distance = 999
            label = 'X'
            tie_found = False
            for i, coordinate in destinations.items():
                x2, y2 = coordinate
                distance = manhattan_distance(x, y, x2, y2)
                if distance == min_distance:
                    tie_found = True
                elif distance < min_distance:
                    min_distance = distance
                    label = i
                    tie_found = False
            grid[x][y] = '.' if tie_found else label
    return grid


def manhattan_distance(x1, y1, x2, y2):
    return abs(x2 - x1) + abs(y2 - y1)


def find_infinite_area_labels(grid, bounding_box, destinations):
    infinite_areas = set()
    # Scan each side of the bounding box. The labels on the boundary will "leak" into infinity.
    for x in range(bounding_box.min_x, bounding_box.max_x):
        label = grid[x][bounding_box.min_y]
        if label.isdecimal():
            infinite_areas.add(label)
    for x in range(bounding_box.min_x, bounding_box.max_x):
        label = grid[x][bounding_box.max_y - 1]
        if label.isdecimal():
            infinite_areas.add(label)
    for y in range(bounding_box.min_y, bounding_box.max_y):
        label = grid[bounding_box.min_x][y]
        if label.isdecimal():
            infinite_areas.add(label)
    for y in range(bounding_box.min_y, bounding_box.max_y):
        label = grid[bounding_box.max_x - 1][y]
        if label.isdecimal():
            infinite_areas.add(label)

    return infinite_areas


def find_sizes_of_finite_areas(grid, candidate_areas):
    areas = {}
    for candidate in candidate_areas:
        areas[candidate] = 0

    for x in range(len(grid)):
        for y in range(len(grid[0])):
            label = grid[x][y]
            if label in candidate_areas:
                areas[label] += 1

    return areas


def find_largest_area(candidate_area_sizes):
    destination = 'X'
    largest_area = 0
    for label, area in candidate_area_sizes.items():
        if area > largest_area:
            destination = label
            largest_area = area
    return destination, largest_area


class BoundingBox(object):
    def __init__(self, min_x, min_y, max_x, max_y):
        self.min_x = min_x
        self.min_y = min_y
        self.max_x = max_x
        self.max_y = max_y

    def height(self):
        return self.max_y - self.min_y

    def width(self):
        return self.max_x - self.min_x


if __name__ == '__main__':
    main()
