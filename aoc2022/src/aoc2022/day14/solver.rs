use std::collections::HashSet;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day14 {
    ceiling: usize,
    floor: usize,
    left_edge: usize,
    right_edge: usize,
    segments: Vec<Segment>,
}

impl Day14 {
    pub fn new() -> Self {
        Day14 {
            ceiling: usize::MAX,
            floor: 0,
            left_edge: usize::MAX,
            right_edge: 0,
            segments: vec![],
        }
    }
}

impl Solver for Day14 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let segments = input
            .trim()
            .lines()
            .flat_map(|line| {
                let corners = line
                    .split(" -> ")
                    .map(|raw_corner| {
                        raw_corner
                            .split(",")
                            .map(str::parse::<usize>)
                            .map(Result::unwrap)
                            .collect::<Vec<_>>()
                    })
                    .map(|v| Coordinate(v[0], v[1]))
                    .collect::<Vec<_>>();
                corners
                    .windows(2)
                    .map(|window| {
                        if let [start, end] = window {
                            Segment(*start, *end)
                        } else {
                            panic!("invalid window");
                        }
                    })
                    .collect::<Vec<_>>()
            })
            .collect::<Vec<_>>();

        use std::cmp::{max, min};

        for segment in segments.iter() {
            let Segment(Coordinate(a_j, a_i), Coordinate(b_j, b_i)) = segment;
            self.ceiling = min(self.ceiling, min(*a_i, *b_i));
            self.floor = max(self.floor, max(*a_i, *b_i));
            self.left_edge = min(self.left_edge, min(*a_j, *b_j));
            self.right_edge = max(self.right_edge, max(*a_j, *b_j));
        }

        self.segments = segments;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut cave_slice =
            CaveSlice::from_segments(&self.segments, self.floor, self.left_edge, self.right_edge);

        while cave_slice.add_sand() {}

        Ok(cave_slice.accumulated_sand().to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut cave_slice = InfiniteCaveSlice::from_segments(&self.segments, self.floor);

        while cave_slice.add_sand() {}

        Ok(cave_slice.accumulated_sand().to_string())
    }
}

struct Segment(Coordinate, Coordinate);

#[derive(Clone, Copy, Hash, Eq, PartialEq)]
struct Coordinate(usize, usize);

const EMPTY: char = '.';
const ROCK: char = '#';
const SAND: char = 'X';

struct CaveSlice {
    depth: usize,
    width: usize,
    n_sand: usize,
    repr: Vec<Vec<char>>,
    sand_origin: Coordinate,
}

impl CaveSlice {
    fn from_segments(
        segments: &Vec<Segment>,
        floor: usize,
        left_edge: usize,
        right_edge: usize,
    ) -> Self {
        let height = floor + 2;
        let width = right_edge - left_edge + 1;

        let mut repr: Vec<Vec<char>> = vec![vec![EMPTY; width]; height];
        use std::cmp::{max, min};
        for segment in segments.iter() {
            let Segment(Coordinate(s_j, s_i), Coordinate(d_j, d_i)) = segment;

            for i in min(*s_i, *d_i)..=max(*s_i, *d_i) {
                for j in min(*s_j, *d_j)..=max(*s_j, *d_j) {
                    repr[i][j - left_edge] = ROCK;
                }
            }
        }

        CaveSlice {
            depth: height,
            width,
            n_sand: 0,
            repr,
            sand_origin: Coordinate(500 - left_edge, 0),
        }
    }

    fn add_sand(&mut self) -> bool {
        let Coordinate(mut j, _) = self.sand_origin;
        let mut settled = false;

        for i in 0..self.depth {
            if self.repr[i + 1][j] == EMPTY {
                continue;
            } else if j == 0 {
                break;
            } else if self.repr[i + 1][j - 1] == EMPTY {
                j -= 1;
                continue;
            } else if j == self.width - 1 {
                break;
            } else if self.repr[i + 1][j + 1] == EMPTY {
                j += 1;
                continue;
            }

            settled = true;
            self.repr[i][j] = SAND;
            self.n_sand += 1;
            break;
        }

        settled
    }

    fn accumulated_sand(&self) -> usize {
        self.n_sand
    }
}

struct InfiniteCaveSlice {
    depth: usize,
    n_sand: usize,
    occupied: HashSet<Coordinate>,
    sand_origin: Coordinate,
}

impl InfiniteCaveSlice {
    fn from_segments(segments: &Vec<Segment>, floor: usize) -> Self {
        let height = floor + 1;

        let mut occupied: HashSet<Coordinate> = HashSet::new();
        use std::cmp::{max, min};
        for segment in segments.iter() {
            let Segment(Coordinate(s_j, s_i), Coordinate(d_j, d_i)) = segment;

            for i in min(*s_i, *d_i)..=max(*s_i, *d_i) {
                for j in min(*s_j, *d_j)..=max(*s_j, *d_j) {
                    occupied.insert(Coordinate(i, j));
                }
            }
        }

        InfiniteCaveSlice {
            depth: height,
            n_sand: 0,
            occupied,
            sand_origin: Coordinate(0, 500),
        }
    }

    fn add_sand(&mut self) -> bool {
        let Coordinate(_, mut j) = self.sand_origin;

        let mut settled_above_floor = false;
        for i in 0..self.depth {
            if !self.occupied.contains(&Coordinate(i + 1, j)) {
                continue;
            } else if !self.occupied.contains(&Coordinate(i + 1, j - 1)) {
                j -= 1;
                continue;
            } else if !self.occupied.contains(&Coordinate(i + 1, j + 1)) {
                j += 1;
                continue;
            }

            // println!("Resting place: {}, {}", i, j);

            settled_above_floor = true;
            self.occupied.insert(Coordinate(i, j));
            self.n_sand += 1;
            break;
        }

        if !settled_above_floor {
            self.occupied.insert(Coordinate(self.depth, j));
            self.n_sand += 1;
        }

        !self.occupied.contains(&self.sand_origin)
    }

    fn accumulated_sand(&self) -> usize {
        self.n_sand
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
            498,4 -> 498,6 -> 496,6
            503,4 -> 502,4 -> 502,9 -> 494,9
        "});
        let mut solver = Day14::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("24", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            498,4 -> 498,6 -> 496,6
            503,4 -> 502,4 -> 502,9 -> 494,9
        "});
        let mut solver = Day14::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("93", result);

        Ok(())
    }
}
