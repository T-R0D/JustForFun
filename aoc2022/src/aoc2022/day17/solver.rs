use std::collections::HashMap;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day17 {
    jet_pattern: Vec<JetDirection>,
}

impl Day17 {
    pub fn new() -> Self {
        Self {
            jet_pattern: Vec::new(),
        }
    }
}

impl Solver for Day17 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        for c in input.trim().bytes() {
            self.jet_pattern.push(match c {
                b'<' => JetDirection::Left,
                b'>' => JetDirection::Right,
                _ => unreachable!("Unexpected char in input: {}", c),
            });
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut simulator = TunnelSimulator::new(&self.jet_pattern, 0);

        let pile_height = simulator.simulate_rock_fall(N_ROCKS);

        Ok(pile_height.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut rocks_to_simulate = ELEPHANT_N_ROCKS;

        let mut simulator = TunnelSimulator::new(&self.jet_pattern, 0);
        let cyclic_structure = simulator.find_cyclic_structure();

        rocks_to_simulate -= cyclic_structure.offset_n_rocks + cyclic_structure.cycle_n_rocks;

        let cycles_to_simulate = rocks_to_simulate / cyclic_structure.cycle_n_rocks;
        let skipped_simulation_height =
            cycles_to_simulate * cyclic_structure.cycle_height_increment;
        let remaining_rocks_to_simulate = rocks_to_simulate % cyclic_structure.cycle_n_rocks;

        let total_pile_height =
            simulator.simulate_rock_fall(remaining_rocks_to_simulate) + skipped_simulation_height;

        Ok(total_pile_height.to_string())
    }
}

const N_ROCK_KINDS: usize = 5;
const CHAMBER_WIDTH: usize = 7;
const FALL_START_HORIZONTAL_OFFSET: usize = 2;
const FALL_START_VERTICAL_OFFSET: usize = 3;
const N_ROCKS: usize = 2022;
const ELEPHANT_N_ROCKS: usize = 1_000_000_000_000;

struct CyclicStructure {
    offset_n_rocks: usize,
    cycle_height_increment: usize,
    cycle_n_rocks: usize,
}

struct CycleInfo {
    n_rocks_fallen: usize,
    pile_height: usize,
}

#[derive(Clone)]
enum JetDirection {
    Left,
    Right,
}

struct Vector(isize, isize);

#[derive(Clone)]
struct Coordinate(usize, usize);

impl Coordinate {
    fn add(&self, vector: &Vector) -> Coordinate {
        Coordinate(
            (self.0 as isize + vector.0) as usize,
            (self.1 as isize + vector.1) as usize,
        )
    }
}

#[derive(Clone)]
struct Rock {
    tag: u8,
    segments: Vec<Coordinate>,
    extreme_indices: [usize; 4],
}

impl Rock {
    fn new_bar(lower_left: &Coordinate) -> Self {
        let &Coordinate(i, j) = lower_left;
        Self {
            tag: 1,
            segments: vec![
                Coordinate(i, j),
                Coordinate(i, j + 1),
                Coordinate(i, j + 2),
                Coordinate(i, j + 3),
            ],
            extreme_indices: [0, 0, 0, 3],
        }
    }

    fn new_cross(lower_left: &Coordinate) -> Self {
        let &Coordinate(i, j) = lower_left;
        Self {
            tag: 2,
            segments: vec![
                Coordinate(i + 2, j + 1),
                Coordinate(i, j + 1),
                Coordinate(i + 1, j),
                Coordinate(i + 1, j + 2),
                Coordinate(i + 1, j + 1),
            ],
            extreme_indices: [0, 1, 2, 3],
        }
    }

    fn new_el(lower_left: &Coordinate) -> Self {
        let &Coordinate(i, j) = lower_left;
        Self {
            tag: 3,
            segments: vec![
                Coordinate(i, j),
                Coordinate(i, j + 1),
                Coordinate(i, j + 2),
                Coordinate(i + 1, j + 2),
                Coordinate(i + 2, j + 2),
            ],
            extreme_indices: [4, 0, 0, 3],
        }
    }
    fn new_beam(lower_left: &Coordinate) -> Self {
        let &Coordinate(i, j) = lower_left;
        Self {
            tag: 4,
            segments: vec![
                Coordinate(i, j),
                Coordinate(i + 1, j),
                Coordinate(i + 2, j),
                Coordinate(i + 3, j),
            ],
            extreme_indices: [3, 0, 0, 0],
        }
    }
    fn new_block(lower_left: &Coordinate) -> Self {
        let &Coordinate(i, j) = lower_left;
        Self {
            tag: 5,
            segments: vec![
                Coordinate(i, j),
                Coordinate(i + 1, j),
                Coordinate(i, j + 1),
                Coordinate(i + 1, j + 1),
            ],
            extreme_indices: [1, 0, 0, 2],
        }
    }

    fn top_height(&self) -> usize {
        self.segments[self.extreme_indices[0]].0 + 1
    }

    fn bottom_height(&self) -> usize {
        self.segments[self.extreme_indices[1]].0
    }

    fn leftmost_edge(&self) -> usize {
        self.segments[self.extreme_indices[2]].1
    }

