use std::collections::HashMap;

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day14 {
    platform: Platform,
}

impl Day14 {
    pub fn new() -> Self {
        Day14 {
            platform: Platform {
                data: vec![],
                m: 0,
                n: 0,
            },
        }
    }
}

impl Solver for Day14 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.platform = Platform::try_from_grid(input.trim())?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut platform = self.platform.clone();
        platform.tilt_north();
        let total_load = platform.evaluate_load_on_north_support();
        Ok(total_load.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut repr_to_steps = HashMap::<String, usize>::new();
        let mut steps_to_load = HashMap::<usize, usize>::new();
        let mut offset = 0_usize;
        let mut cycle_period = 0_usize;
        let mut platform = self.platform.clone();

        for cycle in 0..N_CYCLES_TO_BE_SAFE {
            let repr = platform.to_string();
            if let Some(&steps) = repr_to_steps.get(&repr) {
                offset = steps;
                cycle_period = cycle - offset;
                break;
            }

            repr_to_steps.insert(repr, cycle);
            steps_to_load.insert(cycle, platform.evaluate_load_on_north_support());

            platform.tilt_north();
            platform.tilt_west();
            platform.tilt_south();
            platform.tilt_east();
        }

        let last_cycle_position = offset + ((N_CYCLES_TO_BE_SAFE - offset) % cycle_period);

        if let Some(final_load) = steps_to_load.get(&last_cycle_position) {
            Ok(final_load.to_string())
        } else {
            Err(format!("{last_cycle_position} not found in the lookup"))
        }
    }
}

const N_CYCLES_TO_BE_SAFE: usize = 1_000_000_000;

#[derive(Clone)]
struct Platform {
    data: Vec<Vec<Contents>>,
    m: usize,
    n: usize,
}

impl Platform {
    fn try_from_grid(grid: &str) -> Result<Self, String> {
        let data = grid
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(Contents::try_from_char)
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;
        let m = data.len();
        let n = data[0].len();
        Ok(Self { data, m, n })
    }

    fn evaluate_load_on_north_support(&self) -> usize {
        let mut total_load = 0_usize;
        for i in 0..self.m {
            let leverage_multiplier = self.m - i;
            for j in 0..self.n {
                if self.data[i][j] == Contents::RoundRock {
                    total_load += leverage_multiplier;
                }
            }
        }
        total_load
    }

    fn to_string(&self) -> String {
        let mut s = String::with_capacity(self.m * self.n);
        for row in self.data.iter() {
            for c in row.iter() {
                s.push(match c {
                    Contents::CubeRock => '#',
                    Contents::Empty => '.',
                    Contents::RoundRock => 'O',
                });
            }
        }
        s
    }

    fn tilt_north(&mut self) {
        for j in 0..self.n {
            let mut landing_spot = 0_usize;
            for i in 0..self.m {
                let contents = &self.data[i][j];

                match contents {
                    Contents::CubeRock => {
                        landing_spot = i + 1;
                    }
                    Contents::RoundRock => {
                        if landing_spot < i {
                            self.data[landing_spot][j] = Contents::RoundRock;
                            self.data[i][j] = Contents::Empty;
                        }
                        landing_spot = landing_spot + 1;
                    }
                    Contents::Empty => (),
                }
            }
        }
    }

    fn tilt_south(&mut self) {
        for j in 0..self.n {
            let mut landing_spot = self.m - 1;
            for i in (0..self.m).rev() {
                let contents = &self.data[i][j];

                match contents {
                    Contents::CubeRock => {
                        if i > 0 {
                            landing_spot = i - 1;
                        }
                    }
                    Contents::RoundRock => {
                        if landing_spot > i {
                            self.data[landing_spot][j] = Contents::RoundRock;
                            self.data[i][j] = Contents::Empty;
                        }

                        if landing_spot > 0 {
                            landing_spot = landing_spot - 1;
                        }
                    }
                    Contents::Empty => (),
                }
            }
        }
    }

    fn tilt_west(&mut self) {
        for i in 0..self.m {
            let mut landing_spot = 0_usize;
            for j in 0..self.n {
                let contents = &self.data[i][j];

                match contents {
                    Contents::CubeRock => {
                        landing_spot = j + 1;
                    }
                    Contents::RoundRock => {
                        if landing_spot < j {
                            self.data[i][landing_spot] = Contents::RoundRock;
                            self.data[i][j] = Contents::Empty;
                        }
                        landing_spot = landing_spot + 1;
                    }
                    Contents::Empty => (),
                }
            }
        }
    }

    fn tilt_east(&mut self) {
        for i in 0..self.m {
            let mut landing_spot = self.n - 1;
            for j in (0..self.n).rev() {
                let contents = &self.data[i][j];

                match contents {
                    Contents::CubeRock => {
                        if j > 0 {
                            landing_spot = j - 1;
                        }
                    }
                    Contents::RoundRock => {
                        if landing_spot > j {
                            self.data[i][landing_spot] = Contents::RoundRock;
                            self.data[i][j] = Contents::Empty;
                        }

                        if landing_spot > 0 {
                            landing_spot = landing_spot - 1;
                        }
                    }
                    Contents::Empty => (),
                }
            }
        }
    }
}

#[derive(Clone, Copy, PartialEq, Eq)]
enum Contents {
    RoundRock,
    CubeRock,
    Empty,
}

impl Contents {
    fn try_from_char(c: char) -> Result<Self, String> {
        match c {
            'O' => Ok(Self::RoundRock),
            '#' => Ok(Self::CubeRock),
            '.' => Ok(Self::Empty),
            _ => Err(format!("{c} is not a valid object")),
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
            O....#....
            O.OO#....#
            .....##...
            OO.#O....O
            .O.....O#.
            O.#..O.#.#
            ..O..#O..O
            .......O..
            #....###..
            #OO..#....
        "});
        let mut solver = Day14::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("136", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            O....#....
            O.OO#....#
            .....##...
            OO.#O....O
            .O.....O#.
            O.#..O.#.#
            ..O..#O..O
            .......O..
            #....###..
            #OO..#....
        "});
        let mut solver = Day14::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("64", result);

        Ok(())
    }
}
