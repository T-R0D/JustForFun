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


class State(object):
    def __init__(self):
        pass

    def generate_next_states(self):
        raise NotImplementedError()

    def is_goal(self):
        raise NotImplementedError()

    def __str__(self):
        raise NotImplementedError()


def graph_search(initial_state: State, frontier=None):
    explored = set()
    frontier = frontier
    if frontier == None:
        frontier = queue.deque()
    frontier.append((initial_state, []))
    explored.add(str(initial_state))

    while frontier:
        current_state, path = frontier.popleft()
        explored.add(str(current_state))

        if current_state.is_goal():
            return current_state, path

        for state in current_state.generate_next_states():
            if not str(state) in explored:
                frontier.append((state, path + [str(current_state)]))

                # print(repr(state))
                #
                # print('ex: {} fr: {} d: {}'.format(len(explored), len(frontier), len(frontier[0][1]) if frontier else None))
                # print('\n\n'.join([str(x[0]) for x in frontier]))
                # frontier.clear()

    return None, None
