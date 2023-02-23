use std::collections::{HashMap, HashSet, VecDeque};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day24 {
    blizzards: HashMap<Coordinate, Vec<Blizzard>>,
    max_i: usize,
    max_j: usize,
}

impl Day24 {
    pub fn new() -> Self {
        Self {
            blizzards: HashMap::new(),
            max_i: 0,
            max_j: 0,
        }
    }
}

impl Solver for Day24 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.max_i = input.trim().lines().count() - 2;
        for (i, line) in input.trim().lines().enumerate() {
            self.max_j = line.len() - 2;
            for (j, b) in line.bytes().enumerate() {
                if b == b'^' {
                    self.blizzards.insert(
                        Coordinate { i, j },
                        vec![Blizzard::new(Dir::North, Coordinate { i, j })],
                    );
                } else if b == b'v' {
                    self.blizzards.insert(
                        Coordinate { i, j },
                        vec![Blizzard::new(Dir::South, Coordinate { i, j })],
                    );
                } else if b == b'>' {
                    self.blizzards.insert(
                        Coordinate { i, j },
                        vec![Blizzard::new(Dir::East, Coordinate { i, j })],
                    );
                } else if b == b'<' {
                    self.blizzards.insert(
                        Coordinate { i, j },
                        vec![Blizzard::new(Dir::West, Coordinate { i, j })],
                    );
                }
            }
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let blizzard_states = generate_blizzard_states(&self.blizzards, self.max_i, self.max_j);
        let time_to_navigate = navigate_blizzards(
            &Coordinate { i: 0, j: 1 },
            &Coordinate {
                i: self.max_i + 1,
                j: self.max_j,
            },
            0,
            &blizzard_states,
            self.max_i,
            self.max_j,
        );
        Ok(time_to_navigate.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let blizzard_states = generate_blizzard_states(&self.blizzards, self.max_i, self.max_j);
        let valley_entrance = Coordinate { i: 0, j: 1 };
        let valley_exit = Coordinate {
            i: self.max_i + 1,
            j: self.max_j,
        };

        let mut expedition_time = 0;

        let time_to_get_out = navigate_blizzards(
            &valley_entrance,
            &valley_exit,
            expedition_time,
            &blizzard_states,
            self.max_i,
            self.max_j,
        );
        expedition_time += time_to_get_out;

        let time_to_get_snacks = navigate_blizzards(
            &valley_exit,
            &valley_entrance,
            expedition_time,
            &blizzard_states,
            self.max_i,
            self.max_j,
        );
        expedition_time += time_to_get_snacks;

        let time_to_get_out_again = navigate_blizzards(
            &valley_entrance,
            &valley_exit,
            expedition_time,
            &blizzard_states,
            self.max_i,
            self.max_j,
        );
        expedition_time += time_to_get_out_again;

        Ok(expedition_time.to_string())
    }
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
enum Dir {
    North,
    South,
    East,
    West,
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
struct Coordinate {
    i: usize,
    j: usize,
}

impl Coordinate {
    fn blizz_next(&self, dir: &Dir, max_i: usize, max_j: usize) -> Self {
        match dir {
            Dir::North => {
                if self.i == 1 {
                    Self {
                        i: max_i,
                        j: self.j,
                    }
                } else {
                    Self {
                        i: self.i - 1,
                        j: self.j,
                    }
                }
            }
            Dir::South => {
                if self.i == max_i {
                    Self { i: 1, j: self.j }
                } else {
                    Self {
                        i: self.i + 1,
                        j: self.j,
                    }
                }
            }
            Dir::East => {
                if self.j == max_j {
                    Self { i: self.i, j: 1 }
                } else {
                    Self {
                        i: self.i,
                        j: self.j + 1,
                    }
                }
            }
            Dir::West => {
                if self.j == 1 {
                    Self {
                        i: self.i,
                        j: max_j,
                    }
                } else {
                    Self {
                        i: self.i,
                        j: self.j - 1,
                    }
                }
            }
        }
    }
    fn expedition_next(&self, dir: &Dir) -> Self {
        match dir {
            Dir::North => Self {
                i: self.i - 1,
                j: self.j,
            },
            Dir::South => Self {
                i: self.i + 1,
                j: self.j,
            },
            Dir::East => Self {
                i: self.i,
                j: self.j + 1,
            },
            Dir::West => Self {
                i: self.i,
                j: self.j - 1,
            },
        }
    }
}

#[derive(Clone, Eq, Hash, PartialEq)]
struct Blizzard {
    loc: Coordinate,
    dir: Dir,
}

impl Blizzard {
    fn new(dir: Dir, loc: Coordinate) -> Self {
        Self { loc, dir }
    }

    fn next(&self, max_i: usize, max_j: usize) -> Self {
        Self {
            loc: self.loc.blizz_next(&self.dir, max_i, max_j),
            dir: self.dir,
        }
    }
}

#[derive(Clone, Eq, Hash, PartialEq)]
struct SearchState {
    loc: Coordinate,
    t: usize,
}

impl SearchState {
    fn next_steps(&self, max_i: usize, max_j: usize) -> Vec<SearchState> {
        let t = self.t + 1;
        let mut next_states = Vec::<SearchState>::new();
        for dir in [Dir::North, Dir::South, Dir::East, Dir::West] {
            let state = match dir {
                Dir::North => {
                    if self.loc.i > 1 || (self.loc.i == 1 && self.loc.j == 1) {
                        SearchState {
                            loc: self.loc.expedition_next(&dir),
                            t,
                        }
                    } else {
                        continue;
                    }
                }
                Dir::South => {
                    if self.loc.i < max_i || (self.loc.i == max_i && self.loc.j == max_j) {
                        SearchState {
                            loc: self.loc.expedition_next(&dir),
                            t,
                        }
                    } else {
                        continue;
                    }
                }
                Dir::East => {
                    if self.loc.j == max_j || self.loc.i == 0 || self.loc.i == max_i + 1 {
                        continue;
                    }
                    SearchState {
                        loc: self.loc.expedition_next(&dir),
                        t,
                    }
                }
                Dir::West => {
                    if self.loc.j == 1 || self.loc.i == 0 || self.loc.i == max_i + 1 {
                        continue;
                    }
                    SearchState {
                        loc: self.loc.expedition_next(&dir),
                        t,
                    }
                }
            };

            next_states.push(state);
        }
        next_states
    }
}

fn generate_blizzard_states(
    initial_state: &HashMap<Coordinate, Vec<Blizzard>>,
    max_i: usize,
    max_j: usize,
) -> Vec<HashMap<Coordinate, Vec<Blizzard>>> {
    let n_blizzard_states = least_common_multiple(max_i, max_j);
    let mut blizzard_states =
        Vec::<HashMap<Coordinate, Vec<Blizzard>>>::from([initial_state.clone()]);

    for t in 1..n_blizzard_states {
        blizzard_states.push(gen_next_blizzard_state(&blizzard_states[t-1], max_i, max_j));
    }

    blizzard_states
}

fn navigate_blizzards(
    src: &Coordinate,
    dst: &Coordinate,
    start_t: usize,
    blizzard_states: &Vec<HashMap<Coordinate, Vec<Blizzard>>>,
    max_i: usize,
    max_j: usize,
) -> usize {
    let initial_state = SearchState {
        t: start_t,
        loc: src.clone(),
    };
    let mut frontier = VecDeque::<SearchState>::from([initial_state]);
    let mut seen = HashSet::<(Coordinate, usize)>::new();
    let mut path = HashMap::<(Coordinate, usize), (Coordinate, usize)>::new();
    let n_blizzard_states = blizzard_states.len();

    while let Some(current_state) = frontier.pop_front() {
        if current_state.loc == *dst {
            // let mut state = &(current_state.loc, current_state.t);
            // while let Some(previous) = path.get(&state) {
            //     println!("T {} - {} {}", previous.1, previous.0.i, previous.0.j);
            //     println!("{}", blizzard_states[previous.1 % n_blizzard_states].to_string(max_i, max_j));
            //     state = previous;
            // }

            return current_state.t - start_t;
        }

        if seen.contains(&(current_state.loc, current_state.t % n_blizzard_states)) {
            continue;
        }

        let next_t = current_state.t + 1;

        let next_blizzards = &blizzard_states[next_t % n_blizzard_states];

        'next_step_search: for next_step in current_state.next_steps(max_i, max_j).iter() {
            if seen.contains(&(next_step.loc, next_step.t % n_blizzard_states)) {
                continue 'next_step_search;
            }

            if next_blizzards.contains_key(&next_step.loc) {
                continue 'next_step_search;
            }

            frontier.push_back(next_step.clone());
            path.insert(
                (next_step.loc, next_step.t),
                (current_state.loc, current_state.t),
            );
        }

        // Wait option.
        if !next_blizzards.contains_key(&current_state.loc)
            && !seen.contains(&(current_state.loc, (current_state.t + 1) % n_blizzard_states))
        {
            frontier.push_back(SearchState {
                t: next_t,
                loc: current_state.loc.clone(),
            });
            path.insert(
                (current_state.loc, next_t),
                (current_state.loc, current_state.t),
            );
        }

        seen.insert((current_state.loc, current_state.t % n_blizzard_states));
    }

    unreachable!();
}

fn gen_next_blizzard_state(
    prev: &HashMap<Coordinate, Vec<Blizzard>>,
    max_i: usize,
    max_j: usize,
) -> HashMap<Coordinate, Vec<Blizzard>> {
    let mut next = HashMap::<Coordinate, Vec<Blizzard>>::with_capacity(prev.len());

    for blizzards in prev.values() {
        for blizzard in blizzards.iter() {
            let next_blizzard = blizzard.next(max_i, max_j);
            next.entry(next_blizzard.loc)
                .or_insert(Vec::new())
                .push(next_blizzard);
        }
    }

    next
}

fn least_common_multiple(x: usize, y: usize) -> usize {
    let a = std::cmp::max(x, y);
    let b = std::cmp::min(x, y);
    for candidate in (a..=(x * y)).step_by(a) {
        if candidate % b == 0 {
            return candidate;
        }
    }
    return x * y;
}

trait DebugString {
    fn to_string(&self, max_i: usize, max_j: usize) -> String;
}

impl DebugString for HashMap<Coordinate, Vec<Blizzard>> {
    fn to_string(&self, max_i: usize, max_j: usize) -> String {
        let mut s = String::new();
        for i in 0..=max_i {
            'inner: for j in 0..=max_j {
                let loc = &Coordinate { i, j };
                if !self.contains_key(loc) {
                    s.push('.');
                    continue 'inner;
                }

                let blizzards = self.get(loc).unwrap();
                if blizzards.len() == 1 {
                    s.push(match blizzards[0].dir {
                        Dir::North => '^',
                        Dir::South => 'v',
                        Dir::East => '>',
                        Dir::West => '<',
                    })
                } else {
                    s.push((blizzards.len() as u8 + b'0') as char);
                }
            }
            s.push('\n');
        }

        s
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
            #.######
            #>>.<^<#
            #.<..<<#
            #>v.><>#
            #<^v^^>#
            ######.#
        "});
        let mut solver = Day24::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("18", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            #.######
            #>>.<^<#
            #.<..<<#
            #>v.><>#
            #<^v^^>#
            ######.#
        "});
        let mut solver = Day24::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("54", result);

        Ok(())
    }
}
