import json


class NumberTally(object):
	def __init__(self):
		self.tally = 0

	def add(self, number_str):
		self.tally += int(number_str)


def main():
	with open('input.txt') as f:
		_input = f.read()

	part_1(_input)
	part_2(_input)


def part_1(_input):
	number_tally = NumberTally()

	decoder = json.JSONDecoder(parse_int=number_tally.add)

	decoder.decode(_input)

	print("The sum of all of the numbers in the document is {}".format(number_tally.tally))

def part_2(_input):
	document = json.loads(_input)

	total = get_sum_of_json_without_red(document)

	print("The sum of all the numbers without 'reds' is {}".format(total))


def get_sum_of_json_without_red(root_item):
	total = 0

	if isinstance(root_item, dict):

		for item in root_item.values():
			if item == "red":
				return 0

			total += get_sum_of_json_without_red(item)

	elif isinstance(root_item, list):
		for item in root_item:
			total += get_sum_of_json_without_red(item)

	elif isinstance(root_item, (int, float)):
		total = root_item

	# print(total)

	return total


if __name__ == '__main__':
	main()
