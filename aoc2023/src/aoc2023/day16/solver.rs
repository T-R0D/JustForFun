use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day16 {
    contraption: Contraption,
}

impl Day16 {
    pub fn new() -> Self {
        Day16 {
            contraption: Contraption {
                widget_layout: vec![],
                m: 0,
                n: 0,
            },
        }
    }
}

impl Solver for Day16 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.contraption = Contraption::try_from_block(input.as_str())?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let n_energized_tiles = self
            .contraption
            .energize(GridCoordinate { i: 0, j: 0 }, BeamDirection::Right);

        Ok(n_energized_tiles.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut best_energized_count = 0_usize;

        for i in 0..self.contraption.m {
            let starting_coordinate = GridCoordinate { i, j: 0 };
            let starting_beam_direction = BeamDirection::Right;

            let n_energized = self
                .contraption
                .energize(starting_coordinate, starting_beam_direction);

            if n_energized > best_energized_count {
                best_energized_count = n_energized;
            }
        }

        for i in 0..self.contraption.m {
            let starting_coordinate = GridCoordinate {
                i,
                j: self.contraption.n - 1,
            };
            let starting_beam_direction = BeamDirection::Left;

            let n_energized = self
                .contraption
                .energize(starting_coordinate, starting_beam_direction);

            if n_energized > best_energized_count {
                best_energized_count = n_energized;
            }
        }

        for j in 0..self.contraption.n {
            let starting_coordinate = GridCoordinate { i: 0, j };
            let starting_beam_direction = BeamDirection::Down;

            let n_energized = self
                .contraption
                .energize(starting_coordinate, starting_beam_direction);

            if n_energized > best_energized_count {
                best_energized_count = n_energized;
            }
        }

        for j in 0..self.contraption.n {
            let starting_coordinate = GridCoordinate {
                i: self.contraption.m - 1,
                j,
            };
            let starting_beam_direction = BeamDirection::Up;

            let n_energized = self
                .contraption
                .energize(starting_coordinate, starting_beam_direction);

            if n_energized > best_energized_count {
                best_energized_count = n_energized;
            }
        }

        Ok(best_energized_count.to_string())
    }
}

#[derive(Clone)]
struct Contraption {
    widget_layout: Vec<Vec<Widget>>,
    m: usize,
    n: usize,
}

