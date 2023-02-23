use std::collections::HashSet;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day09 {
    instructions: Vec<Instruction>,
}

impl Day09 {
    pub fn new() -> Self {
        Day09 {
            instructions: vec![],
        }
    }
}

impl Solver for Day09 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.instructions = input
            .trim()
            .split("\n")
            .map(|line| {
                let parts = line.split(" ").collect::<Vec<_>>();
                let direction = match parts[0] {
                    "U" => Direction::Up,
                    "D" => Direction::Down,
                    "L" => Direction::Left,
                    "R" => Direction::Right,
                    _ => panic!("Not a recognized direction"),
                };
                let magnitude = parts[1].parse::<u32>().expect("should be a u8 number");
                Instruction {
                    direction,
                    magnitude,
                }
            })
            .collect();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut tracker = RopeTracker::new(2);

        for instruction in self.instructions.iter() {
            tracker.move_rope(instruction);
        }

        Ok(tracker.visited_by_tail.len().to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut tracker = RopeTracker::new(10);

        for instruction in self.instructions.iter() {
            tracker.move_rope(instruction);
        }

        Ok(tracker.visited_by_tail.len().to_string())
    }
}

struct Instruction {
    direction: Direction,
    magnitude: u32,
}

enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Coordinate(i32, i32);

impl Coordinate {
    fn mv(&self, vector: [i32; 2]) -> Self {
        let Coordinate(i, j) = self;
        Coordinate(i + vector[0], j + vector[1])
    }

    fn follow_mv(&self, head: &Self) -> Self {
        let Coordinate(i, j) = *self;
        let Coordinate(h_i, h_j) = *head;

        let d_i = h_i - i;
        let d_j = h_j - j;
        let manhattan_distance = d_i.abs() + d_j.abs();

        if manhattan_distance < 2 {
            return Coordinate(i, j);
        }

        if d_i.abs() == d_j.abs() {
            return Coordinate(
                i + (d_i.signum() * (d_i.abs() - 1)),
                j + (d_j.signum() * (d_j.abs() - 1)),
            );
        }

        match (d_i, d_j) {
            (2, _) => Coordinate(i + 1, j + d_j),
            (-2, _) => Coordinate(i - 1, j + d_j),
            (_, 2) => Coordinate(i + d_i, j + 1),
            (_, -2) => Coordinate(i + d_i, j - 1),
            _ => panic!("How did the head get so far away? ({h_i}, {h_j}) <- ({i}, {j})",),
        }
    }
}

struct RopeTracker {
    knots: Vec<Coordinate>,
    visited_by_tail: HashSet<Coordinate>,
}

impl RopeTracker {
    fn new(n: usize) -> Self {
        RopeTracker {
            knots: vec![Coordinate(0, 0); n],
            visited_by_tail: HashSet::from([Coordinate(0, 0)]),
        }
    }

    fn move_rope(&mut self, instruction: &Instruction) {
        let mut knots = self.knots.clone();
        let vector: [i32; 2] = match instruction.direction {
            Direction::Up => [1, 0],
            Direction::Down => [-1, 0],
            Direction::Left => [0, -1],
            Direction::Right => [0, 1],
        };

        for _ in 0..instruction.magnitude {
            knots[0] = knots[0].mv(vector);
            for i in 1..knots.len() {
                knots[i] = knots[i].follow_mv(&knots[i - 1]);
            }
            self.visited_by_tail.insert(knots.last().unwrap().clone());
        }

        self.knots = knots;
    }
}

trait Debugger {
    fn grid_string(&self) -> String;
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
            R 4
            U 4
            L 3
            D 1
            R 4
            D 1
            L 5
            R 2
        "});
        let mut solver = Day09::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("13", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            R 4
            U 4
            L 3
            D 1
            R 4
            D 1
            L 5
            R 2
        "});
        let mut solver = Day09::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("1", result);

        Ok(())
    }

    #[test]
    fn part_two_big_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            R 5
            U 8
            L 8
            D 3
            R 17
            D 10
            L 25
            U 20
        "});
        let mut solver = Day09::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("36", result);

        Ok(())
    }
}
