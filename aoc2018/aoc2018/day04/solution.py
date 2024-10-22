import collections
import datetime


def main():
    puzzle_input = parse_input()
    print(part_one(puzzle_input))
    print(part_two(puzzle_input))


def parse_input():
    logs = []
    with open('input.txt') as f:
        for line in f:
            timestamp_str, activity = line.split('] ')
            timestamp_str = timestamp_str[1:]
            logs.append((timestamp_str, activity))

    def get_timestamp(log):
        return log[0]
    logs = list(sorted(logs, key=get_timestamp))

    return logs


def part_one(puzzle_input):
    minutes_slept_by_guard = collections.Counter()
    minutes_slept_table = {}
    current_guard = -1
    sleep_time = None
    for log in puzzle_input:
        timestamp = datetime.datetime.strptime(log[0], '%Y-%m-%d %H:%M')
        event = log[1]
        if 'begins shift' in event:
            current_guard = extract_guard_id(event)
        elif 'falls asleep' in event:
            sleep_time = timestamp
        else:
            minutes_slept = (timestamp - sleep_time).total_seconds() // 60
            minutes_slept_by_guard[current_guard] += minutes_slept

            histogram = minutes_slept_table.get(current_guard, collections.Counter())
            if sleep_time.minute >= 0:
                for i in range(sleep_time.minute, timestamp.minute):
                    histogram[i] += 1
            else:
                if timestamp.minute > sleep_time.minute:
                    for i in range(sleep_time.minute, timestamp.minute):
                        histogram[i] += 1
                else:
                    for i in range(sleep_time.minute, 0):
                        histogram[i] += 1
                    for i in range(0, timestamp.minute):
                        histogram[i] += 1
            minutes_slept_table[current_guard] = histogram

    sleepiest_guard = minutes_slept_by_guard.most_common(1)[0][0]
    histogram = minutes_slept_table[sleepiest_guard]
    sleepiest_minute = histogram.most_common(1)[0][0]

    return str(sleepiest_guard * sleepiest_minute)

def part_two(puzzle_input):
    minutes_slept_table = {}
    current_guard = -1
    sleep_time = None
    for log in puzzle_input:
        timestamp = datetime.datetime.strptime(log[0], '%Y-%m-%d %H:%M')
        event = log[1]
        if 'begins shift' in event:
            current_guard = extract_guard_id(event)
        elif 'falls asleep' in event:
            sleep_time = timestamp
        else:
            histogram = minutes_slept_table.get(current_guard, collections.Counter())
            if sleep_time.minute >= 0:
                for i in range(sleep_time.minute, timestamp.minute):
                    histogram[i] += 1
            else:
                if timestamp.minute > sleep_time.minute:
                    for i in range(sleep_time.minute, timestamp.minute):
                        histogram[i] += 1
                else:
                    for i in range(sleep_time.minute, 0):
                        histogram[i] += 1
                    for i in range(0, timestamp.minute):
                        histogram[i] += 1
            minutes_slept_table[current_guard] = histogram

    regularly_asleep_guard = -1
    global_sleepiest_minute = -1
    sleepy_minute_frequency = -1
    for guard_id, histogram in minutes_slept_table.items():
        sleepiest_minute, frequency = histogram.most_common(1)[0]
        if frequency > sleepy_minute_frequency:
            global_sleepiest_minute = sleepiest_minute
            regularly_asleep_guard = guard_id
            sleepy_minute_frequency = frequency

    return str(regularly_asleep_guard * global_sleepiest_minute)


def extract_guard_id(shift_start_str):
    number_str = shift_start_str.split(' ')[1]
    return int(number_str.replace('#', ''))


if __name__ == '__main__':
    main()
