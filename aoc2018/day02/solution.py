import collections


def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))


def parse_input():
    ids = []
    with open('input.txt') as f:
        ids = [line for line in f]
    return ids


def part_one(puzzle_input):
    n_duplicates = 0
    n_triplicates = 0
    for id in puzzle_input:
        if (has_n_letters(id, 2)):
            n_duplicates += 1
        if (has_n_letters(id, 3)):
            n_triplicates += 1
    return str(n_duplicates * n_triplicates)


def part_two(puzzle_input):
    for id1 in puzzle_input:
        for id2 in puzzle_input[1:]:
            n_different, differing_letters, same_letters = find_difference_in_ids(id1, id2)
            if n_different == 1:
                return ''.join(same_letters)
    return 'Solution not found...'


def has_n_letters(id1, n):
    counts = collections.Counter(id1)
    for count in counts.values():
        if count == n:
            return True
    return False


def find_difference_in_ids(id1, id2):
    n_different = 0
    differing_letters = {}
    same_letters = []
    for a, b in zip(id1, id2):
        if a == b:
            same_letters.append(a)
        else:
            n_different += 1
            differing_letters[a] = b
    return n_different, differing_letters, same_letters


if __name__ == '__main__':
    main()
