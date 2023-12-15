use std::collections::{HashMap, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day15 {
    instructions: Vec<String>,
}

impl Day15 {
    pub fn new() -> Self {
        Day15 {
            instructions: vec![],
        }
    }
}

impl Solver for Day15 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let instructions_lines = input
            .trim()
            .lines()
            .map(|line| line.split(",").map(String::from).collect::<Vec<_>>())
            .collect::<Vec<_>>();

        if instructions_lines.len() != 1 {
            return Err(format!(
                "input had wrong number of lines ({})",
                instructions_lines.len()
            ));
        }

        self.instructions = instructions_lines[0].clone();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let hash_total = self
            .instructions
            .iter()
            .map(holiday_ascii_string_helper)
            .sum::<usize>();

        Ok(hash_total.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let ops = self
            .instructions
            .iter()
            .map(Operation::try_from_string)
            .collect::<Result<Vec<_>, String>>()?;

        let mut lens_boxes = Vec::<LensBox>::with_capacity(N_BOXES);
        for id in 0..N_BOXES {
            lens_boxes.push(LensBox::new(id));
        }

        let mut label_to_box_id = HashMap::<String, usize>::new();

        for operation in ops.iter() {
            match operation {
                Operation::Remove { label } => {
                    let box_id = label_to_box_id
                        .entry(label.clone())
                        .or_insert(holiday_ascii_string_helper(label));

                    lens_boxes[*box_id].remove_lens(&label);
                }
                Operation::Insert {
                    label,
                    focal_length,
                } => {
                    let box_id = label_to_box_id
                        .entry(label.clone())
                        .or_insert(holiday_ascii_string_helper(label));

                    lens_boxes[*box_id].insert_lens(&label, *focal_length);
                }
            }
        }

        let total_focusing_power = lens_boxes
            .iter()
            .map(|lens_box| lens_box.focusing_power())
            .sum::<usize>();

        Ok(total_focusing_power.to_string())
    }
}

const N_BOXES: usize = 256;
const N_FOCAL_LENGTHS: usize = 9;

fn holiday_ascii_string_helper(s: &String) -> usize {
    let mut current_value = 0_usize;
    for &c in s.as_bytes().iter() {
        current_value += usize::try_from(c).unwrap();
        current_value *= 17;
        current_value %= 256;
    }
    current_value
}

#[derive(PartialEq, Eq)]
enum Operation {
    Remove { label: String },
    Insert { label: String, focal_length: usize },
}

impl Operation {
    fn try_from_string(s: &String) -> Result<Self, String> {
        if s.ends_with("-") {
            Ok(Self::Remove {
                label: s.replace("-", ""),
            })
        } else {
            let parts = s.split("=").collect::<Vec<_>>();
            if parts.len() != 2 {
                return Err(format!("{s} could not be parsed as insert operation"));
            }

            match parts[1].parse::<usize>() {
                Ok(value) => Ok(Self::Insert {
                    label: parts[0].to_string(),
                    focal_length: value,
                }),
                Err(err) => Err(err.to_string()),
            }
        }
    }
}

struct LensBox {
    id: usize,
    lenses: VecDeque<Lens>,
}

impl LensBox {
    fn new(id: usize) -> Self {
        Self {
            id,
            lenses: VecDeque::with_capacity(N_FOCAL_LENGTHS),
        }
    }

    fn insert_lens(&mut self, label: &String, focal_length: usize) {
        if let Some(lens_index) = self.lenses.iter().position(|lens| lens.label == *label) {
            self.lenses[lens_index].focal_length = focal_length;
            return;
        } else {
            self.lenses.push_back(Lens {
                label: label.clone(),
                focal_length,
            });
        }
    }

    fn remove_lens(&mut self, label: &String) {
        self.lenses.retain(|lens| lens.label != *label);
    }

    fn focusing_power(&self) -> usize {
        self.lenses
            .iter()
            .enumerate()
            .map(|(i, lens)| (self.id + 1) * (i + 1) * lens.focal_length)
            .sum::<usize>()
    }
}

struct Lens {
    label: String,
    focal_length: usize,
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
            rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
        "});
        let mut solver = Day15::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("1320", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
        "});
        let mut solver = Day15::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("145", result);

        Ok(())
    }
}
