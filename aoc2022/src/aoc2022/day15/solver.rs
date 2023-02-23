use std::collections::HashSet;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day15 {
    sensors: HashSet<Sensor>,
    beacons: HashSet<MatrixCoordinate>,
}

impl Day15 {
    pub fn new() -> Self {
        Day15 {
            sensors: HashSet::new(),
            beacons: HashSet::new(),
        }
    }
}

const SAMPLE_SEARCH_ROW: i64 = 2_000_000;
const KNOWN_LOWER_BOUND: i64 = 0;
const KNOWN_UPPER_BOUND: i64 = 4_000_000;

impl Solver for Day15 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        for line in input.trim().lines() {
            let processed_line = line
                .replace("Sensor at ", "")
                .replace(" closest beacon is at ", "");
            let coordinates = processed_line
                .split(":")
                .map(MatrixCoordinate::from_str)
                .collect::<Vec<_>>();

            self.sensors
                .insert(Sensor::new(&coordinates[0], &coordinates[1]));
            self.beacons.insert(coordinates[1].clone());
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let brute_force = false;
        if brute_force {
            let mut certainly_non_beacon_columns: HashSet<i64> = HashSet::new();

            for sensor in self.sensors.iter() {
                let effective_reach = sensor.loc.manhattan_distance(&sensor.closest_beacon_loc);
                let reach_to_target_row = (SAMPLE_SEARCH_ROW - sensor.loc.0).abs();

                if effective_reach < reach_to_target_row {
                    continue;
                }

                let spread = effective_reach - reach_to_target_row;

                for j in (sensor.loc.1 - spread)..=(sensor.loc.1 + spread) {
                    let candidate = MatrixCoordinate(SAMPLE_SEARCH_ROW, j);
                    if !self.beacons.contains(&candidate) {
                        certainly_non_beacon_columns.insert(j);
                    }
                }
            }

            return Ok(certainly_non_beacon_columns.len().to_string());
        }

        let mut covered_intervals = IntervalContainer::new();

        for sensor in self.sensors.iter() {
            let effective_reach = sensor.loc.manhattan_distance(&sensor.closest_beacon_loc);
            let reach_to_target_row = (SAMPLE_SEARCH_ROW - sensor.loc.0).abs();

            if effective_reach < reach_to_target_row {
                continue;
            }

            let spread = effective_reach - reach_to_target_row;

            covered_intervals.add_interval(&InclusiveInterval(
                sensor.loc.1 - spread,
                sensor.loc.1 + spread,
            ));
        }

        covered_intervals.fix_up();

        let mut certainly_non_beacon_columns: HashSet<i64> = HashSet::new();
        for column in covered_intervals.covered_locations() {
            let candidate = MatrixCoordinate(SAMPLE_SEARCH_ROW, column);
            if !self.beacons.contains(&candidate) {
                certainly_non_beacon_columns.insert(column);
            }
        }

        Ok(certainly_non_beacon_columns.len().to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        for i in KNOWN_LOWER_BOUND..=KNOWN_UPPER_BOUND {
            let mut covered_intervals = IntervalContainer::new();

            for sensor in self.sensors.iter() {
                let effective_reach = sensor.loc.manhattan_distance(&sensor.closest_beacon_loc);
                let reach_to_target_row = (i - sensor.loc.0).abs();

                if effective_reach < reach_to_target_row {
                    continue;
                }

                let spread = effective_reach - reach_to_target_row;

                covered_intervals.add_interval(&InclusiveInterval(
                    sensor.loc.1 - spread,
                    sensor.loc.1 + spread,
                ));
            }

            covered_intervals.fix_up();
            covered_intervals.restrict(KNOWN_LOWER_BOUND, KNOWN_UPPER_BOUND);
            if covered_intervals.covered_count()
                < (KNOWN_UPPER_BOUND - KNOWN_LOWER_BOUND + 1) as usize
            {
                for (x, j) in covered_intervals.covered_locations().iter().enumerate() {
                    if x as i64 != *j {
                        return Ok(((x * 4_000_000) as i64 + i).to_string());
                    }
                }
            }
        }

        Err(String::from("Unable to find an empty location..."))
    }
}

#[derive(Clone, Eq, Hash, PartialEq)]
struct MatrixCoordinate(i64, i64);

impl MatrixCoordinate {
    fn from_str(s: &str) -> Self {
        let coords = s
            .replace("x=", "")
            .replace("y=", "")
            .split(", ")
            .map(str::parse::<i64>)
            .map(Result::unwrap)
            .collect::<Vec<_>>();
        MatrixCoordinate(coords[1], coords[0])
    }
    fn manhattan_distance(&self, other: &MatrixCoordinate) -> i64 {
        (self.0 - other.0).abs() + (self.1 - other.1).abs()
    }
}

#[derive(Eq, Hash, PartialEq)]
struct Sensor {
    loc: MatrixCoordinate,
    closest_beacon_loc: MatrixCoordinate,
}

