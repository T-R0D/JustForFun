use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day12 {
    spring_records: Vec<ConditionRecord>,
}

impl Day12 {
    pub fn new() -> Self {
        Day12 {
            spring_records: vec![],
        }
    }
}

impl Solver for Day12 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.spring_records = input
            .trim()
            .lines()
            .map(|line| {
                let parts = line.split(" ").collect::<Vec<_>>();
                if parts.len() != 2 {
                    return Err(String::from("line was invalid format"));
                }

                let spring_row_statuses = parts[0]
                    .chars()
                    .map(Status::try_from_char)
                    .collect::<Result<Vec<_>, String>>()?;
                let damaged_group_sizes = parts[1]
                    .split(",")
                    .map(|item| match item.parse::<usize>() {
                        Ok(value) => Ok(value),
                        Err(err) => Err(err.to_string()),
                    })
                    .collect::<Result<Vec<_>, String>>()?;

                Ok(ConditionRecord {
                    spring_row_statuses,
                    damaged_group_sizes,
                })
            })
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let arrangement_count_sum = self
            .spring_records
            .iter()
            .map(|record| record.count_damage_arrangements_by_backtracking())
            .sum::<usize>();

        Ok(arrangement_count_sum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let arrangement_count_sum = self
            .spring_records
            .iter()
            .map(|record| record.unfold())
            .map(|record| record.count_arrangements_by_table())
            .sum::<usize>();

        Ok(arrangement_count_sum.to_string())
    }
}

struct ConditionRecord {
    spring_row_statuses: Vec<Status>,
    damaged_group_sizes: Vec<usize>,
}

impl ConditionRecord {
    fn count_damage_arrangements_by_backtracking(&self) -> usize {
        let mut placement_stack = Vec::<Placement>::with_capacity(self.damaged_group_sizes.len());

        let initial_attempt = Placement {
            group_index: 0,
            start_index: 0,
        };
        if let None = self.find_placement(&mut placement_stack, &initial_attempt) {
            return 0;
        }

        let mut backtrack = false;
        let mut ways_to_place = 0_usize;
        while let Some(last_placement) = placement_stack.last() {
            if backtrack {
                if let Some(placement) = placement_stack.pop() {
                    let next_attempt = Placement {
                        start_index: placement.start_index + 1,
                        group_index: placement.group_index,
                    };
                    backtrack = match self.find_placement(&mut placement_stack, &next_attempt) {
                        Some(_) => false,
                        None => true,
                    };
                    continue;
                }
                backtrack = false;
                continue;
            }

            if placement_stack.len() == self.damaged_group_sizes.len() {
                backtrack = true;

                let mut damaged_springs_uncovered = false;
                let remaining_scan_start = last_placement.start_index
                    + self.damaged_group_sizes[last_placement.group_index];
                for i in remaining_scan_start..self.spring_row_statuses.len() {
                    if self.spring_row_statuses[i] == Status::Damaged {
                        damaged_springs_uncovered = true;
                        break;
                    }
                }

                if !damaged_springs_uncovered {
                    ways_to_place += 1;
                }

                continue;
            }

            let next_attempt = Placement {
                start_index: last_placement.start_index
                    + self.damaged_group_sizes[last_placement.group_index]
                    + 1,
                group_index: last_placement.group_index + 1,
            };

            backtrack = match self.find_placement(&mut placement_stack, &next_attempt) {
                Some(_) => false,
                None => true,
            };
        }

        ways_to_place
    }

    fn find_placement(
        &self,
        placement_stack: &mut Vec<Placement>,
        initial_attempt: &Placement,
    ) -> Option<Placement> {
        let group_size = self.damaged_group_sizes[initial_attempt.group_index];
        let first_start_attempt = initial_attempt.start_index;
        let last_start_attempt =
            self.spring_row_statuses.len() - self.damaged_group_sizes[initial_attempt.group_index];

        'placement_scan: for start in first_start_attempt..=last_start_attempt {
            if start > 0 && self.spring_row_statuses[start - 1] == Status::Damaged {
                break 'placement_scan;
            }

            let end = start + group_size;
            if end > self.spring_row_statuses.len() {
                break 'placement_scan;
            }

            if end < self.spring_row_statuses.len()
                && self.spring_row_statuses[end] == Status::Damaged
            {
                continue 'placement_scan;
            }

            for &status in self.spring_row_statuses.iter().skip(start).take(group_size) {
                if status == Status::Operational {
                    continue 'placement_scan;
                }
            }

            let placement = Placement {
                start_index: start,
                group_index: initial_attempt.group_index,
            };
            placement_stack.push(placement);
            return Some(placement);
        }

