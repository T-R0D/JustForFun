def part_one(puzzle_input):
    layout_list, rules_tuples = parse_initial_layout_and_rules(puzzle_input)

    current_generation, negative_bound, positive_bound = layout_to_int_of_bits(
        layout_list
    )
    rules = rules_to_binary_string_to_bool_map(rules_tuples)

    for _ in range(20):
        current_generation, negative_bound, positive_bound = get_next_generation(
            current_generation, negative_bound, positive_bound, rules
        )

    score = count_crop_score(current_generation, negative_bound, positive_bound)

    return str(score)


def part_two(puzzle_input):
    # An interesting thing happened (at least with my input) - a "stable"
    # pattern emerged! Basically, once this pattern emerged, the pattern of
    # plants stayed the same and just shifted over by 1 spot each round.
    # I observed this by printing out 150 generations of plants, though the
    # pattern emerged earlier than that (maybe 70 rounds?). Anyway, the
    # resulting algorithm was to just simulate until that pattern emerged, then
    # increment the score for every remaining generation to simulate by the
    # number of plants in the pattern.

    layout_list, rules_tuples = parse_initial_layout_and_rules(puzzle_input)

    current_generation, negative_bound, positive_bound = layout_to_int_of_bits(
        layout_list
    )
    rules = rules_to_binary_string_to_bool_map(rules_tuples)

    gen = 0
    previous_generation = 0
    while gen < 50_000_000_000 and (current_generation >> 1) != previous_generation:
        previous_generation = current_generation
        current_generation, negative_bound, positive_bound = get_next_generation(
            current_generation, negative_bound, positive_bound, rules
        )
        gen += 1

    n_plants = count_plants(current_generation, negative_bound, positive_bound)
    remaining_generations = 50_000_000_000 - gen

    current_score = count_crop_score(current_generation, negative_bound, positive_bound)

    end_score = current_score + (remaining_generations * n_plants)

    return str(end_score)


PLANT = "#"
EMPTY = "."


def parse_initial_layout_and_rules(puzzle_input):
    layout_str, rules_str = puzzle_input.split("\n\n")

    layout = list(layout_str.replace("initial state: ", ""))

    rules = []
    for line in rules_str.split("\n"):
        pattern, outcome = line.split(" => ")
        rules.append((list(pattern), outcome))

    return layout, rules


def layout_to_int_of_bits(layout):
    int_of_bits = 0
    for x in reversed(layout):
        int_of_bits <<= 1
        if x == PLANT:
            int_of_bits |= 1

    negative_bound = 5
    int_of_bits <<= negative_bound

    return int_of_bits, negative_bound, len(layout) + 2


def rules_to_binary_string_to_bool_map(rules):
    binary_string_to_bool = {}
    for pattern, outcome in rules:
        key = 0
        for x in reversed(pattern):
            key <<= 1
            if x == PLANT:
                key |= 1

        binary_string_to_bool[key] = outcome == PLANT

    return binary_string_to_bool


N_GENERATIONS = 20


def get_next_generation(current_generation, negative_bound, positive_bound, rules):
    n_relevant_bits = positive_bound + negative_bound + 1

    next_generation = 0
    for i in range(n_relevant_bits + 1):
        pattern = (current_generation >> i) & 0x1F
        if rules.get(pattern, False):
            next_generation |= 1 << (i + 2)

    new_negative_bound = negative_bound
    for i in range(5):
        if next_generation & (1 << i) > 0:
            new_negative_bound += 1
            next_generation <<= 1

    new_positive_bound = positive_bound
    for i in range(
        negative_bound + positive_bound - 2, negative_bound + positive_bound + 1
    ):
        if next_generation & (1 << i) > 0:
            new_positive_bound += 1

    return next_generation, new_negative_bound, new_positive_bound


def count_crop_score(layout, negative_bound, positive_bound):
    score = 0

    n_relevant_bits = positive_bound + negative_bound + 1
    for i in range(n_relevant_bits + 1):
        if layout & (1 << i) > 0:
            score += i - negative_bound

    return score


def count_plants(layout, negative_bound, positive_bound):
    n_plants = 0
    for i in range(negative_bound + positive_bound + 1):
        if (layout >> i) & 1 > 0:
            n_plants += 1

    return n_plants
