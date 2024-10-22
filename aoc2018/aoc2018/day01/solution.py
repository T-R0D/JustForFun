def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))


def parse_input():
    changes = []
    with open('input.txt') as f:
        changes = [int(line) for line in f]
    return changes


def part_one(puzzle_input):
    acc = 0
    for x in puzzle_input:
        acc += x

    return str(acc)


def part_two(puzzle_input):
    acc = 0
    frequencies = {0}
    repeat_found = False
    first_repeat = 0
    while not repeat_found:
        for x in puzzle_input:
            acc += x
            if acc in frequencies:
                first_repeat = acc
                repeat_found = True
                break
            frequencies.add(acc)

    return str(first_repeat)


if __name__ == '__main__':
    main()