impl Sensor {
    fn new(loc: &MatrixCoordinate, beacon: &MatrixCoordinate) -> Self {
        Sensor {
            loc: loc.clone(),
            closest_beacon_loc: beacon.clone(),
        }
    }
}

#[derive(Clone)]
struct InclusiveInterval(i64, i64);

struct IntervalContainer {
    intervals: Vec<InclusiveInterval>,
}

impl IntervalContainer {
    fn new() -> Self {
        IntervalContainer {
            intervals: Vec::new(),
        }
    }

    fn add_interval(&mut self, interval: &InclusiveInterval) {
        self.intervals.push(interval.clone());
    }

    fn fix_up(&mut self) {
        if self.intervals.len() < 1 {
            return;
        }

        self.intervals.sort_by(|a, b| {
            let res = a.0.cmp(&b.0);
            if res.is_eq() {
                return a.1.cmp(&b.1);
            }
            res
        });

        let mut new_intervals: Vec<InclusiveInterval> = Vec::with_capacity(self.intervals.len());
        let mut low = self.intervals[0].0;
        let mut high = self.intervals[0].1;
        for interval in self.intervals.iter().skip(1) {
            if interval.0 > high {
                new_intervals.push(InclusiveInterval(low, high));
                low = interval.0;
                high = interval.1;
            } else if interval.1 > high {
                high = interval.1;
            }
        }
        new_intervals.push(InclusiveInterval(low, high));
        self.intervals = new_intervals;
    }

    fn restrict(&mut self, first: i64, last: i64) {
        let mut new_intervals: Vec<InclusiveInterval> = Vec::with_capacity(self.intervals.len());
        for interval in self.intervals.iter() {
            if interval.0 > last {
                break;
            }

            if interval.1 < first {
                continue;
            }

            let new_interval = InclusiveInterval(
                std::cmp::max(interval.0, first),
                std::cmp::min(interval.1, last),
            );
            new_intervals.push(new_interval);

            if interval.1 > last {
                break;
            }
        }

        self.intervals = new_intervals;
    }

    fn covered_count(&mut self) -> usize {
        let mut n_covered = 0;
        for interval in self.intervals.iter() {
            n_covered += (interval.1 - interval.0 + 1) as usize
        }
        n_covered
    }

    fn covered_locations(&mut self) -> Vec<i64> {
        let mut covered_locations = Vec::new();
        for interval in self.intervals.iter() {
            for j in interval.0..=interval.1 {
                covered_locations.push(j);
            }
        }
        covered_locations
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[cfg(test)]
    use indoc::indoc;
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    // Need to figure out how I want to do this tests without changing the
    // hard-coded global constants. Until then, the tests fail ¯\_(ツ)_/¯

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Sensor at x=2, y=18: closest beacon is at x=-2, y=15
            Sensor at x=9, y=16: closest beacon is at x=10, y=16
            Sensor at x=13, y=2: closest beacon is at x=15, y=3
            Sensor at x=12, y=14: closest beacon is at x=10, y=16
            Sensor at x=10, y=20: closest beacon is at x=10, y=16
            Sensor at x=14, y=17: closest beacon is at x=10, y=16
            Sensor at x=8, y=7: closest beacon is at x=2, y=10
            Sensor at x=2, y=0: closest beacon is at x=2, y=10
            Sensor at x=0, y=11: closest beacon is at x=2, y=10
            Sensor at x=20, y=14: closest beacon is at x=25, y=17
            Sensor at x=17, y=20: closest beacon is at x=21, y=22
            Sensor at x=16, y=7: closest beacon is at x=15, y=3
            Sensor at x=14, y=3: closest beacon is at x=15, y=3
            Sensor at x=20, y=1: closest beacon is at x=15, y=3
        "});
        let mut solver = Day15::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("26", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Sensor at x=2, y=18: closest beacon is at x=-2, y=15
            Sensor at x=9, y=16: closest beacon is at x=10, y=16
            Sensor at x=13, y=2: closest beacon is at x=15, y=3
            Sensor at x=12, y=14: closest beacon is at x=10, y=16
            Sensor at x=10, y=20: closest beacon is at x=10, y=16
            Sensor at x=14, y=17: closest beacon is at x=10, y=16
            Sensor at x=8, y=7: closest beacon is at x=2, y=10
            Sensor at x=2, y=0: closest beacon is at x=2, y=10
            Sensor at x=0, y=11: closest beacon is at x=2, y=10
            Sensor at x=20, y=14: closest beacon is at x=25, y=17
            Sensor at x=17, y=20: closest beacon is at x=21, y=22
            Sensor at x=16, y=7: closest beacon is at x=15, y=3
            Sensor at x=14, y=3: closest beacon is at x=15, y=3
            Sensor at x=20, y=1: closest beacon is at x=15, y=3
        "});
        let mut solver = Day15::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("56000011", result);

        Ok(())
    }
}
