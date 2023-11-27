use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day23 {}

impl Day23 {
    pub fn new() -> Self {
        Day23 {}
    }
}

impl Solver for Day23 {
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
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("");
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("");
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("", result);

        Ok(())
    }
}
