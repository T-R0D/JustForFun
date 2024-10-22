import collections
import heapq


def part_one(puzzle_input):
    dependencies = parse_dependency_list(puzzle_input)

    construction_order = determine_construction_order(dependencies)

    return "".join(construction_order)


def part_two(puzzle_input):
    dependencies = parse_dependency_list(puzzle_input)

    construction_time = determine_construction_time(dependencies, 5, 60)

    return str(construction_time)


def parse_dependency_list(puzzle_input):
    dependencies = {}
    for line in puzzle_input.split("\n"):
        requirement, next_step = (
            line.replace("Step ", "")
            .replace("must be finished before step ", "")
            .replace(" can begin.", "")
            .split()
        )
        dependencies[requirement] = dependencies.get(requirement, []) + [next_step]
        if next_step not in dependencies:
            dependencies[next_step] = []

    return dependencies


def determine_construction_order(dependencies):
    reversed_dependencies = reverse_dependencies(dependencies)
    remaining_steps = find_start_labels(reversed_dependencies)
    heapq.heapify(remaining_steps)
    construction_order = []
    while remaining_steps:
        step = heapq.heappop(remaining_steps)
        construction_order.append(step)

        for next_step in dependencies[step]:
            reversed_dependencies[next_step].discard(step)

            if reversed_dependencies[next_step]:
                continue

            heapq.heappush(remaining_steps, next_step)

    return construction_order


def reverse_dependencies(dependencies):
    reversed_dependencies = {}
    for step, dependent_steps in dependencies.items():
        if step not in reversed_dependencies:
            reversed_dependencies[step] = set()

        for dependent_step in dependent_steps:
            reversed_dependencies[dependent_step] = reversed_dependencies.get(
                dependent_step, set()
            ) | {step}

    return reversed_dependencies


def find_start_labels(reversed_dependencies):
    candidates = []
    for step, previous_steps in reversed_dependencies.items():
        if not previous_steps:
            candidates.append(step)

    return candidates


def determine_construction_time(dependencies, n_workers, step_startup_cost):
    reversed_dependencies = reverse_dependencies(dependencies)

    available_workers = collections.deque(x for x in range(n_workers))

    remaining_steps = find_start_labels(reversed_dependencies)
    heapq.heapify(remaining_steps)

    current_work_queue = []

    current_time = 0

    def step_construction_time(label):
        return ord(label) - ord("A") + 1 + step_startup_cost

    while remaining_steps or current_work_queue:
        while remaining_steps and available_workers:
            next_worker = available_workers.pop()
            next_step = heapq.heappop(remaining_steps)
            start_time = current_time
            completion_time = current_time + step_construction_time(next_step)
            heapq.heappush(
                current_work_queue,
                (completion_time, start_time, next_step, next_worker),
            )

        current_time, start_time, step, worker = heapq.heappop(current_work_queue)

        available_workers.appendleft(worker)

        for next_step in dependencies[step]:
            reversed_dependencies[next_step].discard(step)

            if reversed_dependencies[next_step]:
                continue

            heapq.heappush(remaining_steps, next_step)

    return current_time
