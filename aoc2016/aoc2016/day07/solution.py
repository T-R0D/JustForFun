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
    n_tls_capable = 0
    for address in puzzle_input.split('\n'):
        if supports_tls(address):
            n_tls_capable += 1
    return str(n_tls_capable)


def part_two(puzzle_input):
    n_ssl_capable = 0
    for address in puzzle_input.split('\n'):
        if supports_ssl(address):
            n_ssl_capable += 1
    return str(n_ssl_capable)


def get_supernet_and_hypernet_sequences(address):
    sequences = address.replace(']','[').split('[')
    return sequences[::2], sequences[1::2]

def supports_tls(address):
    supernet_sequences, hypernet_sequences = get_supernet_and_hypernet_sequences(address)
    return any(map(contains_abba, supernet_sequences)) and not any(map(contains_abba, hypernet_sequences))


def contains_abba(sequence):
    if len(sequence) >= 4:
        for a, b, c, d in zip(sequence, sequence[1:], sequence[2:], sequence[3:]):
            if a != b and a == d and b == c:
                return True
    return False


def supports_ssl(address):
    supernet_sequences, hypernet_sequences = get_supernet_and_hypernet_sequences(address)
    aba_list = set()
    for sequence in supernet_sequences:
        aba_list.update(find_all_aba(sequence))

    for sequence in hypernet_sequences:
        if contains_corresponding_bab(sequence, aba_list):
            return True

    return False


def find_all_aba(supernet_sequence):
    aba_list = set()
    if len(supernet_sequence) >= 3:
        for a, b, c in zip(supernet_sequence, supernet_sequence[1:], supernet_sequence[2:]):
            if a != b and a == c:
                aba_list.add((a, b))
    return aba_list


def contains_corresponding_bab(hypernet_sequence, aba_list):
    if len(hypernet_sequence) >= 3:
        for aba in aba_list:
            a, b = aba
            for x, y, z in zip(hypernet_sequence, hypernet_sequence[1:], hypernet_sequence[2:]):
                if x == b and y == a and x == z:
                    return True
    return False
