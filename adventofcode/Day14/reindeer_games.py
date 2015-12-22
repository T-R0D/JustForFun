
class Reindeer(object):
    def __init__(self, name, speed, work_time, rest_time):
        self.name = name
        self.speed = int(speed)
        self.work_time = int(work_time)
        self.rest_time = int(rest_time)
        self.cycle_time = self.work_time + self.rest_time

        self.state = 'flying'
        self.seconds_in_state = 0
        self.distance_flown = 0
        self.score = 0

    def get_travel_distance(self, time):
        full_cycles = time // self.cycle_time
        partial_cycle_time = time % self.cycle_time

        return ((full_cycles * self.work_time) + min(partial_cycle_time, self.work_time)) * self.speed

    def fly_for_one_second(self):
        if self.seconds_in_state % self.cycle_time < self.work_time:
            self.distance_flown += self.speed

        self.seconds_in_state += 1

        return self.distance_flown


def main():
    reindeer = []
    with open('input.txt') as input_file:
        for line in input_file:
            parts = line.strip().split(' ')

            reindeer.append(Reindeer(parts[0], parts[3], parts[6], parts[-2]))

    part_1(reindeer)
    part_2(reindeer)


def part_1(_input):
    best = 0, 'dum-dum'

    for reindeer in _input:
        dist = reindeer.get_travel_distance(2503)
        if dist > best[0]:
            best = (dist, reindeer.name)

    print("{1} traveled the farthest with a distance of {0} km".format(*best))


def part_2(_input):
    best = 0, _input[0]

    for i in range(0, 2503):


        for reindeer in _input:
            dist = reindeer.fly_for_one_second()
            if dist > best[0]:
                best = (dist, reindeer)

        best[1].score += 1

    best = 0, _input[0]
    for reindeer in _input:
        if reindeer.score > best[0]:
            best = (reindeer.score, reindeer)

    print("{} scored the best with {} points and flew {} km".format(best[1].name, best[1].score, best[1].distance_flown))


if __name__ == '__main__':
    main()
