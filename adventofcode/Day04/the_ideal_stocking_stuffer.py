import hashlib


def main():
	input = ''
	with open('input.txt') as input_file:
		_input = input_file.read().strip()

	part_1(_input)
	part_2(_input)


def part_1(_input):
	_hash = ''
	i = 0
	while _hash[:5] != '00000':
		i += 1
		string = '{}{}'.format(_input, i).encode()
		_hash = hashlib.md5(string).hexdigest()

	print('The hash key is {}'.format(i))


def part_2(_input):
	_hash = ''
	i = 0
	while _hash[:6] != '000000':
		i += 1
		string = '{}{}'.format(_input, i).encode()
		_hash = hashlib.md5(string).hexdigest()

	print('The hash key is {}'.format(i))


if __name__ == '__main__':
	main()
