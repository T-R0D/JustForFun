

def main():
	part_1()
	part_2()


def part_1():
	with open('input.txt') as i_file:
		total_area = 0
		for line in i_file:
			dimensions = line.split('x')
			length, width, height = map(lambda x: int(x), dimensions)
			area = surface_area_of_rectangle_prism(length, width, height)
			area += area_of_smallest_side(length, width, height)
			total_area += area

		print('The elves need {} square feet of wrapping paper.'.format(total_area))

def surface_area_of_rectangle_prism(length, width, height):
	return 2 * (length * width + length * height + width  *	height)


def area_of_smallest_side(length, width, height):
	return min(length * width, length * height, width * height)


def part_2():
	with open('input.txt') as i_file:
		total_length = 0
		for line in i_file:
			dimensions = line.split('x')
			length, width, height = map(lambda x: int(x), dimensions)
			new_length = perimeter_of_smallest_face(length, width, height)
			new_length += volume_of_present(length, width, height)
			total_length += new_length

		print('The elves need {} feet of ribbon.'.format(total_length))


def perimeter_of_smallest_face(length, width, height):
	return min(2 * length + 2 * width, 
		       2 * length + 2 * height, 
		       2 * width + 2 * height)


def volume_of_present(length, width, height):
	return length * width * height


if __name__ == '__main__':
	main()
