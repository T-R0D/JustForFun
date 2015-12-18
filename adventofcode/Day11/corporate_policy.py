import re


class Password(object):
	def __init__(self, password):
		self.password = [ord(c) for c in password]


	def next(self):
		i = len(self.password) - 1

		carry = True
		while carry:
			self.password[i], carry = Password.__increment_char(self.password[i])

			if carry:
				i -= 1

			if i < 0:
				raise Exception('over-carry')

	def __str__(self):
		ret = ''
		for c in self.password:
			ret += chr(c)

		return ret

	@staticmethod
	def __increment_char(c):
		c += 1

		carry = False

		if c > ord('z'):
			c = ord('a')
			carry = True
		elif c < ord('a'):
			c = ord('a')
		elif chr(c) in {'i', 'o', 'l'}:
			c += 1

		return c, carry


def main():
	_input = None
	with open('input.txt') as input_file:
		_input = input_file.read().strip()

	part_1(_input)
	part_2(_input)


def part_1(_input):
	p = Password(_input.replace('i', 'j').replace('o', 'p').replace('l', 'k'))
	p.next()

	while not password_is_acceptable(str(p)):
		p.next()

	print("Santa's next password is {}".format(p))


def password_is_acceptable(password):
	# print(password)
	# print('contains straigt', contains_straight(password))
	# print('contains_pair', contains_pair(password))
	# print('contains_illegal_chars', contains_illegal_chars(password, '[iol]'))

	return contains_straight(password) and contains_pairs(password) and not contains_illegal_chars(password, '[iol]')


def contains_straight(password):
	if len(password) >= 3:

		for c in range(2, len(password)):
			i = password[c - 2]
			j = password[c - 1]
			k = password[c]

			if (ord(i) == (ord(j) - 1)) and (ord(j) == (ord(k) - 1)):
				return True

	return False


def contains_pairs(password):
	pair_indices = []

	for c in range(1, len(password)):
		if password[c - 1] == password[c]:
			if pair_indices:
				if c - 1 > pair_indices[-1]:
					pair_indices.append(c)
			else:
				pair_indices.append(c)

	return len(pair_indices) >= 2


def contains_illegal_chars(password, illegals):
	return re.search(illegals, password)


def part_2(_input):
	p = Password(_input.replace('i', 'j').replace('o', 'p').replace('l', 'k'))

	for _ in range(0, 2):
		p.next()
		while not password_is_acceptable(str(p)):
			p.next()

	print("Santa's next, next password is {}".format(p))	


if __name__ == '__main__':
	main()
