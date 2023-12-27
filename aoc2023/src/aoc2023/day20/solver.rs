use std::collections::{HashMap, HashSet, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day20 {
    adjacency_lists: HashMap<String, Vec<String>>,
    modules: Vec<Module>,
    module_layout: HashMap<String, usize>,
}

impl Day20 {
    pub fn new() -> Self {
        Day20 {
            adjacency_lists: HashMap::new(),
            module_layout: HashMap::<String, usize>::new(),
            modules: vec![],
        }
    }
}

impl Solver for Day20 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let block = input.trim();

        self.module_layout = map_module_layout(block);

        self.modules = block
            .lines()
            .map(Module::try_from_line)
            .collect::<Result<Vec<_>, String>>()?;

        self.adjacency_lists = HashMap::<String, Vec<String>>::new();
        for module in self.modules.iter() {
            let core = match module {
                Module::Broadcast(x) => &x.core,
                Module::Conjunction(x) => &x.core,
                Module::FlipFlop(x) => &x.core,
            };

            for downstream in core.downstreams.iter() {
                self.adjacency_lists
                    .entry(core.label.to_string())
                    .or_insert(Vec::<String>::new())
                    .push(downstream.to_string());
            }
        }

        for module in self.modules.iter_mut() {
            match module {
                Module::Conjunction(inner) => {
                    inner.clear_memory(&self.adjacency_lists, &self.module_layout);
                }
                _ => continue,
            }
        }

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut module_harness =
            ModuleHarness::new(self.modules.clone(), self.module_layout.clone());

        let mut low_pulses_generated = 0_usize;
        let mut high_pulses_generated = 0_usize;

        for _ in 0..WARMUP_ITERATIONS {
            let (lows, highs) = module_harness.run_cycle();

            low_pulses_generated += lows;
            high_pulses_generated += highs;
        }

        Ok((low_pulses_generated * high_pulses_generated).to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let module_predecessors = reverse_adjacency_graph(&self.adjacency_lists);

        let mut module_harness =
            ModuleHarness::new(self.modules.clone(), self.module_layout.clone());

        let predecessors = match module_predecessors.get(RX_MODULE_LABEL) {
            None => Err(format!("RX module had no inputs!")),
            Some(rx_predecessors) => {
                let &rx_predecessor = rx_predecessors.iter().collect::<Vec<_>>().first().unwrap();

                match module_predecessors.get(rx_predecessor) {
                    None => Err(format!("{} had no inputs?", rx_predecessor)),
                    Some(preds) => Ok(preds.iter().map(|s| s.to_string()).collect::<Vec<_>>()),
                }
            }
        }?;

        let time_to_first_high_pulse_sent =
            module_harness.cycle_to_find_time_to_first_high_sent(&predecessors);

        let projected_time_to_rx_activation = predecessors
            .iter()
            .map(
                |predecessor| match time_to_first_high_pulse_sent.get(predecessor) {
                    None => Err(format!("{} had no time to first low?", predecessor)),
                    Some(&time) => Ok(time),
                },
            )
            .collect::<Result<Vec<usize>, String>>()?
            .iter()
            .fold(1_usize, |acc, &x| least_common_multiple(acc, x));

        Ok(projected_time_to_rx_activation.to_string())
    }
}

fn map_module_layout(block: &str) -> HashMap<String, usize> {
    block
        .lines()
        .enumerate()
        .map(|(i, line)| {
            let label_part = line.split(" -> ").collect::<Vec<_>>()[0];
            let intermediate_label = label_part.replace(FLIP_FLOP_LABEL_PREFIX, "");
            let label = intermediate_label.replace(CONJUNCTION_LABEL_PREFIX, "");
            (label, i)
        })
        .collect::<HashMap<String, usize>>()
}

fn reverse_adjacency_graph(
    adjacency_graph: &HashMap<String, Vec<String>>,
) -> HashMap<String, HashSet<String>> {
    let mut reversed_graph = HashMap::<String, HashSet<String>>::new();

    for (src, dsts) in adjacency_graph.iter() {
        for dst in dsts.iter() {
            reversed_graph
                .entry(dst.to_string())
                .or_insert(HashSet::<String>::new())
                .insert(src.to_string());
        }
    }

    reversed_graph
}

struct ModuleHarness {
    modules: Vec<Module>,
    module_locations: HashMap<String, usize>,
    rx_module_activated: bool,
}

impl ModuleHarness {
    fn new(modules: Vec<Module>, module_locations: HashMap<String, usize>) -> Self {
        Self {
            modules,
            module_locations,
            rx_module_activated: false,
        }
    }

