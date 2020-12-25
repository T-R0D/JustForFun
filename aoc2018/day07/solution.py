def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))


def parse_input():
    dependencies  = []
    with open('input.txt') as f:
        dependencies = [(line.split(' ')[1], line.split(' ')[-3]) for line in f]
    return dependencies

def part_one(puzzle_input):
    pass


def part_two(puzzle_input):
    pass


if __name__ == '__main__':
    main()
