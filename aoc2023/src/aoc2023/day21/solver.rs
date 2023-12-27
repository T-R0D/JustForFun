use std::collections::{HashMap, HashSet, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day21 {
    garden: Vec<Vec<Feature>>,
    m: usize,
    n: usize,
    start: GridCoordinate,
}

impl Day21 {
    pub fn new() -> Self {
        Day21 {
            garden: vec![],
            m: 0,
            n: 0,
            start: GridCoordinate { i: 0, j: 0 },
        }
    }
}

impl Solver for Day21 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let mut start = GridCoordinate { i: 0, j: 0 };

        self.garden = input
            .trim()
            .lines()
            .enumerate()
            .map(|(i, line)| {
                line.chars()
                    .enumerate()
                    .map(|(j, c)| {
                        let feature = Feature::try_from_char(c)?;
                        if feature == Feature::Start {
                            start = GridCoordinate { i, j };
                        }
                        Ok(feature)
                    })
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        self.start = start;

        self.m = self.garden.len();
        self.n = self.garden[0].len();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let reachable_plots = self.find_reachable_plots(TARGET_STEPS);

        Ok(reachable_plots.len().to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let covered_plots = self.count_reachable_plots_in_infinity_garden(ACTUAL_TARGET_STEPS);

        Ok(covered_plots.to_string())
    }
}

impl Day21 {
    fn find_reachable_plots(&self, target_steps: usize) -> Vec<GridCoordinate> {
        let mut reachable_plots = Vec::<GridCoordinate>::new();

        let remainder = target_steps % 2;
        let mut frontier = VecDeque::<(GridCoordinate, usize)>::new();
        frontier.push_back((self.start.clone(), 0));
        let mut seen = HashSet::<GridCoordinate>::new();
        while let Some((coordinate, steps_taken)) = frontier.pop_front() {
            if steps_taken > target_steps {
                break;
            }

            if seen.contains(&coordinate) {
                continue;
            }
            seen.insert(coordinate);

            if self.garden[coordinate.i][coordinate.j] == Feature::Rock {
                continue;
            }

            if (steps_taken == 0 && steps_taken == remainder)
                || (steps_taken > 0 && steps_taken % 2 == remainder)
            {
                reachable_plots.push(coordinate);
            }

            let new_steps_taken = steps_taken + 1;
            if coordinate.i > 0 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i - 1,
                        j: coordinate.j,
                    },
                    new_steps_taken,
                ));
            }

            if coordinate.i < self.m - 1 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i + 1,
                        j: coordinate.j,
                    },
                    new_steps_taken,
                ));
            }

            if coordinate.j > 0 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i,
                        j: coordinate.j - 1,
                    },
                    new_steps_taken,
                ));
            }

            if coordinate.j < self.n - 1 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i,
                        j: coordinate.j + 1,
                    },
                    new_steps_taken,
                ));
            }
        }

        reachable_plots
    }

    fn count_reachable_plots_in_infinity_garden(&self, target_steps: usize) -> usize {
        let plot_to_distance_from_start = self.complete_tile_walk();

        let inner_diamond_distance = self.m / 2;

        let plots_per_even_tile = plot_to_distance_from_start
            .values()
            .filter(|&&distance| distance % 2 == 0)
            .count();
        let plots_per_even_tile_corners = plot_to_distance_from_start
            .values()
            .filter(|&&distance| distance > inner_diamond_distance && distance % 2 == 0)
            .count();
        let plots_per_odd_tile = plot_to_distance_from_start
            .values()
            .filter(|&&distance| distance % 2 == 1)
            .count();
        let plots_per_odd_tile_corners = plot_to_distance_from_start
            .values()
            .filter(|&&distance| distance > inner_diamond_distance && distance % 2 == 1)
            .count();

        let n = (target_steps - (self.m / 2)) / self.m;

        let even_tiles = if n % 2 == 0 { n * n } else { (n + 1) * (n + 1) };
        let odd_tiles = if n % 2 == 0 { (n + 1) * (n + 1) } else { n * n };
        let even_corners = if n % 2 == 0 { n } else { n + 1 };
        let odd_corners = if n % 2 == 0 { n + 1 } else { n };

        let mut plots_reachable = (even_tiles * plots_per_even_tile)
            + (odd_tiles * plots_per_odd_tile);

        if n % 2 == 0 {
            plots_reachable -= odd_corners * plots_per_odd_tile_corners;
            plots_reachable += even_corners * plots_per_even_tile_corners;
        } else {
            plots_reachable -= even_corners * plots_per_even_tile_corners;
            plots_reachable += odd_corners * plots_per_odd_tile_corners;
        }

        plots_reachable
    }

    fn complete_tile_walk(&self) -> HashMap<GridCoordinate, usize> {
        let mut reachable_plots = HashMap::<GridCoordinate, usize>::new();

        let mut frontier = VecDeque::<(GridCoordinate, usize)>::new();
        frontier.push_back((self.start.clone(), 0));
        let mut seen = HashSet::<GridCoordinate>::new();
        while let Some((coordinate, steps_taken)) = frontier.pop_front() {
            if seen.contains(&coordinate) {
                continue;
            }
            seen.insert(coordinate);

            if self.garden[coordinate.i][coordinate.j] == Feature::Rock {
                continue;
            }

            reachable_plots.insert(coordinate, steps_taken);

            let new_steps_taken = steps_taken + 1;
            if coordinate.i > 0 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i - 1,
                        j: coordinate.j,
                    },
                    new_steps_taken,
                ));
            }

            if coordinate.i < self.m - 1 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i + 1,
                        j: coordinate.j,
                    },
                    new_steps_taken,
                ));
            }

            if coordinate.j > 0 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i,
                        j: coordinate.j - 1,
                    },
                    new_steps_taken,
                ));
            }

            if coordinate.j < self.n - 1 {
                frontier.push_back((
                    GridCoordinate {
                        i: coordinate.i,
                        j: coordinate.j + 1,
                    },
                    new_steps_taken,
                ));
            }
        }

        reachable_plots
    }
}

const TARGET_STEPS: usize = 64;
const ACTUAL_TARGET_STEPS: usize = 26501365;

#[derive(Clone, Copy, PartialEq, Eq)]
enum Feature {
    Start,
    Garden,
    Rock,
}

impl Feature {
    fn try_from_char(c: char) -> Result<Self, String> {
        match c {
            'S' => Ok(Self::Start),
            '.' => Ok(Self::Garden),
            '#' => Ok(Self::Rock),
            _ => Err(format!("{} is not a valid map feature", c)),
        }
    }
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
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
            ...........
            .....###.#.
            .###.##..#.
            ..#.#...#..
            ....#.#....
            .##..S####.
            .##..#...#.
            .......##..
            .##.#.####.
            .##..##.##.
            ...........
        "});
        let mut solver = Day21::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.find_reachable_plots(6).len().to_string();

        // Assert.
        assert_eq!("16", result);

        Ok(())
    }
}
