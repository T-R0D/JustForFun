def main():
	_input = None
	with open('input.txt') as input_file:
		_input = input_file.read().strip()

	part_1(_input)
	part_2(_input)


def part_1(_input):
	sequence = _input
	for _ in range(0, 40):
		sequence = look_and_say(sequence)

	print('After 40 iterations, the length of the input is {}'.format(len(sequence)))

def look_and_say(s):
	number = 'not a number yet'
	count = 0

	looks = []
	for i in s:
		if i == number:
			count += 1
		else:
			looks.append((count, number))
			number = i
			count = 1
	looks.append((count, number))

	say = ''
	for look in looks[1: ]:
		say += str(look[0]) + str(look[1])

	return say


def part_2(_input):
	sequence = _input
	for _ in range(0, 50):
		sequence = look_and_say(sequence)

	print('After 50 iterations, the length of the input is {}'.format(len(sequence)))


if __name__ == '__main__':
	main()
