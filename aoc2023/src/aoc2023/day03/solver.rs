use std::collections::HashMap;

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day03 {
    schematic_image: Vec<Vec<char>>,
    schematic_numbers: Vec<SchematicNumber>,
    symbol_locations: HashMap<(usize, usize), char>,
}

impl Day03 {
    pub fn new() -> Self {
        Day03 {
            schematic_image: vec![],
            schematic_numbers: vec![],
            symbol_locations: HashMap::new(),
        }
    }
}

impl Solver for Day03 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.schematic_image = input
            .trim()
            .lines()
            .map(|line| line.chars().collect::<Vec<_>>())
            .collect::<Vec<_>>();

        let mut schematic_numbers = Vec::<SchematicNumber>::new();
        let mut symbol_locations = HashMap::<(usize, usize), char>::new();
        for (i, row) in self.schematic_image.iter().enumerate() {
            let mut start_coords: Option<(usize, usize)> = None;
            let mut running_number: u32 = 0;
            'col_iter: for (j, c) in row.iter().enumerate() {
                if let Some(digit) = c.to_digit(10) {
                    if start_coords.is_none() {
                        start_coords = Some((i, j));
                    }

                    running_number *= 10;
                    running_number += digit;
                    continue 'col_iter;
                }

                if let Some((i2, j2)) = start_coords {
                    schematic_numbers.push(SchematicNumber {
                        start_loc: (i2, j2),
                        end_loc: (i2, j - 1),
                        value: running_number,
                    });
                    start_coords = None;
                    running_number = 0;
                }

                if *c != '.' {
                    symbol_locations.insert((i, j), *c);
                }
            }

            if let Some((i2, j2)) = start_coords {
                schematic_numbers.push(SchematicNumber {
                    start_loc: (i2, j2),
                    end_loc: (i2, row.len() - 1),
                    value: running_number,
                });
            }
        }

        self.schematic_numbers = schematic_numbers;
        self.symbol_locations = symbol_locations;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let part_number_sum = self
            .schematic_numbers
            .iter()
            .filter_map(|schematic_number| {
                let number_row = schematic_number.start_loc.0;
                let number_start_col = schematic_number.start_loc.1;
                let number_end_col = schematic_number.end_loc.1;
                let ((i_start, j_start), (i_end, j_end)) = schematic_number.bounding_box(self.schematic_image.len(), self.schematic_image[0].len());

                for i in i_start..=i_end {
                    'col_iter: for j in j_start..=j_end {
                        if i == number_row && (number_start_col <= j && j <= number_end_col) {
                            continue 'col_iter;
                        }

                        if self.symbol_locations.contains_key(&(i, j)) {
                            // println!("{} is a part number", schematic_number.value);
                            return Some(schematic_number.value);
                        }
                    }
                }

                // println!("{} is NOT a part number", schematic_number.value);
                None
            })
            .sum::<u32>();

        Ok(part_number_sum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let gear_symbol_locations = self
            .symbol_locations
            .iter()
            .filter_map(|(loc, symbol)| {
                if *symbol == '*' {
                    return Some(loc);
                }
                None
            })
            .collect::<Vec<_>>();

        let mut gear_ratio_sum = 0;
        for (gear_i, gear_j) in gear_symbol_locations.iter() {
            let touching_part_numbers = self
                .schematic_numbers
                .iter()
                .filter_map(|schematic_number| {
                    let ((i, j), (i2, j2)) = schematic_number
                        .bounding_box(self.schematic_image.len(), self.schematic_image[0].len());

                    if (i <= *gear_i && *gear_i <= i2) && (j <= *gear_j && *gear_j <= j2) {
                        return Some(schematic_number.value);
                    }

                    None
                })
                .collect::<Vec<_>>();

            if touching_part_numbers.len() == 2 {
                gear_ratio_sum += touching_part_numbers.iter().product::<u32>();
            }
        }

        Ok(gear_ratio_sum.to_string())
    }
}

struct SchematicNumber {
    start_loc: (usize, usize),
    end_loc: (usize, usize),
    value: u32,
}

impl SchematicNumber {
    fn bounding_box(
        &self,
        schematic_height: usize,
        schematic_width: usize,
    ) -> ((usize, usize), (usize, usize)) {
        let number_row = self.start_loc.0;
        let number_start_col = self.start_loc.1;
        let number_end_col = self.end_loc.1;

        let first_row = if number_row > 0 { number_row - 1 } else { 0 };
        let last_row = if number_row < schematic_height - 1 {
            number_row + 1
        } else {
            number_row
        };
        let first_col = if number_start_col > 0 {
            number_start_col - 1
        } else {
            number_start_col
        };
        let last_col = if number_end_col < schematic_width - 1 {
            number_end_col + 1
        } else {
            number_end_col
        };

        ((first_row, first_col), (last_row, last_col))
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
            467..114..
            ...*......
            ..35..633.
            ......#...
            617*......
            .....+.58.
            ..592.....
            ......755.
            ...$.*....
            .664.598..
        "});
        let mut solver = Day03::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("4361", result);

        Ok(())
    }

    #[test]
    fn part_one_with_a_last_col_number_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            467..114..
            ...*......
            ..35..633.
            ......#...
            ........$1
            617*......
            .....+.58.
            ..592.....
            ......755.
            ...$.*....
            .664.598..
        "});
        let mut solver = Day03::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("4362", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            467..114..
            ...*......
            ..35..633.
            ......#...
            617*......
            .....+.58.
            ..592.....
            ......755.
            ...$.*....
            .664.598..
        "});
        let mut solver = Day03::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("467835", result);

        Ok(())
    }
}
