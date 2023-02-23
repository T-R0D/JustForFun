use std::collections::HashMap;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day16 {
    map: HashMap<String, ValveRoom>,
}

impl Day16 {
    pub fn new() -> Self {
        Day16 {
            map: HashMap::new(),
        }
    }
}

const START_LABEL: &str = "AA";
const OPEN_TIME: u32 = 1;
const TIME_LIMIT: u32 = 30;
const TIME_TO_TEACH_ELEPHANT: u32 = 4;

impl Solver for Day16 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        for (name, valve_room) in input.trim().lines().map(parse_line) {
            self.map.insert(name, valve_room);
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let (label_to_index, travel_graph) = compute_all_pairs_shortest_paths(&self.map);
        let rooms_with_working_valves = compute_rooms_with_viable_valves(&self.map);
        let indices_with_working_valves = rooms_with_working_valves
            .iter()
            .map(|label| *label_to_index.get(label).unwrap())
            .collect::<Vec<_>>();
        let mut index_to_flow = HashMap::new();
        for label in rooms_with_working_valves.iter() {
            let room = self.map.get(label).unwrap();
            let index = label_to_index.get(label).unwrap();
            index_to_flow.insert(*index, room.flow_per_minute);
        }
        let start_index = *label_to_index.get(&String::from(START_LABEL)).unwrap();

        let max_pressure_released = find_best_pressure_release(
            start_index,
            &travel_graph,
            &indices_with_working_valves,
            &index_to_flow,
        );

        Ok(max_pressure_released.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let (label_to_index, travel_graph) = compute_all_pairs_shortest_paths(&self.map);
        let rooms_with_working_valves = compute_rooms_with_viable_valves(&self.map);
        let mut indices_with_working_valves = rooms_with_working_valves
            .iter()
            .map(|label| *label_to_index.get(label).unwrap())
            .collect::<Vec<_>>();
        indices_with_working_valves.sort();
        let mut index_to_flow = HashMap::new();
        for label in rooms_with_working_valves.iter() {
            let room = self.map.get(label).unwrap();
            let index = label_to_index.get(label).unwrap();
            index_to_flow.insert(*index, room.flow_per_minute);
        }
        let start_index = *label_to_index.get(&String::from(START_LABEL)).unwrap();

        let max_pressure_released = find_best_pressure_release_with_partner(
            start_index,
            &travel_graph,
            &indices_with_working_valves,
            &index_to_flow,
        );

        Ok(max_pressure_released.to_string())
    }
}

fn parse_line(line: &str) -> (String, ValveRoom) {
    let first_reduction = line.replace("Valve ", "");
    let second_reduction = first_reduction
        .replace("valves", "valve")
        .replace("leads", "lead")
        .replace("tunnels", "tunnel");
    let parts = second_reduction
        .split("; tunnel lead to valve ")
        .collect::<Vec<_>>();

    let name_and_flow_rate = parts[0].split(" has flow rate=").collect::<Vec<_>>();
    let name = String::from(name_and_flow_rate[0]);
    let flow_rate = name_and_flow_rate[1].parse::<u32>().unwrap();

    let neighbors = parts[1].split(", ").map(String::from).collect::<Vec<_>>();

    (
        name,
        ValveRoom {
            flow_per_minute: flow_rate,
            neighbors,
        },
    )
}

struct ValveRoom {
    flow_per_minute: u32,
    neighbors: Vec<String>,
}

