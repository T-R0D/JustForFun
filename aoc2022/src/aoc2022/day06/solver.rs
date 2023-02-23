use std::{vec, collections::VecDeque};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day06 {
    data_stream: Vec<char>,
}

impl Day06 {
    pub fn new() -> Self {
        Day06 {
            data_stream: vec![],
        }
    }

    fn window_has_only_unique(window_contents: [u8; 26]) -> bool {
        for x in window_contents {
            if x > 1 {
                return false;
            }
        }
        true
    }
}

fn lower_ascii_char_to_u8(c: char) -> u8 {
    (c as u32 - 'a' as u32) as u8
}

impl Solver for Day06 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.data_stream = input.trim().to_string().chars().collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut window: VecDeque<char> = VecDeque::new();
        let mut window_contents: [u8;26] = [0; 26];
        for i in 0..4 {
            let c = self.data_stream[i];
            window.push_back(c);
            window_contents[lower_ascii_char_to_u8(c) as usize] += 1;
        }

        let mut start_of_message = 0;
        for i in 4..self.data_stream.len() {
            if Day06::window_has_only_unique(window_contents) {
                start_of_message = i;
                break;
            }

            let ejected = window.pop_front().expect("window should not be empty");
            window_contents[lower_ascii_char_to_u8(ejected) as usize] -= 1;
            let c = self.data_stream[i];
            window.push_back(c);
            window_contents[lower_ascii_char_to_u8(c) as usize] += 1;
        }

        Ok(start_of_message.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut window: VecDeque<char> = VecDeque::new();
        let mut window_contents: [u8;26] = [0; 26];
        for i in 0..14 {
            let c = self.data_stream[i];
            window.push_back(c);
            window_contents[lower_ascii_char_to_u8(c) as usize] += 1;
        }

        let mut start_of_message = 0;
        for i in 14..self.data_stream.len() {
            if Day06::window_has_only_unique(window_contents) {
                start_of_message = i;
                break;
            }

            let ejected = window.pop_front().expect("window should not be empty");
            window_contents[lower_ascii_char_to_u8(ejected) as usize] -= 1;
            let c = self.data_stream[i];
            window.push_back(c);
            window_contents[lower_ascii_char_to_u8(c) as usize] += 1;
        }

        Ok(start_of_message.to_string())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("mjqjpqmgbljsphdztnvjfqwrcgsmlb");
        let mut solver = Day06::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("7", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("mjqjpqmgbljsphdztnvjfqwrcgsmlb");
        let mut solver = Day06::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("19", result);

        Ok(())
    }
}
