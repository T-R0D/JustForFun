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
    _, nodes = parse_node_grid(puzzle_input)

    nodes.sort(key=lambda x: x.used)
    availables = sorted(node.available for node in nodes)

    n_compatible = 0
    j = 0
    for node in nodes:
        if node.used == 0:
            continue

        while j < len(availables) and availables[j] < node.used:
            j += 1

        if j == len(availables):
            break

        n_compatible += len(availables) - j
        if node.available >= node.used:
            n_compatible -= 1

    return str(n_compatible)


def part_two(puzzle_input):
    node_grid, nodes = parse_node_grid(puzzle_input)

    target_node_coordinates = (len(node_grid) - 1, 0)
    free_node_coordinates = (0, 0)
    for node in nodes:
        if node.used_percent == 0:
            free_node_coordinates = (node.x, node.y)
            break
    goal_node_coordinates = (0, 0)

    initial_state = DataShiftState(
        node_grid,
        goal_node_coordinates,
        target_node_coordinates,
        free_node_coordinates,
        node_grid[free_node_coordinates[0]][free_node_coordinates[1]].available,
    )

    _, path = search.graph_search(initial_state)

    return str(len(path) - 1)


def parse_node_grid(puzzle_input):
    nodes = []
    max_x, max_y = 0, 0
    for line in puzzle_input.split("\n")[2:]:
        node = Node.from_line(line)
        nodes.append(node)
        max_x = max(max_x, node.x)
        max_y = max(max_y, node.y)

    node_grid = [[None for _ in range(max_y + 1)] for _ in range(max_x + 1)]
    for node in nodes:
        node_grid[node.x][node.y] = node

    return node_grid, nodes


class Node:
    def __init__(self, x, y, size, used, available, used_percent):
        self.x = x
        self.y = y
        self.size = size
        self.used = used
        self.available = available
        self.used_percent = used_percent

    @classmethod
    def from_line(cls, line):
        filesystem, size, used, available, used_percent = line.split()
        x, y = (
            int(v)
            for v in filesystem.replace("/dev/grid/node-x", "")
            .replace("-y", " ")
            .split()
        )
        return cls(
            x,
            y,
            int(size.replace("T", "")),
            int(used.replace("T", "")),
            int(available.replace("T", "")),
            int(used_percent.replace("%", "")),
        )


class DataShiftState:
    def __init__(
        self,
        node_grid,
        goal_coordinates,
        target_data_coordinates,
        free_node_coordinates,
        free_node_space,
    ):
        self.node_grid = node_grid
        self.goal_coordinates = goal_coordinates
        self.target_data_coordinates = target_data_coordinates
        self.free_node_coordinates = free_node_coordinates
        self.free_node_space = free_node_space

    def generate_next_states(self):
        next_states = []
        x, y = self.free_node_coordinates
        for dx, dy in ((0, -1), (0, 1), (-1, 0), (1, 0)):
            new_x, new_y = x + dx, y + dy
            if (
                new_x < 0
                or new_y < 0
                or new_x >= len(self.node_grid)
                or new_y >= len(self.node_grid[0])
            ):
                continue

            if self.node_grid[new_x][new_y].used > self.free_node_space:
                continue

            new_free_node_coordinates = (new_x, new_y)
            target_data_coordinates = self.target_data_coordinates
            if new_free_node_coordinates == self.target_data_coordinates:
                target_data_coordinates = (x, y)

            next_states.append(
                DataShiftState(
                    self.node_grid,
                    self.goal_coordinates,
                    target_data_coordinates,
                    new_free_node_coordinates,
                    self.node_grid[new_x][new_y].size,
                )
            )

        return next_states

    def is_goal(self):
        return self.target_data_coordinates == self.goal_coordinates

    def key(self):
        return f"{self.target_data_coordinates}{self.free_node_coordinates}"

    def heuristic(self):
        return abs(
            self.target_data_coordinates[0] - self.free_node_coordinates[0]
        ) + abs(self.target_data_coordinates[1] - self.free_node_coordinates[1])