        None
    }

    fn unfold(&self) -> Self {
        let mut spring_row_statuses = self.spring_row_statuses.clone();
        let mut damaged_group_sizes = self.damaged_group_sizes.clone();
        for _ in 0..4 {
            spring_row_statuses.push(Status::Unknown);
            spring_row_statuses.append(&mut self.spring_row_statuses.clone());
            damaged_group_sizes.append(&mut self.damaged_group_sizes.clone());
        }

        Self {
            spring_row_statuses,
            damaged_group_sizes,
        }
    }

    fn count_arrangements_by_table(&self) -> usize {
        let m_groups = self.damaged_group_sizes.len();
        let n_spaces = self.spring_row_statuses.len();
        let mut min_start_indexes = vec![0; m_groups];
        for i in 1..m_groups {
            let previous_group_size = self.damaged_group_sizes[i - 1];
            min_start_indexes[i] = min_start_indexes[i - 1] + previous_group_size + 1;
        }
        let mut max_start_indexes =
            vec![n_spaces - self.damaged_group_sizes[self.damaged_group_sizes.len() - 1]; m_groups];
        for i in (0..(m_groups - 1)).rev() {
            let this_group_size = self.damaged_group_sizes[i];
            max_start_indexes[i] = max_start_indexes[i + 1] - (this_group_size + 1);
        }

        let mut table = vec![vec![0_usize; n_spaces]; m_groups];
        let mut row_maxes = vec![0; m_groups];
        for i in (0..m_groups).rev() {
            let group_size = self.damaged_group_sizes[i];
            let min_start = min_start_indexes[i];
            let max_start = max_start_indexes[i];
            let mut row_sum = 0_usize;
            'next_position: for j in (min_start..=max_start).rev() {
                let end = j + group_size;

                for k in j..end {
                    if self.spring_row_statuses[k] == Status::Operational {
                        continue 'next_position;
                    }
                }

                if i == 0 {
                    for k in 0..j {
                        if self.spring_row_statuses[k] == Status::Damaged {
                            continue 'next_position;
                        }
                    }
                }

                if i == m_groups - 1 && end < n_spaces {
                    for k in end..n_spaces {
                        if self.spring_row_statuses[k] == Status::Damaged {
                            break 'next_position;
                        }
                    }
                }

                if j < max_start
                    && end < n_spaces
                    && self.spring_row_statuses[end] == Status::Damaged
                {
                    if i < m_groups - 1 {
                        continue 'next_position;
                    }

                    break 'next_position;
                }

                if j > min_start && j > 0 && self.spring_row_statuses[j - 1] == Status::Damaged {
                    continue 'next_position;
                }

                let ways_to_place = if i + 1 >= m_groups {
                    1
                } else {
                    let mut x = 0_usize;
                    'count: for k in (end + 1)..=max_start_indexes[i + 1] {
                        if self.spring_row_statuses[k - 1] == Status::Damaged
                            && table[i + 1][k] == 0
                        {
                            break 'count;
                        }
                        x += table[i + 1][k];
                    }

                    x
                };

                table[i][j] = ways_to_place;
                row_sum += ways_to_place;
            }

            row_maxes[i] = row_sum;
        }

        row_maxes[0]
    }
}

#[derive(Clone, Copy, PartialEq, Debug)]
enum Status {
    Operational,
    Damaged,
    Unknown,
}

impl Status {
    fn try_from_char(c: char) -> Result<Status, String> {
        match c {
            '.' => Ok(Status::Operational),
            '#' => Ok(Status::Damaged),
            '?' => Ok(Status::Unknown),
            _ => Err(format!("{c} is not a valid gear status")),
        }
    }
}

#[derive(Clone, Copy, Debug)]
struct Placement {
    group_index: usize,
    start_index: usize,
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
            ???.### 1,1,3
            .??..??...?##. 1,1,3
            ?#?#?#?#?#?#?#? 1,3,1,6
            ????.#...#... 4,1,1
            ????.######..#####. 1,6,5
            ?###???????? 3,2,1
        "});
        let mut solver = Day12::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("21", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            ???.### 1,1,3
            .??..??...?##. 1,1,3
            ?#?#?#?#?#?#?#? 1,3,1,6
            ????.#...#... 4,1,1
            ????.######..#####. 1,6,5
            ?###???????? 3,2,1
        "});
        let mut solver = Day12::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("525152", result);

        Ok(())
    }
}
