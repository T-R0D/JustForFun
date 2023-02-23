use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day10 {
    instructions: Vec<Instruction>,
}

impl Day10 {
    pub fn new() -> Self {
        Day10 {
            instructions: vec![],
        }
    }
}

impl Solver for Day10 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.instructions = input
            .trim()
            .lines()
            .map(|line| {
                if line == "noop" {
                    return Instruction::Noop;
                } else if line.starts_with("addx") {
                    let val_str = line.split(" ").collect::<Vec<_>>()[1];
                    return Instruction::Addx(val_str.parse::<i32>().expect("shoud be an integer"));
                }

                panic!("'{line}' was not a parseable instruction")
            })
            .collect();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let diagnostic_score = run_diagnostic(&self.instructions);

        Ok(diagnostic_score.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let crt_output = run_render_routine(&self.instructions);

        Ok(crt_output)
    }
}

enum Instruction {
    Noop,
    Addx(i32),
}

enum CpuState {
    Idle,
    MidAdd(i32),
}

fn run_diagnostic(instructions: &Vec<Instruction>) -> i32 {
    let mut signal_strength_diagnostic = 0;

    let mut pc = 0;
    let mut state = CpuState::Idle;
    let mut register_x: i32 = 1;

    for cycle in 1..=220 {
        let (v, next_state) = match state {
            CpuState::Idle => {
                let instruction = &instructions[pc];
                pc += 1;
                let next = match instruction {
                    Instruction::Noop => CpuState::Idle,
                    Instruction::Addx(addend) => CpuState::MidAdd(*addend),
                };
                (0, next)
            }
            CpuState::MidAdd(addend) => (addend, CpuState::Idle),
        };

        if cycle % 40 == 20 {
            let signal_strength = register_x * cycle;
            signal_strength_diagnostic += signal_strength;
        }

        register_x += v;
        state = next_state;
    }

    signal_strength_diagnostic
}

fn run_render_routine(instructions: &Vec<Instruction>) -> String {
    let mut crt_output = String::with_capacity(240);

    let mut pc = 0;
    let mut state = CpuState::Idle;
    let mut register_x: i32 = 1;

    for cycle in 1..=240 {
        let (v, next_state) = match state {
            CpuState::Idle => {
                let instruction = &instructions[pc];
                pc += 1;
                let next = match instruction {
                    Instruction::Noop => CpuState::Idle,
                    Instruction::Addx(addend) => CpuState::MidAdd(*addend),
                };
                (0, next)
            }
            CpuState::MidAdd(addend) => (addend, CpuState::Idle),
        };

        let horizontal_position = (cycle - 1) % 40;
        let row = cycle / 40;

        if horizontal_position == 0 && row > 0 {
            crt_output.push('\n');
        }

        let pixel =
            if register_x - 1 <= horizontal_position && horizontal_position <= register_x + 1 {
                '#'
            } else {
                '.'
            };
        crt_output.push(pixel);

        register_x += v;
        state = next_state;
    }
    crt_output.push('\n');

    crt_output
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
            addx 15
            addx -11
            addx 6
            addx -3
            addx 5
            addx -1
            addx -8
            addx 13
            addx 4
            noop
            addx -1
            addx 5
            addx -1
            addx 5
            addx -1
            addx 5
            addx -1
            addx 5
            addx -1
            addx -35
            addx 1
            addx 24
            addx -19
            addx 1
            addx 16
            addx -11
            noop
            noop
            addx 21
            addx -15
            noop
            noop
            addx -3
            addx 9
            addx 1
            addx -3
            addx 8
            addx 1
            addx 5
            noop
            noop
            noop
            noop
            noop
            addx -36
            noop
            addx 1
            addx 7
            noop
            noop
            noop
            addx 2
            addx 6
            noop
            noop
            noop
            noop
            noop
            addx 1
            noop
            noop
            addx 7
            addx 1
            noop
            addx -13
            addx 13
            addx 7
            noop
            addx 1
            addx -33
            noop
            noop
            noop
            addx 2
            noop
            noop
            noop
            addx 8
            noop
            addx -1
            addx 2
            addx 1
            noop
            addx 17
            addx -9
            addx 1
            addx 1
            addx -3
            addx 11
            noop
            noop
            addx 1
            noop
            addx 1
            noop
            noop
            addx -13
            addx -19
            addx 1
            addx 3
            addx 26
            addx -30
            addx 12
            addx -1
            addx 3
            addx 1
            noop
            noop
            noop
            addx -9
            addx 18
            addx 1
            addx 2
            noop
            noop
            addx 9
            noop
            noop
            noop
            addx -1
            addx 2
            addx -37
            addx 1
            addx 3
            noop
            addx 15
            addx -21
            addx 22
            addx -6
            addx 1
            noop
            addx 2
            addx 1
            noop
            addx -10
            noop
            noop
            addx 20
            addx 1
            addx 2
            addx 2
            addx -6
            addx -11
            noop
            noop
            noop
        "});
        let mut solver = Day10::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("13140", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            addx 15
            addx -11
            addx 6
            addx -3
            addx 5
            addx -1
            addx -8
            addx 13
            addx 4
            noop
            addx -1
            addx 5
            addx -1
            addx 5
            addx -1
            addx 5
            addx -1
            addx 5
            addx -1
            addx -35
            addx 1
            addx 24
            addx -19
            addx 1
            addx 16
            addx -11
            noop
            noop
            addx 21
            addx -15
            noop
            noop
            addx -3
            addx 9
            addx 1
            addx -3
            addx 8
            addx 1
            addx 5
            noop
            noop
            noop
            noop
            noop
            addx -36
            noop
            addx 1
            addx 7
            noop
            noop
            noop
            addx 2
            addx 6
            noop
            noop
            noop
            noop
            noop
            addx 1
            noop
            noop
            addx 7
            addx 1
            noop
            addx -13
            addx 13
            addx 7
            noop
            addx 1
            addx -33
            noop
            noop
            noop
            addx 2
            noop
            noop
            noop
            addx 8
            noop
            addx -1
            addx 2
            addx 1
            noop
            addx 17
            addx -9
            addx 1
            addx 1
            addx -3
            addx 11
            noop
            noop
            addx 1
            noop
            addx 1
            noop
            noop
            addx -13
            addx -19
            addx 1
            addx 3
            addx 26
            addx -30
            addx 12
            addx -1
            addx 3
            addx 1
            noop
            noop
            noop
            addx -9
            addx 18
            addx 1
            addx 2
            noop
            noop
            addx 9
            noop
            noop
            noop
            addx -1
            addx 2
            addx -37
            addx 1
            addx 3
            noop
            addx 15
            addx -21
            addx 22
            addx -6
            addx 1
            noop
            addx 2
            addx 1
            noop
            addx -10
            noop
            noop
            addx 20
            addx 1
            addx 2
            addx 2
            addx -6
            addx -11
            noop
            noop
            noop
        "});
        let mut solver = Day10::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        let expected_output = String::from(indoc! {"
            ##..##..##..##..##..##..##..##..##..##..
            ###...###...###...###...###...###...###.
            ####....####....####....####....####....
            #####.....#####.....#####.....#####.....
            ######......######......######......####
            #######.......#######.......#######.....
        "});
        assert_eq!(expected_output, result);

        Ok(())
    }
}
