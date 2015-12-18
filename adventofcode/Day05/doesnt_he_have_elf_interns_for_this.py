import re

def main():
	with open('input.txt') as input_file:
		part_1(input_file)
	with open('input.txt') as input_file:
		part_2(input_file)


def part_1(input_file):
	nice  = 0

	for line in input_file:
		if string_is_nice(line):
			nice += 1

	print("There are {} nice words in the list.".format(nice))


def string_is_nice(string):
	previous_c = ''
	contains_consecutives = False
	vowels = {'a', 'e', 'i', 'o', 'u'}
	vowel_count = 0
	naughty_substrings = {'ab', 'cd', 'pq', 'xy'}
	contains_no_naughty_substring = True

	for c in string:
		pair = previous_c + c
		if pair in naughty_substrings:
			contains_no_naughty_substring = False
			break

		if c == previous_c:
			contains_consecutives = True

		if c in vowels:
			vowel_count += 1

		previous_c = c

	return contains_consecutives and vowel_count >= 3 and contains_no_naughty_substring

def part_2(input_file):
	nice  = 0

	for line in input_file:
		if string_is_nice_2(line.strip()):
			nice += 1

	print("There are {} nice words (by the new definition) in the list.".format(nice))


def string_is_nice_2(string):
	c_0 = ''
	c_1 = ''
	has_pair_triplet = False

	for c in string:
		if c_0 == c:
			has_pair_triplet = True
			break

		c_0 = c_1
		c_1 = c

	has_matching_pairs = False
	for i in range(2, len(string) + 1 - 2):
		for j in range(i + 2, len(string) + 1):
			if string[i-2:i] == string[j-2:j]:
				has_matching_pairs = True
				break

	return has_pair_triplet and has_matching_pairs



if __name__ == '__main__':
	# print(string_is_nice_2('qjhvhtzxzqqjkmpb'))
	# print(string_is_nice_2('xbhjakklmbhsdmdt'))
	# print(string_is_nice_2('xxyxx'))
	# print(string_is_nice_2('uurcxstgmygtbstg'))
	# print(string_is_nice_2('ieodomkazucvgmuy'))
	main()
