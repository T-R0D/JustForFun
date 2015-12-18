def main():
	_input = None
	with open('input.txt') as input_file:
		_input = [line.strip() for line in input_file]

	part_1(_input)
	part_2(_input)

def part_1(_input):
	total = 0
	for line in _input:
		total += len(line) - eval('len({})'.format(line))

	print('The total difference is: {}'.format(total))

def part_2(_input):
	total = 0
	for line in _input:
		total += get_re_encoded_length(line) - len(line)

	print('The total difference with the re-encoded strings is {}'.format(total))

def get_re_encoded_length(s):
	double_quotes = s.count('"')
	slashes = s.count('\\')
	escaped_quotes = s.count('\\"')
	ends_with_slash_quote = 1 if s.endswith('\\"') else 0
	return len(s) + slashes + (2 * double_quotes) - escaped_quotes + ends_with_slash_quote


if __name__ == '__main__':
	main()
