use std::collections::{HashMap, HashSet, VecDeque};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day19 {
    parts: Vec<Part>,
    workflows: HashMap<String, Vec<ProcessSpec>>,
}

impl Day19 {
    pub fn new() -> Self {
        Day19 {
            parts: vec![],
            workflows: HashMap::new(),
        }
    }
}

impl Solver for Day19 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let parts = input.trim().split("\n\n").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(String::from("file did not have 2 sections"));
        }

        self.workflows = try_workflows_from_block(parts[0])?;

        self.parts = parts[1]
            .lines()
            .map(Part::try_from_line)
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let workflow_map = self
            .workflows
            .iter()
            .map(|(label, op_specs)| {
                (
                    label.clone(),
                    op_specs
                        .iter()
                        .map(|spec| spec.to_predicate())
                        .collect::<Vec<_>>(),
                )
            })
            .collect::<HashMap<String, Vec<Box<WorkflowPredicate>>>>();

        let n_accepted_parts = self
            .parts
            .iter()
            .map(|part| match try_process_part(&workflow_map, part) {
                Ok(ProcessResult::Accepted) => part.total_rating(),
                _ => 0,
            })
            .sum::<usize>();

        Ok(n_accepted_parts.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let reversed_process_graph = build_reversed_process_flow_graph(&self.workflows);

        let total_part_rating_comninations =
            count_acceptable_part_rating_combinations(&self.workflows, &reversed_process_graph);

        Ok(total_part_rating_comninations.to_string())
    }
}

const INPUT_WORKFLOW_LABEL: &str = "in";
const ACCEPTED_LABEL: &str = "A";
const REJECTED_LABEL: &str = "R";

const RATING_MIN: usize = 1;
const RATING_MAX: usize = 4000;

fn try_workflows_from_block(block: &str) -> Result<HashMap<String, Vec<ProcessSpec>>, String> {
    let mut workflows = HashMap::<String, Vec<ProcessSpec>>::new();

    for line in block.trim().lines() {
        let trimmed_line = line.replace("}", "");
        let parts = trimmed_line.split("{").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("malformed line: {line}"));
        }

        let label = String::from(parts[0]);

        let workflow = parts[1]
            .split(",")
            .map(ProcessSpec::try_from_str)
            .collect::<Result<Vec<_>, String>>()?;

        workflows.insert(label.clone(), workflow);
    }

    Ok(workflows)
}

fn try_process_part(
    workflows: &HashMap<String, Vec<Box<WorkflowPredicate>>>,
    part: &Part,
) -> Result<ProcessResult, String> {
    let mut next_workflow_label = String::from(INPUT_WORKFLOW_LABEL);
    let iterations = 0_usize;

    loop {
        if iterations > 10_000 {
            return Err(format!("A part processed for too long and broke things"));
        }

        let workflow = match workflows.get(&next_workflow_label) {
            Some(workflow) => workflow,
            _ => return Err(format!("{} not found", next_workflow_label)),
        };

        'workflow_processing: for predicate in workflow.iter() {
            match predicate(part) {
                ProcessResult::PassAlong => continue 'workflow_processing,
                ProcessResult::Next(label) => {
                    next_workflow_label = label;
                    break 'workflow_processing;
                }
                terminal_result @ _ => return Ok(terminal_result),
            }
        }
    }
}

fn build_reversed_process_flow_graph(
    workflows: &HashMap<String, Vec<ProcessSpec>>,
) -> HashMap<String, HashSet<ProcessLocator>> {
    let mut process_adjacencies = HashMap::<String, HashSet<ProcessLocator>>::new();

    for (ref src, ops) in workflows.iter() {
        for (i, op) in ops.iter().enumerate() {
            let neighbor_label = op.affirmative_result.to_label();

            process_adjacencies
                .entry(neighbor_label.to_string())
                .or_insert(HashSet::<ProcessLocator>::new())
                .insert(ProcessLocator {
                    label: src.to_string(),
                    step: i,
                });
        }
    }

    process_adjacencies
}