    fn run_cycle(&mut self) -> (usize, usize) {
        let mut pulse_queue = VecDeque::<Pulse>::new();
        pulse_queue.push_back(Pulse {
            signal: LOW_SIGNAL,
            src: String::from(BUTTON_LABEL),
            dst: String::from(BROADCASTER_LABEL),
        });

        let mut low_pulses_generated = 0_usize;
        let mut high_pulses_generated = 0_usize;

        while let Some(pulse) = pulse_queue.pop_front() {
            if pulse.signal == LOW_SIGNAL {
                low_pulses_generated += 1;
            } else {
                high_pulses_generated += 1;
            }

            if pulse.dst == RX_MODULE_LABEL {
                if pulse.signal == LOW_SIGNAL {
                    self.rx_module_activated = true;
                }
                continue;
            }

            let target_index = self.module_locations.get(&pulse.dst);
            if target_index.is_none() {
                continue;
            }

            let target = &mut self.modules[*target_index.unwrap()];

            let next_pulses = match target {
                Module::Broadcast(internal) => {
                    internal.handle_pulse(&self.module_locations, &pulse)
                }
                Module::Conjunction(internal) => {
                    internal.handle_pulse(&self.module_locations, &pulse)
                }
                Module::FlipFlop(internal) => internal.handle_pulse(&self.module_locations, &pulse),
            };

            for next_pulse in next_pulses.iter() {
                pulse_queue.push_back(next_pulse.clone());
            }
        }

        (low_pulses_generated, high_pulses_generated)
    }

    fn cycle_to_find_time_to_first_high_sent(
        &mut self,
        key_modules: &Vec<String>,
    ) -> HashMap<String, usize> {
        let mut time_to_first_high_sent = HashMap::<String, usize>::new();

        let mut t = 0_usize;
        loop {
            t += 1;

            let mut pulse_queue = VecDeque::<Pulse>::new();
            pulse_queue.push_back(Pulse {
                signal: LOW_SIGNAL,
                src: String::from(BUTTON_LABEL),
                dst: String::from(BROADCASTER_LABEL),
            });

            while let Some(pulse) = pulse_queue.pop_front() {
                if pulse.dst == RX_MODULE_LABEL {
                    if pulse.signal == HIGH_SIGNAL
                        && !time_to_first_high_sent.contains_key(&pulse.src)
                    {
                        time_to_first_high_sent.insert(pulse.src.to_string(), t);
                    }
                    continue;
                }

                let target_index = self.module_locations.get(&pulse.dst);
                if target_index.is_none() {
                    if pulse.signal == HIGH_SIGNAL
                        && !time_to_first_high_sent.contains_key(&pulse.src)
                    {
                        time_to_first_high_sent.insert(pulse.src.to_string(), t);
                    }

                    continue;
                }

                let target = &mut self.modules[*target_index.unwrap()];

                let next_pulses = match target {
                    Module::Broadcast(internal) => {
                        internal.handle_pulse(&self.module_locations, &pulse)
                    }
                    Module::Conjunction(internal) => {
                        internal.handle_pulse(&self.module_locations, &pulse)
                    }
                    Module::FlipFlop(internal) => {
                        internal.handle_pulse(&self.module_locations, &pulse)
                    }
                };

                let mut high_pulse_sent = false;
                for next_pulse in next_pulses.iter() {
                    pulse_queue.push_back(next_pulse.clone());

                    if pulse.signal == HIGH_SIGNAL {
                        high_pulse_sent = true;
                    }
                }

                if high_pulse_sent && !time_to_first_high_sent.contains_key(&pulse.src) {
                    time_to_first_high_sent.insert(pulse.src.to_string(), t);
                }
            }

            if key_modules
                .iter()
                .all(|module| time_to_first_high_sent.contains_key(module))
            {
                break;
            }
        }

        time_to_first_high_sent
    }
}

const WARMUP_ITERATIONS: usize = 1_000;

type Signal = usize;

const LOW_SIGNAL: Signal = 0;
const HIGH_SIGNAL: Signal = 1;

const BUTTON_LABEL: &str = "__button__";
const BROADCASTER_LABEL: &str = "broadcaster";
const FLIP_FLOP_LABEL_PREFIX: &str = "%";
const CONJUNCTION_LABEL_PREFIX: &str = "&";
const RX_MODULE_LABEL: &str = "rx";

#[derive(Clone)]
struct Pulse {
    signal: Signal,
    src: String,
    dst: String,
}

#[derive(Clone)]
enum Module {
    FlipFlop(FlipFlopInternal),
    Conjunction(ConjunctionInternal),
    Broadcast(BroadcastInternal),
}

