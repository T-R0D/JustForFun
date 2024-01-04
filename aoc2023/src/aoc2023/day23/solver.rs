use std::collections::{HashMap, HashSet, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day23 {
    grid: Vec<Vec<Feature>>,
    m: usize,
    n: usize,
    start: GridCoordinate,
    end: GridCoordinate,
}

impl Day23 {
    pub fn new() -> Self {
        Day23 {
            grid: vec![],
            m: 0,
            n: 0,
            start: GridCoordinate { i: 0, j: 0 },
            end: GridCoordinate { i: 0, j: 0 },
        }
    }
}

impl Solver for Day23 {
    fn consume_input(&mut self, _input: &String) -> AoCResult {
        self.grid = _input
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(Feature::try_from_char)
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        self.m = self.grid.len();
        self.n = self.grid[0].len();

        self.start = GridCoordinate {
            i: 0,
            j: self.grid[0]
                .iter()
                .position(|f| *f == Feature::Path)
                .unwrap(),
        };
        self.end = GridCoordinate {
            i: self.m - 1,
            j: self.grid[self.m - 1]
                .iter()
                .position(|f| *f == Feature::Path)
                .unwrap(),
        };

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        // let longest_path_steps = self.find_longest_path_len_accounting_for_slippery_slopes();

        // Ok(longest_path_steps.to_string())

        let (_, vertex_indices) = self.find_path_vertices(|current_loc: &GridCoordinate| {
            self.slippery_slope_next_locations(current_loc)
        });

        let adjacency_graph = self
            .construct_adjacency_graph(&vertex_indices, |current_loc: &GridCoordinate| {
                self.slippery_slope_next_locations(current_loc)
            });

        let (src, dst) = match (
            vertex_indices.get(&self.start),
            vertex_indices.get(&self.end),
        ) {
            (Some(&src), Some(&dst)) => Ok((src, dst)),
            _ => Err(String::from("src and dst indices could not be found")),
        }?;

        let longest_path_steps = Day23::find_longest_path_len(&adjacency_graph, src, dst);

        Ok(longest_path_steps.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let (_, vertex_indices) = self.find_path_vertices(|current_loc: &GridCoordinate| {
            self.high_traction_next_locations(current_loc)
        });

        let adjacency_graph = self
            .construct_adjacency_graph(&vertex_indices, |current_loc: &GridCoordinate| {
                self.high_traction_next_locations(current_loc)
            });

        let (src, dst) = match (
            vertex_indices.get(&self.start),
            vertex_indices.get(&self.end),
        ) {
            (Some(&src), Some(&dst)) => Ok((src, dst)),
            _ => Err(String::from("src and dst indices could not be found")),
        }?;

        let longest_path_steps = Day23::find_longest_path_len(&adjacency_graph, src, dst);

        Ok(longest_path_steps.to_string())
    }
}

impl Day23 {
    fn find_path_vertices(
        &self,
        next_locations_fn: impl Fn(&GridCoordinate) -> Vec<GridCoordinate>,
    ) -> (Vec<GridCoordinate>, HashMap<GridCoordinate, usize>) {
        let mut frontier = VecDeque::<GridCoordinate>::from_iter([self.start.clone()].into_iter());
        let mut seen = HashSet::<GridCoordinate>::new();

        let mut path_vertices = vec![self.start.clone(), self.end.clone()];

        while let Some(current_loc) = frontier.pop_front() {
            if seen.contains(&current_loc) {
                continue;
            }

            let next_states = next_locations_fn(&current_loc);

            if next_states.len() > 2 && !path_vertices.contains(&current_loc) {
                path_vertices.push(current_loc.clone());
            }

            frontier.extend(next_states.into_iter());

            seen.insert(current_loc.clone());
        }

        let vertex_indices = HashMap::from_iter(
            path_vertices
                .clone()
                .iter()
                .enumerate()
                .map(|(i, coordinate)| (coordinate.clone(), i))
                .collect::<HashMap<GridCoordinate, usize>>(),
        );

        (path_vertices, vertex_indices)
    }

    fn construct_adjacency_graph(
        &self,
        vertex_indices: &HashMap<GridCoordinate, usize>,
        next_locations_fn: impl Fn(&GridCoordinate) -> Vec<GridCoordinate>,
    ) -> Vec<Vec<Option<usize>>> {
        let v = vertex_indices.len();
        let mut adjacency_graph: Vec<Vec<Option<usize>>> = vec![vec![None; v]; v];

        type MappingState = (GridCoordinate, usize);

        for (vertex, &src) in vertex_indices.iter() {
            let mut frontier =
                VecDeque::<MappingState>::from_iter([(vertex.clone(), 0_usize)].into_iter());
            let mut seen = HashSet::<GridCoordinate>::new();

            while let Some((current_loc, steps)) = frontier.pop_front() {
                if seen.contains(&current_loc) {
                    continue;
                }

                if let Some(&dst) = vertex_indices.get(&current_loc) {
                    adjacency_graph[src][dst] = Some(steps);
                    if src != dst {
                        continue;
                    }
                }

                let next_states = next_locations_fn(&current_loc)
                    .iter()
                    .map(|loc| (loc.clone(), steps + 1))
                    .collect::<Vec<_>>();

                frontier.extend(next_states.into_iter());

                seen.insert(current_loc.clone());
            }
        }

        adjacency_graph
    }

    fn slippery_slope_next_locations(&self, current_loc: &GridCoordinate) -> Vec<GridCoordinate> {
        match self.grid[current_loc.i][current_loc.j] {
            Feature::Path => {
                let mut next_locs = Vec::<GridCoordinate>::with_capacity(4);

                if current_loc.i > 0
                    && self.grid[current_loc.i - 1][current_loc.j] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i - 1,
                        j: current_loc.j,
                    });
                }
                if current_loc.i < self.m - 1
                    && self.grid[current_loc.i + 1][current_loc.j] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i + 1,
                        j: current_loc.j,
                    });
                }
                if current_loc.j > 0
                    && self.grid[current_loc.i][current_loc.j - 1] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i,
                        j: current_loc.j - 1,
                    });
                }
                if current_loc.j < self.n - 1
                    && self.grid[current_loc.i][current_loc.j + 1] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i,
                        j: current_loc.j + 1,
                    });
                }

                next_locs
            }
            Feature::Forest => vec![],
            Feature::NorthSlope => {
                let mut next_locs = Vec::<GridCoordinate>::with_capacity(1);

                if current_loc.i > 0
                    && self.grid[current_loc.i - 1][current_loc.j] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i - 1,
                        j: current_loc.j,
                    });
                }

                next_locs
            }
            Feature::SouthSlope => {
                let mut next_locs = Vec::<GridCoordinate>::with_capacity(1);

                if current_loc.i < self.m - 1
                    && self.grid[current_loc.i + 1][current_loc.j] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i + 1,
                        j: current_loc.j,
                    });
                }

                next_locs
            }
            Feature::EastSlope => {
                let mut next_locs = Vec::<GridCoordinate>::with_capacity(1);

                if current_loc.j < self.n - 1
                    && self.grid[current_loc.i][current_loc.j + 1] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i,
                        j: current_loc.j + 1,
                    });
                }

                next_locs
            }
            Feature::WestSlope => {
                let mut next_locs = Vec::<GridCoordinate>::with_capacity(1);

                if current_loc.j > 0
                    && self.grid[current_loc.i][current_loc.j - 1] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i,
                        j: current_loc.j - 1,
                    });
                }

                next_locs
            }
        }
    }

    fn high_traction_next_locations(&self, current_loc: &GridCoordinate) -> Vec<GridCoordinate> {
        match self.grid[current_loc.i][current_loc.j] {
            Feature::Path
            | Feature::NorthSlope
            | Feature::SouthSlope
            | Feature::EastSlope
            | Feature::WestSlope => {
                let mut next_locs = Vec::<GridCoordinate>::with_capacity(4);

                if current_loc.i > 0
                    && self.grid[current_loc.i - 1][current_loc.j] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i - 1,
                        j: current_loc.j,
                    });
                }
                if current_loc.i < self.m - 1
                    && self.grid[current_loc.i + 1][current_loc.j] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i + 1,
                        j: current_loc.j,
                    });
                }
                if current_loc.j > 0
                    && self.grid[current_loc.i][current_loc.j - 1] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i,
                        j: current_loc.j - 1,
                    });
                }
                if current_loc.j < self.n - 1
                    && self.grid[current_loc.i][current_loc.j + 1] != Feature::Forest
                {
                    next_locs.push(GridCoordinate {
                        i: current_loc.i,
                        j: current_loc.j + 1,
                    });
                }

                next_locs
            }
            Feature::Forest => vec![],
        }
    }

    fn find_longest_path_len(
        adjacency_graph: &Vec<Vec<Option<usize>>>,
        src: usize,
        dst: usize,
    ) -> usize {
        type SearchState = (usize, HashSet<usize>, usize);

        let mut frontier = Vec::<SearchState>::from_iter([(src, HashSet::new(), 0)]);
        let mut longest_path_steps = 0_usize;

        while let Some((current, seen, steps_taken)) = frontier.pop() {
            if seen.contains(&current) {
                continue;
            }

            if current == dst {
                longest_path_steps = usize::max(longest_path_steps, steps_taken);
                continue;
            }

            let next_states = adjacency_graph[current]
                .iter()
                .enumerate()
                .filter_map(|(j, weight)| {
                    if current == j {
                        return None;
                    }

                    match weight {
                        None => None,
                        Some(new_steps) => {
                            let mut new_seen = seen.clone();
                            new_seen.insert(current);
                            Some((j, new_seen, steps_taken + new_steps))
                        }
                    }
                })
                .collect::<Vec<SearchState>>();

            frontier.extend(next_states);
        }

        longest_path_steps
    }
}

