use std::collections::{HashSet, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day10 {
    start: GridCoordinate,
    pipe_map: Vec<Vec<PipeSegment>>,
    map_height: usize,
    map_width: usize,
}

impl Day10 {
    pub fn new() -> Self {
        Day10 {
            start: GridCoordinate { i: 0, j: 0 },
            pipe_map: vec![],
            map_height: 0,
            map_width: 0,
        }
    }
}

impl Solver for Day10 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.pipe_map = input
            .trim()
            .lines()
            .enumerate()
            .map(|(i, line)| {
                line.chars()
                    .enumerate()
                    .map(|(j, c)| {
                        let segment = PipeSegment::try_from_char(c)?;

                        if segment == PipeSegment::Start {
                            self.start = GridCoordinate { i, j };
                        }

                        Ok(segment)
                    })
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        self.map_height = self.pipe_map.len();
        self.map_width = self.pipe_map[0].len();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut next_direction = self.find_start_direction();
        let mut current_location = self.start.clone();
        let mut steps = 0;
        loop {
            current_location = next_direction.try_next_coordinates(
                &current_location,
                self.map_height,
                self.map_width,
            )?;
            steps += 1;
            if current_location == self.start {
                break;
            }
            next_direction = self.pipe_map[current_location.i][current_location.j]
                .try_transform_direction(&next_direction)?;
        }

        let furthest_point_in_steps = if steps % 2 == 0 {
            steps / 2
        } else {
            (steps / 2) + 1
        };

        Ok(furthest_point_in_steps.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut spaced_grid = SpacedGrid::from_compact_dimensions(self.map_height, self.map_width);
        let mut next_direction = self.find_start_direction();
        let mut compact_location = self.start.clone();
        let mut spaced_location = SpacedGrid::compact_coord_to_spaced_coord(&self.start);

        spaced_grid.mark_loop_location(&spaced_location);
        loop {
            compact_location = next_direction.try_next_coordinates(
                &compact_location,
                self.map_height,
                self.map_width,
            )?;

            for _ in 0..2 {
                spaced_location = next_direction.try_next_coordinates(
                    &spaced_location,
                    spaced_grid.height,
                    spaced_grid.width,
                )?;
                spaced_grid.mark_loop_location(&spaced_location);
            }

            if compact_location == self.start {
                break;
            }

            next_direction = self.pipe_map[compact_location.i][compact_location.j]
                .try_transform_direction(&next_direction)?;
        }

        spaced_grid.mark_unenclosed_tiles();

        let n_enclosed_tiles = spaced_grid.count_enclosed_tiles();

        Ok(n_enclosed_tiles.to_string())
    }
}

impl Day10 {
    fn find_start_direction(&self) -> Direction {
        let directions = [
            Direction::North,
            Direction::South,
            Direction::East,
            Direction::West,
        ];

        for &direction in directions.iter() {
            match direction.try_next_coordinates(&self.start, self.map_height, self.map_width) {
                Ok(coordinates) => {
                    let next_segment = &self.pipe_map[coordinates.i][coordinates.j];
                    if DIRECTION_TO_POSSIBLE_SEGMENTS[direction.index()]
                        .iter()
                        .any(|possible_segment| next_segment == possible_segment)
                    {
                        return direction;
                    }
                }
                Err(_) => continue,
            }
        }

        unreachable!();
    }
}

const DIRECTION_TO_POSSIBLE_SEGMENTS: [[PipeSegment; 3]; 4] = [
    [
        PipeSegment::SouthAndWest,
        PipeSegment::NorthAndSouth,
        PipeSegment::SouthAndEast,
    ], // North
    [
        PipeSegment::NorthAndWest,
        PipeSegment::NorthAndSouth,
        PipeSegment::NorthAndEast,
    ], // South
    [
        PipeSegment::NorthAndWest,
        PipeSegment::EastAndWest,
        PipeSegment::SouthAndWest,
    ], // East
    [
        PipeSegment::NorthAndEast,
        PipeSegment::EastAndWest,
        PipeSegment::SouthAndEast,
    ], // West
];

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
struct GridCoordinate {
    i: usize,
    j: usize,
}

#[derive(Clone, Copy, PartialEq)]
enum Direction {
    North,
    South,
    East,
    West,
}

impl Direction {
    pub fn index(&self) -> usize {
        *self as usize
    }

    pub fn try_next_coordinates(
        &self,
        current: &GridCoordinate,
        grid_height: usize,
        grid_width: usize,
    ) -> Result<GridCoordinate, String> {
        match self {
            Direction::North => {
                if current.i == 0 {
                    Err(String::from("can't move North"))
                } else {
                    Ok(GridCoordinate {
                        i: current.i - 1,
                        j: current.j,
                    })
                }
            }
            Direction::South => {
                if current.i == grid_height - 1 {
                    Err(String::from("can't move South"))
                } else {
                    Ok(GridCoordinate {
                        i: current.i + 1,
                        j: current.j,
                    })
                }
            }
            Direction::East => {
                if current.j == grid_width - 1 {
                    Err(String::from("can't move East"))
                } else {
                    Ok(GridCoordinate {
                        i: current.i,
                        j: current.j + 1,
                    })
                }
            }
            Direction::West => {
                if current.j == 0 {
                    Err(String::from("can't move West"))
                } else {
                    Ok(GridCoordinate {
                        i: current.i,
                        j: current.j - 1,
                    })
                }
            }
        }
    }
}

#[derive(PartialEq)]
enum PipeSegment {
    Start,
    NorthAndSouth,
    EastAndWest,
    NorthAndEast,
    NorthAndWest,
    SouthAndWest,
    SouthAndEast,
    None,
}

impl PipeSegment {
    pub fn try_from_char(c: char) -> Result<Self, String> {
        match c {
            '|' => Ok(Self::NorthAndSouth),
            '-' => Ok(Self::EastAndWest),
            'L' => Ok(Self::NorthAndEast),
            'J' => Ok(Self::NorthAndWest),
            '7' => Ok(Self::SouthAndWest),
            'F' => Ok(Self::SouthAndEast),
            '.' => Ok(Self::None),
            'S' => Ok(Self::Start),
            _ => Err(format!("{c} is not a valid pipe segment")),
        }
    }
}

impl PipeSegment {
    pub fn try_transform_direction(&self, in_direction: &Direction) -> Result<Direction, String> {
        match self {
            PipeSegment::NorthAndSouth => match in_direction {
                Direction::North => Ok(Direction::North),
                Direction::South => Ok(Direction::South),
                _ => Err(String::from("invalid in direction to NorthSouth pipe")),
            },
            PipeSegment::EastAndWest => match in_direction {
                Direction::East => Ok(Direction::East),
                Direction::West => Ok(Direction::West),
                _ => Err(String::from("invalid in direction to EastWest pipe")),
            },
            PipeSegment::NorthAndEast => match in_direction {
                Direction::South => Ok(Direction::East),
                Direction::West => Ok(Direction::North),
                _ => Err(String::from("invalid in direction to NorthAndEast pipe")),
            },
            PipeSegment::NorthAndWest => match in_direction {
                Direction::East => Ok(Direction::North),
                Direction::South => Ok(Direction::West),
                _ => Err(String::from("invalid in direction to NorthAndWest pipe")),
            },
            PipeSegment::SouthAndWest => match in_direction {
                Direction::East => Ok(Direction::South),
                Direction::North => Ok(Direction::West),
                _ => Err(String::from("invalid in direction to SouthAndWest pipe")),
            },
            PipeSegment::SouthAndEast => match in_direction {
                Direction::North => Ok(Direction::East),
                Direction::West => Ok(Direction::South),
                _ => Err(String::from("invalid in direction to SouthAndEast pipe")),
            },
            _ => Err(format!("cannout re-route")),
        }
    }
}

#[derive(Clone, Copy, PartialEq)]
enum Tag {
    Unmarked,
    Loop,
    Unenclosed,
    Enclosed,
}

struct SpacedGrid {
    data: Vec<Vec<Tag>>,
    height: usize,
    width: usize,
}

impl SpacedGrid {
    pub fn from_compact_dimensions(height: usize, width: usize) -> Self {
        let spaced_height = height * 2 + 1;
        let spaced_width = width * 2 + 1;
        SpacedGrid {
            data: vec![vec![Tag::Unmarked; spaced_width]; spaced_height],
            height: spaced_height,
            width: spaced_width,
        }
    }

    pub fn compact_coord_to_spaced_coord(compact_coord: &GridCoordinate) -> GridCoordinate {
        GridCoordinate {
            i: compact_coord.i * 2 + 1,
            j: compact_coord.j * 2 + 1,
        }
    }

    pub fn mark_loop_location(&mut self, coordinate: &GridCoordinate) {
        self.data[coordinate.i][coordinate.j] = Tag::Loop;
    }

    pub fn mark_unenclosed_tiles(&mut self) {
        let mut seen = HashSet::<GridCoordinate>::new();
        let mut frontier = VecDeque::<GridCoordinate>::new();
        frontier.push_back(GridCoordinate { i: 0, j: 0 });

        'search: while let Some(candidate) = frontier.pop_front() {
            if seen.contains(&candidate) {
                continue 'search;
            }

            self.data[candidate.i][candidate.j] = Tag::Unenclosed;

            let i_range = if candidate.i == 0 {
                0..=1
            } else if candidate.i == self.height - 1 {
                (candidate.i - 1)..=candidate.i
            } else {
                (candidate.i - 1)..=(candidate.i + 1)
            };
            let j_range = if candidate.j == 0 {
                0..=1
            } else if candidate.j == self.width - 1 {
                (candidate.j - 1)..=candidate.j
            } else {
                (candidate.j - 1)..=(candidate.j + 1)
            };

            for new_i in i_range {
                'j: for new_j in j_range.clone() {
                    if new_i == candidate.i && new_j == candidate.j {
                        continue 'j;
                    }

                    let next_candidate = GridCoordinate { i: new_i, j: new_j };

                    if self.data[next_candidate.i][next_candidate.j] != Tag::Loop
                        && self.data[next_candidate.i][next_candidate.j] != Tag::Unenclosed
                    {
                        frontier.push_back(next_candidate)
                    }
                }
            }

            seen.insert(candidate);
        }
    }

    pub fn count_enclosed_tiles(&mut self) -> usize {
        let mut n_enclosed_tiles = 0_usize;

        for i in (1..self.height).step_by(2) {
            for j in (1..self.width).step_by(2) {
                if self.data[i][j] == Tag::Unmarked {
                    self.data[i][j] = Tag::Enclosed;
                    n_enclosed_tiles += 1;
                }
            }
        }

        n_enclosed_tiles
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
            ..F7.
            .FJ|.
            SJ.L7
            |F--J
            LJ...
        "});
        let mut solver = Day10::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("8", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            FF7FSF7F7F7F7F7F---7
            L|LJ||||||||||||F--J
            FL-7LJLJ||||||LJL-77
            F--JF--7||LJLJ7F7FJ-
            L---JF-JLJ.||-FJLJJ7
            |F|F-JF---7F7-L7L|7|
            |FFJF7L7F-JF7|JL---7
            7-L-JL7||F7|L7F-7F7|
            L.L7LFJ|||||FJL7||LJ
            L7JLJL-JLJLJL--JLJ.L
        "});
        let mut solver = Day10::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("10", result);

        Ok(())
    }
}
