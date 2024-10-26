import collections


def part_one(puzzle_input):
    target = parse_target_number_of_recipes(puzzle_input)

    next_ten_recipe_scores = find_ten_recipe_scores_after_target_number_of_recipes(
        int(target)
    )

    return "".join(str(score) for score in next_ten_recipe_scores)


def part_two(puzzle_input):
    target = parse_target_number_of_recipes(puzzle_input)

    n_recipes_before_target = find_n_recipes_before_target_score_sequence(target)

    return str(n_recipes_before_target)


def parse_target_number_of_recipes(puzzle_input):
    return puzzle_input


def find_ten_recipe_scores_after_target_number_of_recipes(target):
    held_recipe_ids = [0, 1]
    scoreboard = [3, 7]
    while len(scoreboard) < target + 10:
        new_recipe_score = sum(scoreboard[id] for id in held_recipe_ids)

        if new_recipe_score <= 9:
            scoreboard.append(new_recipe_score)
        else:
            scoreboard.append(new_recipe_score // 10)
            scoreboard.append(new_recipe_score % 10)

        held_recipe_ids = [
            (id + 1 + scoreboard[id]) % len(scoreboard) for id in held_recipe_ids
        ]

    next_ten_recipe_scores = scoreboard[target : target + 10]
    return next_ten_recipe_scores


def find_n_recipes_before_target_score_sequence(target):
    target_sequence = [int(x) for x in target]

    held_recipe_ids = [0, 1]
    scoreboard = [3, 7]

    window = collections.deque(scoreboard)
    n = 0

    while True:
        new_recipe_score = sum(scoreboard[id] for id in held_recipe_ids)

        new_scores = []
        if new_recipe_score <= 9:
            new_scores.append(new_recipe_score)
        else:
            new_scores.append(new_recipe_score // 10)
            new_scores.append(new_recipe_score % 10)

        for new_score in new_scores:
            scoreboard.append(new_score)

            window.append(new_score)
            if len(window) > len(target_sequence):
                window.popleft()
                n += 1

            if list(window) == target_sequence:
                return n

        held_recipe_ids = [
            (id + 1 + scoreboard[id]) % len(scoreboard) for id in held_recipe_ids
        ]
