// TODO: Figure out how to not hardcode.
// TODO: Figure out how to switch between cube size 4 and 50 for tests.

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day22 {
    actions: Vec<Action>,
    map: Vec<Vec<u8>>,
}

impl Day22 {
    pub fn new() -> Self {
        Self {
            actions: Vec::new(),
            map: Vec::new(),
        }
    }
}

impl Solver for Day22 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let parts = input.split("\n\n").collect::<Vec<_>>();

        self.map = parse_map(parts[0]);
        self.actions = parse_actions(parts[1]);

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut navigator = PathLockNavigator::new(&self.map);

        for action in self.actions.iter() {
            navigator.take_action(action);
        }

        let password = navigator.compute_password();

        Ok(password.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut navigator = CubeLockNavigator::new(&self.map, 50);

        for action in self.actions.iter() {
            navigator.take_action(action);
        }

        let password = navigator.compute_password();

        Ok(password.to_string())
    }
}

fn parse_map(map_spec: &str) -> Vec<Vec<u8>> {
    let mut map = Vec::<Vec<u8>>::new();

    let mut n_cols = 0;
    for line in map_spec.lines() {
        if line.len() > n_cols {
            n_cols = line.len();
        }
    }

    for line in map_spec.lines() {
        let mut row = Vec::<u8>::new();
        let mut j = 0;
        for c in line.bytes() {
            row.push(c);
            j += 1;
        }
        while j < n_cols {
            row.push(EMPTY);
            j += 1;
        }
        map.push(row);
    }

    map
}

fn parse_actions(action_spec: &str) -> Vec<Action> {
    let mut actions = Vec::<Action>::new();
    let mut num_buf = Vec::<u8>::new();

    for c in action_spec.trim().bytes() {
        if c == b'L' {
            if num_buf.len() > 0 {
                let magnitude = String::from_utf8(num_buf.clone())
                    .unwrap()
                    .parse::<usize>()
                    .unwrap();
                actions.push(Action::Move(magnitude));
                num_buf.clear();
            }

            actions.push(Action::TurnLeft);
        } else if c == b'R' {
            if num_buf.len() > 0 {
                let magnitude = String::from_utf8(num_buf.clone())
                    .unwrap()
                    .parse::<usize>()
                    .unwrap();
                actions.push(Action::Move(magnitude));
                num_buf.clear();
            }

            actions.push(Action::TurnRight);
        } else if c.is_ascii_digit() {
            num_buf.push(c);
        } else {
            panic!("{c} not parseable!");
        }
    }
    if num_buf.len() > 0 {
        let magnitude = String::from_utf8(num_buf.clone())
            .unwrap()
            .parse::<usize>()
            .unwrap();
        actions.push(Action::Move(magnitude));
        num_buf.clear();
    }

    actions
}

#[derive(Debug)]
enum Action {
    Move(usize),
    TurnLeft,
    TurnRight,
}

#[derive(Clone, Copy, Eq, PartialEq)]
enum Facing {
    Up,
    Down,
    Right,
    Left,
}

impl Facing {
    fn turn(&self, action: &Action) -> Facing {
        match action {
            Action::Move(_) => unreachable!(),
            Action::TurnLeft => match self {
                Facing::Up => Facing::Left,
                Facing::Down => Facing::Right,
                Facing::Right => Facing::Up,
                Facing::Left => Facing::Down,
            },
            Action::TurnRight => match self {
                Facing::Up => Facing::Right,
                Facing::Down => Facing::Left,
                Facing::Right => Facing::Down,
                Facing::Left => Facing::Up,
            },
        }
    }
}

const EMPTY: u8 = b' ';
const STEP: u8 = b'.';
const WALL: u8 = b'#';

struct PathLockNavigator {
    map: Vec<Vec<u8>>,
    n_cols: usize,
    n_rows: usize,
    cur_i: usize,
    cur_j: usize,
    facing: Facing,
}

