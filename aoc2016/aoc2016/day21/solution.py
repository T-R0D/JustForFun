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


def part_one(puzzle_input):
    instructions = parse_instruction_list(puzzle_input)
    result = scramble_password(MY_PASSWORD, instructions)
    return result


def part_two(puzzle_input):
    instructions = parse_instruction_list(puzzle_input)
    result = unscramble_password(SCAMBLED_PASSWORD, instructions)
    return result


MY_PASSWORD = "abcdefgh"
SCAMBLED_PASSWORD = "fbgdceah"


def parse_instruction_list(puzzle_input):
    return [instruction_from_line(line) for line in puzzle_input.split("\n")]


def instruction_from_line(line):
    if line.startswith("swap position"):
        a, b = (
            int(x)
            for x in line.replace("swap position ", "")
            .replace("with position ", "")
            .split(" ")
        )
        return ("swap position", a, b)
    elif line.startswith("swap letter"):
        a, b = (
            x
            for x in line.replace("swap letter ", "")
            .replace("with letter ", "")
            .split(" ")
        )
        return ("swap letter", a, b)
    elif line.startswith("rotate based"):
        a = line.replace("rotate based on position of letter ", "")
        return ("rotate based", a, None)
    elif line.startswith("rotate"):
        a, b = (
            line.replace("rotate ", "")
            .replace(" steps", "")
            .replace(" step", "")
            .split(" ")
        )
        return ("rotate", a, int(b))
    elif line.startswith("reverse"):
        a, b = (
            int(x)
            for x in line.replace("reverse positions ", "")
            .replace("through ", "")
            .split(" ")
        )
        return ("reverse", a, b)
    elif line.startswith("move"):
        a, b = (
            int(x)
            for x in line.replace("move position ", "")
            .replace("to position ", "")
            .split(" ")
        )
        return ("move", a, b)
    else:
        raise ValueError(f"unable to parse line: {line}")


def swap_position(runes, a, b):
    runes[a], runes[b] = runes[b], runes[a]


def swap_letter(runes, a, b):
    i, j = 0, 0
    for x, c in enumerate(runes):
        if c == a:
            i = x
        elif c == b:
            j = x

    swap_position(runes, i, j)


def rotate(runes, direction, magnitude):
    n = len(runes)
    multiplier = 1 if direction == "right" else -1
    rotation = (n + (multiplier * magnitude)) % n
    cloned_runes = [r for r in runes]
    for i in range(len(runes)):
        target = (i + rotation) % n
        runes[target] = cloned_runes[i]


def rotate_based(runes, a, _b):
    magnitude = 1
    for i, x in enumerate(runes):
        if x == a:
            magnitude += i
            if i >= 4:
                magnitude += 1

            break

    rotate(runes, "right", magnitude)


def reverse(runes, a, b):
    for i in range((abs(b - a) // 2) + 1):
        runes[a + i], runes[b - i] = runes[b - i], runes[a + i]


def move(runes, a, b):
    if a < b:
        mover = runes[a]
        for i in range(a, b):
            runes[i] = runes[i + 1]
        runes[b] = mover
    else:
        mover = runes[a]
        for i in range(a, b, -1):
            runes[i] = runes[i - 1]
        runes[b] = mover


OPERATIONS = {
    "swap position": swap_position,
    "swap letter": swap_letter,
    "rotate": rotate,
    "rotate based": rotate_based,
    "reverse": reverse,
    "move": move,
}


def scramble_password(password, instructions):
    runes = list(password)

    for kind, a, b in instructions:
        OPERATIONS[kind](runes, a, b)

    return "".join(runes)


def unrotate(runes, direction, magnitude):
    rotate(runes, "right" if direction == "left" else "left", magnitude)


def unrotate_based(runes, a, _b):
    start = 0
    for i, x in enumerate(runes):
        if x == a:
            start = i
            break

    magnitude = 0
    if start & 1 == 0:
        magnitude = {2: -2, 4: -1, 6: 0, 0: 1}[start]
    else:
        magnitude = start // 2 + 1

    rotate(runes, "left", magnitude)


def unmove(runes, a, b):
    move(runes, b, a)


REVERSE_OPERATIONS = {
    "swap position": swap_position,
    "swap letter": swap_letter,
    "rotate": unrotate,
    "rotate based": unrotate_based,
    "reverse": reverse,
    "move": unmove,
}


def unscramble_password(scrambled_password, instructions):
    runes = list(scrambled_password)

    for kind, a, b in reversed(instructions):
        REVERSE_OPERATIONS[kind](runes, a, b)

    return "".join(runes)