    fn rightmost_edge(&self) -> usize {
        self.segments[self.extreme_indices[3]].1
    }

    fn project(&self, vector: &Vector) -> Rock {
        let mut projection = self.clone();
        projection.segments = projection
            .segments
            .iter()
            .map(|segment| segment.add(vector))
            .collect();
        projection
    }
}

struct TunnelSimulator {
    jet_pattern: Vec<JetDirection>,
    rock_factories: Vec<fn(&Coordinate) -> Rock>,
    chamber: Vec<Vec<u8>>,
    jet_cycle: usize,
    pile_height: usize,
    move_left: Vector,
    move_right: Vector,
    move_down: Vector,
}

const EMPTY: u8 = 0;

impl TunnelSimulator {
    fn new(jet_pattern: &Vec<JetDirection>, jet_start_cyle: usize) -> Self {
        Self {
            jet_pattern: jet_pattern.clone(),
            rock_factories: vec![
                |low| Rock::new_bar(low),
                |low| Rock::new_cross(low),
                |low| Rock::new_el(low),
                |low| Rock::new_beam(low),
                |low| Rock::new_block(low),
            ],
            chamber: Vec::from(vec![
                vec![EMPTY; CHAMBER_WIDTH];
                FALL_START_VERTICAL_OFFSET + 4 // 4 is max height of any rock.
            ]),
            jet_cycle: jet_start_cyle,
            pile_height: 0,
            move_left: Vector(0, -1),
            move_right: Vector(0, 1),
            move_down: Vector(-1, 0),
        }
    }

    fn simulate_rock_fall(&mut self, n_rocks: usize) -> usize {
        for i in 0..n_rocks {
            let rock = self.rock_factories[i % self.rock_factories.len()](&Coordinate(
                self.pile_height + FALL_START_VERTICAL_OFFSET,
                FALL_START_HORIZONTAL_OFFSET,
            ));

            self.simulate_one_rock_fall(&rock);
        }

        self.pile_height
    }

    fn find_cyclic_structure(&mut self) -> CyclicStructure {
        let mut jet_cycle_to_pile_heights: HashMap<usize, Vec<CycleInfo>> = HashMap::new();

        let mut n_rocks = 0;
        loop {
            for i in 0..N_ROCK_KINDS {
                let rock = self.rock_factories[i % self.rock_factories.len()](&Coordinate(
                    self.pile_height + FALL_START_VERTICAL_OFFSET,
                    FALL_START_HORIZONTAL_OFFSET,
                ));

                self.simulate_one_rock_fall(&rock);
            }
            n_rocks += N_ROCK_KINDS;

            let cycle_info = CycleInfo {
                n_rocks_fallen: n_rocks,
                pile_height: self.pile_height,
            };
            jet_cycle_to_pile_heights
                .entry(self.jet_cycle)
                .or_insert(Vec::new())
                .push(cycle_info);

            let infos = jet_cycle_to_pile_heights.get(&self.jet_cycle).unwrap();
            if infos.len() >= 2 {
                let CycleInfo {
                    pile_height: offset_height,
                    n_rocks_fallen: offset_n_rocks,
                } = infos[0];

                return CyclicStructure {
                    offset_n_rocks: offset_n_rocks,
                    cycle_height_increment: infos[1].pile_height - offset_height,
                    cycle_n_rocks: infos[1].n_rocks_fallen - offset_n_rocks,
                };
            }
        }
    }

    fn simulate_one_rock_fall(&mut self, starting_rock: &Rock) {
        let mut rock = starting_rock.clone();

        'rock_journey: loop {
            // Jets push rock.
            let jet_direction = &self.jet_pattern[self.jet_cycle];
            self.jet_cycle = (self.jet_cycle + 1) % self.jet_pattern.len();
            match jet_direction {
                JetDirection::Left => {
                    if rock.leftmost_edge() > 0 {
                        let projection = rock.project(&self.move_left);
                        if !self.rock_is_obstructed(&projection) {
                            rock = projection;
                        }
                    }
                }
                JetDirection::Right => {
                    if rock.rightmost_edge() < CHAMBER_WIDTH - 1 {
                        let projection = rock.project(&self.move_right);
                        if !self.rock_is_obstructed(&projection) {
                            rock = projection;
                        }
                    }
                }
            };

            // Gravity pulls rock.
            if rock.bottom_height() == 0 {
                break 'rock_journey;
            }

            let projection = rock.project(&self.move_down);
            if self.rock_is_obstructed(&projection) {
                break 'rock_journey;
            }

            rock = projection;
        }

        for &Coordinate(i, j) in rock.segments.iter() {
            self.chamber[i][j] = rock.tag;
        }

        if rock.top_height() > self.pile_height {
            let diff = rock.top_height() - self.pile_height;

            for _ in 0..diff {
                self.chamber.push(vec![EMPTY; CHAMBER_WIDTH]);
            }

            self.pile_height += diff;
        }
    }

    fn rock_is_obstructed(&self, rock: &Rock) -> bool {
        for &Coordinate(i, j) in rock.segments.iter() {
            if self.chamber[i][j] != EMPTY {
                return true;
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>");
        let mut solver = Day17::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("3068", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>");
        let mut solver = Day17::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("1514285714288", result);

        Ok(())
    }
}
