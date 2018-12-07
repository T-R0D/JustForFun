def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))


def parse_input():
    polymer = ''
    with open('input.txt') as f:
        polymer = f.read()
    return polymer

def part_one(puzzle_input):
    result = fully_reduce_polymer(puzzle_input)
    return str(len(result))


def part_two(puzzle_input):
    original_result = fully_reduce_polymer(puzzle_input)
    original_result = ''.join(original_result)
    possible_removals = set(original_result.lower())

    best_removal = ''
    shortest_polymer_len = len(original_result)
    for removal in possible_removals:
        single_removal_polymer = original_result.replace(removal, '').replace(removal.upper(), '')
        result = fully_reduce_polymer(single_removal_polymer)
        if len(result) < shortest_polymer_len:
            shortest_polymer_len = len(result)
            best_removal = removal

    return str(shortest_polymer_len)


def fully_reduce_polymer(original_polymer):
    if (len(original_polymer) < 2):
        return str(len(original_polymer))

    polymer = original_polymer
    elimination_made = True
    while elimination_made:
        result = reduce_polymer(polymer)
        elimination_made = len(result) < len(polymer)
        polymer = result

    return polymer


def reduce_polymer(polymer):
    result = []
    i = 1
    while i < len(polymer):
        a = polymer[i - 1]
        b = polymer[i]
        if same_letter_different_case(a, b):
            i += 2
            elimination_made = True
        else:
            result.append(a)
            i += 1
    # If the last pair was not an elimination, then i should point just past the end of src. Add this last
    # molecule that would otherwise be missed.
    if i == len(polymer):
        result.append(polymer[-1])

    return result


def same_letter_different_case(a, b):
    return (ord(a) + ord(' ')) == ord(b) or (ord(a) - ord(' ')) == ord(b)


if __name__ == '__main__':
    main()
