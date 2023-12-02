use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day02 {
    game_results: Vec<Vec<CubeCount>>,
}

impl Day02 {
    pub fn new() -> Self {
        Day02 {
            game_results: vec![],
        }
    }
}

impl Solver for Day02 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let lines = input.trim().split("\n").collect::<Vec<_>>();
        let game_results = lines
            .iter()
            .map(|line| {
                let line_parts = line.split(": ").collect::<Vec<_>>();
                if line_parts.len() != 2 {
                    return Err(format!("Line with contents '{line}' was invalid"));
                }

                let reveals = line_parts[1]
                    .split("; ")
                    .map(|reveal_description| {
                        CubeCount::from_text_description(&String::from(reveal_description))
                    })
                    .collect::<Result<_, _>>()?;

                Ok(reveals)
            })
            .collect::<Result<Vec<_>, String>>()?;

        self.game_results = game_results;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let target_cube_count = CubeCount {
            r: 12,
            g: 13,
            b: 14,
        };

        let valid_game_ids = self
            .game_results
            .iter()
            .enumerate()
            .filter_map(|(i, reveals)| {
                for reveal in reveals.iter() {
                    if reveal.r > target_cube_count.r
                        || reveal.g > target_cube_count.g
                        || reveal.b > target_cube_count.b
                    {
                        return None;
                    }
                }
                // Games are at index ID - 1.
                Some(i + 1)
            })
            .collect::<Vec<_>>();

        let valid_game_id_sum = valid_game_ids.iter().sum::<usize>();

        Ok(valid_game_id_sum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let game_powers = self
            .game_results
            .iter()
            .map(|reveals| {
                reveals.iter().fold(
                    CubeCount {
                        r: 0,
                        g: 0,
                        b: 0,
                    },
                    |acc, reveal| CubeCount::fold(&acc, reveal),
                )
            })
            .map(|min_cube_count| min_cube_count.power())
            .collect::<Vec<_>>();

        let power_sum = game_powers.iter().sum::<u32>();

        Ok(power_sum.to_string())
    }
}

struct CubeCount {
    pub r: u32,
    pub g: u32,
    pub b: u32,
}

impl CubeCount {
    pub fn from_text_description(description: &String) -> Result<Self, String> {
        let mut cube_count = Self { r: 0, g: 0, b: 0 };

        let tuples = description.split(", ").collect::<Vec<_>>();
        for tuple in tuples.iter() {
            let parts = tuple.split(" ").collect::<Vec<_>>();
            if parts.len() != 2 {
                return Err(format!("{tuple} could not be parsed"));
            }
            let count_result = parts[0].parse::<u32>();
            let color = parts[1];
            match (count_result, color) {
                (Ok(count), "red") => {
                    cube_count.r = count;
                    Ok(())
                }
                (Ok(count), "green") => {
                    cube_count.g = count;
                    Ok(())
                }
                (Ok(count), "blue") => {
                    cube_count.b = count;
                    Ok(())
                }
                _ => Err(format!("Tuple '{tuple}' was not valid")),
            }?
        }

        Ok(cube_count)
    }

    pub fn fold(one: &Self, another: &Self) -> Self {
        Self {
            r: u32::max(one.r, another.r),
            g: u32::max(one.g, another.g),
            b: u32::max(one.b, another.b),
        }
    }

    pub fn power(&self) -> u32 {
        self.r * self.g * self.b
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
            Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
            Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
            Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
            Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
            Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
        "});
        let mut solver = Day02::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("8", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
            Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
            Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
            Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
            Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
        "});
        let mut solver = Day02::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("2286", result);

        Ok(())
    }
}
