# This file is part of aoc2016.
#
# aoc2016 is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# aoc2016 is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with aoc2016.  If not, see <http://www.gnu.org/licenses/>.
import aoc2016.common.search as search


def part_one(puzzle_input):
    grid, targets = parse_grid(puzzle_input)

    adjacency_matrix = find_adjacency_matrix(grid, targets)

    def count_steps_to_cover(
        adjacency_matrix,
        current_location,
        remaining_targets,
    ):
        if not remaining_targets:
            return 0

        steps_required = 999_999_999
        for target in remaining_targets:
            steps_to_target = adjacency_matrix[current_location][target]
            steps_to_cover_the_rest = count_steps_to_cover(
                adjacency_matrix,
                target,
                remaining_targets.difference({target}),
            )

            steps_required = min(
                steps_required, steps_to_target + steps_to_cover_the_rest
            )

        return steps_required

    steps_required = count_steps_to_cover(
        adjacency_matrix, 0, set(targets.keys()) - {0}
    )

    return str(steps_required)


def part_two(puzzle_input):
    grid, targets = parse_grid(puzzle_input)
    n_locations = len(targets.keys())
    targets[n_locations] = targets[0]

    adjacency_matrix = find_adjacency_matrix(grid, targets)

    def count_steps_to_cover(
        adjacency_matrix, current_location, remaining_targets, last_location
    ):
        if not remaining_targets:
            return adjacency_matrix[current_location][last_location]

        steps_required = 999_999_999
        for target in remaining_targets:
            steps_to_target = adjacency_matrix[current_location][target]
            steps_to_cover_the_rest = count_steps_to_cover(
                adjacency_matrix,
                target,
                remaining_targets.difference({target}),
                last_location
            )

            steps_required = min(
                steps_required, steps_to_target + steps_to_cover_the_rest
            )

        return steps_required

    steps_required = count_steps_to_cover(
        adjacency_matrix, 0, set(targets.keys()) - {0, n_locations}, n_locations
    )

    return str(steps_required)


def parse_grid(puzzle_input):
    grid = []
    targets = {}
    for i, line in enumerate(puzzle_input.split("\n")):
        row = []
        for j, cell in enumerate(list(line)):
            row.append(cell)
            try:
                x = int(cell)
                targets[x] = (i, j)
            except:
                pass

        grid.append(row)

    return grid, targets


def find_adjacency_matrix(grid, targets):
    n = len(list(targets.keys()))
    adjacency_matrix = [[-1 for _ in range(n)] for _ in range(n)]

    for label_a, location_a in targets.items():
        for label_b, location_b in targets.items():
            if label_a == label_b:
                continue

            if adjacency_matrix[label_a][label_b] != -1:
                continue

            initial_state = SearchState(grid, location_a, location_b)
            _, path = search.graph_search(initial_state)
            n_steps = len(path) - 1

            adjacency_matrix[label_a][label_b] = n_steps
            adjacency_matrix[label_b][label_a] = n_steps

    return adjacency_matrix


class SearchState:
    def __init__(self, grid, current_location, destination):
        self.grid = grid
        self.current_location = current_location
        self.destination = destination

    def generate_next_states(self):
        i, j = self.current_location
        next_states = []
        for di, dj in ((0, -1), (0, 1), (-1, 0), (1, 0)):
            new_i, new_j = i + di, j + dj

            if (
                new_i < 0
                or new_j < 0
                or len(self.grid) <= new_i
                or len(self.grid[0]) <= new_j
            ):
                continue

            if self.grid[new_i][new_j] == "#":
                continue

            next_states.append(SearchState(self.grid, (new_i, new_j), self.destination))

        return next_states

    def is_goal(self):
        return self.current_location == self.destination

    def key(self):
        return str(self.current_location)