impl PathLockNavigator {
    fn new(map: &Vec<Vec<u8>>) -> Self {
        let mut start_i = 0;
        let mut start_j = 0;
        'find_start: for (i, row) in map.iter().enumerate() {
            for (j, &tile) in row.iter().enumerate() {
                if tile == STEP {
                    start_i = i;
                    start_j = j;
                    break 'find_start;
                }
            }
        }

        Self {
            map: map.clone(),
            facing: Facing::Right,
            cur_i: start_i,
            cur_j: start_j,
            n_cols: map[0].len(),
            n_rows: map.len(),
        }
    }

    fn take_action(&mut self, action: &Action) {
        match action {
            Action::TurnLeft | Action::TurnRight => {
                self.facing = self.facing.turn(action);
            }
            Action::Move(magnitude) => match self.facing {
                Facing::Up => self.move_vertically(*magnitude, true),
                Facing::Down => self.move_vertically(*magnitude, false),
                Facing::Right => self.move_horizontally(*magnitude, false),
                Facing::Left => self.move_horizontally(*magnitude, true),
            },
        }
    }

    fn move_horizontally(&mut self, magnitude: usize, move_left: bool) {
        let mut increment = 1;
        if move_left {
            increment = -1;
        }

        let mut j = self.cur_j as i32;
        'stepping: for _ in 0..magnitude {
            j += increment;
            if j < 0 {
                j += self.n_cols as i32;
            }

            while self.map[self.cur_i][(j % self.n_cols as i32) as usize] == EMPTY {
                j += increment;
                if j < 0 {
                    j += self.n_cols as i32;
                }
            }

            if self.map[self.cur_i][(j % self.n_cols as i32) as usize] == WALL {
                break 'stepping;
            }

            self.cur_j = (j % self.n_cols as i32) as usize;
        }
    }

    fn move_vertically(&mut self, magnitude: usize, move_up: bool) {
        let mut increment = 1;
        if move_up {
            increment = -1;
        }

        let mut i = self.cur_i as i32;
        'stepping: for _ in 0..magnitude {
            i += increment;
            if i < 0 {
                i += self.n_rows as i32;
            }

            while self.map[(i % self.n_rows as i32) as usize][self.cur_j] == EMPTY {
                i += increment;
                if i < 0 {
                    i += self.n_rows as i32;
                }
            }

            if self.map[(i % self.n_rows as i32) as usize][self.cur_j] == WALL {
                break 'stepping;
            }

            self.cur_i = (i % self.n_rows as i32) as usize;
        }
    }

    fn compute_password(&self) -> usize {
        ((self.cur_i + 1) * 1000)
            + ((self.cur_j + 1) * 4)
            + match self.facing {
                Facing::Up => 3,
                Facing::Down => 1,
                Facing::Right => 0,
                Facing::Left => 2,
            }
    }
}

struct FaceMapping {
    up: (usize, ResultingSide),
    right: (usize, ResultingSide),
    down: (usize, ResultingSide),
    left: (usize, ResultingSide),
}

struct ResultingSide {
    side: Facing,
    facing: Facing,
    inverted: bool,
}

struct CubeLockNavigator {
    cur_i: usize,
    cur_j: usize,
    cur_face: usize,
    facing: Facing,
    cube_faces: Vec<Vec<Vec<u8>>>,
    face_mappings: Vec<FaceMapping>,
    face_size: usize,
}

impl CubeLockNavigator {
    fn new(map: &Vec<Vec<u8>>, face_size: usize) -> Self {
        let (cube_faces, face_mappings) = Self::identify_cube_sides(map, face_size);

        Self {
            cur_i: 0,
            cur_j: 0,
            cur_face: 0,
            facing: Facing::Right,
            cube_faces,
            face_mappings,
            face_size,
        }
    }

