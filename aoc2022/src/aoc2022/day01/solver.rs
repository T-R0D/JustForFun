use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day01 {
    carries: Vec<Vec<u32>>,
}

impl Day01 {
    pub fn new() -> Self {
        Day01 { carries: vec![] }
    }
}

impl Solver for Day01 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let elf_payload_strs = input
            .split("\n\n")
            .collect::<Vec<_>>()
            .iter()
            .filter_map(|s| match *s {
                "" => None,
                _ => Some(String::from(*s)),
            })
            .collect::<Vec<_>>();

        let elf_payloads = elf_payload_strs
            .iter()
            .map(|s| {
                s.split("\n")
                    .collect::<Vec<_>>()
                    .iter()
                    .filter_map(|i| match *i {
                        "" => None,
                        _ => Some(i.parse::<u32>().expect("Item should be a positive integer")),
                    })
                    .collect::<Vec<_>>()
            })
            .collect::<Vec<_>>();

        self.carries = elf_payloads;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let totals = self
            .carries
            .iter()
            .map(|payload| payload.iter().sum())
            .collect::<Vec<u32>>();

        let highest_caloric_payload = totals.iter().max();

        Ok(highest_caloric_payload.unwrap_or(&0).to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut totals = self
            .carries
            .iter()
            .map(|payload| payload.iter().sum())
            .collect::<Vec<u32>>();
        totals.sort();
        let top_three_caloric_collections = totals.split_off(totals.len() - 3);
        let total_calories: u32 = top_three_caloric_collections.iter().sum();

        Ok(total_calories.to_string())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = String::from(indoc! {r"
            1000
            2000
            3000

            4000

            5000
            6000

            7000
            8000
            9000

            10000
        "});
        let mut solver = Day01::new();
        solver.consume_input(&input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("24000", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = String::from(indoc! {r"
            1000
            2000
            3000

            4000

            5000
            6000

            7000
            8000
            9000

            10000
        "});
        let mut solver = Day01::new();
        solver.consume_input(&input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("45000", result);

        Ok(())
    }
}
