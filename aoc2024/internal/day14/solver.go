package day14

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	return configurablePartOne(input, [2]int{spaceM, spaceN})
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	robots, err := parseInitialRobotSummaries(input)
	if err != nil {
		return "", err
	}

	mapSizes := [2]int{spaceM, spaceN}
	t := simulateRobotMovementUntilEasterEggFound(robots, mapSizes)

	return strconv.Itoa(t), nil
}

const (
	simulatedTimeSeconds = 100

	spaceM = 103
	spaceN = 101
)

func configurablePartOne(input string, mapSizes [2]int) (string, error) {
	robots, err := parseInitialRobotSummaries(input)
	if err != nil {
		return "", err
	}

	finalRobots := simulateRobotMovement(robots, mapSizes, simulatedTimeSeconds)
	quadrantCounts := countRobotsInQuadrants(finalRobots, mapSizes)
	safetyFactor := findSafetyFactor(quadrantCounts)

	return strconv.Itoa(safetyFactor), nil
}

type robotSummary struct {
	location [2]int
	velocity [2]int
}

func parseInitialRobotSummaries(input string) ([]robotSummary, error) {
	lines := strings.Split(input, "\n")
	robots := make([]robotSummary, 0, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return []robotSummary{}, fmt.Errorf("line %d did not have 2 parts", i)
		}

		robot := robotSummary{}
		for j, label := range []rune{'p', 'v'} {
			numberStrings := strings.Split(
				strings.ReplaceAll(parts[j], fmt.Sprintf("%c=", label), ""), ",")
			if len(numberStrings) != 2 {
				return []robotSummary{}, fmt.Errorf("line %d, part %c did not have 2 parts", i, label)
			}

			values := [2]int{}
			for k, numberString := range numberStrings {
				value, err := strconv.Atoi(numberString)
				if err != nil {
					return []robotSummary{}, fmt.Errorf("line %d, part %c, number %d did not have 2 parts", i, label, k)
				}
				values[2-k-1] = value
			}

			if label == 'p' {
				robot.location = values
			} else {
				robot.velocity = values
			}
		}

		robots = append(robots, robot)
	}

	return robots, nil
}

func simulateRobotMovement(
	robots []robotSummary, mapSizes [2]int, nSeconds int) []robotSummary {

	nextRobots := robots
	for range nSeconds {
		nextRobots = simulateRobotMovementsForOneSecond(nextRobots, mapSizes)
	}

	return nextRobots
}

func simulateRobotMovementUntilEasterEggFound(robots []robotSummary, mapSizes [2]int) int {
	tilesToCover := easterEggImageToMembershipSet(
		easterEggChristmasTreeImage, [2]int{easterEggI, easterEggJ})

	t := 0
	nextRobots := robots
	for {
		nextRobots = simulateRobotMovementsForOneSecond(nextRobots, mapSizes)
		t += 1

		if robotsInEasterEggArrangement(nextRobots, tilesToCover) {
			break
		}
	}

	return t
}

func simulateRobotMovementsForOneSecond(robots []robotSummary, mapSizes [2]int) []robotSummary {
	nextRobots := make([]robotSummary, 0, len(robots))
	for _, robot := range robots {
		nextRobot := robotSummary{velocity: robot.velocity}

		for i, size := range mapSizes {
			nextRobot.location[i] = (robot.location[i] + size + robot.velocity[i]) % size
		}

		nextRobots = append(nextRobots, nextRobot)
	}

	return nextRobots
}

func countRobotsInQuadrants(robots []robotSummary, mapSizes [2]int) [4]int {
	quadrantCounts := [4]int{0, 0, 0, 0}
	midlines := [2]int{}
	for i, size := range mapSizes {
		midlines[i] = size / 2
	}

EVALUATE_ROBOTS:
	for _, robot := range robots {
		quadrantAssignment := 0x00
		for i, position := range robot.location {
			midline := midlines[i]
			if position == midline {
				continue EVALUATE_ROBOTS
			} else if position > midline {
				quadrantAssignment |= (1 << i)
			}
		}
		quadrantCounts[quadrantAssignment] += 1
	}

	return quadrantCounts
}

func findSafetyFactor(quadrantCounts [4]int) int {
	safetyFactor := 1
	for _, count := range quadrantCounts {
		safetyFactor *= count
	}

	return safetyFactor
}

func easterEggImageToMembershipSet(image string, offsets [2]int) map[[2]int]struct{} {
	membership := map[[2]int]struct{}{}
	i, j := 0, 0
	for _, r := range image {
		switch r {
		case '\n':
			i, j = i+1, 0
		case '1':
			location := [2]int{i + offsets[0], j + offsets[1]}
			membership[location] = struct{}{}
			fallthrough
		default:
			j += 1
		}
	}

	return membership
}

func robotsInEasterEggArrangement(
	robots []robotSummary, easterEggMembershipSet map[[2]int]struct{}) bool {

	remainingTilesToCover := maps.Clone(easterEggMembershipSet)
	for _, robot := range robots {
		delete(remainingTilesToCover, robot.location)
	}

	return len(remainingTilesToCover) == 0
}

const (
	easterEggI = 27
	easterEggJ = 38
)

const easterEggChristmasTreeImage = `1111111111111111111111111111111
1                             1
1                             1
1                             1
1                             1
1              1              1
1             111             1
1            11111            1
1           1111111           1
1          111111111          1
1            11111            1
1           1111111           1
1          111111111          1
1         11111111111         1
1        1111111111111        1
1          111111111          1
1         11111111111         1
1        1111111111111        1
1       111111111111111       1
1      11111111111111111      1
1        1111111111111        1
1       111111111111111       1
1      11111111111111111      1
1     1111111111111111111     1
1    111111111111111111111    1
1             111             1
1             111             1
1             111             1
1                             1
1                             1
1                             1
1                             1
1111111111111111111111111111111`
