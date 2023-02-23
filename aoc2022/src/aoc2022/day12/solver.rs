use std::collections::{HashSet, VecDeque};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day12 {
    map: Vec<Vec<u8>>,
    target: Coordinate,
    start: Coordinate,
}

impl Day12 {
    pub fn new() -> Self {
        Day12 {
            map: vec![],
            target: Coordinate(0, 0),
            start: Coordinate(0, 0),
        }
    }
}

impl Solver for Day12 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let lines = input
            .trim()
            .split("\n")
            .map(|line| String::from(line))
            .collect::<Vec<_>>();
        let mut map: Vec<Vec<u8>> = Vec::with_capacity(lines.len());
        for (i, line) in lines.iter().enumerate() {
            let mut row: Vec<u8> = Vec::with_capacity(line.len());
            for (j, c) in line.as_bytes().iter().enumerate() {
                if *c == b'S' {
                    self.start = Coordinate(i, j);
                    row.push(b'a' - b'a');
                } else if *c == b'E' {
                    self.target = Coordinate(i, j);
                    row.push(b'z' - b'a');
                } else {
                    row.push(*c - b'a');
                }
            }
            map.push(row);
        }
        self.map = map;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut seen: HashSet<Coordinate> = HashSet::new();
        let mut frontier: VecDeque<SearchState> = VecDeque::from([SearchState {
            coordinate: self.start,
            steps: 0,
        }]);

        'search_loop: while let Some(to_consider) = frontier.pop_front() {
            let Coordinate(p_i, p_j) = to_consider.coordinate;

            if let Some(_) = seen.get(&to_consider.coordinate) {
                continue 'search_loop;
            }

            if to_consider.coordinate == self.target {
                return Ok((to_consider.steps).to_string());
            }

            let steps = to_consider.steps + 1;
            let current_elevation = self.map[p_i][p_j];
            for i in -1i32..=1 {
                'j_increment: for j in -1i32..=1 {
                    // This check got a little weird dancing around usize underflow.
                    // We check the values before transforming them to the neighboring
                    // coordinate.
                    if i.abs() == j.abs()
                    || (p_i == 0 && i == -1)
                    || (self.map.len() == p_i + 1 && i == 1)
                    || (p_j == 0 && j == -1)
                    || (self.map[1].len() <= p_j + 1 && j == 1)
                    {
                        continue 'j_increment;
                    }

                    let new_i = ((p_i as i32) + i) as usize;
                    let new_j = ((p_j as i32) + j) as usize;

                    let next_elevation = self.map[new_i][new_j];
                    if (next_elevation as i8 - current_elevation as i8) > 1 {
                        continue 'j_increment;
                    }

                    if let Some(_) = seen.get(&Coordinate(new_i, new_j)) {
                        continue 'j_increment;
                    }

                    frontier.push_back(SearchState {
                        coordinate: Coordinate(new_i, new_j),
                        steps,
                    });
                }

                seen.insert(to_consider.coordinate);
            }

            seen.insert(to_consider.coordinate);
        }

        Err(String::from("Unable to find solution..."))
    }

    fn solve_part_2(&self) -> AoCResult {
        let start = self.target;
        let mut seen: HashSet<Coordinate> = HashSet::new();
        let mut frontier: VecDeque<SearchState> = VecDeque::from([SearchState {
            coordinate: start,
            steps: 0,
        }]);

        'search_loop: while let Some(to_consider) = frontier.pop_front() {
            let Coordinate(p_i, p_j) = to_consider.coordinate;

            if let Some(_) = seen.get(&to_consider.coordinate) {
                continue 'search_loop;
            }

            let current_elevation = self.map[p_i][p_j];
            if current_elevation == 0 {
                return Ok(to_consider.steps.to_string());
            }

            let steps = to_consider.steps + 1;
            for i in -1i32..=1 {
                'j_increment: for j in -1i32..=1 {
                    // This check got a little weird dancing around usize underflow.
                    // We check the values before transforming them to the neighboring
                    // coordinate.
                    if i.abs() == j.abs()
                        || (p_i == 0 && i == -1)
                        || (self.map.len() == p_i + 1 && i == 1)
                        || (p_j == 0 && j == -1)
                        || (self.map[1].len() <= p_j + 1 && j == 1)
                    {
                        continue 'j_increment;
                    }

                    let new_i = ((p_i as i32) + i) as usize;
                    let new_j = ((p_j as i32) + j) as usize;

                    let next_elevation = self.map[new_i][new_j];
                    if (current_elevation as i8 - next_elevation as i8) > 1 {
                        continue 'j_increment;
                    }

                    if let Some(_) = seen.get(&Coordinate(new_i, new_j)) {
                        continue 'j_increment;
                    }

                    frontier.push_back(SearchState {
                        coordinate: Coordinate(new_i, new_j),
                        steps,
                    });
                }

                seen.insert(to_consider.coordinate);
            }

            seen.insert(to_consider.coordinate);
        }

        Err(String::from("Unable to find a valid path..."))
    }
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
struct Coordinate(usize, usize);

struct SearchState {
    coordinate: Coordinate,
    steps: usize,
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
            Sabqponm
            abcryxxl
            accszExk
            acctuvwj
            abdefghi
        "});
        let mut solver = Day12::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("31", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Sabqponm
            abcryxxl
            accszExk
            acctuvwj
            abdefghi
        "});
        let mut solver = Day12::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("29", result);

        Ok(())
    }
}
