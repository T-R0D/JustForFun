use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day04 {
    scratch_cards: Vec<ScratchCard>,
}

impl Day04 {
    pub fn new() -> Self {
        Day04 {
            scratch_cards: vec![],
        }
    }
}

impl Solver for Day04 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let lines = input.trim().lines().collect::<Vec<_>>();

        self.scratch_cards = lines
            .iter()
            .enumerate()
            .map(|(i, line)| {
                let label_and_contents = line.split(": ").collect::<Vec<_>>();
                if label_and_contents.len() != 2 {
                    return Err(format!(
                        "Line {i} ({line}) could not have contents extracted"
                    ));
                }

                let number_groups = label_and_contents[1].split(" |").collect::<Vec<_>>();
                if label_and_contents.len() != 2 {
                    return Err(format!("Line {i} ({line}) did not have 2 number groups"));
                }

                let winning_nums = number_groups[0]
                    .split(" ")
                    .filter_map(|num_str| match num_str.parse::<u32>() {
                        Ok(value) => Some(value),
                        Err(_) => None,
                    })
                    .collect::<Vec<_>>();

                let received_nums = number_groups[1]
                    .split(" ")
                    .filter_map(|num_str| match num_str.parse::<u32>() {
                        Ok(value) => Some(value),
                        Err(_) => None,
                    })
                    .collect::<Vec<_>>();

                Ok(ScratchCard {
                    winning_nums,
                    received_nums,
                })
            })
            .collect::<Result<Vec<ScratchCard>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let card_score_total = self
            .scratch_cards
            .iter()
            .map(|card| card.score())
            .sum::<u32>();

        Ok(card_score_total.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut cards_held = vec![1_usize; self.scratch_cards.len()];
        
        for (i, card) in self.scratch_cards.iter().enumerate() {
            let n_held = cards_held[i];
            let next_cards_won = card.winning_num_matches();
            for j in (i+1)..=(i+usize::try_from(next_cards_won).unwrap()) {
                cards_held[j] += n_held;
            }
        }

        let total_cards_held = cards_held.iter().sum::<usize>();

        Ok(total_cards_held.to_string())
    }
}

struct ScratchCard {
    received_nums: Vec<u32>,
    winning_nums: Vec<u32>,
}

impl ScratchCard {
    fn winning_num_matches(&self) -> u32 {
        self.winning_nums.iter().fold(0, |acc, winning_num| {
            if self.received_nums.contains(winning_num) {
                acc + 1
            } else {
                acc
            }
        })
    }

    fn score(&self) -> u32 {
        let n_matches = self.winning_num_matches();

        if n_matches <= 1 {
            n_matches
        } else {
            2_u32.pow(n_matches - 1)
        }
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
            Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
            Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
            Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
            Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
            Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
            Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
        "});
        let mut solver = Day04::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("13", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
            Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
            Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
            Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
            Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
            Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
        "});
        let mut solver = Day04::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("30", result);

        Ok(())
    }
}
