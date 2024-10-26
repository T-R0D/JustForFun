# Not 99,77, (109, 49), (120, 73), (23,30)
def part_one(puzzle_input):
    grid, carts = parse_track_map(puzzle_input)

    t = 0
    collisions = []
    while True:
        carts, collisions = simulate_tick(grid, carts)
        t += 1

        if len(collisions) > 0:
            break

    return str(ij_coordinate_to_xy(collisions[0][0], collisions[0][1]))


def part_two(puzzle_input):
    grid, carts = parse_track_map(puzzle_input)

    t = 0
    while len(carts) > 1:
        carts, _ = simulate_tick(grid, carts)
        t += 1

    if not carts:
        return "Somehow, there isn't a cart left"

    last_cart_location = list(carts.keys())[0]

    return str(ij_coordinate_to_xy(last_cart_location[0], last_cart_location[1]))


def parse_track_map(puzzle_input):
    carts = {}
    grid = []

    max_row_len = 0
    for i, line in enumerate(puzzle_input.split("\n")):
        row = []
        for j, rune in enumerate(line):
            if rune in ("<", ">", "^", "v"):
                carts[(i, j)] = Cart((i, j), rune)

                if rune in ("<", ">"):
                    row.append("-")
                else:
                    row.append("|")

            else:
                row.append(rune)

        if len(row) > max_row_len:
            max_row_len = len(row)
        else:
            row.extend(" " for _ in range(max_row_len - len(row)))

        grid.append(row)

    return grid, carts


class Cart:
    def __init__(self, location, heading):
        self.location = location
        self.heading = heading
        self.turn_counter = 0

    def tick(self, current_track_piece):
        i, j = self.location

        if current_track_piece == "-":
            if self.heading == "<":
                self.location = (i, j - 1)
            elif self.heading == ">":
                self.location = (i, j + 1)
            else:
                raise ValueError(f"{(i, j)} -> {self.heading} -> {current_track_piece}")
        elif current_track_piece == "|":
            if self.heading == "^":
                self.location = (i - 1, j)
            elif self.heading == "v":
                self.location = (i + 1, j)
            else:
                raise ValueError(f"{(i, j)} -> {self.heading} -> {current_track_piece}")
        elif current_track_piece == "/":
            if self.heading == "^":
                self.location = (i, j + 1)
                self.heading = ">"
            elif self.heading == "v":
                self.location = (i, j - 1)
                self.heading = "<"
            elif self.heading == "<":
                self.location = (i + 1, j)
                self.heading = "v"
            elif self.heading == ">":
                self.location = (i - 1, j)
                self.heading = "^"
            else:
                raise ValueError(f"{(i, j)} -> {self.heading} -> {current_track_piece}")
        elif current_track_piece == "\\":
            if self.heading == "^":
                self.location = (i, j - 1)
                self.heading = "<"
            elif self.heading == "v":
                self.location = (i, j + 1)
                self.heading = ">"
            elif self.heading == "<":
                self.location = (i - 1, j)
                self.heading = "^"
            elif self.heading == ">":
                self.location = (i + 1, j)
                self.heading = "v"
            else:
                raise ValueError(f"{(i, j)} -> {self.heading} -> {current_track_piece}")
        elif current_track_piece == "+":
            if self.turn_counter == 1:
                self.heading = self.heading
            elif self.turn_counter == 0 and self.heading == "^":
                self.heading = "<"
            elif self.turn_counter == 2 and self.heading == "^":
                self.heading = ">"
            elif self.turn_counter == 0 and self.heading == "v":
                self.heading = ">"
            elif self.turn_counter == 2 and self.heading == "v":
                self.heading = "<"
            elif self.turn_counter == 0 and self.heading == "<":
                self.heading = "v"
            elif self.turn_counter == 2 and self.heading == "<":
                self.heading = "^"
            elif self.turn_counter == 0 and self.heading == ">":
                self.heading = "^"
            elif self.turn_counter == 2 and self.heading == ">":
                self.heading = "v"

            self.turn_counter = (self.turn_counter + 1) % 3

            if self.heading == "^":
                self.location = (i - 1, j)
            elif self.heading == "v":
                self.location = (i + 1, j)
            elif self.heading == "<":
                self.location = (i, j - 1)
            elif self.heading == ">":
                self.location = (i, j + 1)
            else:
                raise ValueError(f"{(i, j)} -> {self.heading} -> {current_track_piece}")

        return self.location


def simulate_tick(grid, carts):
    next_carts = {}
    collisions = []

    carts_clone = {k: v for k, v in carts.items()}
    cart_order = sorted(carts_clone.keys())
    for cart_location in cart_order:
        if not cart_location in carts_clone:
            continue

        cart = carts_clone.pop(cart_location)
        current_track_piece = grid[cart_location[0]][cart_location[1]]
        new_cart_location = cart.tick(current_track_piece)

        if new_cart_location in carts_clone:
            collisions.append(new_cart_location)
            carts_clone.pop(new_cart_location)
            continue

        elif new_cart_location in next_carts:
            collisions.append(new_cart_location)
            next_carts.pop(new_cart_location)
            continue

        next_carts[new_cart_location] = cart

    return next_carts, collisions


def ij_coordinate_to_xy(i, j):
    return j, i
