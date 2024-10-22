def part_one(puzzle_input):
    lights = parse_lights(puzzle_input)

    light_lookup = {light.position(): light for light in lights}
    for _ in range(999_999_999):
        image, ok = animate_lights(light_lookup)

        if ok:
            return image

        light_lookup = {
            (next_light := light.next_position()).position(): next_light
            for light in light_lookup.values()
        }


def part_two(puzzle_input):
    lights = parse_lights(puzzle_input)

    light_lookup = {light.position(): light for light in lights}
    for t in range(999_999_999):
        _, ok = animate_lights(light_lookup)

        if ok:
            return t

        light_lookup = {
            (next_light := light.next_position()).position(): next_light
            for light in light_lookup.values()
        }


def parse_lights(puzzle_input):
    return [Light.from_line(line) for line in puzzle_input.split("\n")]


class Light:
    def __init__(self, x, y, w, v):
        self.x = x
        self.y = y
        self.w = w
        self.v = v

    @classmethod
    def from_line(cls, line):
        x, y, w, v = (
            int(x)
            for x in line.replace("position=<", "")
            .replace("> velocity=<", " ")
            .replace(">", "")
            .replace(",", " ")
            .split()
        )

        return cls(x, y, w, v)

    def next_position(self, tick=1):
        x = self.x + tick * self.w
        y = self.y + tick * self.v

        return Light(x, y, self.w, self.v)

    def position(self):
        return (self.x, self.y)


def animate_lights(light_lookup):
    min_x, max_x = 999_999, 0
    min_y, max_y = 999_999, 0

    for x, y in light_lookup.keys():
        min_x = min(min_x, x)
        max_x = max(max_x, x)
        min_y = min(min_y, y)
        max_y = max(max_y, y)

    x_offset = 0 - min_x
    y_offset = 0 - min_y

    # What I determined the message window size to be, approximately.
    if max_x - min_x > 65 or max_y - min_y > 12:
        return "", False

    rows = []
    for y in range(0, max_y + y_offset + 1):
        row = []
        for x in range(0, max_x + x_offset + 1):
            if (x - x_offset, y - y_offset) in light_lookup:
                row.append("*")
            else:
                row.append(" ")
        rows.append("".join(row))

    return "\n".join(rows), True
