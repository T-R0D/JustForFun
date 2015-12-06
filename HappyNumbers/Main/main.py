import matplotlib.pyplot as pyplot


class HappinessResult(object):
    def __init__(self, number, is_happy=False, iterations_to_happiness=-1, iterations_to_single_digit=-1,
                 iterations_to_repeat=-1):
        self.number = number
        self.is_happy = is_happy
        self.iterations_to_happiness = iterations_to_happiness
        self.iterations_to_single_digit = iterations_to_single_digit
        self.iterations_to_repeat = iterations_to_repeat


def main():
    """ This is a test to find out some properties of (un)happy numbers.

    I started mostly to find out (1) if unhappy numbers were more likely to repeat or go to a single digit number first,
    and (2) find out if there could be infinite looping.

    (1) It turns out that the single digit number appears before any repeats in the summing part of determining a
        happy number. This makes sense since once at a single digit, then we see the repeats; also the sum of a numbers
        digits __must__ be less than the original number (unless I've made a bad oversight)

    (2) Again, since the sum of a number's digits must be less than the number itself, I doubt we would see any kinds
        of long/infinite sequences. In fact, in the range of [1, 999999), there were no numbers that took more than
        4 iterations in the happiness finding process. Maybe prove this more formally?

    (3) I probably could have done this with pure math, but it was good to bang around with some python for something
        fun and not required.
    """

    results = [happy_test(n) for n in range(1, 1000000)]

    happies = [result.number for result in results if result.is_happy]

    print([result.number for result in results if (result.is_happy and result.iterations_to_happiness >= 4)])

    print(len(happies))
    single_digit_times = [result.iterations_to_happiness for result in results if result.is_happy]

    pyplot.plot(happies, single_digit_times, 'r+', label='iterations to happiness')
    sdmax = max(single_digit_times)
    pyplot.axis([1, max(happies) + 1, 0, sdmax + 1])
    pyplot.show()

    print('num happy: {}'.format(len(happies)))


    # unhappy_numbers = [result.number for result in results]
    # single_digit_times = [result.iterations_to_single_digit for result in results]
    # repeat_times = [result.iterations_to_repeat for result in results]
    #
    # cantankerous_numbers = [x for x in results if x.iterations_to_repeat < x.iterations_to_single_digit]
    # print(len(cantankerous_numbers))
    #
    # pyplot.plot(unhappy_numbers, single_digit_times, 'r,', label='single digit time')
    # pyplot.plot(unhappy_numbers, repeat_times, 'b,', label='repeat time')
    # sdmax = max(single_digit_times)
    # rmax = max(repeat_times)
    # pyplot.axis([1, max(unhappy_numbers) + 1, 0, max(sdmax, rmax) + 1])

    # pyplot.plot(single_digit_times, label = 'iterations to reach single digit')
    # pyplot.plot(repeat_times, label = 'iterations to reach a repeated term')
    # pyplot.show()


def happy_test(number):
    og = number
    observed_numbers = []
    is_single_digit = number < 10
    is_repeat = False
    i = 0
    single_digit_iterations = 0
    repeat_iterations = 0
    while not (is_single_digit and is_repeat):
        i += 1
        observed_numbers.append(number)
        number = sum_digits(number)

        if not is_single_digit and number < 10:
            is_single_digit = True
            single_digit_iterations = i
        if not is_repeat and number in observed_numbers:
            is_repeat = True
            repeat_iterations = i

    if number == 1:
        return HappinessResult(number=og, is_happy=True, iterations_to_happiness=i)
    else:
        return HappinessResult(number=og, is_happy=False, iterations_to_single_digit=single_digit_iterations,
                               iterations_to_repeat=repeat_iterations)


def sum_digits(number):
    num_string = '{}'.format(number)
    sum = 0

    for letter in num_string:
        sum += int(letter)

    return sum


if __name__ == '__main__':
    main()