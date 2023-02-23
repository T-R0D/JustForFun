use std::collections::{HashSet, VecDeque};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day18 {
    voxel_list: Vec<Coordinate>,
}

impl Day18 {
    pub fn new() -> Self {
        Self {
            voxel_list: Vec::new(),
        }
    }
}

impl Solver for Day18 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.voxel_list = input
            .trim()
            .lines()
            .map(Coordinate::from_line)
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let voxel_lookup = self
            .voxel_list
            .iter()
            .map(Coordinate::clone)
            .collect::<HashSet<_>>();

        let mut shared_sides = 0;
        for voxel in self.voxel_list.iter() {
            for neighbor in voxel.adjacent_neighbors().iter() {
                if voxel_lookup.contains(neighbor) {
                    shared_sides += 1;
                }
            }
        }

        let total_sides = self.voxel_list.len() * 6;
        let exposed_sides = total_sides - shared_sides;

        Ok(exposed_sides.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        // for x in 0..voxel_space.len() {
        //     let slice = &voxel_space[x];

        //     println!("x={x}");
        //     let mut s = String::new();
        //     for y in 0..slice.len() {
        //         for z in 0..slice[0].len() {
        //             if voxel_space[x][y][z] == EMPTY {
        //                 s.extend("  . ".chars());
        //             } else {
        //                 s.extend(format!("{: ^4}", voxel_space[x][y][z]).chars());
        //             }
        //         }
        //         s.push('\n');
        //     }

        //     println!("{s}");
        // }

        let external_surface_area = count_exposed_exterior_sides(&self.voxel_list);

        Ok(external_surface_area.to_string())
    }
}

const EMPTY: u8 = u8::MAX;

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
struct Coordinate {
    x: usize,
    y: usize,
    z: usize,
}

impl Coordinate {
    fn new(x: usize, y: usize, z: usize) -> Self {
        Coordinate { x, y, z }
    }

    fn from_line(line: &str) -> Self {
        let parts = line
            .split(",")
            .map(str::parse::<usize>)
            .map(Result::unwrap)
            .collect::<Vec<_>>();
        Coordinate::new(parts[0], parts[1], parts[2])
    }

    fn adjacent_neighbors(&self) -> Vec<Coordinate> {
        let mut neighbors = Vec::with_capacity(6);

        for [x, y, z] in [
            [-1, 0, 0],
            [1, 0, 0],
            [0, -1, 0],
            [0, 1, 0],
            [0, 0, -1],
            [0, 0, 1],
        ] {
            if x == -1 && self.x == 0 {
                continue;
            }

            if y == -1 && self.y == 0 {
                continue;
            }

            if z == -1 && self.z == 0 {
                continue;
            }

            neighbors.push(Coordinate::new(
                (self.x as i32 + x) as usize,
                (self.y as i32 + y) as usize,
                (self.z as i32 + z) as usize,
            ));
        }

        neighbors
    }
}

fn construct_3d_model_of_droplet(voxels: &Vec<Coordinate>) -> Vec<Vec<Vec<u8>>> {
    let (x_max, y_max, z_max) = voxels.iter().fold(
        (0, 0, 0),
        |(x_max, y_max, z_max), &Coordinate { x, y, z }| {
            (
                std::cmp::max(x, x_max),
                std::cmp::max(y, y_max),
                std::cmp::max(z, z_max),
            )
        },
    );
    // In this grid we pad the edges with 1 space to make computations (like
    // search) easier down the line. The space is expected to be relatively
    // small (18^3), so this shouldn't be too much of a RAM burden.
    let mut voxel_space: Vec<Vec<Vec<u8>>> =
        vec![vec![vec![EMPTY; z_max + 3]; y_max + 3]; x_max + 3];
    for (i, &Coordinate { x, y, z }) in voxels.iter().enumerate() {
        // +1 to move any droplets away from the axes to make computations with
        // usizes easier down the line.
        voxel_space[x + 1][y + 1][z + 1] = i as u8;
    }

    voxel_space
}

fn count_exposed_exterior_sides(voxel_list: &Vec<Coordinate>) -> usize {
    let voxel_space = construct_3d_model_of_droplet(voxel_list);
    let x_bound = voxel_space.len();
    let y_bound = voxel_space[0].len();
    let z_bound = voxel_space[0][0].len();

    let mut frontier: VecDeque<Coordinate> = VecDeque::from([Coordinate { x: 0, y: 0, z: 0 }]);
    let mut seen: HashSet<Coordinate> = HashSet::new();

    let mut covered_sides = 0;

    while let Some(candidate) = frontier.pop_front() {
        if seen.contains(&candidate) {
            continue;
        }

        'neighbor_search: for neighbor in candidate.adjacent_neighbors().iter() {
            if seen.contains(neighbor)
                || neighbor.x >= x_bound
                || neighbor.y >= y_bound
                || neighbor.z >= z_bound
            {
                continue 'neighbor_search;
            }

            if voxel_space[neighbor.x][neighbor.y][neighbor.z] != EMPTY {
                covered_sides += 1;
                continue 'neighbor_search;
            }

            frontier.push_back(neighbor.clone());
        }

        seen.insert(candidate);
    }

    covered_sides
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
            2,2,2
            1,2,2
            3,2,2
            2,1,2
            2,3,2
            2,2,1
            2,2,3
            2,2,4
            2,2,6
            1,2,5
            3,2,5
            2,1,5
            2,3,5
        "});
        let mut solver = Day18::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("64", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            2,2,2
            1,2,2
            3,2,2
            2,1,2
            2,3,2
            2,2,1
            2,2,3
            2,2,4
            2,2,6
            1,2,5
            3,2,5
            2,1,5
            2,3,5
        "});
        let mut solver = Day18::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("58", result);

        Ok(())
    }
}
