use regex::{Error, Regex};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

const NUMBER_VALUES: [(&str, u32); 9] = [
    (r"one", 1),
    (r"two", 2),
    (r"three", 3),
    (r"four", 4),
    (r"five", 5),
    (r"six", 6),
    (r"seven", 7),
    (r"eight", 8),
    (r"nine", 9),
];

pub struct Day01 {
    forward_digit_regexes: Vec<(Regex, u32)>,
    reverse_digit_regexes: Vec<(Regex, u32)>,
    text_lines: Vec<String>,
}

impl Day01 {
    pub fn new() -> Self {
        Day01 {
            forward_digit_regexes: vec![],
            reverse_digit_regexes: vec![],
            text_lines: vec![],
        }
    }
}

impl Solver for Day01 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let forward_regexes_and_values_result = NUMBER_VALUES
            .iter()
            .map(|s_and_v| self.new_regex_value_tuple(s_and_v))
            .collect::<Result<Vec<_>, _>>();

        self.forward_digit_regexes = match forward_regexes_and_values_result {
            Ok(regexes_and_values) => regexes_and_values,
            Err(e) => Err(e.to_string())?,
        };

        let reverse_regexes_and_values_result = NUMBER_VALUES
            .iter()
            .map(|s_and_v| {
                let (s, v) = s_and_v;
                let reversed_s = s.chars().rev().collect::<String>();
                self.new_regex_value_tuple(&(reversed_s.as_str(), *v))
            })
            .collect::<Result<Vec<_>, _>>();

        self.reverse_digit_regexes = match reverse_regexes_and_values_result {
            Ok(regexes_and_values) => regexes_and_values,
            Err(e) => Err(e.to_string())?,
        };

        self.text_lines = input
            .trim()
            .split("\n")
            .map(|line| line.to_string())
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let line_digits = self
            .text_lines
            .iter()
            .map(|line| {
                line.chars()
                    .filter_map(|c| {
                        if c.is_ascii_digit() {
                            c.to_digit(10)
                        } else {
                            None
                        }
                    })
                    .collect::<Vec<_>>()
            })
            .collect::<Vec<_>>();

        let calibration_sum = line_digits
            .iter()
            .map(|line| match (line.first(), line.last()) {
                (Some(a), Some(b)) => a * 10 + b,
                _ => 0,
            })
            .sum::<u32>();

        Ok(calibration_sum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let first_and_last_digits = self
            .text_lines
            .iter()
            .map(|line| {
                let mut first_digit = 0;
                let mut buffer = Vec::<char>::new();
                for c in line.chars() {
                    if let Some(digit) = c.to_digit(10) {
                        first_digit = digit;
                        break;
                    }
                    buffer.push(c);

                    if let Some(digit) =
                        self.word_to_forward_digit(&buffer.iter().collect::<String>())
                    {
                        first_digit = digit;
                        break;
                    }
                }

                let mut last_digit = 0;
                buffer.clear();
                for c in line.chars().rev() {
                    if let Some(digit) = c.to_digit(10) {
                        last_digit = digit;
                        break;
                    }

                    buffer.push(c);

                    if let Some(digit) =
                        self.word_to_reverse_digit(&buffer.iter().collect::<String>())
                    {
                        last_digit = digit;
                        break;
                    }
                }

                vec![first_digit, last_digit]
            })
            .collect::<Vec<_>>();

        let calibration_sum = first_and_last_digits
            .iter()
            .map(|line| match (line.first(), line.last()) {
                (Some(a), Some(b)) => a * 10 + b,
                _ => 0,
            })
            .sum::<u32>();

        Ok(calibration_sum.to_string())
    }
}

impl Day01 {
    fn new_regex_value_tuple(&self, s_and_v: &(&str, u32)) -> Result<(Regex, u32), Error> {
        let (s, v) = s_and_v;
        let regex = Regex::new(s)?;
        Ok((regex, *v))
    }

    fn word_to_forward_digit(&self, word: &String) -> Option<u32> {
        for (regex, value) in self.forward_digit_regexes.iter() {
            if regex.is_match(word) {
                return Some(*value);
            }
        }

        None
    }

    fn word_to_reverse_digit(&self, word: &String) -> Option<u32> {
        for (regex, value) in self.reverse_digit_regexes.iter() {
            if regex.is_match(word) {
                return Some(*value);
            }
        }

        None
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
        let input = &String::from(indoc! { r"
            1abc2
            pqr3stu8vwx
            a1b2c3d4e5f
            treb7uchet
        "});
        let mut solver = Day01::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("142", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            two1nine
            eightwothree
            abcone2threexyz
            xtwone3four
            4nineeightseven2
            zoneight234
            7pqrstsixteen
        "});
        let mut solver = Day01::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("281", result);

        Ok(())
    }

    #[test]
    fn part_two_tricky_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            oneight
            seven
            twone
        "});
        let mut solver = Day01::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("116", result);

        Ok(())
    }
}
