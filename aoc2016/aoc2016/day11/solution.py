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
import re
from aoc2016.common.search import graph_search


def part_one(puzzle_input):
    element_to_floor_locations = get_element_to_floor_locations(puzzle_input)

    initial_state = SlottedState.from_dict(element_to_floor_locations, 0, 4)

    _, path = graph_search(initial_state)
    if not path:
        return "No solution found"

    return str(len(path) - 1)


def part_two(puzzle_input):
    element_to_floor_locations = get_element_to_floor_locations(puzzle_input)
    element_to_floor_locations["elerium"] = [0, 0]
    element_to_floor_locations["dilithium"] = [0, 0]

    initial_state = SlottedState.from_dict(element_to_floor_locations, 0, 4)

    _, path = graph_search(initial_state)
    if not path:
        return "No solution found"

    return str(len(path) - 1)


def get_element_to_floor_locations(puzzle_input):
    element_to_floor_locations = {}
    for i, line in enumerate(puzzle_input.split("\n")):
        chips, rtgs = identify_components(line)

        for rtg in rtgs:
            if rtg not in element_to_floor_locations:
                element_to_floor_locations[rtg] = [-1, -1]
            element_to_floor_locations[rtg][0] = i

        for chip in chips:
            if chip not in element_to_floor_locations:
                element_to_floor_locations[chip] = [-1, -1]
            element_to_floor_locations[chip][1] = i

    return element_to_floor_locations


def identify_components(line):
    microchips = []
    rtgs = []
    if not "nothing relevant" in line:
        microchips = re.findall("\w+(?=-compatible microchip)", line)
        rtgs = re.findall("\w+(?= generator)", line)
    return microchips, rtgs


class SlottedState:
    def __init__(self, elevator_pos, item_locations, n_floors):
        self.elevator_pos = elevator_pos
        self.item_locations = item_locations
        self.n_floors = n_floors

    @classmethod
    def from_dict(cls, layout, elevator_pos, n_floors):
        n_elements = len(layout)
        item_locations = [0 for _ in range(2 * n_elements)]
        for i, element in enumerate(sorted(layout.keys())):
            item_locations[2 * i] = layout[element][0]
            item_locations[(2 * i) + 1] = layout[element][1]

        return cls(elevator_pos, item_locations, n_floors)

    def generate_next_states(self):
        next_states = []

        next_elevator_positions = []
        if self.elevator_pos > 0 and any(item_loc < self.elevator_pos for item_loc in self.item_locations):
            next_elevator_positions.append(self.elevator_pos - 1)
        
        if self.elevator_pos < self.n_floors - 1:
            next_elevator_positions.append(self.elevator_pos + 1)

        rtgs_on_floor = [
            i
            for i in range(0, len(self.item_locations), 2)
            if self.item_locations[i] == self.elevator_pos
        ]
        chips_on_floor = [
            i
            for i in range(1, len(self.item_locations), 2)
            if self.item_locations[i] == self.elevator_pos
        ]
        pairs_on_floor = [
            (i, i + 1)
            for i in range(0, len(self.item_locations), 2)
            if self.item_locations[i] == self.item_locations[i + 1]
            and self.item_locations[i] == self.elevator_pos
        ]

        for new_elevator_pos in next_elevator_positions:
            one_item_next_states = []
            two_item_next_states = []

            # Take one microchip.
            for chip_index in chips_on_floor:
                new_item_locations = [x for x in self.item_locations]
                new_item_locations[chip_index] = new_elevator_pos
                next_state = SlottedState(
                    new_elevator_pos, new_item_locations, self.n_floors
                )
                if not next_state.is_good():
                    continue
                one_item_next_states.append(next_state)

            # Take one RTG.
            for rtg_index in rtgs_on_floor:
                new_item_locations = [x for x in self.item_locations]
                new_item_locations[rtg_index] = new_elevator_pos
                next_state = SlottedState(
                    new_elevator_pos, new_item_locations, self.n_floors
                )
                if not next_state.is_good():
                    continue
                one_item_next_states.append(next_state)

            # Take two microchips.
            for chip_a_index, chip_b_index in zip(chips_on_floor, chips_on_floor[1:]):
                new_item_locations = [x for x in self.item_locations]
                new_item_locations[chip_a_index] = new_elevator_pos
                new_item_locations[chip_b_index] = new_elevator_pos
                next_state = SlottedState(
                    new_elevator_pos, new_item_locations, self.n_floors
                )
                if not next_state.is_good():
                    continue
                two_item_next_states.append(next_state)

            # Take two RTGs.
            for rtg_a_index, rtg_b_index in zip(rtgs_on_floor, rtgs_on_floor[1:]):
                new_item_locations = [x for x in self.item_locations]
                new_item_locations[rtg_a_index] = new_elevator_pos
                new_item_locations[rtg_b_index] = new_elevator_pos
                next_state = SlottedState(
                    new_elevator_pos, new_item_locations, self.n_floors
                )
                if not next_state.is_good():
                    continue
                two_item_next_states.append(next_state)

            # Take a matching RTG and microchip.
            for a, b in pairs_on_floor:
                new_item_locations = [x for x in self.item_locations]
                new_item_locations[a] = new_elevator_pos
                new_item_locations[b] = new_elevator_pos
                next_state = SlottedState(
                    new_elevator_pos, new_item_locations, self.n_floors
                )
                if not next_state.is_good():
                    continue
                two_item_next_states.append(next_state)

            if self.elevator_pos < new_elevator_pos:
                next_states.extend(two_item_next_states)
                if not two_item_next_states:
                    next_states.extend(one_item_next_states)
            elif self.elevator_pos > new_elevator_pos:
                next_states.extend(one_item_next_states)
                if not one_item_next_states:
                    next_states.extend(two_item_next_states)

        return next_states

    def is_goal(self):
        return self.elevator_pos == self.n_floors - 1 and all(
            loc == self.n_floors - 1 for loc in self.item_locations
        )

    def key(self):
        component_pairs = list(
            sorted(
                (self.item_locations[i], self.item_locations[i + 1])
                for i in range(0, len(self.item_locations), 2)
            )
        )
        return f"{self.elevator_pos} | {component_pairs}"

    def is_good(self):
        floor_has_rtg = {i: False for i in range(self.n_floors)}
        for i in range(0, len(self.item_locations), 2):
            floor_has_rtg[self.item_locations[i]] = True

        for i in range(0, len(self.item_locations), 2):
            if (
                self.item_locations[i] != self.item_locations[i + 1]
                and floor_has_rtg[self.item_locations[i + 1]]
            ):
                return False

        return True
