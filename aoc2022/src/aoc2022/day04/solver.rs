use crate::aoc2022::solver::interface::{AoCResult, Solver};

#[derive(Clone, Copy)]
struct SectionRange {
    first: u8,
    last: u8,
}

pub struct Day04 {
    section_assignments: Vec<(SectionRange, SectionRange)>,
}

impl Day04 {
    pub fn new() -> Self {
        Day04 {
            section_assignments: vec![],
        }
    }
}

impl Solver for Day04 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.section_assignments = input
            .trim()
            .split("\n")
            .collect::<Vec<_>>()
            .iter()
            .map(|line| {
                let assignments = line.split(",").collect::<Vec<_>>();
                let parsed_assignments = assignments
                    .iter()
                    .map(|assignment| {
                        let parts = assignment
                            .split("-")
                            .collect::<Vec<_>>()
                            .iter()
                            .map(|&part| {
                                part.parse::<u8>().expect("Parts should be valid integers")
                            })
                            .collect::<Vec<_>>();
                        SectionRange {
                            first: parts[0],
                            last: parts[1],
                        }
                    })
                    .collect::<Vec<_>>();

                (parsed_assignments[0], parsed_assignments[1])
            })
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut enclosed_assignments: u32 = 0;

        for (a, b) in self.section_assignments.iter() {
            let (first, second) = if a.first == b.first {
                if a.last < b.last {
                    (b, a)
                } else {
                    (a, b)
                }
            } else if a.first < b.first {
                (a, b)
            } else {
                (b, a)
            };

            if second.last <= first.last {
                enclosed_assignments += 1;
            }
        }

        Ok(enclosed_assignments.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut enclosed_assignments: u32 = 0;

        for (a, b) in self.section_assignments.iter() {
            let (first, second) = if a.first <= b.first { (a, b) } else { (b, a) };

            if first.first <= second.first && second.first <= first.last {
                enclosed_assignments += 1;
            }
        }

        Ok(enclosed_assignments.to_string())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            2-4,6-8
            2-3,4-5
            5-7,7-9
            2-8,3-7
            6-6,4-6
            2-6,4-8
        "});
        let mut solver = Day04::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("2", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            2-4,6-8
            2-3,4-5
            5-7,7-9
            2-8,3-7
            6-6,4-6
            2-6,4-8
        "});
        let mut solver = Day04::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("4", result);

        Ok(())
    }
}
