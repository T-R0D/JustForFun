use crate::aoc2022::solver::interface::{AoCResult, Solver};

enum Play {
    Rock,
    Paper,
    Scissors,
}

pub struct Day02 {
    round_plays: Vec<(Play, Play)>,
}

impl Day02 {
    pub fn new() -> Self {
        Day02 {
            round_plays: vec![],
        }
    }
}

impl Solver for Day02 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let rounds = input.trim().split("\n").collect::<Vec<_>>();
        let round_plays = rounds
            .iter()
            .map(|s| {
                let plays = s.split(" ").collect::<Vec<_>>();
                if plays.len() != 2 {
                    return Err(format!(
                        "{} features too many players",
                        plays.join(" ").to_string()
                    ));
                }

                let play_1 = match plays[0] {
                    "A" => Play::Rock,
                    "B" => Play::Paper,
                    "C" => Play::Scissors,
                    _ => {
                        return Err(format!("{} is not a valid play for player 1", plays[0]));
                    }
                };

                let play_2 = match plays[1] {
                    "X" => Play::Rock,
                    "Y" => Play::Paper,
                    "Z" => Play::Scissors,
                    _ => {
                        return Err(format!("{} is not a valid play for player 2", plays[1]));
                    }
                };

                Ok((play_1, play_2))
            })
            .collect::<Result<Vec<_>, _>>()?;

        self.round_plays = round_plays;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let play_scores: [u32; 3] = [1, 2, 3];
        let outcome_scores: [u32; 3] = [0, 3, 6];

        let mut total_score: u32 = 0;
        for (opponent_play, self_play) in self.round_plays.iter() {
            let play_score = play_scores[match self_play {
                Play::Rock => 0,
                Play::Paper => 1,
                Play::Scissors => 2,
            }];

            let outcome_score = outcome_scores[match (opponent_play, self_play) {
                (Play::Rock, Play::Rock) => 1,
                (Play::Rock, Play::Paper) => 2,
                (Play::Rock, Play::Scissors) => 0,
                (Play::Paper, Play::Rock) => 0,
                (Play::Paper, Play::Paper) => 1,
                (Play::Paper, Play::Scissors) => 2,
                (Play::Scissors, Play::Rock) => 2,
                (Play::Scissors, Play::Paper) => 0,
                (Play::Scissors, Play::Scissors) => 1,
            }];

            total_score += play_score + outcome_score;
        }

        Ok(total_score.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let play_scores: [u32; 3] = [1, 2, 3];
        let outcome_scores: [u32; 3] = [0, 3, 6];

        let mut total_score: u32 = 0;
        for (opponent_play, outcome_needed) in self.round_plays.iter() {
            let (actual_play, outcome_score) = match (opponent_play, outcome_needed) {
                (Play::Rock, Play::Rock) => (Play::Scissors, outcome_scores[0]),
                (Play::Rock, Play::Paper) => (Play::Rock, outcome_scores[1]),
                (Play::Rock, Play::Scissors) => (Play::Paper, outcome_scores[2]),
                (Play::Paper, Play::Rock) => (Play::Rock, outcome_scores[0]),
                (Play::Paper, Play::Paper) => (Play::Paper, outcome_scores[1]),
                (Play::Paper, Play::Scissors) => (Play::Scissors, outcome_scores[2]),
                (Play::Scissors, Play::Rock) => (Play::Paper, outcome_scores[0]),
                (Play::Scissors, Play::Paper) => (Play::Scissors, outcome_scores[1]),
                (Play::Scissors, Play::Scissors) => (Play::Rock, outcome_scores[2]),
            };

            let play_score = play_scores[match actual_play {
                Play::Rock => 0,
                Play::Paper => 1,
                Play::Scissors => 2,
            }];

            total_score += play_score + outcome_score;
        }

        Ok(total_score.to_string())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {r"
            A Y
            B X
            C Z
        "});
        let mut solver = Day02::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("15", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {r"
            A Y
            B X
            C Z
        "});
        let mut solver = Day02::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("12", result);

        Ok(())
    }
}
