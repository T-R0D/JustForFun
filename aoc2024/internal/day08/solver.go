package day08

import (
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	region := parseAntennaMap(input)

	antinodes := mapAntinodes(region)
	nUniqueLocations := countUniqueLocations(antinodes)

	return fmt.Sprintf("%d", nUniqueLocations), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	region := parseAntennaMap(input)

	antinodes := mapAntinodesAccountingForResonantHarmonics(region)
	nUniqueLocations := countUniqueLocations(antinodes)

	return fmt.Sprintf("%d", nUniqueLocations), nil
}

type antennaRegion struct {
	height   int
	width    int
	antennas []antennaSpec
}
type antennaSpec struct {
	frequency rune
	location  [2]int
}

func parseAntennaMap(input string) antennaRegion {
	antennas := []antennaSpec{}
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := 0
	for i, line := range lines {
		width = len(line)
		for j, r := range line {
			if r != '.' {
				antenna := antennaSpec{
					frequency: r,
					location:  [2]int{i, j},
				}
				antennas = append(antennas, antenna)
			}
		}
	}

	return antennaRegion{antennas: antennas, height: height, width: width}
}

func mapAntinodes(region antennaRegion) map[rune][][2]int {
	antinodes := map[rune][][2]int{}
	for i, antennaA := range region.antennas {
		for j, antennaB := range region.antennas {
			if i == j {
				continue
			}

			if antennaA.frequency != antennaB.frequency {
				continue
			}

			dx, dy := antennaA.location[0]-antennaB.location[0], antennaA.location[1]-antennaB.location[1]
			antinodeLocation := [2]int{antennaA.location[0] + dx, antennaA.location[1] + dy}
			if antinodeLocation[0] < 0 || region.height <= antinodeLocation[0] ||
				antinodeLocation[1] < 0 || region.width <= antinodeLocation[1] {

				continue
			}

			if antinodesForFrequency, ok := antinodes[antennaA.frequency]; ok {
				antinodes[antennaA.frequency] = append(antinodesForFrequency, antinodeLocation)
			} else {
				antinodes[antennaA.frequency] = [][2]int{antinodeLocation}
			}
		}
	}

	return antinodes
}

func mapAntinodesAccountingForResonantHarmonics(region antennaRegion) map[rune][][2]int {
	antinodes := map[rune][][2]int{}
	for i, antennaA := range region.antennas {
		for j:= i + 1; j < len(region.antennas); j += 1 {
			antennaB := region.antennas[j]

			if antennaA.frequency != antennaB.frequency {
				continue
			}

			discoveredAntinodes := mapAntinodesAlongLine(antennaA, antennaB, region)
			if existingLocations, ok := antinodes[antennaA.frequency]; ok {
				antinodes[antennaA.frequency] = append(existingLocations, discoveredAntinodes...)
			} else {
				antinodes[antennaA.frequency] = discoveredAntinodes
			}
		}
	}

	return antinodes
}

func mapAntinodesAlongLine(a antennaSpec, b antennaSpec, region antennaRegion) [][2]int {
	dx, dy := a.location[0] - b.location[0], a.location[1] - b.location[1]

	antinodeLocations := [][2]int{a.location}
	for i := 1; ; i+=1{
		candidate := [2]int{a.location[0] + i * dx, a.location[1] + i * dy}
		
		if candidate[0] < 0 || region.height <= candidate[0] ||
			candidate[1] < 0 || region.width <= candidate[1] {

			break
		}

		antinodeLocations = append(antinodeLocations, candidate)
	}

	for i := 1; ; i+=1{
		candidate := [2]int{a.location[0] + -i * dx, a.location[1] + -i * dy}
		
		if candidate[0] < 0 || region.height <= candidate[0] ||
			candidate[1] < 0 || region.width <= candidate[1] {

			break
		}

		antinodeLocations = append(antinodeLocations, candidate)
	}

	return antinodeLocations
}

func countUniqueLocations(antinodes map[rune][][2]int) int {
	locations := map[[2]int]struct{}{}
	for _, locationsForFrequency := range antinodes {
		for _, location := range locationsForFrequency {
			locations[location] = struct{}{}
		}
	}

	return len(locations)
}
