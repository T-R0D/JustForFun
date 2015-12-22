def main():
    sue_data = []
    with open('input.txt') as input_file:
        for line in input_file:
            parts = line.strip().replace(':', '').replace(',', '').split(' ')
            
            sue_dict = {'n': int(parts[1])}
            for i in range(2, len(parts), 2):
                sue_dict[parts[i]] = int(parts[i + 1])

            sue_data.append(sue_dict)

    MFCSAM_output = {
        'children': 3,
        'cats': 7,
        'samoyeds': 2,
        'pomeranians': 3,
        'akitas': 0,
        'vizslas': 0,
        'goldfish': 5,
        'trees': 3,
        'cars': 2,
        'perfumes': 1
    }

    part_1(sue_data, MFCSAM_output)
    part_2(sue_data, MFCSAM_output)


def part_1(sue_data, MFCSAM_output):
    remaining_sues = sue_data
    for key, value in MFCSAM_output.items():
        def keep(x):
            v = x.get(key, None)
            return v is None or v == value

        remaining_sues = list(filter(keep, remaining_sues))


    remaining_sue = remaining_sues[0]
    print("The number of the Aunt Sue that sent the gift is: {}".format(remaining_sue['n']))


def part_2(sue_data, MFCSAM_output):
    remaining_sues = sue_data
    for key, value in MFCSAM_output.items():
        if key in ('cats', 'trees'):
            def keep(x):
                v = x.get(key, None)
                return v is None or v > value
        elif key in ('pomeranians', 'goldfish'):
            def keep(x):
                v = x.get(key, None)
                return v is None or v < value
        else:
            def keep(x):
                v = x.get(key, None)
                return v is None or v == value

        remaining_sues = list(filter(keep, remaining_sues))


    remaining_sue = remaining_sues[0]
    print("The number of the Aunt Sue that sent the gift is: {}".format(remaining_sue['n']))


if __name__ == '__main__':
    main()
