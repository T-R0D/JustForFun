def main():
	part_1()
	part_2()


def part_1():
	with open('input.txt') as i_file:
		instructions = i_file.read()
		ups = instructions.count('(')
		downs = instructions.count(')')

		print('Go to floor {}'.format(ups - downs))

def part_2():
	instructions = ''
	with open('input.txt') as i_file:
		instructions = list(i_file.read())

	floor = 0
	i = 0
	for instruction in instructions:
		if instruction == '(':
			floor += 1
		else:
			floor -= 1

		if floor < 0:
			print('The basement is entered on the {}th instruction'.format(i + 1))
			# i + 1 because the answer is expecting 1 based indexing
			break

		i += 1

	if floor >= 0:
		print('The basement is never entered.')

	

if __name__ == '__main__':
	main()
