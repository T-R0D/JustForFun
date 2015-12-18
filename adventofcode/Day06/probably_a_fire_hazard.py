class Action(object):
	def __init__(self, text_line):
		parts = text_line.split(' ')

		self.lower_x, self.lower_y = parts[-3].split(',')
		self.lower_x = int(self.lower_x)
		self.lower_y = int(self.lower_y)
		self.upper_x, self.upper_y = parts[-1].split(',')
		self.upper_x = int(self.upper_x)
		self.upper_y = int(self.upper_y)
		self.action = parts[-4]

	def __str__(self):
		return '{} lights ({}, {}) - ({}, {})'.format(self.action, self.lower_x, self.lower_y, self.upper_x, self.upper_y)

	def __repr__(self):
		return str(self)


def main():
	_input = None

	with open('input.txt') as input_file:
		_input = [Action(line) for line in input_file]
	
	part_1(_input)
	part_2(_input)


def part_1(_input):
	lights = bytearray(1000 * 1000)

	for action in _input:
		for i in range(action.lower_x, action.upper_x + 1):
			for j in range(action.lower_y, action.upper_y + 1):
				if action.action == 'on':
					lights[i * 1000 + j] |= 0x01

				elif action.action == 'off':
					lights[i * 1000 + j] &= 0x00

				else:
					lights[i * 1000 + j] ^= 0x01

	lights_on = 0
	for i in range(0, 1000):
		for j in range(0, 1000):
			lights_on += lights[i * 1000 + j]

	print('{} lights are on at the end of the sequence.'.format(lights_on))


def part_2(_input):
	lights = bytearray(1000 * 1000)

	for action in _input:
		for i in range(action.lower_x, action.upper_x + 1):
			for j in range(action.lower_y, action.upper_y + 1):
				if action.action == 'on':
					lights[i * 1000 + j] += 1

				elif action.action == 'off':
					lights[i * 1000 + j] = max(lights[i * 1000 + j] - 1, 0)

				else:
					lights[i * 1000 + j] += 2

	total_brightness = 0
	for i in range(0, 1000):
		for j in range(0, 1000):
			total_brightness += lights[i * 1000 + j]

	print('The total brightness of the lights is {}'.format(total_brightness))


if __name__ == '__main__':
	main()