impl Contraption {
    fn try_from_block(block: &str) -> Result<Self, String> {
        let widget_layout = block
            .trim()
            .lines()
            .map(|line| {
                line.chars()
                    .map(Widget::try_from_char)
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        let m = widget_layout.len();
        let n = widget_layout[0].len();

        Ok(Self {
            widget_layout,
            m,
            n,
        })
    }

    fn energize(&self, src_coordinate: GridCoordinate, src_direction: BeamDirection) -> usize {
        let mut energized_layout = vec![vec![0_usize; self.n]; self.m];
        let mut energized_tiles = 0_usize;
        let mut frontier = Vec::<(GridCoordinate, BeamDirection)>::new();
        frontier.push((src_coordinate, src_direction));

        while let Some(operation) = frontier.pop() {
            let (GridCoordinate { i, j }, beam_direction) = operation;
            let widget = &self.widget_layout[i][j];
            let current_cell = energized_layout[i][j];

            match widget {
                Widget::Empty | Widget::VerticalSplitter | Widget::HorizontalSplitter => {
                    if current_cell & (beam_direction.flag() | beam_direction.opposite().flag()) > 0
                    {
                        continue;
                    }
                }
                w @ (Widget::ForwardSlashMirror | Widget::BackSlashMirror) => {
                    let opposite_beam_direction = w
                        .next_beam_direction(beam_direction)
                        .first()
                        .unwrap()
                        .opposite();
                    if current_cell & (beam_direction.flag() | opposite_beam_direction.flag()) > 0 {
                        continue;
                    }
                }
            }

            if current_cell == 0 {
                energized_tiles += 1;
            }

            energized_layout[i][j] |= beam_direction.flag();

            let mut nexts = widget
                .next_beam_direction(beam_direction)
                .iter()
                .filter_map(|&next_beam_direction| {
                    if let Some(coordinate) =
                        next_beam_direction.try_next_coordinate(i, j, self.m, self.n)
                    {
                        Some((coordinate, next_beam_direction))
                    } else {
                        None
                    }
                })
                .collect::<Vec<_>>();

            frontier.append(&mut nexts);
        }

        energized_tiles
    }
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
enum BeamDirection {
    Up = 0b0001,
    Down = 0b0010,
    Left = 0b0100,
    Right = 0b1000,
}

impl BeamDirection {
    fn flag(&self) -> usize {
        *self as usize
    }

    fn try_next_coordinate(
        &self,
        i: usize,
        j: usize,
        m: usize,
        n: usize,
    ) -> Option<GridCoordinate> {
        match self {
            Self::Up => {
                if i > 0 {
                    Some(GridCoordinate { i: i - 1, j })
                } else {
                    None
                }
            }
            Self::Down => {
                if i < m - 1 {
                    Some(GridCoordinate { i: i + 1, j })
                } else {
                    None
                }
            }
            Self::Left => {
                if j > 0 {
                    Some(GridCoordinate { i, j: j - 1 })
                } else {
                    None
                }
            }
            Self::Right => {
                if j < n - 1 {
                    Some(GridCoordinate { i, j: j + 1 })
                } else {
                    None
                }
            }
        }
    }

    fn opposite(&self) -> Self {
        match self {
            Self::Up => Self::Down,
            Self::Down => Self::Up,
            Self::Left => Self::Right,
            Self::Right => Self::Left,
        }
    }
}

#[derive(Clone, PartialEq, Eq)]
enum Widget {
    ForwardSlashMirror,
    BackSlashMirror,
    VerticalSplitter,
    HorizontalSplitter,
    Empty,
}

impl Widget {
    fn try_from_char(c: char) -> Result<Self, String> {
        match c {
            '/' => Ok(Self::ForwardSlashMirror),
            '\\' => Ok(Self::BackSlashMirror),
            '|' => Ok(Self::VerticalSplitter),
            '-' => Ok(Self::HorizontalSplitter),
            '.' => Ok(Self::Empty),
            _ => Err(format!("{c} is not a valid widget")),
        }
    }

    fn next_beam_direction(&self, beam_direction: BeamDirection) -> Vec<BeamDirection> {
        match self {
            Widget::ForwardSlashMirror => match beam_direction {
                BeamDirection::Up => vec![BeamDirection::Right],
                BeamDirection::Down => vec![BeamDirection::Left],
                BeamDirection::Left => vec![BeamDirection::Down],
                BeamDirection::Right => vec![BeamDirection::Up],
            },
            Widget::BackSlashMirror => match beam_direction {
                BeamDirection::Up => vec![BeamDirection::Left],
                BeamDirection::Down => vec![BeamDirection::Right],
                BeamDirection::Left => vec![BeamDirection::Up],
                BeamDirection::Right => vec![BeamDirection::Down],
            },
            Widget::VerticalSplitter => match beam_direction {
                BeamDirection::Up => vec![beam_direction],
                BeamDirection::Down => vec![beam_direction],
                BeamDirection::Left => vec![BeamDirection::Up, BeamDirection::Down],
                BeamDirection::Right => vec![BeamDirection::Up, BeamDirection::Down],
            },
            Widget::HorizontalSplitter => match beam_direction {
                BeamDirection::Up => vec![BeamDirection::Right, BeamDirection::Left],
                BeamDirection::Down => vec![BeamDirection::Right, BeamDirection::Left],
                BeamDirection::Left => vec![beam_direction],
                BeamDirection::Right => vec![beam_direction],
            },
            Widget::Empty => vec![beam_direction],
        }
    }
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
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
        let input = &String::from(indoc! {r"
            .|...\....
            |.-.\.....
            .....|-...
            ........|.
            ..........
            .........\
            ..../.\\..
            .-.-/..|..
            .|....-|.\
            ..//.|....
        "});
        let mut solver = Day16::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("46", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {r"
            .|...\....
            |.-.\.....
            .....|-...
            ........|.
            ..........
            .........\
            ..../.\\..
            .-.-/..|..
            .|....-|.\
            ..//.|....
        "});
        let mut solver = Day16::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("51", result);

        Ok(())
    }
}
