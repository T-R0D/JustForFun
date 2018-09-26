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
import re
from aoc2016.common.search import State, graph_search


class RoomState(State):
    def __init__(self, elevator, floors):
        super().__init__()
        self.elevator = elevator
        self.floors = floors

    def current_floor(self):
        return self.floors[self.elevator]

    def generate_next_states(self):
        next_states = []
        elevator_positions = {max(self.elevator - 1, 0), min(self.elevator + 1, len(self.floors) - 1)}
        elevator_positions = elevator_positions.difference([self.elevator])

        for elevator_position in elevator_positions:
            current_floor = self.current_floor()

            # Take one RTG.
            for rtg in current_floor.rtgs:
                floors = copy.deepcopy(self.floors)
                floors[self.elevator].rtgs.remove(rtg)
                floors[elevator_position].rtgs.append(rtg)
                state = RoomState(elevator_position, floors)
                if not state.is_bad():
                    next_states.append(state)

            # Take two RTGs.
            rtgs_to_move = {tuple(sorted([rtg1, rtg2]))
                            for rtg1 in current_floor.rtgs
                            for rtg2 in current_floor.rtgs
                            if rtg1 != rtg2}
            for rtg1, rtg2 in rtgs_to_move:
                floors = copy.deepcopy(self.floors)
                floors[self.elevator].rtgs.remove(rtg1)
                floors[elevator_position].rtgs.append(rtg1)
                floors[self.elevator].rtgs.remove(rtg2)
                floors[elevator_position].rtgs.append(rtg2)
                state = RoomState(elevator_position, floors)
                if not state.is_bad():
                    next_states.append(state)


            # Take one microchip.
            for chip in current_floor.chips:
                floors = copy.deepcopy(self.floors)
                floors[self.elevator].chips.remove(chip)
                floors[elevator_position].chips.append(chip)
                state = RoomState(elevator_position, floors)
                if not state.is_bad():
                    next_states.append(state)

            # Take two chips.
            chips_to_move = {tuple(sorted([chip1, chip2]))
                            for chip1 in current_floor.chips
                            for chip2 in current_floor.chips
                            if chip1 != chip2}
            for chip1, chip2 in chips_to_move:
                floors = copy.deepcopy(self.floors)
                floors[self.elevator].chips.remove(chip1)
                floors[elevator_position].chips.append(chip1)
                floors[self.elevator].chips.remove(chip2)
                floors[elevator_position].chips.append(chip2)
                state = RoomState(elevator_position, floors)
                if not state.is_bad():
                    next_states.append(state)

            # Take a matching RTG and microchip.
            for rtg in current_floor.rtgs:
                if rtg in current_floor.chips:
                    floors = copy.deepcopy(self.floors)
                    floors[self.elevator].rtgs.remove(rtg)
                    floors[elevator_position].rtgs.append(rtg)
                    floors[self.elevator].chips.remove(rtg)
                    floors[elevator_position].chips.append(rtg)
                    state = RoomState(elevator_position, floors)
                    if not state.is_bad():
                        next_states.append(state)

        return next_states

    def is_bad(self):
        for floor in self.floors:
            if floor.is_bad():
                return True

        return False

    def is_goal(self):
        last_floor_index = len(self.floors) - 1
        if self.elevator != last_floor_index:
            return False

        for floor in self.floors[: -1]:
            if not floor.empty():
                return False

        return True

    def __str__(self):
        strings = []
        for i, floor in enumerate(self.floors):
            strings.append('{} -> {}'.format('E' if self.elevator == i else ' ', str(floor)))
        return '\n'.join(strings)

    def __repr__(self):
        return str(self)


class Floor(object):
    def __init__(self, rtgs, chips):
        self.rtgs = rtgs
        self.chips = chips

    def is_bad(self):
        if self.rtgs:
            for chip in self.chips:
                if chip not in self.rtgs:
                    return True

        return False

    def empty(self):
        return not self.rtgs and not self.chips

    def __str__(self):
        return 'RTGS: {}, Chips: {}'.format(sorted(self.rtgs), sorted(self.chips))

    def __repr__(self):
        return str(self)


