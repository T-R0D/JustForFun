use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day25 {
    fuel_requirements: Vec<String>,
}

impl Day25 {
    pub fn new() -> Self {
        Self {
            fuel_requirements: Vec::new(),
        }
    }
}

impl Solver for Day25 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.fuel_requirements = input.trim().lines().map(String::from).collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let total_fuel = self
            .fuel_requirements
            .iter()
            .map(snafu_to_decimal)
            .sum::<i64>();

        let total_fuel_in_snafu = decimal_to_snafu(total_fuel);

        Ok(total_fuel_in_snafu)
    }

    fn solve_part_2(&self) -> AoCResult {
        Ok(String::from("Merry Christmas!"))
    }
}

const SNAFU_BASE: i64 = 5;

fn snafu_to_decimal(snafu_number: &String) -> i64 {
    let mut place_value = 1;
    let mut total = 0;

    for digit in snafu_number.chars().rev() {
        total += place_value
            * match digit {
                '0' => 0,
                '1' => 1,
                '2' => 2,
                '-' => -1,
                '=' => -2,
                _ => unreachable!(),
            };

        place_value *= SNAFU_BASE;
    }

    total
}

fn decimal_to_snafu(x: i64) -> String {
    let mut snafu_number = String::new();

    let mut working_num = x;
    let mut carry = 0;

    while carry > 0 || working_num > 0 {
        let digit = (working_num % SNAFU_BASE) + carry;
        working_num /= SNAFU_BASE;

        carry = if digit > 2 { 1 } else { 0 };

        let snafu_digit = match digit % SNAFU_BASE {
            0 => '0',
            1 => '1',
            2 => '2',
            3 => '=',
            4 => '-',
            _ => unreachable!(),
        };
        snafu_number.push(snafu_digit);
    }

    snafu_number.chars().rev().collect()
}

/*
29 -> 5, 4 + 0 -> 5 | 1 | -
5 | 1 | - -> 1, 0 + 1 -> 1 | 0 | 1-
1 | 0 | 1- -> 0, 1 + 0 -> 0 | 0 | 11-


24 -> 4, 4 + 0 -> 4 | 1 | -
4 | 1 | - -> 0, 4 + 1 -> 0 | 1 | 0-
0 | 1 | 0- -> 0, 0 + 1 -> 0 | 0 | 10-


 */

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
            1=-0-2
            12111
            2=0=
            21
            2=01
            111
            20012
            112
            1=-1=
            1-12
            12
            1=
            122
        "});
        let mut solver = Day25::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("2=-1=0", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("");
        let mut solver = Day25::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("Merry Christmas!", result);

        Ok(())
    }
}
