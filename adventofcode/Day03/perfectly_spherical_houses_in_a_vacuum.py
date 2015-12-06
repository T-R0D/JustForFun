def main():
	path = ''
	with open('input.txt') as input_file:
		path = input_file.read()

	part_1(path)
	part_2(path)


def part_1(path):
	x = 0
	y = 0
	visited_houses = {coordinates_to_str(x, y)}

	for c in path:
		if c == '^':
			y += 1

		elif c == 'v':
			y -= 1

		elif c == '<':
			x -= 1

		elif c == '>':
			x += 1

		visited_houses.add(coordinates_to_str(x, y))

	print('{} unique houses get presents.'.format(len(visited_houses)))


def coordinates_to_str(x, y):
	return '{},{}'.format(x, y)


def part_2(path):
	real_x = 0
	real_y = 0

	robo_x = 0
	robo_y = 0

	visited_houses = {coordinates_to_str(real_x, real_y)}

	t = 0
	for c in path:
		if t % 2 == 0:
			real_x, real_y = get_updated_coordinates(real_x, real_y, c)
			visited_houses.add(coordinates_to_str(real_x, real_y))
		else:
			robo_x, robo_y = get_updated_coordinates(robo_x, robo_y, c)
			visited_houses.add(coordinates_to_str(robo_x, robo_y))

		t += 1

	print('With Robo-Santa, {} unique houses get presents.'.format(len(visited_houses)))


def get_updated_coordinates(x, y, instruction):
	if instruction == '^':
		return x, y + 1

	elif instruction == 'v':
		return x, y - 1

	elif instruction == '<':
		return x - 1, y

	elif instruction == '>':
		return x + 1, y


if __name__ == '__main__':
	main()
