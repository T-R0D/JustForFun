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

import copy
import hashlib
import queue

from aoc2016.common.search import State, graph_search

class MazeState(State):
    DOOR_OPEN_VALUES = set('bcdef')

    def __init__(self, passcode, position, path_so_far):
        super().__init__()

        self.passcode = passcode
        self.position = position
        self.path_so_far = path_so_far

    def generate_next_states(self):
        door_signals = MazeState.hash_passcode_and_path(self.passcode, self.path_so_far)[:4]
        next_states = []
        # Up.
        if door_signals[0] in self.DOOR_OPEN_VALUES:
            if self.position[1] > 0:
                next_states += [
                    MazeState(self.passcode, (self.position[0], self.position[1] - 1), self.path_so_far + ['U'])
                ]
        # Down.
        if door_signals[1] in self.DOOR_OPEN_VALUES:
            if self.position[1] < 3:
                next_states += [
                    MazeState(self.passcode, (self.position[0], self.position[1] + 1), self.path_so_far + ['D'])
                ]
        # Left.
        if door_signals[2] in self.DOOR_OPEN_VALUES:
            if self.position[0] > 0:
                next_states += [
                    MazeState(self.passcode, (self.position[0] - 1, self.position[1]), self.path_so_far + ['L'])
                ]
        # Right.
        if door_signals[3] in self.DOOR_OPEN_VALUES:
            if self.position[0] < 3:
                next_states += [
                    MazeState(self.passcode, (self.position[0] + 1, self.position[1]), self.path_so_far + ['R'])
                ]

        return next_states

    def is_goal(self):
        return (3, 3) == self.position

    def get_path(self):
        return ''.join(self.path_so_far)

    @staticmethod
    def hash_passcode_and_path(passcode, path_so_far):
        data = '{}{}'.format(passcode, ''.join(path_so_far)).encode('utf-8')
        m = hashlib.md5()
        m.update(data)
        return m.hexdigest()

def part_one(puzzle_input):
    initial_state = MazeState(passcode=puzzle_input, position=(0, 0), path_so_far=[])
    explored = set()
    frontier = queue.deque()
    frontier.append(initial_state)
    explored.add(initial_state.get_path())

    while frontier:
        current_state = frontier.popleft()
        explored.add(current_state.get_path())

        if current_state.is_goal():
            return ''.join(current_state.get_path())

        for state in current_state.generate_next_states():
            if state.get_path() not in explored:
                frontier.append(state)

    return None



def part_two(puzzle_input):
    initial_state = MazeState(passcode=puzzle_input, position=(0, 0), path_so_far=[])
    explored = set()
    frontier = queue.deque()
    frontier.append(initial_state)
    explored.add(initial_state.get_path())

    max_path = []

    while frontier:
        current_state = frontier.pop()
        explored.add(current_state.get_path())

        if current_state.is_goal():
            path = current_state.get_path()
            if len(path) > len(max_path):
                max_path = path
            continue

        for state in current_state.generate_next_states():
            if state.get_path() not in explored:
                frontier.append(state)

    return len(max_path) if max_path else None

