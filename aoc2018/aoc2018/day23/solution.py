import heapq


def part_one(puzzle_input):
    nanobots = parse_nanobots(puzzle_input)

    best_sensor_nanobot = sorted(
        nanobots, key=lambda bot: bot.sensor_radius, reverse=True
    )[0]

    nanobots_within_range = 0
    for nanobot in nanobots:
        separation_distance = manhattan_distance(
            nanobot.coordinates(), best_sensor_nanobot.coordinates()
        )
        if separation_distance <= best_sensor_nanobot.sensor_radius:
            nanobots_within_range += 1

    return str(nanobots_within_range)


def part_two(puzzle_input):
    nanobots = parse_nanobots(puzzle_input)

    most_sensed_location = search_for_most_sensed_location(nanobots)

    return str(manhattan_distance(most_sensed_location, (0, 0, 0)))


def parse_nanobots(puzzle_input):
    nanobots = []
    for line in puzzle_input.split("\n"):
        coordinates_str, radius_str = line.split(">, r=")
        x, y, z = tuple(int(x) for x in coordinates_str.replace("pos=<", "").split(","))
        sensor_radius = int(radius_str)

        nanobots.append(Nanobot(x, y, z, sensor_radius))

    return nanobots


class Nanobot:
    def __init__(self, x, y, z, sensor_radius):
        self.x = x
        self.y = y
        self.z = z
        self.sensor_radius = sensor_radius

    def coordinates(self):
        return (self.x, self.y, self.z)


def manhattan_distance(a, b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1]) + abs(a[2] - b[2])


def search_for_most_sensed_location(nanobots):
    # It seems dumb/inefficient to make the starting "radius" a power of 2,
    # but for whatever reason, it aids in correctness (I think because it
    # prevents accidentally excluding volumes due to rounding). It's actually
    # not really that inefficient because it adds maybe one more step/division
    # over using the maximum coordinate as the "radius" due to the basically
    # logarithmic nature of this algorithm.
    box_radius_target = max(
        [
            abs(coordinate) + bot.sensor_radius
            for bot in nanobots
            for coordinate in bot.coordinates()
        ]
    )
    initial_box_radius = 1
    while initial_box_radius < box_radius_target:
        initial_box_radius *= 2

    initial_box_center = (0, 0, 0)

    initial_reachable_bots = [bot for bot in nanobots]

    frontier = [
        (
            -len(initial_reachable_bots),
            -initial_box_radius,
            manhattan_distance(initial_box_center, (0, 0, 0)),
            initial_box_center,
        )
    ]

    while frontier:
        (
            negative_n_reachable_bots,
            negative_box_radius,
            _,
            box_center,
        ) = heapq.heappop(frontier)

        if negative_n_reachable_bots == 0:
            continue

        if negative_box_radius == 0:
            return box_center

        new_radius = -negative_box_radius // 2

        offsets = [
            (-1, -1, -1),
            (-1, -1, 1),
            (-1, 1, -1),
            (-1, 1, 1),
            (1, -1, -1),
            (1, -1, 1),
            (1, 1, -1),
            (1, 1, 1),
        ]

        if new_radius == 0:
            offsets.append((0, 0, 0))

        for offsets in offsets:
            new_center = tuple(
                box_center[i]
                + ((offsets[i] * new_radius) if new_radius > 0 else offsets[i])
                for i in range(3)
            )
            new_reachable_bots = [
                bot
                for bot in nanobots
                if nanobot_can_reach_cube(bot, new_center, new_radius)
            ]

            heapq.heappush(
                frontier,
                (
                    -len(new_reachable_bots),
                    -new_radius,
                    manhattan_distance(new_center, (0, 0, 0)),
                    new_center,
                ),
            )

    return None


def nanobot_can_reach_cube(nanobot, cube_center, cube_radius):
    overages = 0
    for i in range(3):
        overages += abs(nanobot.coordinates()[i] - cube_center[i]) - cube_radius

    return overages - nanobot.sensor_radius <= 0