fn count_acceptable_part_rating_combinations(
    workflows: &HashMap<String, Vec<ProcessSpec>>,
    reversed_workflow_graph: &HashMap<String, HashSet<ProcessLocator>>,
) -> usize {
    let mut frontier = VecDeque::<(
        ProcessLocator,
        [Range; RatingCategory::NRatingCategories.index()],
    )>::new();
    frontier.push_back((
        ProcessLocator {
            label: String::from(ACCEPTED_LABEL),
            step: 0,
        },
        [Range::new(RATING_MIN, RATING_MAX); RatingCategory::NRatingCategories.index()],
    ));

    let mut acceptable_part_ratings =
        HashSet::<[Range; RatingCategory::NRatingCategories.index()]>::new();

    while let Some((next_process, possible_part_ratings)) = frontier.pop_front() {
        if next_process.label == REJECTED_LABEL {
            continue;
        }

        let mut new_ratings = possible_part_ratings.clone();
        if next_process.label != ACCEPTED_LABEL {
            let workflow = workflows.get(&next_process.label).unwrap();

            let process = &workflow[next_process.step];
            match process.op {
                Op::GreaterThan(category, threshold) => {
                    let rating_range = new_ratings[category.index()];
                    new_ratings[category.index()] = rating_range.raise(threshold + 1);
                }
                Op::LessThan(category, threshold) => {
                    let rating_range = new_ratings[category.index()];
                    new_ratings[category.index()] = rating_range.cap(threshold - 1);
                }
                Op::Any => (),
            }

            for i in (0..next_process.step).rev() {
                let skipped_process = &workflow[i];
                match skipped_process.op {
                    Op::GreaterThan(category, threshold) => {
                        let rating_range = new_ratings[category.index()];
                        new_ratings[category.index()] = rating_range.cap(threshold);
                    }
                    Op::LessThan(category, threshold) => {
                        let rating_range = new_ratings[category.index()];
                        new_ratings[category.index()] = rating_range.raise(threshold);
                    }
                    Op::Any => (),
                }
            }
        }

        if next_process.label == INPUT_WORKFLOW_LABEL {
            acceptable_part_ratings.insert(new_ratings);
            continue;
        }

        for previous_process in reversed_workflow_graph
            .get(&next_process.label)
            .unwrap()
            .iter()
        {
            frontier.push_back((previous_process.clone(), new_ratings.clone()))
        }
    }

    acceptable_part_ratings
        .iter()
        .map(|accepatable_ratings| {
            accepatable_ratings
                .iter()
                .map(|range| range.len())
                .product::<usize>()
        })
        .sum::<usize>()
}

type WorkflowPredicate = dyn Fn(&Part) -> ProcessResult;

struct ProcessSpec {
    op: Op,
    affirmative_result: ProcessResult,
}

impl ProcessSpec {
    fn try_from_str(s: &str) -> Result<Self, String> {
        let parts = s.split(":").collect::<Vec<_>>();
        if parts.len() > 2 {
            return Err(format!("{s} is poorly formatted"));
        }

        let affirmative_result = match parts.last() {
            Some(&ACCEPTED_LABEL) => Ok(ProcessResult::Accepted),
            Some(&REJECTED_LABEL) => Ok(ProcessResult::Rejected),
            Some(&next_workflow_label) => {
                Ok(ProcessResult::Next(String::from(next_workflow_label)))
            }
            None => Err(format!("Missing last part of workflow")),
        }?;

        let op = if parts.len() == 1 {
            Op::Any
        } else {
            let (eval_expr_parts, op_symbol) = match parts.first() {
                Some(&expr) if expr.contains("<") => Ok((expr.split("<").collect::<Vec<_>>(), "<")),
                Some(&expr) if expr.contains(">") => Ok((expr.split(">").collect::<Vec<_>>(), ">")),
                _ => Err(format!("Got an invalid eval expr")),
            }?;
            if eval_expr_parts.len() != 2 {
                return Err(format!("Got an invalid eval expr"));
            }

            let rating_category = RatingCategory::try_from_str(eval_expr_parts[0])?;
            let threshold = match eval_expr_parts[1].parse::<usize>() {
                Ok(value) => Ok(value),
                Err(err) => Err(err.to_string()),
            }?;

            if op_symbol == "<" {
                Op::LessThan(rating_category, threshold)
            } else {
                Op::GreaterThan(rating_category, threshold)
            }
        };

        Ok(Self {
            op,
            affirmative_result,
        })
    }

    fn to_predicate(&self) -> Box<WorkflowPredicate> {
        let affirmative_result = self.affirmative_result.clone();

        match self.op {
            Op::Any => Box::new(move |_: &Part| affirmative_result.clone()),
            Op::GreaterThan(rating_category, threshold) => Box::new(move |part: &Part| {
                let rating = part.ratings[rating_category.index()];
                if rating > threshold {
                    affirmative_result.clone()
                } else {
                    ProcessResult::PassAlong
                }
            }),
            Op::LessThan(rating_category, threshold) => Box::new(move |part: &Part| {
                let rating = part.ratings[rating_category.index()];
                if rating < threshold {
                    affirmative_result.clone()
                } else {
                    ProcessResult::PassAlong
                }
            }),
        }
    }
}

enum Op {
    Any,
    GreaterThan(RatingCategory, usize),
    LessThan(RatingCategory, usize),
}

#[derive(Clone, PartialEq, Eq)]
enum ProcessResult {
    Accepted,
    Rejected,
    PassAlong,
    Next(String),
}

impl ProcessResult {
    fn to_label(&self) -> String {
        match self {
            Self::Accepted => String::from(ACCEPTED_LABEL),
            Self::Rejected => String::from(REJECTED_LABEL),
            Self::PassAlong => String::from("__NOT A LABEL__"),
            Self::Next(label) => label.clone(),
        }
    }
}