impl Module {
    fn try_from_line(line: &str) -> Result<Self, String> {
        let parts = line.split(" -> ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("'{}' did not have a label and downstreams", line));
        }

        let downstream_labels = parts[1].split(", ").map(String::from).collect::<Vec<_>>();
        if downstream_labels.len() < 1 {
            return Err(format!(
                "{} did not contain downstream module labels",
                parts[1]
            ));
        }

        let module = if parts[0] == BROADCASTER_LABEL {
            let core = ModuleCore {
                label: String::from(BROADCASTER_LABEL),
                downstreams: downstream_labels.clone(),
            };
            Self::Broadcast(BroadcastInternal::new(core))
        } else if parts[0].starts_with(FLIP_FLOP_LABEL_PREFIX) {
            let label = parts[0].replace(FLIP_FLOP_LABEL_PREFIX, "");
            let core = ModuleCore {
                label,
                downstreams: downstream_labels.clone(),
            };
            Self::FlipFlop(FlipFlopInternal::new(core))
        } else if parts[0].starts_with(CONJUNCTION_LABEL_PREFIX) {
            let label = parts[0].replace(CONJUNCTION_LABEL_PREFIX, "");
            let core = ModuleCore {
                label,
                downstreams: downstream_labels.clone(),
            };
            Self::Conjunction(ConjunctionInternal::new(core))
        } else {
            return Err(format!("{} is an invalid label", parts[0]));
        };

        Ok(module)
    }
}

#[derive(Clone)]
struct ModuleCore {
    label: String,
    downstreams: Vec<String>,
}

trait ModuleInternal {
    fn handle_pulse(
        &mut self,
        module_locations: &HashMap<String, usize>,
        pulse: &Pulse,
    ) -> Vec<Pulse>;

    fn generate_downstream_pulses(
        &self,
        signal: Signal,
        src: &String,
        downstreams: &Vec<String>,
    ) -> Vec<Pulse> {
        let mut downstream_pulses = Vec::<Pulse>::with_capacity(downstreams.len());
        for downstream in downstreams.iter() {
            downstream_pulses.push(Pulse {
                signal,
                src: src.to_string(),
                dst: downstream.to_string(),
            });
        }
        downstream_pulses
    }
}

#[derive(Clone)]
struct FlipFlopInternal {
    core: ModuleCore,
    state: Signal,
}

impl FlipFlopInternal {
    fn new(core: ModuleCore) -> Self {
        Self {
            core,
            state: LOW_SIGNAL,
        }
    }
}

impl ModuleInternal for FlipFlopInternal {
    fn handle_pulse(&mut self, _: &HashMap<String, usize>, pulse: &Pulse) -> Vec<Pulse> {
        match pulse.signal {
            HIGH_SIGNAL => vec![],
            LOW_SIGNAL => {
                self.state = (self.state + 1) % 2;

                self.generate_downstream_pulses(
                    self.state,
                    &self.core.label,
                    &self.core.downstreams,
                )
            }
            _ => panic!("signal was not a valid value"),
        }
    }
}

#[derive(Clone)]
struct ConjunctionInternal {
    core: ModuleCore,
    state: u128,
}

impl ConjunctionInternal {
    fn new(core: ModuleCore) -> Self {
        Self {
            core,
            state: u128::MAX,
        }
    }

    fn clear_memory(
        &mut self,
        adjacency_lists: &HashMap<String, Vec<String>>,
        module_locations: &HashMap<String, usize>,
    ) {
        let mut upstreams = HashSet::<String>::new();
        for (upstream, downstreams) in adjacency_lists.iter() {
            for downstream in downstreams.iter() {
                if downstream == &self.core.label {
                    upstreams.insert(upstream.to_string());
                }
            }
        }

        for upstream in upstreams.iter() {
            let index = module_locations.get(upstream).unwrap();
            self.state &= !(1 << index);
        }
    }
}

impl ModuleInternal for ConjunctionInternal {
    fn handle_pulse(
        &mut self,
        module_locations: &HashMap<String, usize>,
        pulse: &Pulse,
    ) -> Vec<Pulse> {
        let module_index = module_locations.get(&pulse.src).unwrap();

        self.state &= !(1 << module_index);
        self.state |= (pulse.signal as u128) << module_index;

        let signal = if self.state & u128::MAX == u128::MAX {
            LOW_SIGNAL
        } else {
            HIGH_SIGNAL
        };

        self.generate_downstream_pulses(signal, &self.core.label, &self.core.downstreams)
    }
}

#[derive(Clone)]
struct BroadcastInternal {
    core: ModuleCore,
}

impl BroadcastInternal {
    fn new(core: ModuleCore) -> Self {
        Self { core }
    }
}

impl ModuleInternal for BroadcastInternal {
    fn handle_pulse(&mut self, _: &HashMap<String, usize>, pulse: &Pulse) -> Vec<Pulse> {
        self.generate_downstream_pulses(pulse.signal, &self.core.label, &self.core.downstreams)
    }
}

fn least_common_multiple(a: usize, b: usize) -> usize {
    (a * b) / greatest_common_denominator(a, b)
}

fn greatest_common_denominator(a: usize, b: usize) -> usize {
    let (mut a_1, mut b_1) = (a, b);
    while b_1 > 0 {
        (a_1, b_1) = (b_1, a_1 % b_1);
    }
    a_1
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
            broadcaster -> a, b, c
            %a -> b
            %b -> c
            %c -> inv
            &inv -> a
        "});
        let mut solver = Day20::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("32000000", result);

        Ok(())
    }
}