#[derive(PartialEq, Eq)]
enum Feature {
    Path,
    Forest,
    NorthSlope,
    SouthSlope,
    EastSlope,
    WestSlope,
}

impl Feature {
    fn try_from_char(c: char) -> Result<Self, String> {
        match c {
            '.' => Ok(Self::Path),
            '#' => Ok(Self::Forest),
            '^' => Ok(Self::NorthSlope),
            'v' => Ok(Self::SouthSlope),
            '>' => Ok(Self::EastSlope),
            '<' => Ok(Self::WestSlope),
            _ => Err(format!("{} is not a map feature", c)),
        }
    }
}

#[derive(Clone, PartialEq, Eq, Hash)]
struct GridCoordinate {
    i: usize,
    j: usize,
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
            #.#####################
            #.......#########...###
            #######.#########.#.###
            ###.....#.>.>.###.#.###
            ###v#####.#v#.###.#.###
            ###.>...#.#.#.....#...#
            ###v###.#.#.#########.#
            ###...#.#.#.......#...#
            #####.#.#.#######.#.###
            #.....#.#.#.......#...#
            #.#####.#.#.#########v#
            #.#...#...#...###...>.#
            #.#.#v#######v###.###v#
            #...#.>.#...>.>.#.###.#
            #####v#.#.###v#.#.###.#
            #.....#...#...#.#.#...#
            #.#########.###.#.#.###
            #...###...#...#...#.###
            ###.###.#.###v#####v###
            #...#...#.#.>.>.#.>.###
            #.###.###.#.###.#.#v###
            #.....###...###...#...#
            #####################.#
        "});
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("94", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            #.#####################
            #.......#########...###
            #######.#########.#.###
            ###.....#.>.>.###.#.###
            ###v#####.#v#.###.#.###
            ###.>...#.#.#.....#...#
            ###v###.#.#.#########.#
            ###...#.#.#.......#...#
            #####.#.#.#######.#.###
            #.....#.#.#.......#...#
            #.#####.#.#.#########v#
            #.#...#...#...###...>.#
            #.#.#v#######v###.###v#
            #...#.>.#...>.>.#.###.#
            #####v#.#.###v#.#.###.#
            #.....#...#...#.#.#...#
            #.#########.###.#.#.###
            #...###...#...#...#.###
            ###.###.#.###v#####v###
            #...#...#.#.>.>.#.>.###
            #.###.###.#.###.#.#v###
            #.....###...###...#...#
            #####################.#
        "});
        let mut solver = Day23::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("154", result);

        Ok(())
    }
}
