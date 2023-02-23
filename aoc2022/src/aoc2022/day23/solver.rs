use std::collections::{HashMap, HashSet};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day23 {
    elf_positions: HashSet<Coordinate>,
}

impl Day23 {
    pub fn new() -> Self {
        Self {
            elf_positions: HashSet::new(),
        }
    }
}

impl Solver for Day23 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        for (i, line) in input.trim().lines().enumerate() {
            for (j, c) in line.bytes().enumerate() {
                if c == ELF {
                    self.elf_positions.insert(Coordinate {
                        i: i as i32,
                        j: j as i32,
                    });
                }
            }
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let n_elves = self.elf_positions.len();
        let final_positions = simulate_elf_movement(&self.elf_positions, N_ROUNDS);
        let (upper_left, lower_right) = find_bounding_box(&final_positions);
        let empty_plots = count_empty_plots(&upper_left, &lower_right, n_elves);

        Ok(empty_plots.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        // 920 too high.
        // 921 too high. Also, someone else's answer...

        let rounds_required = simulate_elf_movement_until_rest(&self.elf_positions);

        Ok(rounds_required.to_string())
    }
}

const ELF: u8 = b'#';

const N_ROUNDS: usize = 10;

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
struct Coordinate {
    i: i32,
    j: i32,
}

#[derive(Debug)]
enum Dir {
    North,
    South,
    West,
    East,
}

const DIR_CHOICE: [Dir; 4] = [Dir::North, Dir::South, Dir::West, Dir::East];
const NEIGHBOR_GROUP_INDICES: [[usize; 3]; 4] = [[0, 1, 2], [6, 7, 8], [0, 3, 6], [2, 5, 8]];

fn simulate_elf_movement(
    starting_positions: &HashSet<Coordinate>,
    n_rounds: usize,
) -> HashSet<Coordinate> {
    let mut current_positions = starting_positions.clone();

    for r in 0..n_rounds {
        (current_positions, _) = simulate_one_round_of_elf_movement(&current_positions, r);
    }

    current_positions
}

fn simulate_elf_movement_until_rest(starting_positions: &HashSet<Coordinate>) -> usize {
    let mut current_positions = starting_positions.clone();
    let mut n_moving_elves;

    for r in 0..usize::MAX {
        (current_positions, n_moving_elves) =
            simulate_one_round_of_elf_movement(&current_positions, r);

        if n_moving_elves == 0 {
            return r + 1;
        }
    }

    usize::MAX
}

fn simulate_one_round_of_elf_movement(
    current_positions: &HashSet<Coordinate>,
    round_num: usize,
) -> (HashSet<Coordinate>, usize) {
    // Phase 1: Propose.
    let mut next_positions = HashSet::<Coordinate>::new();
    let mut proposed_position_to_proposers = HashMap::<Coordinate, Vec<Coordinate>>::new();
    let first_dir_choice = round_num % DIR_CHOICE.len();

    'elf_proposals: for elf in current_positions.iter() {
        let neighbors = find_neighbors(elf, current_positions);

        if neighbors.iter().all(|neighbor| neighbor.is_none()) {
            next_positions.insert(*elf);
            continue 'elf_proposals;
        }

        let mut proposal: Option<Coordinate> = None;
        'find_proposal_direction: for dir_choice in
            first_dir_choice..first_dir_choice + DIR_CHOICE.len()
        {
            let dir_index = dir_choice % DIR_CHOICE.len();

            if NEIGHBOR_GROUP_INDICES[dir_index]
                .iter()
                .any(|&check_index| neighbors[check_index].is_some())
            {
                continue 'find_proposal_direction;
            }

            proposal = Some(match DIR_CHOICE[dir_index] {
                Dir::North => Coordinate {
                    i: elf.i - 1,
                    j: elf.j,
                },
                Dir::South => Coordinate {
                    i: elf.i + 1,
                    j: elf.j,
                },
                Dir::West => Coordinate {
                    i: elf.i,
                    j: elf.j - 1,
                },
                Dir::East => Coordinate {
                    i: elf.i,
                    j: elf.j + 1,
                },
            });
            break 'find_proposal_direction;
        }

        if proposal.is_none() {
            next_positions.insert(*elf);
            continue 'elf_proposals;
        }

        proposed_position_to_proposers
            .entry(proposal.unwrap())
            .or_insert(Vec::new())
            .push(*elf);
    }

    // Phase 2: Execute move.
    let mut n_moving_elves = 0;
    for (&proposal, proposers) in proposed_position_to_proposers.iter() {
        if proposers.len() == 1 {
            next_positions.insert(proposal);
            n_moving_elves += 1;
            continue;
        }

        for &failed_proposer in proposers.iter() {
            next_positions.insert(failed_proposer);
        }
    }

    (next_positions, n_moving_elves)
}

