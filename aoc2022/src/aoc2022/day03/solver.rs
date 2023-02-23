use std::collections::HashSet;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day03 {
    rucksacks: Vec<Rucksack>,
}

impl Day03 {
    pub fn new() -> Self {
        Day03 { rucksacks: vec![] }
    }
}

impl Solver for Day03 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let specs = input
            .trim()
            .split("\n")
            .map(|s| s.as_bytes().to_vec())
            .collect::<Vec<_>>();

        self.rucksacks = specs
            .iter()
            .map(|spec| Rucksack::from(spec))
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut priority_score_of_duplicated_items = 0;
        for rucksack in self.rucksacks.iter() {
            let duplicated_items = rucksack.find_duplicated_priority_items();
            priority_score_of_duplicated_items += duplicated_items.iter().sum::<u32>();
        }

        Ok(priority_score_of_duplicated_items.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut badge_priority_sum: u32 = 0;
        for group in self.rucksacks.chunks(3) {
            let unique_item_collections = group
                .iter()
                .map(|rucksack| rucksack.unique_item_priority_scores())
                .collect::<Vec<_>>();

            let intersection = unique_item_collections.iter().fold(unique_item_collections[0].clone(), |acc, s2| &acc & s2);
            for shared_item_priority_score in intersection.iter() {
                badge_priority_sum += shared_item_priority_score;
            }
        }

        Ok(badge_priority_sum.to_string())
    }
}

enum Compartment {
    One,
    Two,
}

struct Rucksack {
    raw_contents: Vec<u8>,
}

impl Rucksack {
    fn from(spec: &Vec<u8>) -> Self {
        Rucksack {
            raw_contents: spec.to_vec(),
        }
    }

    fn item_priority(item: &u8) -> usize {
        let it = *item;
        return if b'a' <= it && it <= b'z' {
            (item - b'a') as usize
        } else {
            (item - b'A' + 26) as usize
        };
    }

    fn priority_score(priority: &usize) -> u32 {
        (priority + 1) as u32
    }

    fn compartment_contents(&self, compartment: Compartment) -> Vec<u8> {
        let i = match compartment {
            Compartment::One => 0,
            Compartment::Two => 1,
        };
        let contents = self
            .raw_contents
            .chunks(self.raw_contents.len() / 2)
            .nth(i)
            .expect("There should be an available chunk");
        contents.to_vec()
    }

    fn find_duplicated_priority_items(&self) -> Vec<u32> {
        let one_contents = self.compartment_contents(Compartment::One);
        let two_contents = self.compartment_contents(Compartment::Two);

        let one_uniques = one_contents.iter().collect::<HashSet<_>>();
        let two_uniques = two_contents.iter().collect::<HashSet<_>>();

        one_uniques
            .intersection(&two_uniques)
            .map(|item| Rucksack::priority_score(&Rucksack::item_priority(item)))
            .collect::<Vec<_>>()
    }

    fn unique_item_priority_scores(&self) -> HashSet<u32> {
        self.raw_contents
            .iter()
            .map(|item| Rucksack::priority_score(&Rucksack::item_priority(item)))
            .collect::<HashSet<_>>()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            vJrwpWtwJgWrhcsFMMfFFhFp
            jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
            PmmdzqPrVvPwwTWBwg
            wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
            ttgJtRGJQctTZtZT
            CrZsJsPPZsGzwwsLwLmpwMDw
        "});
        let mut solver = Day03::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("157", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            vJrwpWtwJgWrhcsFMMfFFhFp
            jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
            PmmdzqPrVvPwwTWBwg
            wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
            ttgJtRGJQctTZtZT
            CrZsJsPPZsGzwwsLwLmpwMDw
        "});
        let mut solver = Day03::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("70", result);

        Ok(())
    }
}
