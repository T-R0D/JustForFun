use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day18 {
    instructions: Vec<DigInstruction>,
}

impl Day18 {
    pub fn new() -> Self {
        Day18 {
            instructions: vec![],
        }
    }
}

impl Solver for Day18 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.instructions = input
            .trim()
            .lines()
            .map(DigInstruction::try_from_line)
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let area = Day18::find_dig_area(&self.instructions);

        Ok(area.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let corrected_instructions = self
            .instructions
            .iter()
            .map(|instruction| instruction.color.to_dig_instruction())
            .collect::<Vec<_>>();

        let area = Day18::find_dig_area(&corrected_instructions);

        Ok(area.to_string())
    }
}

impl Day18 {
    /// Computes the area of a dug pit based on trench digging instructions.
    /// Uses a combination of the Shoelace Formula..
    /// The Shoelace Formula finds the area of a polygon by the formula:
    /// 
    /// ```latex
    /// A = (1/2) \sigma_{i=0}^{n-1}(x_{i+1,} * y_{i} - y_{i+1} * x_{i}).
    /// ```
    /// 
    /// Consider the Shoelace formula as having its vertices centered in the squares of the dug out trench.
    /// This would leave a border of have of the square exterior to the dug out shape whose area is computed
    /// by the Shoelace Formula. The area of this exterior border would be computed by:
    /// 
    /// ```latex
    /// A = trench-length + 1
    /// ```
    /// 
    /// The `+1` comes from the additional quarter of a square exterior space on "projecting" corners. One may 
    /// "but there could be many projecting corners!", but for each additional projecting corner beyond
    /// 4, there will be "inward" corners that will actually need 1/4 unit _less_ border, so things
    /// balance out.
    fn find_dig_area(instructions: &Vec<DigInstruction>) -> isize {
        let mut border_length = 0_isize;
        let mut previous_position = (0_isize, 0_isize);
        let mut shoelace_area = 0_isize;
        for DigInstruction {
            direction,
            magnitude,
            ..
        } in instructions.iter()
        {
            let (i_start, j_start) = previous_position;
            let (i_end, j_end) = match direction {
                Direction::Up => (i_start - magnitude, j_start),
                Direction::Down => (i_start + magnitude, j_start),
                Direction::Left => (i_start, j_start - magnitude),
                Direction::Right => (i_start, j_start + magnitude),
            };

            shoelace_area += i_end * j_start - j_end * i_start;
            border_length += magnitude;
            previous_position = (i_end, j_end);
        }

        // These were double counted.
        shoelace_area /= 2;

        shoelace_area + ((border_length / 2) + 1)
    }
}

struct DigInstruction {
    direction: Direction,
    magnitude: isize,
    color: RGB,
}

impl DigInstruction {
    fn try_from_line(line: &str) -> Result<Self, String> {
        let trimmed_line = line.replace(")", "");
        let parts = trimmed_line.split(" (#").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("{line} was bad format"));
        }

        let move_parts = parts[0].split(" ").collect::<Vec<_>>();
        if move_parts.len() != 2 {
            return Err(format!("{} was bad format for movement", parts[0]));
        }

        let direction = Direction::try_from_str(move_parts[0])?;
        let magnitude = match move_parts[1].parse::<isize>() {
            Ok(value) => Ok(value),
            Err(err) => Err(err.to_string()),
        }?;

        let color = RGB::try_from_str(parts[1])?;

        Ok(Self {
            direction,
            magnitude,
            color,
        })
    }
}

#[derive(Clone, Copy)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    fn try_from_str(c: &str) -> Result<Self, String> {
        match c {
            "U" => Ok(Self::Up),
            "D" => Ok(Self::Down),
            "L" => Ok(Self::Left),
            "R" => Ok(Self::Right),
            _ => Err(format!("{c} is not a valid Direction")),
        }
    }
}

struct RGB {
    r: u8,
    g: u8,
    b: u8,
}

impl RGB {
    fn try_from_str(s: &str) -> Result<Self, String> {
        if s.len() != 6 {
            return Err(format!("{s} is not 6 characters long"));
        }

        let values = s
            .chars()
            .step_by(2)
            .zip(s.chars().skip(1).step_by(2))
            .map(|(a, b)| match (a.to_digit(16), b.to_digit(16)) {
                (Some(hex_a), Some(hex_b)) => Ok(hex_a * 16 + hex_b),
                _ => Err(format!("{a} {b} could not be converted to a hex number")),
            })
            .collect::<Result<Vec<_>, String>>()?;

        Ok(Self {
            r: values[0] as u8,
            g: values[1] as u8,
            b: values[2] as u8,
        })
    }

    fn to_dig_instruction(&self) -> DigInstruction {
        let mut magnitude = 0_isize;

        magnitude |= isize::from(self.r);

        magnitude <<= 8;
        magnitude |= isize::from(self.g);

        magnitude <<= 4;
        magnitude |= isize::from(self.b) >> 4;

        let direction = match self.b & 0x03 {
            0x00 => Direction::Right,
            0x01 => Direction::Down,
            0x02 => Direction::Left,
            0x03 => Direction::Up,
            _ => Direction::Up,
        };

        DigInstruction {
            direction,
            magnitude,
            color: RGB { r: 0, g: 0, b: 0 },
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
            R 6 (#70c710)
            D 5 (#0dc571)
            L 2 (#5713f0)
            D 2 (#d2c081)
            R 2 (#59c680)
            D 2 (#411b91)
            L 5 (#8ceee2)
            U 2 (#caa173)
            L 1 (#1b58a2)
            U 2 (#caa171)
            R 2 (#7807d2)
            U 3 (#a77fa3)
            L 2 (#015232)
            U 2 (#7a21e3)
        "});
        let mut solver = Day18::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("62", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            R 6 (#70c710)
            D 5 (#0dc571)
            L 2 (#5713f0)
            D 2 (#d2c081)
            R 2 (#59c680)
            D 2 (#411b91)
            L 5 (#8ceee2)
            U 2 (#caa173)
            L 1 (#1b58a2)
            U 2 (#caa171)
            R 2 (#7807d2)
            U 3 (#a77fa3)
            L 2 (#015232)
            U 2 (#7a21e3)
        "});
        let mut solver = Day18::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("952408144115", result);

        Ok(())
    }
}
