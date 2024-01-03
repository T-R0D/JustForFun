use std::collections::{HashMap, HashSet, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day22 {
    bricks: Vec<Brick>,
}

impl Day22 {
    pub fn new() -> Self {
        Day22 { bricks: vec![] }
    }
}

impl Solver for Day22 {
    fn consume_input(&mut self, _input: &String) -> AoCResult {
        self.bricks = _input
            .trim()
            .lines()
            .map(Brick::try_from_line)
            .collect::<Result<Vec<_>, String>>()?;
        self.bricks.sort_by_key(|brick| brick.bottom_edge());

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let settled_bricks = compact_bricks_vertically(&self.bricks);

        let (bricks_by_resting_edge, bricks_by_top_edge) = create_brick_lookups(&settled_bricks);

        let mut removable_bricks = 0_usize;
        for base_brick in settled_bricks.iter() {
            let top_edge = base_brick.top_edge();

            let supported_bricks = match bricks_by_resting_edge.get(&(top_edge + 1)) {
                None => vec![],
                Some(bricks_above) => bricks_above
                    .iter()
                    .filter(|&brick_above| Brick::has_restable_overlap(brick_above, base_brick))
                    .collect(),
            };

            let brick_is_removable = supported_bricks.iter().all(|&supported_brick| {
                bricks_by_top_edge
                    .get(&(supported_brick.bottom_edge() - 1))
                    .is_some_and(|supporting_bricks| {
                        supporting_bricks
                            .iter()
                            .filter(|&supporting_brick| {
                                supporting_brick != base_brick
                                    && Brick::has_restable_overlap(
                                        supporting_brick,
                                        supported_brick,
                                    )
                            })
                            .count()
                            > 0
                    })
            });

            if brick_is_removable {
                removable_bricks += 1;
            }
        }

        Ok(removable_bricks.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let settled_bricks = compact_bricks_vertically(&self.bricks);

        let (bricks_by_resting_edge, bricks_by_top_edge) = create_brick_lookups(&settled_bricks);

        let support_graph = {
            let mut support_graph = HashMap::<usize, Vec<usize>>::new();
            for (i, brick) in settled_bricks.iter().enumerate() {
                let support_level = brick.top_edge() + 1;

                let supported_bricks = match bricks_by_resting_edge.get(&support_level) {
                    None => vec![],
                    Some(bricks_above) => bricks_above
                        .iter()
                        .filter(|brick_above| Brick::has_restable_overlap(&brick_above, brick))
                        .collect(),
                };

                let supported_brick_indices = supported_bricks
                    .iter()
                    .filter_map(|&&supported_brick| {
                        settled_bricks.iter().position(|&b| b == supported_brick)
                    })
                    .collect::<Vec<_>>();

                support_graph.insert(i, supported_brick_indices.clone());
            }

            support_graph
        };

        let mut falling_brick_sum = 0_usize;

        for i in 0..settled_bricks.len() {
            falling_brick_sum +=
                count_falling_bricks(&settled_bricks, &support_graph, &bricks_by_top_edge, i);
        }

        Ok(falling_brick_sum.to_string())
    }
}

fn compact_bricks_vertically(bricks: &Vec<Brick>) -> Vec<Brick> {
    let mut compacted_bricks = Vec::<Brick>::with_capacity(bricks.len());

    for falling_brick in bricks.iter() {
        let mut new_resting_level = 1_usize;
        for settled_brick in compacted_bricks.iter() {
            if falling_brick.bottom_edge() > settled_brick.top_edge()
                && Brick::has_restable_overlap(falling_brick, settled_brick)
            {
                new_resting_level = usize::max(new_resting_level, settled_brick.top_edge() + 1);
            }
        }

        compacted_bricks.push(falling_brick.dropped_to(new_resting_level));
    }

    compacted_bricks
}

fn create_brick_lookups(
    settled_bricks: &Vec<Brick>,
) -> (HashMap<usize, Vec<Brick>>, HashMap<usize, Vec<Brick>>) {
    let mut bricks_by_resting_edge = HashMap::<usize, Vec<Brick>>::new();
    let mut bricks_by_top_edge = HashMap::<usize, Vec<Brick>>::new();

    for brick in settled_bricks.iter() {
        let resting_edge = brick.bottom_edge();
        let top_edge = brick.top_edge();

        bricks_by_resting_edge
            .entry(resting_edge)
            .or_insert(vec![])
            .push(brick.clone());
        bricks_by_top_edge
            .entry(top_edge)
            .or_insert(vec![])
            .push(brick.clone());
    }

    (bricks_by_resting_edge, bricks_by_top_edge)
}

fn count_falling_bricks(
    settled_bricks: &Vec<Brick>,
    support_graph: &HashMap<usize, Vec<usize>>,
    bricks_by_top_edge: &HashMap<usize, Vec<Brick>>,
    removed_brick: usize,
) -> usize {
    let mut dropped_brick_indices = VecDeque::<usize>::from_iter([removed_brick].into_iter());
    let mut seen = HashSet::<usize>::new();
    let mut missing_bricks = HashSet::<usize>::from_iter([removed_brick].into_iter());
    let mut n_falling_bricks = 0_usize;
    let empty = Vec::<usize>::new();

    while let Some(falling_brick_index) = dropped_brick_indices.pop_front() {
        if seen.contains(&falling_brick_index) {
            continue;
        }

        // println!("Removing brick {}", falling_brick_index);

        missing_bricks.insert(falling_brick_index);

        let next_brick_indices = match support_graph.get(&falling_brick_index) {
            None => empty.clone(),
            Some(supported_bricks) => supported_bricks
                .iter()
                .filter(|&&supported_brick_index| {
                    let supported_brick = &settled_bricks[supported_brick_index];
                    let maybe_supporting_bricks = match bricks_by_top_edge
                        .get(&(supported_brick.bottom_edge() - 1))
                    {
                        None => empty.clone(),
                        Some(supporting_bricks) => supporting_bricks
                            .iter()
                            .filter_map(|supporting_brick| {
                                let supporting_brick_index = settled_bricks
                                    .iter()
                                    .position(|&b| b == *supporting_brick)
                                    .unwrap();

                                if missing_bricks.contains(&supporting_brick_index) {
                                    return None;
                                }

                                if Brick::has_restable_overlap(supporting_brick, supported_brick) {
                                    return Some(supporting_brick_index);
                                }

                                None
                            })
                            .collect(),
                    };

                    maybe_supporting_bricks.len() == 0
                })
                .map(|x| *x)
                .collect(),
        };

        // println!("{} was supporting {} bricks: {:?}", falling_brick_index, next_brick_indices.len(), next_brick_indices);

        n_falling_bricks += next_brick_indices.len();
        dropped_brick_indices.extend(next_brick_indices);

        seen.insert(falling_brick_index);
    }

    n_falling_bricks
}

#[derive(Clone, Copy, PartialEq, Eq)]
struct Brick {
    one_end: (usize, usize, usize),
    other_end: (usize, usize, usize),
}

impl Brick {
    fn try_from_line(line: &str) -> Result<Self, String> {
        let end_strs = line.split("~").collect::<Vec<_>>();
        let ends = end_strs
            .iter()
            .map(|end| {
                end.split(",")
                    .map(|coordinate| match coordinate.parse::<usize>() {
                        Ok(value) => Ok(value),
                        Err(err) => Err(err.to_string()),
                    })
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        Ok(Self {
            one_end: (ends[0][0], ends[0][1], ends[0][2]),
            other_end: (ends[1][0], ends[1][1], ends[1][2]),
        })
    }

    fn has_restable_overlap(a: &Brick, b: &Brick) -> bool {
        let a_face = a.top_bottom_face();
        let b_face = b.top_bottom_face();

        let has_x_overlap = Brick::range_overlaps(a_face.0, b_face.0);
        let has_y_overlap = Brick::range_overlaps(a_face.1, b_face.1);

        has_x_overlap && has_y_overlap
    }

    fn range_overlaps(range_a: (usize, usize), range_b: (usize, usize)) -> bool {
        if range_a.0 < range_b.0 {
            range_a.1 >= range_b.0
        } else {
            range_b.1 >= range_a.0
        }
    }

    fn bottom_edge(&self) -> usize {
        usize::min(self.one_end.2, self.other_end.2)
    }

    fn top_edge(&self) -> usize {
        usize::max(self.one_end.2, self.other_end.2)
    }

    fn top_bottom_face(&self) -> ((usize, usize), (usize, usize)) {
        return (
            (
                usize::min(self.one_end.0, self.other_end.0),
                usize::max(self.one_end.0, self.other_end.0),
            ),
            (
                usize::min(self.one_end.1, self.other_end.1),
                usize::max(self.one_end.1, self.other_end.1),
            ),
        );
    }

    fn dropped_to(&self, resting_level: usize) -> Self {
        let drop = self.bottom_edge() - resting_level;
        Self {
            one_end: (self.one_end.0, self.one_end.1, self.one_end.2 - drop),
            other_end: (self.other_end.0, self.other_end.1, self.other_end.2 - drop),
        }
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
            1,0,1~1,2,1
            0,0,2~2,0,2
            0,2,3~2,2,3
            0,0,4~0,2,4
            2,0,5~2,2,5
            0,1,6~2,1,6
            1,1,8~1,1,9
        "});
        let mut solver = Day22::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("5", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            1,0,1~1,2,1
            0,0,2~2,0,2
            0,2,3~2,2,3
            0,0,4~0,2,4
            2,0,5~2,2,5
            0,1,6~2,1,6
            1,1,8~1,1,9
        "});
        let mut solver = Day22::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("7", result);

        Ok(())
    }
}