    fn identify_cube_sides(
        map: &Vec<Vec<u8>>,
        face_size: usize,
    ) -> (Vec<Vec<Vec<u8>>>, Vec<FaceMapping>) {
        let mut cube_sides = Vec::<Vec<Vec<u8>>>::with_capacity(6);
        let mut face_nums = Vec::<usize>::new();

        let rows = map.len() / face_size;
        let cols = map[0].len() / face_size;

        for face_i in 0..rows {
            for face_j in 0..cols {
                if map[face_i * face_size][face_j * face_size] != EMPTY {
                    let mut face = Vec::<Vec<u8>>::with_capacity(face_size);
                    for i in 0..face_size {
                        let mut row = Vec::<u8>::with_capacity(face_size);
                        for j in 0..face_size {
                            row.push(map[face_i * face_size + i][face_j * face_size + j]);
                        }
                        face.push(row);
                    }
                    cube_sides.push(face);
                    face_nums.push(face_i * rows + face_j);
                }
            }
        }

        let mut face_mappings = Vec::<FaceMapping>::with_capacity(6);
        if face_size == 50 {
            face_mappings.push(FaceMapping {
                up: (
                    5,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                right: (
                    1,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                down: (
                    2,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
                left: (
                    3,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: true,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    5,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                right: (
                    4,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: true,
                    },
                ),
                down: (
                    2,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
                left: (
                    0,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    0,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                right: (
                    1,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                down: (
                    4,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
                left: (
                    3,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    2,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                right: (
                    4,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                down: (
                    5,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
                left: (
                    0,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: true,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    2,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                right: (
                    1,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: true,
                    },
                ),
                down: (
                    5,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
                left: (
                    3,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    3,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                right: (
                    4,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                down: (
                    1,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
                left: (
                    0,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
            });
        } else {
            face_mappings.push(FaceMapping {
                up: (
                    1,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: true,
                    },
                ),
                right: (
                    5,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: true,
                    },
                ),
                down: (
                    3,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
                left: (
                    2,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    0,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: true,
                    },
                ),
                right: (
                    2,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                down: (
                    4,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: true,
                    },
                ),
                left: (
                    5,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: true,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    0,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                right: (
                    3,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                down: (
                    4,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: true,
                    },
                ),
                left: (
                    1,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    0,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                right: (
                    5,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: true,
                    },
                ),
                down: (
                    4,
                    ResultingSide {
                        side: Facing::Up,
                        facing: Facing::Down,
                        inverted: false,
                    },
                ),
                left: (
                    2,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    3,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: false,
                    },
                ),
                right: (
                    5,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: false,
                    },
                ),
                down: (
                    1,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: true,
                    },
                ),
                left: (
                    2,
                    ResultingSide {
                        side: Facing::Down,
                        facing: Facing::Up,
                        inverted: true,
                    },
                ),
            });
            face_mappings.push(FaceMapping {
                up: (
                    3,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: true,
                    },
                ),
                right: (
                    0,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: true,
                    },
                ),
                down: (
                    1,
                    ResultingSide {
                        side: Facing::Left,
                        facing: Facing::Right,
                        inverted: true,
                    },
                ),
                left: (
                    4,
                    ResultingSide {
                        side: Facing::Right,
                        facing: Facing::Left,
                        inverted: false,
                    },
                ),
            });
        }

        (cube_sides, face_mappings)
    }

    fn take_action(&mut self, action: &Action) {
        match action {
            Action::TurnLeft | Action::TurnRight => {
                self.facing = self.facing.turn(action);
            }
            Action::Move(magnitude) => match self.facing {
                Facing::Up => self.move_vertically(*magnitude, true),
                Facing::Down => self.move_vertically(*magnitude, false),
                Facing::Right => self.move_horizontally(*magnitude, false),
                Facing::Left => self.move_horizontally(*magnitude, true),
            },
        }
    }

    fn move_horizontally(&mut self, magnitude: usize, move_left: bool) {
        let mut increment: i32 = 1;
        if move_left {
            increment = -1;
        }

        'stepping: for s in 0..magnitude {
            let face_mapping = &self.face_mappings[self.cur_face];
            if self.cur_j == 0 && increment == -1 {
                let (next_face, resulting_side) = &(face_mapping.left);

                let new_face = &self.cube_faces[*next_face];
                let (new_i, new_j) = match resulting_side.side {
                    Facing::Up => {
                        if resulting_side.inverted {
                            (0, (self.face_size - 1) - self.cur_i)
                        } else {
                            (0, self.cur_i)
                        }
                    }
                    Facing::Down => {
                        if resulting_side.inverted {
                            (self.face_size - 1, (self.face_size - 1) - self.cur_i)
                        } else {
                            (self.face_size - 1, self.cur_i)
                        }
                    }
                    Facing::Left => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_i, 0)
                        } else {
                            (self.cur_i, 0)
                        }
                    }
                    Facing::Right => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_i, self.face_size - 1)
                        } else {
                            (self.cur_i, self.face_size - 1)
                        }
                    }
                };

                if new_face[new_i][new_j] == WALL {
                    break 'stepping;
                }

                self.cur_i = new_i;
                self.cur_j = new_j;
                self.cur_face = *next_face;
                self.facing = resulting_side.facing;

                match resulting_side.facing {
                    Facing::Up => self.move_vertically(magnitude - (s + 1), true),
                    Facing::Down => self.move_vertically(magnitude - (s + 1), false),
                    Facing::Right => self.move_horizontally(magnitude - (s + 1), false),
                    Facing::Left => self.move_horizontally(magnitude - (s + 1), true),
                }
                break 'stepping;
            } else if self.cur_j == self.face_size - 1 && increment == 1 {
                let (next_face, resulting_side) = &(face_mapping.right);

                let new_face = &self.cube_faces[*next_face];
                let (new_i, new_j) = match resulting_side.side {
                    Facing::Up => {
                        if resulting_side.inverted {
                            (0, (self.face_size - 1) - self.cur_i)
                        } else {
                            (0, self.cur_i)
                        }
                    }
                    Facing::Down => {
                        if resulting_side.inverted {
                            (self.face_size - 1, (self.face_size - 1) - self.cur_i)
                        } else {
                            (self.face_size - 1, self.cur_i)
                        }
                    }
                    Facing::Left => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_i, 0)
                        } else {
                            (self.cur_i, 0)
                        }
                    }
                    Facing::Right => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_i, self.face_size - 1)
                        } else {
                            (self.cur_i, self.face_size - 1)
                        }
                    }
                };

                if new_face[new_i][new_j] == WALL {
                    break 'stepping;
                }

                self.cur_i = new_i;
                self.cur_j = new_j;
                self.cur_face = *next_face;
                self.facing = resulting_side.facing;

                match resulting_side.facing {
                    Facing::Up => self.move_vertically(magnitude - (s + 1), true),
                    Facing::Down => self.move_vertically(magnitude - (s + 1), false),
                    Facing::Right => self.move_horizontally(magnitude - (s + 1), false),
                    Facing::Left => self.move_horizontally(magnitude - (s + 1), true),
                }
                break 'stepping;
            } else {
                if self.cube_faces[self.cur_face][self.cur_i]
                    [((self.cur_j as i32) + increment) as usize]
                    == WALL
                {
                    break 'stepping;
                }
                self.cur_j = ((self.cur_j as i32) + increment) as usize;
            }
        }
    }

    fn move_vertically(&mut self, magnitude: usize, move_up: bool) {
        let mut increment: i32 = 1;
        if move_up {
            increment = -1;
        }

        'stepping: for s in 0..magnitude {
            let face_mapping = &self.face_mappings[self.cur_face];
            if self.cur_i == 0 && increment == -1 {
                let (next_face, resulting_side) = &(face_mapping.up);

                let new_face = &self.cube_faces[*next_face];
                let (new_i, new_j) = match resulting_side.side {
                    Facing::Up => {
                        if resulting_side.inverted {
                            (0, (self.face_size - 1) - self.cur_j)
                        } else {
                            (0, self.cur_j)
                        }
                    }
                    Facing::Down => {
                        if resulting_side.inverted {
                            (self.face_size - 1, (self.face_size - 1) - self.cur_j)
                        } else {
                            (self.face_size - 1, self.cur_j)
                        }
                    }
                    Facing::Left => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_j, 0)
                        } else {
                            (self.cur_j, 0)
                        }
                    }
                    Facing::Right => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_j, self.face_size - 1)
                        } else {
                            (self.cur_j, self.face_size - 1)
                        }
                    }
                };

                if new_face[new_i][new_j] == WALL {
                    break 'stepping;
                }

                self.cur_i = new_i;
                self.cur_j = new_j;
                self.cur_face = *next_face;
                self.facing = resulting_side.facing;

                match resulting_side.facing {
                    Facing::Right => self.move_horizontally(magnitude - (s + 1), false),
                    Facing::Left => self.move_horizontally(magnitude - (s + 1), true),
                    Facing::Up => self.move_vertically(magnitude - (s + 1), true),
                    Facing::Down => self.move_vertically(magnitude - (s + 1), false),
                }
                break 'stepping;
            } else if self.cur_i == self.face_size - 1 && increment == 1 {
                let (next_face, resulting_side) = &(face_mapping.down);

                let new_face = &self.cube_faces[*next_face];
                let (new_i, new_j) = match resulting_side.side {
                    Facing::Up => {
                        if resulting_side.inverted {
                            (0, self.face_size - self.cur_j)
                        } else {
                            (0, self.cur_j)
                        }
                    }
                    Facing::Down => {
                        if resulting_side.inverted {
                            (self.face_size - 1, (self.face_size - 1) - self.cur_j)
                        } else {
                            (self.face_size - 1, self.cur_j)
                        }
                    }
                    Facing::Left => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_j, 0)
                        } else {
                            (self.cur_j, 0)
                        }
                    }
                    Facing::Right => {
                        if resulting_side.inverted {
                            ((self.face_size - 1) - self.cur_j, self.face_size - 1)
                        } else {
                            (self.cur_j, self.face_size - 1)
                        }
                    }
                };

                if new_face[new_i][new_j] == WALL {
                    break 'stepping;
                }

                self.cur_i = new_i;
                self.cur_j = new_j;
                self.cur_face = *next_face;
                self.facing = resulting_side.facing;

                match resulting_side.facing {
                    Facing::Right => self.move_horizontally(magnitude - (s + 1), false),
                    Facing::Left => self.move_horizontally(magnitude - (s + 1), true),
                    Facing::Up => self.move_vertically(magnitude - (s + 1), true),
                    Facing::Down => self.move_vertically(magnitude - (s + 1), false),
                }
                break 'stepping;
            } else {
                if self.cube_faces[self.cur_face][((self.cur_i as i32) + increment) as usize]
                    [self.cur_j]
                    == WALL
                {
                    break 'stepping;
                }
                self.cur_i = (self.cur_i as i32 + increment) as usize;
            }
        }
    }

    fn compute_password(&self) -> usize {
        let (i, j) = self.get_map_coordinates();

        ((i + 1) * 1000)
            + ((j + 1) * 4)
            + match self.facing {
                Facing::Up => 3,
                Facing::Down => 1,
                Facing::Right => 0,
                Facing::Left => 2,
            }
    }

    fn get_map_coordinates(&self) -> (usize, usize) {
        let (i_multiplier, j_multiplier) = if self.face_size == 50 {
            let segments = [1, 2, 4, 6, 7, 9];
            ((segments[self.cur_face] / 3), (segments[self.cur_face] % 3))
        } else {
            let segments = [2, 4, 5, 6, 10, 11];
            ((segments[self.cur_face] / 4), (segments[self.cur_face] % 4))
        };

        let i = self.cur_i + (i_multiplier * self.face_size);
        let j = self.cur_j + (j_multiplier * self.face_size);

        (i, j)
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
                    ...#
                    .#..
                    #...
                    ....
            ...#.......#
            ........#...
            ..#....#....
            ..........#.
                    ...#....
                    .....#..
                    .#......
                    ......#.

            10R5L5R10L4R5L5
        "});
        let mut solver = Day22::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("6032", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
                    ...#
                    .#..
                    #...
                    ....
            ...#.......#
            ........#...
            ..#....#....
            ..........#.
                    ...#....
                    .....#..
                    .#......
                    ......#.

            10R5L5R10L4R5L5
        "});
        let mut solver = Day22::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("5031", result);

        Ok(())
    }
}
