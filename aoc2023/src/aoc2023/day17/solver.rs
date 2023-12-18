use std::{
    cmp::Ordering,
    collections::{BinaryHeap, HashSet},
};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day17 {
    heat_loss_map: HeatLossMap,
}

impl Day17 {
    pub fn new() -> Self {
        Day17 {
            heat_loss_map: HeatLossMap {
                data: vec![],
                m: 0,
                n: 0,
            },
        }
    }
}

impl Solver for Day17 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.heat_loss_map = HeatLossMap::try_from_block(input)?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let heat_loss = self
            .heat_loss_map
            .find_lowest_loss_path_loss(MIN_CRUCIBLE_RUN_LENGTH, MAX_CRUCIBLE_RUN_LENGTH);

        Ok(heat_loss.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let heat_loss = self.heat_loss_map.find_lowest_loss_path_loss(
            MIN_ULTRA_CRUCIBLE_RUN_LENGTH,
            MAX_ULTRA_CRUCIBLE_RUN_LENGTH,
        );

        Ok(heat_loss.to_string())
    }
}

const MIN_CRUCIBLE_RUN_LENGTH: usize = 0;
const MAX_CRUCIBLE_RUN_LENGTH: usize = 3;

const MIN_ULTRA_CRUCIBLE_RUN_LENGTH: usize = 4;
const MAX_ULTRA_CRUCIBLE_RUN_LENGTH: usize = 10;

struct HeatLossMap {
    data: Vec<Vec<u32>>,
    m: usize,
    n: usize,
}

impl HeatLossMap {
    fn try_from_block(block: &String) -> Result<Self, String> {
        let data = block
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(|c| c.to_digit(10).ok_or(format!("{c} is not a numeric value")))
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;
        let m = data.len();
        let n = data[0].len();

        Ok(Self { data, m, n })
    }

    fn find_lowest_loss_path_loss(&self, min_run_len: usize, max_run_len: usize) -> u32 {
        let mut heat_loss_traversals = vec![vec![u32::MAX; self.n]; self.m];
        heat_loss_traversals[0][0] = 0;
        let mut frontier = BinaryHeap::<SearchState>::new();
        frontier.push(SearchState {
            i: 0,
            j: 0,
            run: 0,
            heading: Heading::East,
            heat_loss_so_far: 0,
        });
        frontier.push(SearchState {
            i: 0,
            j: 0,
            run: 0,
            heading: Heading::South,
            heat_loss_so_far: 0,
        });

        let mut seen = HashSet::<(usize, usize, usize, usize)>::new();

        let move_generators: [Box<dyn Fn(&SearchState) -> Option<SearchState>>;
            Heading::NHeadings.index()] = [
            self.new_try_next_move_fn(&NextMoveFnSpec {
                min_run_len,
                max_run_len,
                target_heading: Heading::North,
            }),
            self.new_try_next_move_fn(&NextMoveFnSpec {
                min_run_len,
                max_run_len,
                target_heading: Heading::South,
            }),
            self.new_try_next_move_fn(&NextMoveFnSpec {
                min_run_len,
                max_run_len,
                target_heading: Heading::East,
            }),
            self.new_try_next_move_fn(&NextMoveFnSpec {
                min_run_len,
                max_run_len,
                target_heading: Heading::West,
            }),
        ];

        while let Some(state) = frontier.pop() {
            if heat_loss_traversals[state.i][state.j] > state.heat_loss_so_far {
                heat_loss_traversals[state.i][state.j] = state.heat_loss_so_far;
            }

            if state.i == self.m - 1 && state.j == self.n - 1 {
                break;
            }

            if seen.contains(&(state.i, state.j, state.heading.index(), state.run)) {
                continue;
            }

            move_generators
                .iter()
                .filter_map(|try_next_state| try_next_state(&state))
                .for_each(|new_state| frontier.push(new_state));

            seen.insert((state.i, state.j, state.heading.index(), state.run));
        }

        heat_loss_traversals[self.m - 1][self.n - 1]
    }

    fn new_try_next_move_fn<'a>(
        &'a self,
        spec: &NextMoveFnSpec,
    ) -> Box<dyn Fn(&SearchState) -> Option<SearchState> + 'a> {
        let NextMoveFnSpec {
            min_run_len,
            max_run_len,
            target_heading,
        } = spec.clone();

        Box::new(move |state: &SearchState| {
            let i = match target_heading {
                Heading::North => {
                    if state.i < 1 {
                        return None;
                    }
                    state.i - 1
                }
                Heading::South => {
                    if state.i >= self.m - 1 {
                        return None;
                    }
                    state.i + 1
                }
                _ => state.i,
            };

            let j = match target_heading {
                Heading::West => {
                    if state.j < 1 {
                        return None;
                    }
                    state.j - 1
                }
                Heading::East => {
                    if state.j >= self.n - 1 {
                        return None;
                    }
                    state.j + 1
                }

                _ => state.j,
            };

            if state.heading == target_heading.opposite() {
                return None;
            }

            if state.heading != target_heading && state.run < min_run_len {
                return None;
            }

            if state.heading == target_heading && state.run == max_run_len {
                return None;
            }

            let run = if state.heading == target_heading {
                state.run + 1
            } else {
                1
            };
            let heat_loss_so_far = state.heat_loss_so_far + self.data[i][j];
            Some(SearchState {
                i,
                j,
                run,
                heading: target_heading,
                heat_loss_so_far,
            })
        })
    }
}

#[derive(Clone, Copy, PartialEq, Eq)]
struct SearchState {
    i: usize,
    j: usize,
    run: usize,
    heading: Heading,
    heat_loss_so_far: u32,
}

impl Ord for SearchState {
    fn cmp(&self, other: &Self) -> Ordering {
        other
            .heat_loss_so_far
            .cmp(&self.heat_loss_so_far)
            .then_with(|| other.run.cmp(&self.run))
    }
}

impl PartialOrd for SearchState {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(Clone, Copy, PartialEq, Eq)]
enum Heading {
    North,
    South,
    East,
    West,
    NHeadings,
}

impl Heading {
    const fn index(&self) -> usize {
        *self as usize
    }

    const fn opposite(&self) -> Heading {
        match self {
            Self::North => Self::South,
            Self::South => Self::North,
            Self::East => Self::West,
            Self::West => Self::East,
            Self::NHeadings => Self::NHeadings,
        }
    }
}

#[derive(Clone, Copy)]
struct NextMoveFnSpec {
    min_run_len: usize,
    max_run_len: usize,
    target_heading: Heading,
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
            2413432311323
            3215453535623
            3255245654254
            3446585845452
            4546657867536
            1438598798454
            4457876987766
            3637877979653
            4654967986887
            4564679986453
            1224686865563
            2546548887735
            4322674655533
        "});
        let mut solver = Day17::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("102", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            2413432311323
            3215453535623
            3255245654254
            3446585845452
            4546657867536
            1438598798454
            4457876987766
            3637877979653
            4654967986887
            4564679986453
            1224686865563
            2546548887735
            4322674655533
        "});
        let mut solver = Day17::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("94", result);

        Ok(())
    }
}
