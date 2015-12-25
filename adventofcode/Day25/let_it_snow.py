def main():
    row = 0
    col = 0
    with open('input.txt') as input_file:
        f = input_file.read().strip()
        parts = f.split(' ')
        row = int(parts[-3].replace(',', ''))
        col = int(parts[-1].replace('.', ''))

    # print(row, col)
    # print(next_code(20151125))

# 20151125, 18749137, 17289845, 30943339, 10071777, 33511524,
# 31916031, 21629792, 16929656 , 7726640, 15514188,  4041754,
# 16080970,  8057251,  1601130,  7981243, 11661866, 16474243,
# 24592653, 32451966, 21345942,  9380097, 10600672, 31527494,
# 77061, 17552253, 28094349,  6899651,  9250759, 31663883,
# 33071741,  6796745, 25397450, 24659492,  1534922, 27995004,

    # for i in range(1, 7):
    #     complete_rows = i
    #     numbers_in_diagonals = (1 + complete_rows) * (complete_rows // 2)
    #     if complete_rows % 2 == 1:
    #         numbers_in_diagonals += (complete_rows // 2) + 1
    #     print(numbers_in_diagonals)

    print(get_n(1, 1))
    print(get_n(2, 2))
    print(get_n(3, 3))
    print(get_n(4, 4))
    print(get_n(3, 1))
    print(get_n(3, 2))
    print(get_n(3, 4))


    part_1(None, row, col)
    part_2(None, 0, 0)


def part_1(table_fragment, row, col):
    n = get_n(row, col)

    code = 20151125
    for i in range(2, n + 1):
        code = next_code(code)

    print("The activation code is {}.".format(code))


def get_n(row, col):
    complete_rows = row + col - 2
    numbers_in_diagonals = (1 + complete_rows) * (complete_rows // 2)
    if complete_rows % 2 == 1:
        numbers_in_diagonals += (complete_rows // 2) + 1

    return numbers_in_diagonals + col


def next_code(previous_code):
    return (252533 * previous_code) % 33554393


def part_2(table_fragment, row, col):
    pass


if __name__ == '__main__':
    main()
