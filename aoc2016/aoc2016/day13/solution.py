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

import queue

def part_one(puzzle_input):
    designer_number = parse_input(puzzle_input)
    room = DilbertLand(designer_number, (31, 39))
    path = room.find_route_to_exit()
    return len(path)


def part_two(puzzle_input):
    designer_number = parse_input(puzzle_input)
    room = DilbertLand(designer_number, (31, 39))
    path = room.find_highest_coverage_path(50)
    return len(path)


def parse_input(puzzle_input):
    return int(puzzle_input)


class DilbertLand(object):
    def __init__(self, designer_number, exit):
        self.designer_number = designer_number
        self.exit = exit
        self.walls = {}
        self.paths = {}
        self.exit_route = None

    def find_route_to_exit(self):
        if self.exit_route:
            return self.exit_route

        frontier = queue.deque([((1, 1), [])])
        while frontier:
            current_position, path = frontier.popleft()
            if current_position == self.exit:
                self.exit_route = path
                return self.exit_route

            next_positions = self.generate_next_positions(current_position)
            for position in next_positions:
                x, y = position
                if (x in self.walls and y in self.walls.get(x, set())) or (
                        x in self.paths and y in self.paths.get(x, set())):
                    continue  # Already explored.
                if self.is_wall(x, y):
                    ys = self.walls.get(x, set())
                    ys.add(y)
                    self.walls[x] = ys
                else:
                    ys = self.paths.get(x, set())
                    ys.add(y)
                    self.paths[x] = ys
                    frontier.append(((x, y), path + [(x, y)]))

    def find_highest_coverage_path(self, path_length):
        highest_coverage = 0
        best_path = []

        class CoverageState(object):
            def __init__(self, position, path, covered):
                self.position = position
                self.path = path
                self.covered = covered

        frontier = queue.deque([CoverageState((1, 1), [], set())])

        while frontier:
            current_state = frontier.pop()

            if len(current_state.path) > path_length:
                continue

            if len(current_state.covered) > highest_coverage:
                highest_coverage = len(current_state.covered)
                best_path = current_state.path

            next_positions = self.generate_next_positions(current_state.position)
            for position in next_positions:
                x, y = position
                if (x in self.walls and y in self.walls.get(x, set())):
                    continue

                if self.is_wall(x, y):
                    ys = self.walls.get(x, set())
                    ys.add(y)
                    self.walls[x] = ys
                    continue

                next_state = CoverageState(
                    position, current_state.path + [position], current_state.covered | {position})
                frontier.append(next_state)


        return best_path

    def generate_next_positions(self, position):
        x, y = position
        next_positions = []

        # North.
        if (y - 1) >= 0:
            next_positions.append((x, y - 1))
        # East.
        next_positions.append((x + 1, y))
        # South.
        next_positions.append((x, y + 1))
        # West.
        if (x - 1) >= 0:
            next_positions.append((x - 1, y))

        return next_positions

    def is_wall(self, x, y):
        magic_number = self.compute_magic_number(x, y)
        return True if self.count_bits(magic_number) & 1 == 1 else False

    def compute_magic_number(self, x, y):
        return (x * x) + (3 * x) + (2 * x * y) + y + (y * y) + self.designer_number

    def count_bits(self, x):
        a = x
        n_bits = 0
        while a > 0:
            n_bits += a & 1
            a >>= 1
        return n_bits
