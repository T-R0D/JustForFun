use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day13 {
    grids: Vec<AshAndRockGrid>,
}

impl Day13 {
    pub fn new() -> Self {
        Day13 { grids: vec![] }
    }
}

impl Solver for Day13 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.grids = input
            .trim()
            .split("\n\n")
            .map(AshAndRockGrid::try_from_block)
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let total_score = self
            .grids
            .iter()
            .map(|grid| {
                if let Some(vertical_symmetry_score) = grid.find_vertical_symmetry(0) {
                    return vertical_symmetry_score;
                }

                if let Some(horizontal_symmetry_score) = grid.find_horizontal_symmetry(0) {
                    return horizontal_symmetry_score * 100;
                }

                0
            })
            .sum::<usize>();

        Ok(total_score.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let total_score = self
            .grids
            .iter()
            .map(|grid| {
                if let Some(vertical_symmetry_score) = grid.find_vertical_symmetry(1) {
                    return vertical_symmetry_score;
                }

                if let Some(horizontal_symmetry_score) = grid.find_horizontal_symmetry(1) {
                    return horizontal_symmetry_score * 100;
                }

                0
            })
            .sum::<usize>();

        Ok(total_score.to_string())
    }
}

struct AshAndRockGrid {
    data: Vec<Vec<Feature>>,
    m: usize,
    n: usize,
}

impl AshAndRockGrid {
    fn try_from_block(block: &str) -> Result<Self, String> {
        let data = block
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(Feature::try_from_char)
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        let m = data.len();
        let n = data[0].len();

        Ok(Self { data, m, n })
    }

    fn find_vertical_symmetry(&self, expected_smudges: usize) -> Option<usize> {
        'next_line: for j1 in 0..(self.n - 1) {
            let j2 = j1 + 1;
            let mut smudges = 0_usize;
            for delta in 0..(usize::min(j1 + 1, self.n - j2)) {
                for i in 0..self.m {
                    if self.data[i][j1 - delta] != self.data[i][j2 + delta] {
                        smudges += 1;

                        if smudges > expected_smudges {
                            continue 'next_line;
                        }
                    }
                }
            }

            if smudges == expected_smudges {
                return Some(j1 + 1);
            }
        }
        None
    }

    fn find_horizontal_symmetry(&self, expected_smudges: usize) -> Option<usize> {
        'next_line: for i1 in 0..(self.m - 1) {
            let i2 = i1 + 1;
            let mut smudges = 0_usize;
            for delta in 0..(usize::min(i1 + 1, self.m - i2)) {
                for j in 0..self.n {
                    if self.data[i1 - delta][j] != self.data[i2 + delta][j] {
                        smudges += 1;

                        if smudges > expected_smudges {
                            continue 'next_line;
                        }
                    }
                }
            }

            if smudges == expected_smudges {
                return Some(i1 + 1);
            }
        }
        None
    }
}

#[derive(PartialEq, Eq)]
enum Feature {
    Ash,
    Rock,
}

impl Feature {
    fn try_from_char(c: char) -> Result<Self, String> {
        match c {
            '.' => Ok(Self::Ash),
            '#' => Ok(Self::Rock),
            _ => Err(format!("{c} is not a landscape feature")),
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
            #.##..##.
            ..#.##.#.
            ##......#
            ##......#
            ..#.##.#.
            ..##..##.
            #.#.##.#.

            #...##..#
            #....#..#
            ..##..###
            #####.##.
            #####.##.
            ..##..###
            #....#..#
        "});
        let mut solver = Day13::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("405", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            #.##..##.
            ..#.##.#.
            ##......#
            ##......#
            ..#.##.#.
            ..##..##.
            #.#.##.#.

            #...##..#
            #....#..#
            ..##..###
            #####.##.
            #####.##.
            ..##..###
            #....#..#
        "});
        let mut solver = Day13::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("400", result);

        Ok(())
    }
}
