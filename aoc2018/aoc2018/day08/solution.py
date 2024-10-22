def part_one(puzzle_input):
    numbers = parse_input_list(puzzle_input)

    metadata_sum = 0
    stack = []
    n_children = None
    n_metadata = None
    for x in numbers:
        if n_children is None:
            n_children = x

        elif n_metadata is None:
            n_metadata = x

            if n_children > 0:
                stack.append((n_children, n_metadata))
                n_children = None
                n_metadata = None

        elif n_metadata > 0:
            metadata_sum += x
            n_metadata -= 1

            if n_metadata == 0 and stack:
                n_children, n_metadata = stack.pop()
                n_children -= 1

                if n_children > 0:
                    stack.append((n_children, n_metadata))
                    n_children = None
                    n_metadata = None

    return str(metadata_sum)


def part_two(puzzle_input):
    numbers = parse_input_list(puzzle_input)

    stack = []
    children_remaining = None
    metadata_remaining = None
    child_values = []
    node_value = 0
    for x in numbers:
        if children_remaining is None:
            children_remaining = x

        elif metadata_remaining is None:
            metadata_remaining = x

            if children_remaining > 0:
                stack.append((children_remaining, metadata_remaining, child_values))
                children_remaining = None
                metadata_remaining = None
                child_values = []
                node_value = 0

        elif metadata_remaining > 0:
            if len(child_values) == 0:
                node_value += x
            elif 1 <= x <= len(child_values):
                node_value += child_values[x - 1]

            metadata_remaining -= 1

            if metadata_remaining == 0 and stack:
                children_remaining, metadata_remaining, child_values = stack.pop()
                children_remaining -= 1
                child_values.append(node_value)
                node_value = 0

                if children_remaining > 0:
                    stack.append((children_remaining, metadata_remaining, child_values))
                    children_remaining = None
                    metadata_remaining = None
                    child_values = []
                    node_value = 0

    return str(node_value)


def parse_input_list(puzzle_input):
    return [int(x) for x in puzzle_input.split()]
