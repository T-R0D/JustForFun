use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day11 {
    empty_cols: Vec<usize>,
    empty_rows: Vec<usize>,
    galaxies: Vec<GridCoordinate>,
}

impl Day11 {
    pub fn new() -> Self {
        Day11 {
            empty_cols: vec![],
            empty_rows: vec![],
            galaxies: vec![],
        }
    }
}

impl Solver for Day11 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let image = input
            .trim()
            .lines()
            .map(|line| line.chars().collect::<Vec<_>>())
            .collect::<Vec<Vec<_>>>();

        for i in 0..image.len() {
            let mut row_empty = true;
            for j in 0..image[0].len() {
                if image[i][j] == '#' {
                    row_empty = false;
                    self.galaxies.push(GridCoordinate { i, j });
                }
            }

            if row_empty {
                self.empty_rows.push(i);
            }
        }

        'col_processing: for j in 0..image[0].len() {
            for i in 0..image.len() {
                if image[i][j] != '.' {
                    continue 'col_processing;
                }
            }

            self.empty_cols.push(j);
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let galaxies_after_expansion =
            self.find_expanded_galactic_locations(DOUBLING_EXPANSION_FACTOR);

        let shortest_path_sum = Day11::sum_all_pairs_shortest_paths(&galaxies_after_expansion);

        Ok(shortest_path_sum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let galaxies_after_expansion = self.find_expanded_galactic_locations(MEGA_EXPANSION_FACTOR);

        let shortest_path_sum = Day11::sum_all_pairs_shortest_paths(&galaxies_after_expansion);

        Ok(shortest_path_sum.to_string())
    }
}

const DOUBLING_EXPANSION_FACTOR: usize = 2;
const MEGA_EXPANSION_FACTOR: usize = 1_000_000;

impl Day11 {
    fn find_expanded_galactic_locations(&self, expansion_factor: usize) -> Vec<GridCoordinate> {
        self.galaxies
            .iter()
            .map(|galaxy| {
                let empty_rows_before = self
                    .empty_rows
                    .iter()
                    .filter(|&&empty_i| empty_i < galaxy.i)
                    .count();
                let expansion_rows = if empty_rows_before > 0 {
                    empty_rows_before * (expansion_factor - 1)
                } else {
                    0
                };
                
                let empty_cols_before = self
                    .empty_cols
                    .iter()
                    .filter(|&&empty_j| empty_j < galaxy.j)
                    .count();
                let expansion_cols = if empty_cols_before > 0 {
                    empty_cols_before * (expansion_factor - 1)
                } else {
                    0
                };

                GridCoordinate {
                    i: galaxy.i + expansion_rows,
                    j: galaxy.j + expansion_cols,
                }
            })
            .collect::<Vec<_>>()
    }

    fn sum_all_pairs_shortest_paths(galaxies: &Vec<GridCoordinate>) -> usize {
        let mut shortest_path_sum = 0_usize;
        for i in 0..galaxies.len() {
            for j in (i + 1)..galaxies.len() {
                shortest_path_sum += Day11::galactic_manhattan_distance(&galaxies[i], &galaxies[j])
            }
        }

        shortest_path_sum
    }

    fn galactic_manhattan_distance(a: &GridCoordinate, b: &GridCoordinate) -> usize {
        (usize::max(a.i, b.i) - usize::min(a.i, b.i))
            + (usize::max(a.j, b.j) - usize::min(a.j, b.j))
    }
}

#[derive(Debug)]
struct GridCoordinate {
    i: usize,
    j: usize,
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
            ...#......
            .......#..
            #.........
            ..........
            ......#...
            .#........
            .........#
            ..........
            .......#..
            #...#.....
        "});
        let mut solver = Day11::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("374", result);

        Ok(())
    }

    #[test]
    fn solves_with_expansion_factor_10() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            ...#......
            .......#..
            #.........
            ..........
            ......#...
            .#........
            .........#
            ..........
            .......#..
            #...#.....
        "});
        let mut solver = Day11::new();
        solver.consume_input(input)?;

        // Act.
        let galaxies_after_expansion = solver.find_expanded_galactic_locations(10);
        let shortest_path_sum = Day11::sum_all_pairs_shortest_paths(&galaxies_after_expansion);
        let result = shortest_path_sum.to_string();

        // Assert.
        assert_eq!("1030", result);

        Ok(())
    }

    #[test]
    fn solves_with_expansion_factor_100() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            ...#......
            .......#..
            #.........
            ..........
            ......#...
            .#........
            .........#
            ..........
            .......#..
            #...#.....
        "});
        let mut solver = Day11::new();
        solver.consume_input(input)?;

        // Act.
        let galaxies_after_expansion = solver.find_expanded_galactic_locations(100);
        let shortest_path_sum = Day11::sum_all_pairs_shortest_paths(&galaxies_after_expansion);
        let result = shortest_path_sum.to_string();

        // Assert.
        assert_eq!("8410", result);

        Ok(())
    }
}
