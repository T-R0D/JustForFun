use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day09 {
    reading_histories: Vec<Vec<i32>>,
}

impl Day09 {
    pub fn new() -> Self {
        Day09 {
            reading_histories: vec![],
        }
    }
}

impl Solver for Day09 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.reading_histories = input
            .trim()
            .lines()
            .map(|line| {
                line.split(" ")
                    .map(|item| match item.parse::<i32>() {
                        Ok(value) => Ok(value),
                        Err(err) => Err(err.to_string()),
                    })
                    .collect::<Result<Vec<_>, String>>()
            })
            .collect::<Result<Vec<Vec<_>>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let predictions = self
            .reading_histories
            .iter()
            .map(|reading_history| Day09::try_compute_prediction(reading_history))
            .collect::<Result<Vec<i32>, String>>()?;

        let prediction_sum = predictions.iter().sum::<i32>();

        Ok(prediction_sum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let extrapolations = self
            .reading_histories
            .iter()
            .map(|reading_history| Day09::try_extrapolate_history(reading_history))
            .collect::<Result<Vec<i32>, String>>()?;

        let extrapolation_sum = extrapolations.iter().sum::<i32>();

        Ok(extrapolation_sum.to_string())
    }
}

impl Day09 {
    fn try_compute_prediction(reading_history: &Vec<i32>) -> Result<i32, String> {
        let mut current_sequence = reading_history.clone();
        let mut last_vals = vec![];

        while current_sequence.len() > 1 && current_sequence.iter().any(|&x| x != 0) {
            last_vals.push(
                *current_sequence
                    .last()
                    .ok_or(String::from("next sequence had no last value"))?,
            );
            current_sequence = current_sequence
                .iter()
                .zip(current_sequence.iter().skip(1))
                .map(|(&a, &b)| b - a)
                .collect::<Vec<_>>();
        }

        if current_sequence.len() <= 1 {
            return Err(String::from(
                "resulting sequence disappeared before becoming all zeroes",
            ));
        }

        let mut prediction = 0;
        for &last_val in last_vals.iter().rev() {
            prediction += last_val;
        }

        Ok(prediction)
    }

    fn try_extrapolate_history(reading_history: &Vec<i32>) -> Result<i32, String> {
        let mut current_sequence = reading_history.clone();
        let mut first_vals = vec![];

        while current_sequence.len() > 1 && current_sequence.iter().any(|&x| x != 0) {
            first_vals.push(
                *current_sequence
                    .first()
                    .ok_or(String::from("next sequence had no last value"))?,
            );
            current_sequence = current_sequence
                .iter()
                .zip(current_sequence.iter().skip(1))
                .map(|(&a, &b)| b - a)
                .collect::<Vec<_>>();

            // println!("{:?}\n{:?}", first_vals, current_sequence);
        }

        if current_sequence.len() <= 1 {
            return Err(String::from(
                "resulting sequence disappeared before becoming all zeroes",
            ));
        }

        let mut prediction = 0;
        for &first_val in first_vals.iter().rev() {
            prediction = first_val - prediction;
        }

        Ok(prediction)
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
            0 3 6 9 12 15
            1 3 6 10 15 21
            10 13 16 21 30 45
        "});
        let mut solver = Day09::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("114", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            0 3 6 9 12 15
            1 3 6 10 15 21
            10 13 16 21 30 45
        "});
        let mut solver = Day09::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("2", result);

        Ok(())
    }
}
