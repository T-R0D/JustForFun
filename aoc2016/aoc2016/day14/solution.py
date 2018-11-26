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

import collections
import hashlib


def part_one(puzzle_input):
    generator = OTPGenerator(salt='qzyelonm')
    otps = generator.get_otps(64)
    return str(otps[-1])


def part_two(puzzle_input):
    def stretched_md5(input):
        h = builtin_md5(input)
        for _ in range(2016):
            h = builtin_md5(h)
        return h
    generator = OTPGenerator(salt='qzyelonm', hash_fn=stretched_md5)
    otps = generator.get_otps(64)
    return str(otps[-1])


def builtin_md5(string):
    h = hashlib.md5()
    h.update(string.encode())
    return h.digest().hex()


class OTPGenerator(object):
    def __init__(self, salt, hash_fn=builtin_md5):
        self.salt = salt
        self.md5 = hash_fn
        self.runs_of_three = collections.deque()
        self.memoized_hashes = {}

    def get_otps(self, n_needed=64, proximity=1000):
        otp_indices = []
        ndx = 0
        self.runs_of_three = collections.deque()
        while len(otp_indices) < n_needed:
            string = self.salt + str(ndx)
            hash_val = self.memoized_hashes.get(string, self.md5(string))
            self.memoized_hashes[string] = hash_val
            repeated_char = self.find_runs(hash_val, 3)
            if repeated_char:
                for i in range(ndx + 1, ndx + proximity + 1):
                    hash_val2 = self.md5(self.salt + str(i))
                    char2 = self.find_runs(hash_val2, 5)
                    if repeated_char == char2:
                        otp_indices.append(ndx)
                        break
            ndx += 1

        return otp_indices

    @staticmethod
    def find_runs(string, run_len):
        end = len(string) - run_len + 1
        for i in range(end):
            char = string[i]
            if all(map(lambda x: x == char, list(string[i:i + run_len]))):
                return char
        return None

    def find_matching_candidate(self, char, index, proximity):
        pass
