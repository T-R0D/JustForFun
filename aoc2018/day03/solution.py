import collections


def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))


def parse_input():
    claims = []
    with open('input.txt') as f:
        for line in f:
            l = line.replace('#', '').replace(' @ ', ',').replace(': ', ',').replace('x', ',').replace('\n', '')
            l = tuple(map(lambda x: int(x), l.split(',')))
            claims.append(l)
    return claims


def part_one(puzzle_input):
    fabric_usage = {}
    for claim in puzzle_input:
        claim_id, x, y, width, height = claim
        for i in range(x, x + width):
            fabric_strip = fabric_usage.get(i, collections.Counter())
            for j in range(y, y + height):
                fabric_strip[j] += 1
            fabric_usage[i] = fabric_strip

    overused_squares = 0
    for fabric_strip in fabric_usage.values():
        for claims in fabric_strip.values():
            if claims > 1:
                overused_squares += 1

    return str(overused_squares)


def part_two(puzzle_input):
    fabric_usage = {}
    for claim in puzzle_input:
        claim_id, x, y, width, height = claim
        for i in range(x, x + width):
            fabric_strip = fabric_usage.get(i, collections.Counter())
            for j in range(y, y + height):
                fabric_strip[j] += 1
            fabric_usage[i] = fabric_strip

    for claim in puzzle_input:
        claim_id, x, y, width, height = claim
        has_overused_square = False
        for i in range(x, x + width):
            fabric_strip = fabric_usage.get(i, collections.Counter())
            for j in range(y, y + height):
                square_usages = fabric_strip[j]
                if square_usages > 1:
                    has_overused_square = True
        if not has_overused_square:
            return str(claim_id)

    return 'No non-overlapping claim found...'

class FabricClaim(object):
    def __init__(self, claim_id, x, y, height, width):
        self.claim_id = cl
        self.x = x

if __name__ == '__main__':
    main()