fn compute_all_pairs_shortest_paths(
    map: &HashMap<String, ValveRoom>,
) -> (HashMap<String, usize>, Vec<Vec<u32>>) {
    let mut label_to_index: HashMap<String, usize> = HashMap::with_capacity(map.len());
    let mut graph: Vec<Vec<u32>> = vec![vec![u32::MAX; map.len()]; map.len()];

    for (i, label) in map.keys().enumerate() {
        label_to_index.insert(label.clone(), i);
        graph[i][i] = 0;
    }

    for (label, valve_room) in map.iter() {
        let src = label_to_index.get(label).unwrap();
        for neighbor in valve_room.neighbors.iter() {
            let dst = label_to_index.get(neighbor).unwrap();
            graph[*src][*dst] = 1;
        }
    }

    for k in 0..map.len() {
        for i in 0..map.len() {
            'dest_candidate: for j in 0..map.len() {
                if graph[i][k] == u32::MAX || graph[k][j] == u32::MAX {
                    continue 'dest_candidate;
                }

                let candidate = graph[i][k] + graph[k][j];
                if graph[i][j] > candidate {
                    graph[i][j] = candidate;
                }
            }
        }
    }

    (label_to_index, graph)
}

fn compute_rooms_with_viable_valves(map: &HashMap<String, ValveRoom>) -> Vec<String> {
    let mut rooms_with_working_valves: Vec<String> = Vec::new();
    for (label, room) in map.iter() {
        if room.flow_per_minute > 0 {
            rooms_with_working_valves.push(label.clone());
        }
    }
    rooms_with_working_valves
}

fn find_best_pressure_release(
    start_index: usize,
    travel_graph: &Vec<Vec<u32>>,
    indices_of_working_valves: &Vec<usize>,
    index_to_flow: &HashMap<usize, u32>,
) -> u32 {
    best_pressure_release_helper(
        start_index,
        TIME_LIMIT,
        indices_of_working_valves,
        index_to_flow,
        travel_graph,
    )
}

fn best_pressure_release_helper(
    current_index: usize,
    remaining_time: u32,
    remaining_valves: &Vec<usize>,
    index_to_flow: &HashMap<usize, u32>,
    travel_graph: &Vec<Vec<u32>>,
) -> u32 {
    let mut best_release = 0;

    for next_index in remaining_valves.iter() {
        let time_to_open = travel_graph[current_index][*next_index] + OPEN_TIME;
        if time_to_open >= remaining_time {
            continue;
        }

        let hypothetical_remaining_time = remaining_time - time_to_open;

        let pressure_released =
            hypothetical_remaining_time * (*index_to_flow.get(next_index).unwrap());

        let valves_left = remaining_valves
            .clone()
            .iter()
            .filter_map(|&index| {
                if index != *next_index {
                    Some(index)
                } else {
                    None
                }
            })
            .collect::<Vec<_>>();

        let sub_solution = best_pressure_release_helper(
            *next_index,
            hypothetical_remaining_time,
            &valves_left,
            index_to_flow,
            travel_graph,
        );

        if best_release < pressure_released + sub_solution {
            best_release = pressure_released + sub_solution;
        }
    }

    best_release
}

fn find_best_pressure_release_with_partner(
    start_index: usize,
    travel_graph: &Vec<Vec<u32>>,
    indices_of_working_valves: &Vec<usize>,
    index_to_flow: &HashMap<usize, u32>,
) -> u32 {
    let mut memo: HashMap<String, u32> = HashMap::new();

    let pressure_released = best_pressure_release_with_partner_helper(
        &mut memo,
        start_index,
        start_index,
        TIME_LIMIT - TIME_TO_TEACH_ELEPHANT,
        TIME_LIMIT - TIME_TO_TEACH_ELEPHANT,
        indices_of_working_valves,
        index_to_flow,
        travel_graph,
    );

    pressure_released
}

