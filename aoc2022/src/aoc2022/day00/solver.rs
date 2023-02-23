/// Day00 is a safe space for noodling.

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day00 {}

impl Day00 {
    pub fn new() -> Self {
        Day00 {}
    }
}

impl Solver for Day00 {
    fn consume_input(&mut self, _input: &String) -> AoCResult {
        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        Err(String::from("Not implemented."))
    }

    fn solve_part_2(&self) -> AoCResult {
        Err(String::from("Not implemented."))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let solver = Day00::new();

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("", result);

        Ok(())
    }
}