fn find_neighbors(
    center: &Coordinate,
    elf_positions: &HashSet<Coordinate>,
) -> [Option<Coordinate>; 9] {
    let mut neighbors: [Option<Coordinate>; 9] = [None; 9];

    for i in -1..=1 {
        for j in -1..=1 {
            if i == 0 && j == 0 {
                continue;
            }

            let candidate = Coordinate {
                i: center.i + i,
                j: center.j + j,
            };
            if elf_positions.contains(&candidate) {
                neighbors[(((i + 1) * 3) + (j + 1)) as usize] = Some(candidate)
            }
        }
    }

    neighbors
}

fn find_bounding_box(elf_positions: &HashSet<Coordinate>) -> (Coordinate, Coordinate) {
    let mut min_i = i32::MAX;
    let mut min_j = i32::MAX;
    let mut max_i = i32::MIN;
    let mut max_j = i32::MIN;

    for &Coordinate { i, j } in elf_positions.iter() {
        if i < min_i {
            min_i = i;
        } else if i > max_i {
            max_i = i;
        }

        if j < min_j {
            min_j = j;
        } else if j > max_j {
            max_j = j;
        }
    }

    (
        Coordinate { i: min_i, j: min_j },
        Coordinate { i: max_i, j: max_j },
    )
}

fn count_empty_plots(upper_left: &Coordinate, lower_right: &Coordinate, n_elves: usize) -> usize {
    let total_plots = ((lower_right.i + 1) - upper_left.i) * ((lower_right.j + 1) - upper_left.j);
    total_plots as usize - n_elves
}

// fn print_current_map(
//     elf_positions: &HashSet<Coordinate>,
//     prev: &HashSet<Coordinate>,
//     new: &HashSet<Coordinate>,
// ) {
//     let (upper_left, lower_right) = find_bounding_box(elf_positions);
//     let i_offset = 0 - upper_left.i;
//     let j_offset = 0 - upper_left.j;

//     let mut s = String::new();
//     for i in 0..=lower_right.i + i_offset {
//         for j in 0..=lower_right.j + j_offset {
//             let candidate = Coordinate {
//                 i: i - i_offset,
//                 j: j - j_offset,
//             };
//             if prev.contains(&candidate) {
//                 s.push('S');
//             } else if new.contains(&candidate) {
//                 s.push('#');
//             } else if elf_positions.contains(&candidate) {
//                 s.push('.');
//             } else {
//                 s.push(' ');
//             }
//         }
//         s.push('\n');
//     }

//     println!("{s}");
// }

#[cfg(test)]
mod tests {
    use super::*;
    #[cfg(test)]
    use indoc::indoc;
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    #[test]
    fn part_one_solves_tiny_example_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            .....
            ..##.
            ..#..
            .....
            ..##.
            .....
        "});
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("25", result);

        Ok(())
    }

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            ....#..
            ..###.#
            #...#.#
            .#...##
            #.###..
            ##.#.##
            .#..#..
        "});
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("110", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_tiny_example_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            .....
            ..##.
            ..#..
            .....
            ..##.
            .....
        "});
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("4", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            ....#..
            ..###.#
            #...#.#
            .#...##
            #.###..
            ##.#.##
            .#..#..
        "});
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("20", result);

        Ok(())
    }
}