class BitRoom(State):
    FLOORS = range(0, 4)
    SLOTS = 11
    RTGS = range(1, SLOTS, 2)
    CHIPS = range(2, SLOTS, 2)

    def __init__(self, current_floor = 0, layout = None):
        super().__init__()
        self.current_floor = current_floor
        self.layout = layout if layout else \
            [1] + [1, 1] + [1, 0] + [1, 0] + [1, 1] + [1, 1] + \
            [0] + [0, 0] + [0, 1] + [0, 1] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0]
            #E    Co       Po       Pr       Ru       T

    def is_valid(self):
        elevator = 0
        for floor in self.FLOORS:
            if (floor * len(self.FLOORS)) == 1:
                elevator = floor

        for floor in self.FLOORS:
            rtgs = self.get_indices_of_items_on_floor(floor, self.RTGS)
            chips = self.get_indices_of_items_on_floor(floor, self.CHIPS)

            if floor == elevator and (not rtgs and not chips):
                return False

            if rtgs:
                for chip in chips:
                    if (chip - 1) not in rtgs:
                        return False

        return True

    def generate_next_states(self):
        next_states = []
        current_floor = self.current_floor

        rtgs = self.get_indices_of_items_on_floor(current_floor, self.RTGS)
        chips = self.get_indices_of_items_on_floor(current_floor, self.CHIPS)

        if current_floor > 0:
            down = -1
            next_states.extend(self.move_one_rtg(down, rtgs))
            next_states.extend(self.move_one_chip(down, chips))
            next_states.extend(self.move_two_rtgs(down, rtgs))
            next_states.extend(self.move_two_chips(down, chips))
            next_states.extend(self.move_an_rtg_and_chip(down, rtgs, chips))

        if current_floor < (3):
            up = 1
            next_states.extend(self.move_one_rtg(direction=1, rtgs=rtgs))
            next_states.extend(self.move_one_chip(direction=1, chips=chips))
            next_states.extend(self.move_two_rtgs(direction=1, rtgs=rtgs))
            next_states.extend(self.move_two_chips(direction=1, chips=chips))
            next_states.extend(self.move_an_rtg_and_chip(direction=1, rtgs=rtgs, chips=chips))

        return next_states

    def move_one_rtg(self, direction, rtgs):
        if direction not in (-1, 1):
            raise Exception('Direction must be -1 or 1')

        next_states = []
        start_offset = self.current_floor * self.SLOTS
        end_offset = start_offset + (direction * self.SLOTS)

        new_floor = self.current_floor + direction
        if new_floor < 0 or 3 < new_floor:
            raise Exception('Can\'t move to floor {}'.format(new_floor))

        for rtg in rtgs:
            layout = copy.deepcopy(self.layout)

            # Move the elevator.
            layout[start_offset] = 0
            layout[end_offset] = 1
            # Move the RTG.
            layout[rtg + start_offset] = 0
            layout[rtg + end_offset] = 1

            state = BitRoom(new_floor, layout)
            if state.is_valid():
                next_states.append(state)

        return next_states

    def move_one_chip(self, direction, chips):
        if direction not in (-1, 1):
            raise Exception('Direction must be -1 or 1')

        next_states = []
        start_offset = self.current_floor * self.SLOTS
        end_offset = start_offset + (direction * self.SLOTS)

        new_floor = self.current_floor + direction
        if new_floor < 0 or 3 < new_floor:
            raise Exception('Can\'t move to floor {}'.format(new_floor))

        for chip in chips:
            layout = copy.deepcopy(self.layout)

            # Move the elevator.
            layout[start_offset] = 0
            layout[end_offset] = 1
            # Move the RTG.
            layout[chip + start_offset] = 0
            layout[chip + end_offset] = 1

            state = BitRoom(new_floor, layout)
            if state.is_valid():
                next_states.append(state)

        return next_states

    def move_two_rtgs(self, direction, rtgs):
        if direction not in (-1, 1):
            raise Exception('Direction must be -1 or 1')

        next_states = []
        start_offset = self.current_floor * self.SLOTS
        end_offset = start_offset + (direction * self.SLOTS)

        new_floor = self.current_floor + direction
        if new_floor < 0 or 3 < new_floor:
            raise Exception('Can\'t move to floor {}'.format(new_floor))

        # Take two RTGs.
        rtgs_to_move = {tuple(sorted([rtg1, rtg2])) for rtg1 in rtgs for rtg2 in rtgs if rtg1 != rtg2}

        for rtg1, rtg2 in rtgs_to_move:
            layout = copy.deepcopy(self.layout)

            # Move the elevator.
            layout[start_offset] = 0
            layout[end_offset] = 1
            # Move the RTGs.
            layout[rtg1 + start_offset] = 0
            layout[rtg1 + end_offset] = 1
            layout[rtg2 + start_offset] = 0
            layout[rtg2 + end_offset] = 1

            state = BitRoom(new_floor, layout)
            if state.is_valid():
                next_states.append(state)

        return next_states

    def move_two_chips(self, direction, chips):
        if direction not in (-1, 1):
            raise Exception('Direction must be -1 or 1')

        next_states = []
        start_offset = self.current_floor * self.SLOTS
        end_offset = start_offset + (direction * self.SLOTS)

        new_floor = self.current_floor + direction
        if new_floor < 0 or 3 < new_floor:
            raise Exception('Can\'t move to floor {}'.format(new_floor))

        # Take two RTGs.
        chips_to_move = {tuple(sorted([chip1, chip2])) for chip1 in chips for chip2 in chips if chip1 != chip2}

        for chip1, chip2 in chips_to_move:
            layout = copy.deepcopy(self.layout)

            # Move the elevator.
            layout[start_offset] = 0
            layout[end_offset] = 1
            # Move the RTGs.
            layout[chip1 + start_offset] = 0
            layout[chip1 + end_offset] = 1
            layout[chip2 + start_offset] = 0
            layout[chip2 + end_offset] = 1

            state = BitRoom(new_floor, layout)
            if state.is_valid():
                next_states.append(state)

        return next_states

    def move_an_rtg_and_chip(self, direction, rtgs, chips):
        if direction not in (-1, 1):
            raise Exception('Direction must be -1 or 1')

        next_states = []
        start_offset = self.current_floor * self.SLOTS
        end_offset = start_offset + (direction * self.SLOTS)

        new_floor = self.current_floor + direction
        if new_floor < 0 or 3 < new_floor:
            raise Exception('Can\'t move to floor {}'.format(new_floor))

        # Just move a chip and an RTG and let the validation step handle checking if they can be moved together.
        # Assuming a valid start state, if things end in a valid state, things are fine.
        for rtg in rtgs:
            for chip in chips:
                layout = copy.deepcopy(self.layout)

                # Move the elevator.
                layout[start_offset] = 0
                layout[end_offset] = 1
                # Move the RTG.
                layout[rtg + start_offset] = 0
                layout[rtg + end_offset] = 1
                # Move the chip.
                layout[chip + start_offset] = 0
                layout[chip + end_offset] = 1

                state = BitRoom(new_floor, layout)
                if state.is_valid():
                    next_states.append(state)

        return next_states

    def is_goal(self):
        return all(map(lambda x: True if x == 0 else False, self.layout[0:11])) and \
        all(map(lambda x: True if x == 0 else False, self.layout[11:22])) and \
        all(map(lambda x: True if x == 0 else False, self.layout[22:33])) and \
        all(map(lambda x: True if x == 0 else False, self.layout[33:44]))

    def get_indices_of_items_on_floor(self, floor, items):
        offset = floor * self.SLOTS
        # print(floor, offset, items[-1]+offset)
        return [(x + offset) for x in items if self.layout[(x + offset)] != 0]

    def __str__(self):
        return str(self.layout)

    def __repr__(self):
        END = 4 * self.SLOTS
        cobalt_rtg =      ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(1, END, self.SLOTS)])
        cobalt_chip =     ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(2, END, self.SLOTS)])
        polonium_rtg =    ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(3, END, self.SLOTS)])
        polonium_chip =   ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(4, END, self.SLOTS)])
        prometheum_rtg =  ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(5, END, self.SLOTS)])
        prometheum_chip = ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(6, END, self.SLOTS)])
        rubidium_rtg =    ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(7, END, self.SLOTS)])
        rubidium_chip =   ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(8, END, self.SLOTS)])
        thulium_rtg =     ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(9, END, self.SLOTS)])
        thulium_chip =    ' | '.join(['X' if self.layout[x] != 0 else ' ' for x in range(10, END, self.SLOTS)])
        elevator =        ' | '.join(['E' if self.layout[x] != 0 else ' ' for x in range(0, END, self.SLOTS)])
        s = '    Cobalt RTG: {}\n' + \
               '          Chip: {}\n' + \
               '  Polonium RTG: {}\n' + \
               '          Chip: {}\n' + \
               'Promethium RTG: {}\n' + \
               '          Chip: {}\n' + \
               '  Rubidium RTG: {}\n' + \
               '          Chip: {}\n' + \
               '   Thulium RTG: {}\n' + \
               '          Chip: {}\n' + \
               '\n' + \
               '      Elevator: {}'
        return s.format(cobalt_rtg, cobalt_chip, polonium_rtg, polonium_chip, prometheum_rtg,
                                           prometheum_chip, rubidium_rtg, rubidium_chip, thulium_rtg, thulium_chip,
                                           elevator)


def part_one(puzzle_input):
    # layout = create_floor_layout(puzzle_input)
    # initial_state = RoomState(0, layout)
    initial_state = BitRoom()
    return graph_search(initial_state)



def part_two(puzzle_input):
    pass

def create_floor_layout(puzzle_input):
    layout = []
    for line in puzzle_input.split('\n'):
        chips, rtgs = identify_components(line)
        layout.append(Floor(rtgs, chips))
    return layout

def identify_components(line):
    microchips = []
    rtgs = []
    if not 'nothing relevant' in line:
        microchips = re.findall('\w+(?=-compatible microchip)', line)
        rtgs =  re.findall('\w+(?= generator)', line)
    return microchips, rtgs
