class Action(object):
	def __init__(self, text_line):
		self.action = None
		self.value = None
		self.lhs = None
		self.rhs = None

		lhs, rhs = text_line.split('->')

		self.output = rhs.strip()
		lhs = lhs.strip().split(' ')
		if len(lhs) == 1:
			self.action = 'load'
			try:
				self.value = int(lhs[0])
			except:
				self.rhs = lhs[0]

		elif len(lhs) == 2:
			self.action = lhs[0]
			self.rhs = lhs[1]

		else:
			self.lhs = lhs[0]
			self.action = lhs[1]
			self.rhs = lhs[2]

		self.text_line = text_line.strip()

		self.dependencies = []  # for tree building

	def __str__(self):
		return self.text_line

	def __repr__(self):
		return str(self)

SHORT_MAX = 0xFFFF

def main():
	_input = None
	with open('input.txt') as input_file:
		_input = [Action(line) for line in input_file]

	part_1(_input)
	part_2(_input)


def part_1(_input):
	action_dict = {}
	for action in _input:
		action_dict[action.output] = action

	wires = {}

	evaluate_action(action_dict['a'], wires, action_dict)

	print('The value of wire "a" is {}'.format(wires['a']))


def evaluate_action(node, wires, action_dict):
	left = None
	right = None

	if node.lhs:
		if represents_int(node.lhs):
			left = int(node.lhs)
		else:
			if not node.lhs in wires:
				evaluate_action(action_dict[node.lhs], wires, action_dict)

			left = wires[node.lhs]

	if node.rhs:
		if represents_int(node.rhs):
			right = int(node.rhs)
		else:
			if not node.rhs in wires:
				evaluate_action(action_dict[node.rhs], wires, action_dict)

			right = wires[node.rhs]

	if node.action == 'load':
		if node.value is not None:
			wires[node.output] = node.value

		else:
			wires[node.output] = right

	elif node.action == 'NOT':
		wires[node.output] = (~right) & SHORT_MAX

	elif node.action == 'AND':
		wires[node.output] = (left & right) & SHORT_MAX
	
	elif node.action == 'OR':
		wires[node.output] = (left | right) & SHORT_MAX
	
	elif node.action == 'LSHIFT':
		wires[node.output] = (left << right) & SHORT_MAX
	
	elif node.action == 'RSHIFT':
		wires[node.output] = (left >> right) & SHORT_MAX


def represents_int(s):
	try:
		x = int(s)
		return True
	except:
		return False


def part_2(_input):
	action_dict = {}
	for action in _input:
		action_dict[action.output] = action

	wires = {}
	evaluate_action(action_dict['a'], wires, action_dict)

	action_dict['b'] = Action('{} -> b'.format(wires['a']))
	wires = {}
	evaluate_action(action_dict['a'], wires, action_dict)
	print('The new value of wire "a" is {}'.format(wires['a']))


if __name__ == '__main__':
	main()