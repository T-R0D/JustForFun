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
from collections import OrderedDict


def part_one(puzzle_input):
    chips, bots = create_processing_graph(puzzle_input)
    paths, outputs = construct_chip_paths(chips, bots)
    return str(find_common_bot(paths, 61, 17))

def part_two(puzzle_input):
    chips, bots = create_processing_graph(puzzle_input)
    paths, outputs = construct_chip_paths(chips, bots)
    return str(find_output_product(outputs, 0, 1, 2))


def create_processing_graph(puzzle_input):
    chips = {}
    bots = {}
    for line in puzzle_input.split('\n'):
        parts = line.split()
        if len(parts) < 7:
            chips[int(parts[1])] = [(parts[-2], int(parts[-1]))]
        else:
            bots[int(parts[1])] = \
                ((parts[5], int(parts[6])), (parts[10], int(parts[11])))
    return chips, bots


def construct_chip_paths(chips, bots):
    jobs = OrderedDict()
    outputs = {}
    for chip in chips.keys():
        obj, number = chips[chip][0]
        if obj == 'output':
            outputs[number] = outputs.get(number, []) + [chip]
            continue
        jobs[number] = jobs.get(number, []) + [chip]
    while jobs:
        job = jobs.popitem(last=False)
        bot, items = job

        if len(items) < 2:
            jobs[bot] = items
            continue

        chip0, chip1 = sorted(items)
        low, high = bots[bot]
        chips[chip0].append(low)
        chips[chip1].append(high)

        if low[0] == 'bot':
            bot = low[1]
            jobs[bot] = jobs.get(bot, []) + [chip0]
        elif low[0] == 'output':
            output = low[1]
            outputs[output] = outputs.get(output, []) + [chip0]

        if high[0] == 'bot':
            bot = high[1]
            jobs[bot] = jobs.get(bot, []) + [chip1]
        elif high[0] == 'output':
            output = high[1]
            outputs[output] = outputs.get(output, []) + [chip1]

    return chips, outputs


def find_common_bot(chips, chip0, chip1):
    path0 = set(chips[chip0])
    path1 = set(chips[chip1])
    return set.intersection(path0, path1).pop()


def find_output_product(outputs, *args):
    product = 1
    for number in args:
        output = outputs.get(number, [])
        for chip in output:
            product *= chip
    return product
