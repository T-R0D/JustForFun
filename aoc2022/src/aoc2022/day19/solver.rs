// Really need to find an optimization for this one...

use std::{
    collections::{HashSet},
    iter::zip,
};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day19 {
    blueprints: Vec<Blueprint>,
}

impl Day19 {
    pub fn new() -> Self {
        Self {
            blueprints: Vec::new(),
        }
    }
}

impl Solver for Day19 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.blueprints = input
            .trim()
            .lines()
            .map(Blueprint::from_line)
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let total_quality_level: u32 = self
            .blueprints
            .iter()
            .map(|blueprint| {
                blueprint.id * evaluate_geode_production(blueprint, SIMULATION_TIME_MINUTES)
            })
            .sum();

        Ok(total_quality_level.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let geodes_producted_product: u32 = self
            .blueprints
            .iter()
            .take(UNEATEN_BLUEPRINT_COUNT)
            .map(|blueprint| evaluate_geode_production(blueprint, LONGER_SIMULATION_TIME_MINUTES))
            .fold(1, |product, geodes_produced| product * geodes_produced);

        Ok(geodes_producted_product.to_string())
    }
}

struct Blueprint {
    id: u32,
    bot_blueprints: [[u32; Resource::NMembers as usize]; Resource::NMembers as usize],
    cost_maxes: [u32; Resource::NMembers as usize],
}

impl Blueprint {
    fn from_line(line: &str) -> Self {
        let reduced_line = line
            .replace("Blueprint ", "")
            .replace(" Each ", "")
            .replace(" robot ", "");
        let parts = reduced_line.split(":").collect::<Vec<_>>();
        let id = parts[0].parse::<u32>().unwrap();
        let bot_blueprints = parts[1]
            .split(".")
            .filter(|&s| s != "")
            .map(|line| {
                let parts = line.split("costs ").collect::<Vec<_>>();
                let robot_kind = Resource::from_str(parts[0]);
                let resource_costs = parts[1]
                    .split(" and ")
                    .map(ResourceCost::from_line)
                    .collect::<Vec<_>>();
                (robot_kind, resource_costs)
            })
            .fold(
                [[0; Resource::NMembers as usize]; Resource::NMembers as usize],
                |mut robot_blueprints, (robot_kind, resource_costs)| {
                    for cost in resource_costs.iter() {
                        robot_blueprints[robot_kind as usize][cost.kind as usize] = cost.quantity;
                    }
                    robot_blueprints
                },
            );
        let mut cost_maxes = [0; Resource::NMembers as usize];
        for blueprint in bot_blueprints.iter() {
            for (i, &cost) in blueprint.iter().enumerate() {
                if cost_maxes[i] < cost {
                    cost_maxes[i] = cost;
                }
            }
        }

        Blueprint {
            id,
            bot_blueprints,
            cost_maxes,
        }
    }
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
enum Resource {
    Geode = 0,
    Obsidian,
    Clay,
    Ore,
    NMembers,
}

impl Resource {
    fn from_str(text: &str) -> Self {
        match text {
            "geode" => Resource::Geode,
            "obsidian" => Resource::Obsidian,
            "clay" => Resource::Clay,
            "ore" => Resource::Ore,
            _ => panic!("'{text}' is not a recognizable resource"),
        }
    }
}

#[derive(Clone, Copy)]
struct ResourceCost {
    kind: Resource,
    quantity: u32,
}

impl ResourceCost {
    fn from_line(line: &str) -> Self {
        let parts = line.split(" ").collect::<Vec<_>>();
        let quantity = parts[0].parse::<u32>().unwrap();
        let kind = Resource::from_str(parts[1]);
        Self { kind, quantity }
    }
}

const SIMULATION_TIME_MINUTES: u8 = 24;
const LONGER_SIMULATION_TIME_MINUTES: u8 = 32;
const UNEATEN_BLUEPRINT_COUNT: usize = 3;

const BOT_KIND_PRIORITY: [Resource; Resource::NMembers as usize] = [
    Resource::Geode,
    Resource::Obsidian,
    Resource::Clay,
    Resource::Ore,
];

#[derive(Clone)]
struct SimState {
    factory_status: Option<Resource>,
    resource_counts: [u32; Resource::NMembers as usize],
    bot_counts: [u32; Resource::NMembers as usize],
}

impl SimState {
    fn generate_bot_build_states(&self, blueprint: &Blueprint) -> Vec<SimState> {
        let mut new_states = Vec::from([self.clone()]);
        if self.factory_status.is_some() {
            return new_states;
        }

        'build_bot_attempt: for &bot_kind in BOT_KIND_PRIORITY.iter() {
            match bot_kind {
                Resource::Geode => (),
                Resource::Obsidian => {
                    if self.bot_counts[Resource::Obsidian as usize]
                        >= blueprint.cost_maxes[Resource::Obsidian as usize]
                    {
                        continue 'build_bot_attempt;
                    }
                }
                Resource::Clay => {
                    if self.bot_counts[Resource::Clay as usize]
                        >= blueprint.cost_maxes[Resource::Clay as usize]
                    {
                        continue 'build_bot_attempt;
                    }
                }
                Resource::Ore => {
                    if self.bot_counts[Resource::Ore as usize]
                        >= blueprint.cost_maxes[Resource::Ore as usize]
                    {
                        continue 'build_bot_attempt;
                    }
                }
                Resource::NMembers => unreachable!(),
            }

            let bot_bill_of_materials = blueprint.bot_blueprints[bot_kind as usize];

            for (&resource_kind, &cost) in
                zip(BOT_KIND_PRIORITY.iter(), bot_bill_of_materials.iter())
            {
                if cost > 0 && self.resource_counts[resource_kind as usize] < cost {
                    continue 'build_bot_attempt;
                }
            }

            let mut new_state = self.clone();
            new_state.factory_status = Some(bot_kind);
            for (&resource_kind, &cost) in
                zip(BOT_KIND_PRIORITY.iter(), bot_bill_of_materials.iter())
            {
                new_state.resource_counts[resource_kind as usize] -= cost;
            }

            new_states.push(new_state);
        }

