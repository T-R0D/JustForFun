use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day08 {
    instruction_list: Vec<Instruction>,
    location_nodes: HashMap<String, LocationNode>,
}

impl Day08 {
    pub fn new() -> Self {
        Day08 {
            instruction_list: vec![],
            location_nodes: HashMap::<String, LocationNode>::new(),
        }
    }
}

impl Solver for Day08 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let parts = input.trim().split("\n\n").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(String::from("input had wrong number of sections"));
        }

        self.instruction_list = parts[0]
            .chars()
            .map(|c| Instruction::try_from_char(c))
            .collect::<Result<Vec<_>, String>>()?;

        self.location_nodes = parts[1]
            .lines()
            .map(|line| {
                let label_and_neighbors = line.split(" = ").collect::<Vec<_>>();
                if label_and_neighbors.len() != 2 {
                    return Err(format!("{line} could not be split into neighbor and label"));
                }

                let label = String::from(label_and_neighbors[0]);

                let neighbors = label_and_neighbors[1]
                    .replace("(", "")
                    .replace(")", "")
                    .split(", ")
                    .map(|part| String::from(part))
                    .collect::<Vec<_>>();
                if neighbors.len() != 2 {
                    return Err(format!("wrong number of neighbors ({})", neighbors.len()));
                }
                let location_node =
                    LocationNode::new(label.clone(), neighbors[0].clone(), neighbors[1].clone());

                Ok((label, location_node))
            })
            .collect::<Result<HashMap<String, LocationNode>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let start_label = String::from(START_LABEL_RAW);
        let end_label = String::from(END_LABEL_RAW);
        let mut current_location_label = start_label;
        let mut i = 0_usize;
        let mut steps_taken = 0_usize;
        while let Some(current_location) = self.location_nodes.get(&current_location_label) {
            if current_location.label == end_label {
                return Ok(steps_taken.to_string());
            }

            current_location_label = current_location.neighbor(self.instruction_list[i]);
            i = (i + 1) % self.instruction_list.len();
            steps_taken += 1;
        }

        Err(String::from("Got lost - final node not found"))
    }

    fn solve_part_2(&self) -> AoCResult {
        let starting_locations = self
            .location_nodes
            .keys()
            .filter(|label| label.ends_with(START_TAG_RAW))
            .collect::<Vec<_>>();

        let periodicities = starting_locations
            .iter()
            .map(|&label| self.find_periodicities_of_path(label))
            .collect::<Result<Vec<_>, String>>()?;

        // Curiously, the offsets and periods are the same for all start->end cycles, so we only need to keep
        // of the period or the offset. It seems that the input was crafted to have the following properties:
        // * Each start reaches only one destination
        // * The length of path from start to destination is the same as the path from destination to itself
        // * (Maybe) Periods are all prime?
        let periods = periodicities.iter().map(|p| p.period).collect::<Vec<_>>();

        // let steps_to_end_on_only_ends = self.try_least_common_multiple(&periods)?;
        let steps_to_end_on_only_ends = periods.iter().fold(1, |acc, &p| Day08::lcm(acc, p));

        Ok(steps_to_end_on_only_ends.to_string())
    }
}

impl Day08 {
    fn find_periodicities_of_path(&self, start_label: &String) -> Result<PathPeriodicity, String> {
        let mut current_location_label = start_label.clone();
        let mut i = 0_usize;
        let mut end_locations_to_steps = HashMap::<String, usize>::new();
        let mut steps_taken = 0_usize;
        let mut first_steps_to_end = 0_usize;

        while let Some(current_location) = self.location_nodes.get(&current_location_label) {
            if current_location.label.ends_with(END_TAG_RAW) {
                let location_entry = end_locations_to_steps.entry(current_location.label.clone());
                match location_entry {
                    Occupied(_) => {
                        return Ok(PathPeriodicity {
                            _offset: first_steps_to_end,
                            period: steps_taken - first_steps_to_end,
                        })
                    }
                    Vacant(entry) => {
                        first_steps_to_end = steps_taken;
                        entry.insert(steps_taken);
                        ()
                    }
                }
            }

            current_location_label = current_location.neighbor(self.instruction_list[i]);
            i = (i + 1) % self.instruction_list.len();
            steps_taken += 1;
        }

        unreachable!()
    }

    fn lcm(a: usize, b: usize) -> usize {
        (a * b) / Day08::gcd(a, b)
    }

    // A shamelessly stolen implementation of Euclid's algorithm.
    // (Looked up after finding the solution with a little more "brute forcey" LCM function.)
    fn gcd(a: usize, b: usize) -> usize {
        let (mut a_2, mut b_2) = (a, b);
        while b_2 > 0 {
            (a_2, b_2) = (b_2, a_2 % b_2);
        }
        a_2
    }
}

#[derive(Clone, Copy)]
enum Instruction {
    Left,
    Right,
    NInstructions,
}

impl Instruction {
    pub fn try_from_char(c: char) -> Result<Instruction, String> {
        match c {
            'L' => Ok(Self::Left),
            'R' => Ok(Self::Right),
            _ => Err(format!("{c} is not a valid instruction")),
        }
    }

    pub fn index(&self) -> usize {
        *self as usize
    }
}

const START_LABEL_RAW: &str = "AAA";
const START_TAG_RAW: &str = "A";
const END_LABEL_RAW: &str = "ZZZ";
const END_TAG_RAW: &str = "Z";

#[derive(Clone)]
struct LocationNode {
    label: String,
    neighbors: [String; Instruction::NInstructions as usize],
}

impl LocationNode {
    pub fn new(
        label: String,
        left_neighbor_label: String,
        right_neighbor_label: String,
    ) -> LocationNode {
        LocationNode {
            label,
            neighbors: [left_neighbor_label, right_neighbor_label],
        }
    }

    pub fn neighbor(&self, instruction: Instruction) -> String {
        self.neighbors[instruction.index()].clone()
    }
}

struct PathPeriodicity {
    period: usize,
    _offset: usize,
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
            LLR

            AAA = (BBB, BBB)
            BBB = (AAA, ZZZ)
            ZZZ = (ZZZ, ZZZ)
        "});
        let mut solver = Day08::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("6", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            LR

            11A = (11B, XXX)
            11B = (XXX, 11Z)
            11Z = (11B, XXX)
            22A = (22B, XXX)
            22B = (22C, 22C)
            22C = (22Z, 22Z)
            22Z = (22B, 22B)
            XXX = (XXX, XXX)
        "});
        let mut solver = Day08::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("6", result);

        Ok(())
    }
}
