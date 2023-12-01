use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day00 {
    input: String,
}

impl Day00 {
    pub fn new() -> Self {
        Day00 {
            input: String::from(""),
        }
    }
}

impl Solver for Day00 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.input = input.to_string();
        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        Ok(self.input.to_owned())
    }

    fn solve_part_2(&self) -> AoCResult {
        Ok(self.input.to_ascii_uppercase())
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
        let input = &String::from("XYZ");
        let mut solver = Day00::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("XYZ", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("abc");
        let mut solver = Day00::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("ABC", result);

        Ok(())
    }
}
