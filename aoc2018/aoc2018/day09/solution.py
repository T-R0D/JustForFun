import collections


def part_one(puzzle_input):
    n_players, last_marble = parse_game_parameters(puzzle_input)

    best_score = simulate_marble_circle_game(n_players, last_marble)

    return str(best_score)


def part_two(puzzle_input):
    n_players, last_marble = parse_game_parameters(puzzle_input)
    last_marble *= 100

    best_score = simulate_circle_game_with_deque(n_players, last_marble)

    return str(best_score)


def parse_game_parameters(puzzle_input):
    return tuple(
        int(x)
        for x in puzzle_input.replace(" points", "").split(
            " players; last marble is worth "
        )
    )


class Marble:
    def __init__(self, value, clockwise, counterclockwise):
        self.value = value
        self.clockwise = clockwise
        self.counterclockwise = counterclockwise

    def __str__(self):
        return f"({self.counterclockwise} <- ({self.value}) -> {self.clockwise})"

    def __repr__(self):
        return str(self)


def simulate_marble_circle_game(n_players, last_marble):
    scores = [0 for _ in range(n_players)]
    marbles = [Marble(value, 0, 0) for value in range(0, last_marble + 1)]
    current = 0

    for new_marble in range(1, last_marble + 1):
        if new_marble % 23 == 0:
            player = (new_marble % n_players) - 1

            cursor = current
            for _ in range(7):
                cursor = marbles[cursor].counterclockwise

            cw = marbles[cursor].clockwise
            ccw = marbles[cursor].counterclockwise

            marbles[cw].counterclockwise = ccw
            marbles[ccw].clockwise = cw
            current = cw

            scores[player] += new_marble + marbles[cursor].value

            continue

        ccw = marbles[current].clockwise
        cw = marbles[ccw].clockwise

        marbles[ccw].clockwise = new_marble
        marbles[cw].counterclockwise = new_marble
        marbles[new_marble].counterclockwise = ccw
        marbles[new_marble].clockwise = cw

        current = new_marble

    return max(scores)


# I kinda stole this one form r/adventofcode. I didn't know about
# `deque.rotate`... Makes for an interesting, very fast solution. Might be
# fun to implement the ring buffer myself, but also... Meh, I think I know how
# It works. For comparison, this solution runs in about half a second, while
# the above solution runs in about 4 and a half seconds (both for part 2).
def simulate_circle_game_with_deque(n_players, last_marble):
    scores = [0 for _ in range(n_players)]
    marble_circle = collections.deque([0])

    for marble in range(1, last_marble + 1):
        if marble % 23 == 0:
            marble_circle.rotate(7)
            removed = marble_circle.pop()
            marble_circle.rotate(-1)

            scores[marble % n_players] += marble + removed

            continue

        marble_circle.rotate(-1)
        marble_circle.append(marble)

    return max(scores)