        new_states
    }

    fn collect_resources(&mut self) {
        for (i, amount) in self.bot_counts.iter().enumerate() {
            self.resource_counts[i] += amount;
        }
    }

    fn to_string(&self) -> String {
        format!(
            "r:[{}] b:[{}]",
            self.resource_counts.map(|x| x.to_string()).join(","),
            self.bot_counts.map(|x| x.to_string()).join(",")
        )
        .to_string()
    }
}

fn evaluate_geode_production(blueprint: &Blueprint, time_minutes: u8) -> u32 {
    let mut initial_state = SimState {
        factory_status: None,
        resource_counts: [0; Resource::NMembers as usize],
        bot_counts: [0; Resource::NMembers as usize],
    };
    initial_state.bot_counts[Resource::Ore as usize] = 1;
    let mut states_to_consider: Vec<SimState> = Vec::new();
    states_to_consider.push(initial_state);
    let mut held_states: HashSet<String> = HashSet::new();

    for _ in 0..time_minutes {
        let mut next_states: Vec<SimState> = Vec::with_capacity(states_to_consider.len());

        while let Some(previous_state) = states_to_consider.pop() {
            // Start build phase.
            let mut new_states = previous_state.generate_bot_build_states(blueprint);

            // Collection phase.
            for state in new_states.iter_mut() {
                state.collect_resources()
            }

            // Complete build phase.
            for state in new_states.iter_mut() {
                if let Some(bot_kind) = state.factory_status {
                    state.factory_status = None;
                    state.bot_counts[bot_kind as usize] += 1;
                }
            }

            'uniqueness_check: for state in new_states.iter() {
                let state_str = state.to_string();
                if held_states.contains(&state_str) {
                    continue 'uniqueness_check;
                }

                held_states.insert(state_str);
                next_states.push(state.clone());
            }
        }

        states_to_consider = next_states;
    }

    let mut max_geodes_found = 0;
    for state in states_to_consider.iter() {
        if state.resource_counts[Resource::Geode as usize] > max_geodes_found {
            max_geodes_found = state.resource_counts[Resource::Geode as usize];
        }
    }

    max_geodes_found
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
            Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
            Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
        "});
        let mut solver = Day19::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("33", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
            Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
        "});
        let mut solver = Day19::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("3472", result);

        Ok(())
    }
}
