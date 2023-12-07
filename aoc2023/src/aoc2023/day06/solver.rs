use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day06 {
    boat_races: Vec<BoatRace>,
}

impl Day06 {
    pub fn new() -> Self {
        Day06 { boat_races: vec![] }
    }
}

impl Solver for Day06 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let lines = input.trim().lines().collect::<Vec<_>>();
        if lines.len() != 2 {
            return Err(String::from("Input is malformed"));
        }

        let durations = lines[0]
            .split(" ")
            .skip(1)
            .filter_map(|item| {
                if item == "" {
                    return None;
                }

                Some(match item.parse::<u64>() {
                    Ok(value) => Ok(value),
                    Err(_) => Err(format!("Failed to parse {item}")),
                })
            })
            .collect::<Result<Vec<u64>, String>>()?;

        let records = lines[1]
            .split(" ")
            .skip(1)
            .filter_map(|item| {
                if item == "" {
                    return None;
                }

                Some(match item.parse::<u64>() {
                    Ok(value) => Ok(value),
                    Err(_) => Err(format!("Failed to parse {item}")),
                })
            })
            .collect::<Result<Vec<u64>, String>>()?;

        if durations.len() != records.len() {
            return Err(format!(
                "Mismatch of durations and records ({} != {})",
                durations.len(),
                records.len()
            ));
        }

        self.boat_races = durations
            .iter()
            .zip(records.iter())
            .map(|(duration, record)| BoatRace {
                duration: *duration,
                record_distance: *record,
            })
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let winning_product = self
            .boat_races
            .iter()
            .map(|race| race.find_record_breaking_strategy_outcomes().len())
            .product::<usize>();

        Ok(winning_product.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let actual_race = self.get_merged_boat_race()?;

        let n_wins = actual_race.find_n_ways_to_win();

        Ok(n_wins.to_string())
    }
}

impl Day06 {
    fn get_merged_boat_race(&self) -> Result<BoatRace, String> {
        let merged_duration = self.concatenate_numbers(
            self.boat_races
                .iter()
                .map(|race| race.duration)
                .collect::<Vec<_>>(),
        )?;
        let merged_record = self.concatenate_numbers(
            self.boat_races
                .iter()
                .map(|race| race.record_distance)
                .collect::<Vec<_>>(),
        )?;
        Ok(BoatRace {
            duration: merged_duration,
            record_distance: merged_record,
        })
    }

    fn concatenate_numbers(&self, numbers: Vec<u64>) -> Result<u64, String> {
        match numbers
            .iter()
            .map(|x| x.to_string())
            .collect::<String>()
            .parse::<u64>()
        {
            Ok(value) => Ok(value),
            Err(err) => Err(format!("Unable to merge numbers: {}", err.to_string())),
        }
    }
}

struct BoatRace {
    duration: u64,
    record_distance: u64,
}

impl BoatRace {
    pub fn find_record_breaking_strategy_outcomes(&self) -> Vec<u64> {
        (1..self.duration)
            .filter_map(|hold_duration| {
                let distance_moved = self.find_distance_for_hold_duration(hold_duration);
                if distance_moved > self.record_distance {
                    return Some(distance_moved);
                }

                None
            })
            .collect::<Vec<_>>()
    }

    pub fn find_n_ways_to_win(&self) -> usize {
        let n_possibilities = self.duration + 1;
        let mut losing_possibilites = 1;
        let midpoint = n_possibilities / 2;
        for hold_duration in 1..=midpoint {
            if self.find_distance_for_hold_duration(hold_duration) > self.record_distance {
                break;
            }

            losing_possibilites += 1;
        }
        
        usize::try_from(n_possibilities - (losing_possibilites * 2)).unwrap()
    }

    fn find_distance_for_hold_duration(&self, hold_duration: u64) -> u64 {
        let charge_speed = hold_duration;
        let moving_time = self.duration - hold_duration;
        moving_time * charge_speed
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[cfg(test)]
    use indoc::indoc;
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Time:      7  15   30
            Distance:  9  40  200
        "});
        let mut solver = Day06::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("288", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Time:      7  15   30
            Distance:  9  40  200
        "});
        let mut solver = Day06::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("71503", result);

        Ok(())
    }
}