fn best_pressure_release_with_partner_helper(
    memo: &mut HashMap<String, u32>,
    current_index_a: usize,
    current_index_b: usize,
    remaining_time_a: u32,
    remaining_time_b: u32,
    remaining_valves: &Vec<usize>,
    index_to_flow: &HashMap<usize, u32>,
    travel_graph: &Vec<Vec<u32>>,
) -> u32 {
    let keys = state_with_partner_to_keys(current_index_a, current_index_b, remaining_time_a, remaining_time_b, remaining_valves);

    for key in keys.iter() {
        if memo.contains_key(key) {
            return *memo.get(key).unwrap();
        }
    }

    let mut best_release = 0;

    for next_index in remaining_valves.iter() {
        let time_to_open_a = travel_graph[current_index_a][*next_index] + OPEN_TIME;
        let time_to_open_b = travel_graph[current_index_b][*next_index] + OPEN_TIME;
        if time_to_open_a >= remaining_time_a && time_to_open_b >= remaining_time_b {
            continue;
        }

        let valves_left = remaining_valves
            .clone()
            .iter()
            .filter_map(|&index| {
                if index != *next_index {
                    Some(index)
                } else {
                    None
                }
            })
            .collect::<Vec<_>>();

        if time_to_open_a < remaining_time_a {
            let hypothetical_remaining_time = remaining_time_a - time_to_open_a;
            let pressure_released =
                hypothetical_remaining_time * (*index_to_flow.get(next_index).unwrap());
            let sub_solution = best_pressure_release_with_partner_helper(
                memo,
                *next_index,
                current_index_b,
                hypothetical_remaining_time,
                remaining_time_b,
                &valves_left,
                index_to_flow,
                travel_graph,
            );

            if best_release < pressure_released + sub_solution {
                best_release = pressure_released + sub_solution;
            }
        }

        if time_to_open_b < remaining_time_b {
            let hypothetical_remaining_time = remaining_time_b - time_to_open_b;
            let pressure_released =
                hypothetical_remaining_time * (*index_to_flow.get(next_index).unwrap());
            let sub_solution = best_pressure_release_with_partner_helper(
                memo,
                current_index_a,
                *next_index,
                remaining_time_a,
                hypothetical_remaining_time,
                &valves_left,
                index_to_flow,
                travel_graph,
            );

            if best_release < pressure_released + sub_solution {
                best_release = pressure_released + sub_solution;
            }
        }
    }

    for key in keys.iter() {
        memo.insert(key.clone(), best_release);
    }

    best_release
}

fn state_with_partner_to_keys(
    index_a: usize,
    index_b: usize,
    remaining_time_a: u32,
    remaining_time_b: u32,
    remaining_valves: &Vec<usize>,
) -> Vec<String> {
    let r = remaining_valves
        .iter()
        .map(usize::to_string)
        .collect::<Vec<_>>()
        .join(",");

        vec![
            String::from(format!(
                "i:{index_a} t:{remaining_time_a} i:{index_b} t:{remaining_time_b} f:{r}"
            )),
            String::from(format!(
                "i:{index_b} t:{remaining_time_b} i:{index_a} t:{remaining_time_a} f:{r}"
            )),
        ]
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
            Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
            Valve BB has flow rate=13; tunnels lead to valves CC, AA
            Valve CC has flow rate=2; tunnels lead to valves DD, BB
            Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
            Valve EE has flow rate=3; tunnels lead to valves FF, DD
            Valve FF has flow rate=0; tunnels lead to valves EE, GG
            Valve GG has flow rate=0; tunnels lead to valves FF, HH
            Valve HH has flow rate=22; tunnel leads to valve GG
            Valve II has flow rate=0; tunnels lead to valves AA, JJ
            Valve JJ has flow rate=21; tunnel leads to valve II
        "});
        let mut solver = Day16::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("1651", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
            Valve BB has flow rate=13; tunnels lead to valves CC, AA
            Valve CC has flow rate=2; tunnels lead to valves DD, BB
            Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
            Valve EE has flow rate=3; tunnels lead to valves FF, DD
            Valve FF has flow rate=0; tunnels lead to valves EE, GG
            Valve GG has flow rate=0; tunnels lead to valves FF, HH
            Valve HH has flow rate=22; tunnel leads to valve GG
            Valve II has flow rate=0; tunnels lead to valves AA, JJ
            Valve JJ has flow rate=21; tunnel leads to valve II
        "});
        let mut solver = Day16::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("1707", result);

        Ok(())
    }
}
