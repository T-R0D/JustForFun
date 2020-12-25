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
import unittest
from aoc2016.common.search import graph_search
from aoc2016.day11 import solution as sut

INPUT = 'The first floor contains a polonium generator, a thulium generator, a thulium-compatible microchip, a promethium generator, a ruthenium generator, a ruthenium-compatible microchip, a cobalt generator, and a cobalt-compatible microchip.'
FULL_INPUT = 'The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.\n' + \
'The second floor contains a hydrogen generator.\n' + \
'The third floor contains a lithium generator.\n' + \
'The fourth floor contains nothing relevant.'


class TestDay11Solution(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_identify_components(self):
        expected_microchips = ['thulium', 'ruthenium', 'cobalt',]
        expected_generators = ['polonium', 'thulium', 'promethium', 'ruthenium', 'cobalt',]
        microchips, generators = sut.identify_components(INPUT)
        self.assertListEqual(expected_microchips, microchips)
        self.assertListEqual(expected_generators, generators)

    def test_identify_components_empty(self):
        expected_microchips = []
        expected_generators = []
        microchips, generators = sut.identify_components(
            'The fourth floor contains nothing relevant.')
        self.assertListEqual(expected_microchips, microchips)
        self.assertListEqual(expected_generators, generators)

    def test_generate_next_states(self):
        floors = sut.create_floor_layout(FULL_INPUT)
        room = sut.RoomState(0, floors)

        states = room.generate_next_states()
        string = ''
        for state in states:
            string += str(state)
            string += '\n ~~~ \n'
        print(string)

    def test_bitroom_repr(self):
        expected = \
            '    Cobalt RTG: X |   |   |  \n' + \
            '          Chip: X |   |   |  \n' + \
            '  Polonium RTG: X |   |   |  \n' + \
            '          Chip:   | X |   |  \n' + \
            'Promethium RTG: X |   |   |  \n' + \
            '          Chip:   | X |   |  \n' + \
            '  Rubidium RTG: X |   |   |  \n' + \
            '          Chip: X |   |   |  \n' + \
            '   Thulium RTG: X |   |   |  \n' + \
            '          Chip: X |   |   |  \n' + \
            '\n' + \
            '      Elevator: E |   |   |  '
        state = sut.BitRoom()
        self.assertEqual(expected, repr(state))

    def test_bitroom_simple_case(self):
        starting_state = sut.BitRoom(current_floor=0, layout=\
            [1] + [0, 1] + [0, 1] + [0, 0, 0, 0, 0, 0] +\
            [0] + [1, 0] + [0, 0] + [0, 0, 0, 0, 0, 0] +\
            [0] + [0, 0] + [1, 0] + [0, 0, 0, 0, 0, 0] +\
            [0] + [0, 0] + [0, 0] + [0, 0, 0, 0, 0, 0])
            #E    H        L
        expected = \
            '    Cobalt RTG:   | X |   |  \n' + \
            '          Chip: X |   |   |  \n' + \
            '  Polonium RTG:   |   | X |  \n' + \
            '          Chip: X |   |   |  \n' + \
            'Promethium RTG:   |   |   |  \n' + \
            '          Chip:   |   |   |  \n' + \
            '  Rubidium RTG:   |   |   |  \n' + \
            '          Chip:   |   |   |  \n' + \
            '   Thulium RTG:   |   |   |  \n' + \
            '          Chip:   |   |   |  \n' + \
            '\n' + \
            '      Elevator: E |   |   |  '

        self.assertEqual(expected, repr(starting_state))

        result = graph_search(starting_state)
        expected = 11
        self.assertEqual(expected, len(result[0]))

    def test_bitroom_generate_next_states(self):
        state = sut.BitRoom()

        next_states = state.generate_next_states()
        print(len(next_states))
        # print('\n ~~~ \n'.join([repr(x) for x in next_states]))
        print(repr(next_states[0]))
        print('\n\n\n')

        print('\n ~~~ \n'.join([repr(x) for x in next_states[0].generate_next_states()]))

    def test_bitroom_move_one_rtg(self):
        state = sut.BitRoom()

        rtgs = state.get_indices_of_items_on_floor(state.current_floor, state.RTG_INDICES)
        chips = state.get_indices_of_items_on_floor(state.current_floor, state.CHIP_INDICES)

        next_states = state.move_one_rtg(1, rtgs)
        next_states = [state.layout for state in next_states]

        self.assertListEqual(next_states, [])

    def test_bitroom_move_one_chip(self):
        state = sut.BitRoom()

        rtgs = state.get_indices_of_items_on_floor(state.current_floor, state.RTG_INDICES)
        chips = state.get_indices_of_items_on_floor(state.current_floor, state.CHIP_INDICES)

        next_states = state.move_one_chip(1, chips)
        next_states = [state.layout for state in next_states]

        self.assertListEqual(next_states, [
            # Co chip.
            [0] + [1, 0] + [1, 0] + [1, 0] + [1, 1] + [1, 1] + \
            [1] + [0, 1] + [0, 1] + [0, 1] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
            # Ru chip.
            [0] + [1, 1] + [1, 0] + [1, 0] + [1, 0] + [1, 1] + \
            [1] + [0, 0] + [0, 1] + [0, 1] + [0, 1] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
            # T chip.
            [0] + [1, 1] + [1, 0] + [1, 0] + [1, 1] + [1, 0] + \
            [1] + [0, 0] + [0, 1] + [0, 1] + [0, 0] + [0, 1] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
        ])

    def test_bitroom_move_two_rtgs(self):
        state = sut.BitRoom()

        rtgs = state.get_indices_of_items_on_floor(state.current_floor, state.RTG_INDICES)
        chips = state.get_indices_of_items_on_floor(state.current_floor, state.CHIP_INDICES)

        next_states = state.move_two_rtgs(1, rtgs)
        next_states = [state.layout for state in next_states]

        self.assertListEqual(next_states, [
            # Po and Pr rtgs.
            [0] + [1, 1] + [0, 0] + [0, 0] + [1, 1] + [1, 1] + \
            [1] + [0, 0] + [1, 1] + [1, 1] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
        ])

    def test_bitroom_move_two_chips(self):
        state = sut.BitRoom()

        rtgs = state.get_indices_of_items_on_floor(state.current_floor, state.RTG_INDICES)
        chips = state.get_indices_of_items_on_floor(state.current_floor, state.CHIP_INDICES)

        next_states = state.move_two_rtgs(1, chips)
        next_states = [state.layout for state in next_states]

        self.assertListEqual(next_states, [
            # Co and Ru chips.
            [0] + [1, 0] + [1, 0] + [1, 0] + [1, 0] + [1, 1] + \
            [1] + [0, 1] + [0, 1] + [0, 1] + [0, 1] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
            # Co and T chips.
            [0] + [1, 0] + [1, 0] + [1, 0] + [1, 1] + [1, 0] + \
            [1] + [0, 1] + [0, 1] + [0, 1] + [0, 0] + [0, 1] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
            # Ru and T chips.
            [0] + [1, 1] + [1, 0] + [1, 0] + [1, 0] + [1, 0] + \
            [1] + [0, 0] + [0, 1] + [0, 1] + [0, 1] + [0, 1] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + \
            [0] + [0, 0] + [0, 0] + [0, 0] + [0, 0] + [0, 0],
            #E    Co       Po       Pr       Ru       T
        ])

    def test_bitroom_move_one_rtg_and_one_chip(self):
        state = sut.BitRoom()

        rtgs = state.get_indices_of_items_on_floor(state.current_floor, state.RTG_INDICES)
        chips = state.get_indices_of_items_on_floor(state.current_floor, state.CHIP_INDICES)

        next_states = state.move_an_rtg_and_chip(1, rtgs, chips)
        next_states = [state.layout for state in next_states]

        self.assertListEqual(next_states, [])

    def test_part_one(self):
        path = sut.part_one(FULL_INPUT)[1]
        self.assertEqual(11, len(path))
