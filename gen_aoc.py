import os

def main():
    print('Generating AoC directory...')
    year = 2018

    root_dir = create_root_dir(year)

    os.chdir(root_dir)
    for day in range(1, 26):
        create_day_dir(root_dir, day)


def create_root_dir(year):
    cwd = os.getcwd()
    aoc_dir = 'aoc{}'.format(year)
    root_dir = os.path.join(cwd, aoc_dir)
    print('Creating root directory "{}"'.format(root_dir))
    os.makedirs(root_dir, exist_ok=True)
    return root_dir


def create_day_dir(root_dir, day):
    day_dir_name = 'day{0:02d}'.format(day)
    day_dir = os.path.join(root_dir, day_dir_name)
    print('Generating subdir {}'.format(day_dir))
    os.makedirs(day_dir, exist_ok=True)

    with open(os.path.join(day_dir, '__init__.py'), mode='w') as f:
        f.write('')

    with open(os.path.join(day_dir, 'solution.py'), mode='w') as f:
        f.write(
'''def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))

def parse_input():
    pass

def part_one(puzzle_input):
    pass

def part_two(puzzle_input):
    pass

if __name__ == '__main__':
    main()
''')

    with open(os.path.join(day_dir, 'problem_statement.txt'), mode='w') as f:
        f.write('')

    with open(os.path.join(day_dir, 'input.txt'), mode='w') as f:
        f.write('')

    create_test_dir(day_dir, day)

def create_test_dir(day_dir, day):
    test_dir = os.path.join(day_dir, 'test')
    os.makedirs(test_dir, exist_ok=True)

    with open(os.path.join(test_dir, '__init__.py'), mode='w') as f:
        f.write('')

    with open(os.path.join(test_dir, 'test_solution.py'), mode='w') as f:
        f.write(
'''import unittest
import day{0:02d}.solution as solution

class TestDay{0:02d}(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_part_one(self):
        pass

    def test_part_two(self):
        pass
''')

if __name__ == '__main__':
    main()
