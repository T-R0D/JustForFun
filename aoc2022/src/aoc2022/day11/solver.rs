use std::collections::VecDeque;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day11 {
    starting_monkey_configuration: Vec<Monkey>,
}

impl Day11 {
    pub fn new() -> Self {
        Day11 {
            starting_monkey_configuration: vec![],
        }
    }
}

impl Solver for Day11 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.starting_monkey_configuration = input
            .trim()
            .split("\n\n")
            .map(Monkey::from_config)
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut n_examined_items =
            run_monkey_keep_away_simulation(&self.starting_monkey_configuration);
        n_examined_items.sort();
        let monkey_business_level = n_examined_items[n_examined_items.len() - 1]
            * n_examined_items[n_examined_items.len() - 2];
        Ok(monkey_business_level.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut n_examined_items = run_monkey_keep_away_simulation_without_relief_factor(
            &self.starting_monkey_configuration,
        );
        n_examined_items.sort();
        let monkey_business_level = n_examined_items[n_examined_items.len() - 1]
            * n_examined_items[n_examined_items.len() - 2];
        Ok(monkey_business_level.to_string())
    }
}

#[derive(Clone)]
struct Monkey {
    held: VecDeque<u64>,
    operation: Op,
    test: u64,
    if_true: usize,
    if_false: usize,
}

#[derive(Clone)]
enum Op {
    Add(u64),
    Mul(u64),
    Square,
}

impl Monkey {
    fn from_config(config: &str) -> Self {
        let lines = config.lines().collect::<Vec<_>>();
        let held = lines[1]
            .replace("  Starting items: ", "")
            .split(", ")
            .map(|item| item.parse::<u64>().expect("items should parse as u64"))
            .collect::<VecDeque<_>>();
        let operation_essence = lines[2].replace("  Operation: new = old ", "");
        let operation_parts = operation_essence.split(" ").collect::<Vec<_>>();
        let operation = match (operation_parts[0], operation_parts[1]) {
            ("+", addend) => Op::Add(addend.parse::<u64>().expect("op should parse as u64")),
            ("*", "old") => Op::Square,
            ("*", multiplicand) => {
                Op::Mul(multiplicand.parse::<u64>().expect("op should parse as u64"))
            }
            _ => panic!(
                "'{} {}' was unparseable",
                operation_parts[0], operation_parts[1]
            ),
        };
        let test = lines[3]
            .replace("  Test: divisible by ", "")
            .parse::<u64>()
            .expect("test should parse as u64");
        let if_true = lines[4]
            .replace("    If true: throw to monkey ", "")
            .parse::<usize>()
            .expect("true target monkey id should parse to usize");
        let if_false = lines[5]
            .replace("    If false: throw to monkey ", "")
            .parse::<usize>()
            .expect("false target monkey id should parse to usize");

        Monkey {
            held,
            operation,
            test: test,
            if_true,
            if_false,
        }
    }
}

fn run_monkey_keep_away_simulation(starting_monkeys: &Vec<Monkey>) -> Vec<usize> {
    let mut n_examined_items: Vec<usize> = vec![0; starting_monkeys.len()];
    let mut monkeys = starting_monkeys.clone();

    for _ in 0..20 {
        for i in 0..monkeys.len() {
            n_examined_items[i] += monkeys[i].held.len();
            while let Some(worry_level) = monkeys[i].held.pop_front() {
                let new_worry_level = match monkeys[i].operation {
                    Op::Add(addend) => worry_level + addend,
                    Op::Mul(multiplicand) => worry_level * multiplicand,
                    Op::Square => worry_level * worry_level,
                } / 3;

                let target = if new_worry_level % monkeys[i].test == 0 {
                    monkeys[i].if_true
                } else {
                    monkeys[i].if_false
                };
                monkeys[target].held.push_back(new_worry_level);
            }
        }
    }

    n_examined_items
}

fn run_monkey_keep_away_simulation_without_relief_factor(
    starting_monkeys: &Vec<Monkey>,
) -> Vec<usize> {
    let mut n_examined_items: Vec<usize> = vec![0; starting_monkeys.len()];
    let mut monkeys = starting_monkeys.clone();
    let mut relief_modulus: u64 = 1;
    for test_val in  monkeys.iter().map(|m| m.test) {
        relief_modulus *= test_val;
    }

    for _ in 0..10000 {
        for i in 0..monkeys.len() {
            n_examined_items[i] += monkeys[i].held.len();
            while let Some(worry_level) = monkeys[i].held.pop_front() {
                let new_worry_level = match monkeys[i].operation {
                    Op::Add(addend) => worry_level + addend,
                    Op::Mul(multiplicand) => worry_level * multiplicand,
                    Op::Square => worry_level * worry_level,
                } % relief_modulus;

                let target = if new_worry_level % monkeys[i].test == 0 {
                    monkeys[i].if_true
                } else {
                    monkeys[i].if_false
                };
                monkeys[target].held.push_back(new_worry_level);
            }
        }
    }

    n_examined_items
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
            Monkey 0:
              Starting items: 79, 98
              Operation: new = old * 19
              Test: divisible by 23
                If true: throw to monkey 2
                If false: throw to monkey 3

            Monkey 1:
              Starting items: 54, 65, 75, 74
              Operation: new = old + 6
              Test: divisible by 19
                If true: throw to monkey 2
                If false: throw to monkey 0

            Monkey 2:
              Starting items: 79, 60, 97
              Operation: new = old * old
              Test: divisible by 13
                If true: throw to monkey 1
                If false: throw to monkey 3

            Monkey 3:
              Starting items: 74
              Operation: new = old + 3
              Test: divisible by 17
                If true: throw to monkey 0
                If false: throw to monkey 1
        "});
        let mut solver = Day11::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("10605", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Monkey 0:
              Starting items: 79, 98
              Operation: new = old * 19
              Test: divisible by 23
                If true: throw to monkey 2
                If false: throw to monkey 3

            Monkey 1:
              Starting items: 54, 65, 75, 74
              Operation: new = old + 6
              Test: divisible by 19
                If true: throw to monkey 2
                If false: throw to monkey 0

            Monkey 2:
              Starting items: 79, 60, 97
              Operation: new = old * old
              Test: divisible by 13
                If true: throw to monkey 1
                If false: throw to monkey 3

            Monkey 3:
              Starting items: 74
              Operation: new = old + 3
              Test: divisible by 17
                If true: throw to monkey 0
                If false: throw to monkey 1
        "});
        let mut solver = Day11::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("2713310158", result);

        Ok(())
    }
}