struct Part {
    ratings: [usize; RatingCategory::NRatingCategories.index()],
}

impl Part {
    fn try_from_line(line: &str) -> Result<Self, String> {
        let mut trimmed_line = line.replace("}", "");
        trimmed_line = trimmed_line.replace("{", "");
        let parts = trimmed_line.split(",").collect::<Vec<_>>();

        if parts.len() != 4 {
            return Err(format!("{line} doesn't have the right number of parts"));
        }

        let mut ratings = [0_usize; RatingCategory::NRatingCategories.index()];
        let mut rated_flags = 0_u8;
        for part in parts.iter() {
            let rating_parts = part.split("=").collect::<Vec<_>>();
            if rating_parts.len() != 2 {
                return Err(format!("{part} was a bad rating"));
            }

            let category = RatingCategory::try_from_str(rating_parts[0])?;
            let value = match rating_parts[1].parse::<usize>() {
                Ok(v) => Ok(v),
                Err(err) => Err(err.to_string()),
            }?;

            ratings[category.index()] = value;
            rated_flags |= 1 << category.index();
        }

        if rated_flags != 0x0F {
            return Err(String::from("not all categories were rated"));
        }

        Ok(Self { ratings })
    }

    fn total_rating(&self) -> usize {
        self.ratings.iter().sum()
    }
}

#[derive(Clone, Copy)]
enum RatingCategory {
    ExtremelyCoolLooking,
    Musical,
    Aerodynamic,
    Shiny,
    NRatingCategories,
}

impl RatingCategory {
    const fn index(&self) -> usize {
        *self as usize
    }

    fn try_from_str(s: &str) -> Result<Self, String> {
        match s {
            "x" => Ok(Self::ExtremelyCoolLooking),
            "m" => Ok(Self::Musical),
            "a" => Ok(Self::Aerodynamic),
            "s" => Ok(Self::Shiny),
            _ => Err(format!("{s} is not a valid part rating category")),
        }
    }
}

#[derive(Clone, PartialEq, Eq, Hash)]
struct ProcessLocator {
    label: String,
    step: usize,
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
struct Range {
    bounds: Option<(usize, usize)>,
}

impl Range {
    fn new(start: usize, end: usize) -> Self {
        if end < start {
            Self { bounds: None }
        } else {
            Self {
                bounds: Some((start, end)),
            }
        }
    }

    fn cap(&self, upper_bound: usize) -> Self {
        match self.bounds {
            Some((start, end)) => {
                if upper_bound < start {
                    Self { bounds: None }
                } else {
                    Self {
                        bounds: Some((start, usize::min(end, upper_bound))),
                    }
                }
            }
            None => Self { bounds: None },
        }
    }

    fn raise(&self, lower_bound: usize) -> Range {
        match self.bounds {
            Some((start, end)) => {
                if lower_bound > end {
                    Self { bounds: None }
                } else {
                    Self {
                        bounds: Some((usize::max(start, lower_bound), end)),
                    }
                }
            }
            None => Self { bounds: None },
        }
    }

    fn len(&self) -> usize {
        match self.bounds {
            Some((start, end)) => end - start + 1,
            None => 0,
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
            px{a<2006:qkq,m>2090:A,rfg}
            pv{a>1716:R,A}
            lnx{m>1548:A,A}
            rfg{s<537:gd,x>2440:R,A}
            qs{s>3448:A,lnx}
            qkq{x<1416:A,crn}
            crn{x>2662:A,R}
            in{s<1351:px,qqz}
            qqz{s>2770:qs,m<1801:hdj,R}
            gd{a>3333:R,R}
            hdj{m>838:A,pv}

            {x=787,m=2655,a=1222,s=2876}
            {x=1679,m=44,a=2067,s=496}
            {x=2036,m=264,a=79,s=2244}
            {x=2461,m=1339,a=466,s=291}
            {x=2127,m=1623,a=2188,s=1013}
        "});
        let mut solver = Day19::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("19114", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            px{a<2006:qkq,m>2090:A,rfg}
            pv{a>1716:R,A}
            lnx{m>1548:A,A}
            rfg{s<537:gd,x>2440:R,A}
            qs{s>3448:A,lnx}
            qkq{x<1416:A,crn}
            crn{x>2662:A,R}
            in{s<1351:px,qqz}
            qqz{s>2770:qs,m<1801:hdj,R}
            gd{a>3333:R,R}
            hdj{m>838:A,pv}

            {x=787,m=2655,a=1222,s=2876}
            {x=1679,m=44,a=2067,s=496}
            {x=2036,m=264,a=79,s=2244}
            {x=2461,m=1339,a=466,s=291}
            {x=2127,m=1623,a=2188,s=1013}
        "});
        let mut solver = Day19::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("167409079868000", result);

        Ok(())
    }
}
