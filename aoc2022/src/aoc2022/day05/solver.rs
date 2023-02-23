use std::collections::VecDeque;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

struct Instruction {
    n: u8,
    src: usize,
    dst: usize,
}

pub struct Day05 {
    stacks: Vec<VecDeque<char>>,
    instructions: Vec<Instruction>,
}

impl Day05 {
    pub fn new() -> Self {
        Day05 {
            stacks: vec![],
            instructions: vec![],
        }
    }
}

impl Day05 {
    fn parse_starting_stacks(input: &String) -> Vec<VecDeque<char>> {
        let lines = input
            .split("\n")
            .collect::<Vec<_>>()
            .iter()
            .map(|line| line.chars().collect::<Vec<_>>())
            .collect::<Vec<_>>();

        let mut stacks: Vec<VecDeque<char>> = Vec::new();
        'mainloop: for (j, c) in lines[lines.len() - 1].iter().enumerate() {
            if *c == ' ' {
                continue 'mainloop;
            }

            let mut stack: VecDeque<char> = VecDeque::new();
            'stackbuild: for i in (0..lines.len() - 1).rev() {
                if lines[i][j] == ' ' {
                    break 'stackbuild;
                }

                stack.push_back(lines[i][j]);
            }

            stacks.push(stack);
        }

        stacks
    }

    fn parse_instructions(input: &String) -> Vec<Instruction> {
        let lines = input
            .split("\n")
            .map(|line| line.split(' ').collect::<Vec<_>>())
            .collect::<Vec<_>>();

        lines
            .iter()
            .map(|line| Instruction {
                n: line[1].parse::<u8>().expect("Should be an int"),
                src: line[3].parse::<usize>().expect("Should be an int") - 1,
                dst: line[5].parse::<usize>().expect("Should be an int") - 1,
            })
            .collect::<Vec<_>>()
    }
}

impl Solver for Day05 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let parts = input.trim().split("\n\n").collect::<Vec<_>>();

        self.stacks = Day05::parse_starting_stacks(&String::from(parts[0]));

        self.instructions = Day05::parse_instructions(&String::from(parts[1]));

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut stacks = self.stacks.clone();
        for instruction in self.instructions.iter() {
            for _ in 0..instruction.n {
                let _crate = stacks[instruction.src]
                    .pop_back()
                    .expect("Should not be empty.");
                stacks[instruction.dst].push_back(_crate);
            }
        }

        let mut tops: Vec<char> = Vec::new();
        for stack in stacks.iter_mut() {
            tops.push(stack.pop_back().expect("Should not be empty"));
        }

        let codeword = tops.iter().collect::<String>();

        Ok(codeword)
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut stacks = self.stacks.clone();
        let mut crate_claw: VecDeque<char> = VecDeque::new();
        for instruction in self.instructions.iter() {
            for _ in 0..instruction.n {
                crate_claw.push_front(stacks[instruction.src]
                    .pop_back()
                    .expect("Should not be empty."));
            }
            for _ in 0..instruction.n {
                stacks[instruction.dst].push_back(crate_claw.pop_front().expect("Crane claw should not be empty."));
            }
        }

        let mut tops: Vec<char> = Vec::new();
        for stack in stacks.iter_mut() {
            tops.push(stack.pop_back().expect("Should not be empty"));
        }

        let codeword = tops.iter().collect::<String>();

        Ok(codeword)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("[ ] [D] [ ]\n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2\n");
        let mut solver = Day05::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("CMZ", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("[ ] [D] [ ]\n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2\n");
        let mut solver = Day05::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("MCD", result);

        Ok(())
    }
}
