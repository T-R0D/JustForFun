import copy


INF = 99999

def main():
	_input = None
	with open('input.txt') as input_file:
		_input = [line.strip() for line in input_file]

	part_1(_input)
	part_2(_input)


class SearchState(object):
	def __init__(self):
		self.path = []
		self.distance = 0

	def __repr__(self):
		return str(self.distance) + ' ' + ' -> '.join(self.path)


def part_1(_input):
	options = {}
	for line in _input:
		src, _, dest, _, dist = line.split(' ')	

		if src not in options:
			options[src] = []
		if dest not in options:
			options[dest] = []

		options[src].append((dest, int(dist)))
		options[dest].append((src, int(dist)))

	the_path = SearchState()
	the_path.distance = 9999999999999999999

	for src in options.keys():
		state = SearchState()
		state.path.append(src)

		stack = [state]

		while stack:
			state = stack.pop()

			if len(state.path) == len(options.keys()) and state.distance < the_path.distance:
				the_path = state

			src = state.path[-1]
			if src in options:
				for dest, dist in options[src]:
					if dest not in state.path:
						new_state = copy.deepcopy(state)
						new_state.path.append(dest)
						new_state.distance += dist
						stack.append(new_state)

	print('The shortest path to visit all of the locations is:\n{}'.format(the_path))


def part_2(_input):
	options = {}
	for line in _input:
		src, _, dest, _, dist = line.split(' ')	

		if src not in options:
			options[src] = []
		if dest not in options:
			options[dest] = []

		options[src].append((dest, int(dist)))
		options[dest].append((src, int(dist)))

	the_path = SearchState()
	the_path.distance = 0

	for src in options.keys():
		state = SearchState()
		state.path.append(src)

		stack = [state]

		while stack:
			state = stack.pop()

			if len(state.path) == len(options.keys()) and state.distance > the_path.distance:
				the_path = state

			src = state.path[-1]
			if src in options:
				for dest, dist in options[src]:
					if dest not in state.path:
						new_state = copy.deepcopy(state)
						new_state.path.append(dest)
						new_state.distance += dist
						stack.append(new_state)

	print('The longest path to visit all of the locations is:\n{}'.format(the_path))


def other(_input):
	location_names = set()

	for line in _input:
		l = line.split(' ')
		location_names.add(l[0])
		location_names.add(l[2])

	n = len(location_names)
	indices = {}
	names = {}
	for i, name in enumerate(location_names):
		indices[i] = name
		names[name] = i

	print(indices)

	graph = [[INF] * n for _ in range(0, n)]
	nexts = [[-1] * n for _ in range(0, n)]

	for row in graph:
		print(row)
	for row in nexts:
		print(row)
	print()

	for line in _input:
		l = line.split(' ')
		i = names[l[0]]
		j = names[l[2]]
		v = int(l[4])
		graph[i][j] = v
		nexts[i][j] = j


	for row in graph:
		print(row)
	for row in nexts:
		print(row)
	print()

	for k in range(0, n):
		for i in range(0, n):
			for j in range(0, n):
				new_path = graph[i][k] + graph[k][j]
				if new_path < graph[i][j]:
					graph[i][j] = new_path
					nexts[i][j] = nexts[i][k]


	paths = []
	for i in range(0, n):
		for j in range(0, n):
			paths.append(reconstruct_path(i, j, nexts, graph))

	def path_len(x):
		return len(x[0])

	paths = list(sorted(paths, key=path_len))

	for path in paths:
		print(path)

	total = 0

	print('The distance of the shortest route is: {}'.format(total))


def reconstruct_path(i, j, nexts, graph):
	original = (i, j)
	if nexts[i][j] == -1:
		return ([], INF, original)

	path = [i]
	dist = graph[i][j]
	while i != j:
		i = nexts[i][j]
		path.append(i)

	return (path, dist, original)


if __name__ == '__main__':
	main()
