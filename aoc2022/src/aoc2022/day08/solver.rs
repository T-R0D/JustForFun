use std::vec;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day08 {
    tree_grid: Vec<Vec<i8>>,
    n_rows: usize,
    n_cols: usize,
}

impl Day08 {
    pub fn new() -> Self {
        Day08 {
            tree_grid: vec![],
            n_rows: 0,
            n_cols: 0,
        }
    }
}

impl Solver for Day08 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let lines = input.trim().lines().collect::<Vec<_>>();

        let (n_rows, n_cols) = (lines.len(), lines.first().unwrap().len());

        let mut grid: Vec<Vec<i8>> = Vec::with_capacity(n_rows);
        for i in 0..n_rows {
            let mut row: Vec<i8> = Vec::with_capacity(n_cols);
            let line = lines[i].chars().collect::<Vec<char>>();
            for j in 0..n_cols {
                row.push(line[j].to_digit(10).unwrap() as i8);
            }
            grid.push(row);
        }

        self.tree_grid = grid;
        self.n_rows = n_rows;
        self.n_cols = n_cols;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut visibilities: Vec<Vec<VisibilityItem>> = Vec::with_capacity(self.n_rows);
        for i in 0..self.n_rows {
            let mut row: Vec<VisibilityItem> = Vec::with_capacity(self.n_cols);
            for j in 0..self.n_cols {
                let visibility = VisibilityItem::new(self.tree_grid[i][j]);
                row.push(visibility);
            }
            visibilities.push(row);
        }

        for i in 0..self.n_rows {
            let mut current_height = -1;
            for j in 0..self.n_cols {
                let mut visibility = &mut visibilities[i][j];
                visibility.west_height = current_height;
                current_height = std::cmp::max(current_height, visibility.this_height);
            }
        }

        for i in 0..self.n_rows {
            let mut current_height = -1;
            for j in (0..self.n_cols).rev() {
                let mut visibility = &mut visibilities[i][j];
                visibility.east_height = current_height;
                current_height = std::cmp::max(current_height, visibility.this_height);
            }
        }

        for j in 0..self.n_cols {
            let mut current_height = -1;
            for i in 0..self.n_rows {
                let mut visibility = &mut visibilities[i][j];
                visibility.north_height = current_height;
                current_height = std::cmp::max(current_height, visibility.this_height);
            }
        }

        for j in 0..self.n_cols {
            let mut current_height = -1;
            for i in (0..self.n_rows).rev() {
                let mut visibility = &mut visibilities[i][j];
                visibility.south_height = current_height;
                current_height = std::cmp::max(current_height, visibility.this_height);
            }
        }

        let mut n_visible = 0;
        for i in 0..self.n_rows {
            for j in 0..self.n_cols {
                if visibilities[i][j].is_visible() {
                    n_visible += 1;
                }
            }
        }

        Ok(n_visible.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut best_scenic_score: u32 = 0;

        for i in 1..self.n_rows - 1 {
            for j in 1..self.n_cols - 1 {
                let this_height = self.tree_grid[i][j];

                let mut n_trees_north = 0;
                for i2 in (0..i).rev() {
                    let that_height = self.tree_grid[i2][j];
                    n_trees_north += 1;
                    if that_height >= this_height {
                        break;
                    }
                }

                let mut n_trees_east = 0;
                for j2 in j + 1..self.n_cols {
                    let that_height = self.tree_grid[i][j2];
                    n_trees_east += 1;
                    if that_height >= this_height {
                        break;
                    }
                }

                let mut n_trees_south = 0;
                for i2 in i + 1..self.n_rows {
                    let that_height = self.tree_grid[i2][j];
                    n_trees_south += 1;
                    if that_height >= this_height {
                        break;
                    }
                }

                let mut n_trees_west = 0;
                for j2 in (0..j).rev() {
                    let that_height = self.tree_grid[i][j2];
                    n_trees_west += 1;
                    if that_height >= this_height {
                        break;
                    }
                }

                let scenic_score = n_trees_north * n_trees_east * n_trees_south * n_trees_west;
                best_scenic_score = std::cmp::max(best_scenic_score, scenic_score);
            }
        }

        Ok(best_scenic_score.to_string())
    }
}

struct VisibilityItem {
    north_height: i8,
    east_height: i8,
    south_height: i8,
    west_height: i8,
    this_height: i8,
}

impl VisibilityItem {
    fn new(this_height: i8) -> Self {
        VisibilityItem {
            north_height: -1,
            east_height: -1,
            south_height: -1,
            west_height: -1,
            this_height: this_height,
        }
    }

    fn is_visible(&self) -> bool {
        return self.this_height > self.north_height
            || self.this_height > self.east_height
            || self.this_height > self.south_height
            || self.this_height > self.west_height;
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
            30373
            25512
            65332
            33549
            35390
        "});
        let mut solver = Day08::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("21", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            30373
            25512
            65332
            33549
            35390
        "});
        let mut solver = Day08::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("8", result);

        Ok(())
    }
}
