use std::collections::HashMap;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day21 {
    monkey_to_job: HashMap<String, Op>,
}

impl Day21 {
    pub fn new() -> Self {
        Self {
            monkey_to_job: HashMap::new(),
        }
    }
}

impl Solver for Day21 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.monkey_to_job = input
            .trim()
            .lines()
            .map(|line| {
                let parts = line.split(": ").collect::<Vec<_>>();
                let name = String::from(parts[0]);

                let yell_num_result = parts[1].parse::<i64>();
                let op = if yell_num_result.is_ok() {
                    Op::Yell(yell_num_result.unwrap())
                } else if parts[1].contains("+") {
                    let op_parts = parts[1].split(" + ").map(String::from).collect::<Vec<_>>();
                    Op::Add(op_parts[0].clone(), op_parts[1].clone())
                } else if parts[1].contains("-") {
                    let op_parts = parts[1].split(" - ").map(String::from).collect::<Vec<_>>();
                    Op::Sub(op_parts[0].clone(), op_parts[1].clone())
                } else if parts[1].contains("*") {
                    let op_parts = parts[1].split(" * ").map(String::from).collect::<Vec<_>>();
                    Op::Mul(op_parts[0].clone(), op_parts[1].clone())
                } else {
                    let op_parts = parts[1].split(" / ").map(String::from).collect::<Vec<_>>();
                    Op::Div(op_parts[0].clone(), op_parts[1].clone())
                };

                (name, op)
            })
            .fold(HashMap::new(), |mut acc, (name, op)| {
                acc.insert(name, op);
                acc
            });

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let root_monkey_name = &String::from(ROOT_MONKEY_NAME);
        let result = evaluate_moneky_expression_tree(&self.monkey_to_job, root_monkey_name);

        Ok(result.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let name_to_value = construct_value_tree(&self.monkey_to_job);

        let human_yell_value = determine_human_yell_value(&name_to_value, &self.monkey_to_job);

        Ok(human_yell_value.to_string())
    }
}

const ROOT_MONKEY_NAME: &str = "root";
const HUMAN_NAME: &str = "humn";

#[derive(Clone, Debug)]
enum Op {
    Yell(i64),
    Add(String, String),
    Sub(String, String),
    Mul(String, String),
    Div(String, String),
}

#[derive(Debug)]
struct EvalState {
    op: Op,
    left: Option<i64>,
    right: Option<i64>,
}

impl EvalState {
    fn new(op: &Op) -> Self {
        Self {
            op: op.clone(),
            left: None,
            right: None,
        }
    }

    fn next_child_op_name(&self) -> String {
        match self.op.clone() {
            Op::Add(a, b) | Op::Sub(a, b) | Op::Mul(a, b) | Op::Div(a, b) => {
                self.pick_next_name(&a, &b)
            }
            Op::Yell(_) => panic!("Yell ops have no child ops"),
        }
    }

    fn pick_next_name(&self, a: &String, b: &String) -> String {
        if self.left.is_none() {
            a.to_string()
        } else {
            b.to_string()
        }
    }

    fn can_be_evaluated(&self) -> bool {
        match self.op {
            Op::Yell(_) => true,
            _ => self.left.is_some() && self.right.is_some(),
        }
    }

    fn evaluate(&self) -> i64 {
        match (self.left, self.right) {
            (Some(a), Some(b)) => match self.op {
                Op::Add(_, _) => a + b,
                Op::Sub(_, _) => a - b,
                Op::Mul(_, _) => a * b,
                Op::Div(_, _) => a / b,
                Op::Yell(_) => unreachable!(),
            },
            _ => match self.op {
                Op::Yell(x) => x,
                _ => panic!("Not ready to evaluate! {:?}", self),
            },
        }
    }
}

fn evaluate_moneky_expression_tree(
    monkey_to_job: &HashMap<String, Op>,
    starting_monkey_name: &String,
) -> i64 {
    let root_op = monkey_to_job.get(starting_monkey_name).unwrap();
    let root_eval_state = EvalState::new(root_op);
    let mut evaluation_stack = Vec::<EvalState>::from([root_eval_state]);

    'evaluation: while let Some(eval_state) = evaluation_stack.pop() {
        if eval_state.can_be_evaluated() {
            let result = eval_state.evaluate();
            match evaluation_stack.last_mut() {
                Some(prev_eval_state) => {
                    if prev_eval_state.left.is_none() {
                        prev_eval_state.left = Some(result);
                    } else {
                        prev_eval_state.right = Some(result);
                    }
                }
                None => {
                    return result;
                }
            }
            continue 'evaluation;
        }

        let next_op_name = eval_state.next_child_op_name();
        let next_eval_state = EvalState::new(monkey_to_job.get(&next_op_name).unwrap());
        evaluation_stack.push(eval_state);
        evaluation_stack.push(next_eval_state);
    }

    unreachable!("Somehow didn't evaluate a result...");
}

#[derive(Debug)]
enum Branch {
    Left,
    Right,
    Done,
}

#[derive(Debug)]
struct ValueNode {
    name: String,
    next_branch: Branch,
    yell: i64,
    left: Option<i64>,
    right: Option<i64>,
}

impl ValueNode {
    fn new(name: &String) -> Self {
        Self {
            name: name.clone(),
            next_branch: Branch::Left,
            yell: 0,
            left: None,
            right: None,
        }
    }

    fn value_is_known(&self) -> bool {
        match self.next_branch {
            Branch::Done => self.left.is_some() && self.right.is_some(),
            _ => false,
        }
    }

    fn update_next_branch(&mut self) {
        self.next_branch = match self.next_branch {
            Branch::Left => Branch::Right,
            Branch::Right => Branch::Done,
            Branch::Done => unreachable!(),
        }
    }
}

fn construct_value_tree(monkey_to_job: &HashMap<String, Op>) -> HashMap<String, ValueNode> {
    let root_eval_state = ValueNode::new(&String::from(ROOT_MONKEY_NAME));
    let mut evaluation_stack = Vec::<ValueNode>::from([root_eval_state]);
    let mut name_to_value = HashMap::<String, ValueNode>::new();

    'search: while let Some(mut node) = evaluation_stack.pop() {
        match node.next_branch {
            Branch::Done => {
                match evaluation_stack.pop() {
                    Some(mut prev_node) => {
                        if node.value_is_known() {
                            match prev_node.next_branch {
                                Branch::Left => {
                                    prev_node.left = Some(node.yell);
                                }
                                Branch::Right => {
                                    prev_node.right = Some(node.yell);
                                }
                                Branch::Done => unreachable!(),
                            }
                        }
                        prev_node.update_next_branch();

                        if prev_node.left.is_some() && prev_node.right.is_some() {
                            let op = monkey_to_job.get(&prev_node.name).unwrap();
                            prev_node.yell = match op {
                                Op::Add(_, _) => prev_node.left.unwrap() + prev_node.right.unwrap(),
                                Op::Sub(_, _) => prev_node.left.unwrap() - prev_node.right.unwrap(),
                                Op::Mul(_, _) => prev_node.left.unwrap() * prev_node.right.unwrap(),
                                Op::Div(_, _) => prev_node.left.unwrap() / prev_node.right.unwrap(),
                                Op::Yell(_) => unreachable!(),
                            }
                        }
                        evaluation_stack.push(prev_node);
                    }
                    None => (),
                }

                name_to_value.insert(node.name.clone(), node);

                continue 'search;
            }
            _ => (),
        }

        if node.name == HUMAN_NAME {
            name_to_value.insert(node.name.clone(), node);

            match evaluation_stack.pop() {
                Some(mut prev_node) => {
                    prev_node.update_next_branch();
                    evaluation_stack.push(prev_node);
                }
                None => (),
            }

            continue 'search;
        }

        let op = monkey_to_job.get(&node.name).unwrap();
        match op {
            Op::Add(left, right)
            | Op::Sub(left, right)
            | Op::Mul(left, right)
            | Op::Div(left, right) => {
                let next_node = match node.next_branch {
                    Branch::Left => ValueNode::new(left),
                    Branch::Right => ValueNode::new(right),
                    Branch::Done => unreachable!(),
                };
                evaluation_stack.push(node);
                evaluation_stack.push(next_node);
            }
            Op::Yell(x) => {
                node.yell = *x;
                node.next_branch = Branch::Done;
                node.left = Some(0);
                node.right = Some(0);
                evaluation_stack.push(node);
            }
        }
    }

    name_to_value
}

fn determine_human_yell_value(
    name_to_value: &HashMap<String, ValueNode>,
    name_to_op: &HashMap<String, Op>,
) -> i64 {
    let root_node = name_to_value.get(&String::from(ROOT_MONKEY_NAME)).unwrap();
    let root_op = name_to_op.get(&root_node.name).unwrap();
    let mut current_value_to_create = match (root_node.left, root_node.right) {
        (Some(value), None) => value,
        (None, Some(value)) => value,
        _ => unreachable!(),
    };

    let mut current_node = match root_op {
        Op::Add(left, right)
        | Op::Sub(left, right)
        | Op::Mul(left, right)
        | Op::Div(left, right) => {
            if root_node.left.is_none() {
                name_to_value.get(left).unwrap()
            } else {
                name_to_value.get(right).unwrap()
            }
        }
        Op::Yell(_) => unreachable!(),
    };

    while current_node.name != HUMAN_NAME {
        let op = name_to_op.get(&current_node.name).unwrap();

        current_value_to_create = match op.clone() {
            Op::Add(left, right) => match (current_node.left, current_node.right) {
                (Some(val), None) => {
                    current_node = name_to_value.get(&right).unwrap();
                    current_value_to_create - val
                }
                (None, Some(val)) => {
                    current_node = name_to_value.get(&left).unwrap();
                    current_value_to_create - val
                }
                _ => unreachable!(),
            },
            Op::Sub(left, right) => match (current_node.left, current_node.right) {
                (Some(val), None) => {
                    current_node = name_to_value.get(&right).unwrap();
                    val - current_value_to_create
                }
                (None, Some(val)) => {
                    current_node = name_to_value.get(&left).unwrap();
                    current_value_to_create + val
                }
                _ => unreachable!(),
            },
            Op::Mul(left, right) => match (current_node.left, current_node.right) {
                (Some(val), None) => {
                    current_node = name_to_value.get(&right).unwrap();
                    current_value_to_create / val
                }
                (None, Some(val)) => {
                    current_node = name_to_value.get(&left).unwrap();
                    current_value_to_create / val
                }
                _ => unreachable!(),
            },
            Op::Div(left, right) => match (current_node.left, current_node.right) {
                (Some(val), None) => {
                    current_node = name_to_value.get(&right).unwrap();
                    val / current_value_to_create
                }
                (None, Some(val)) => {
                    current_node = name_to_value.get(&left).unwrap();
                    current_value_to_create * val
                }
                _ => unreachable!(),
            },
            Op::Yell(_) => unreachable!(),
        }
    }

    current_value_to_create
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
            root: pppw + sjmn
            dbpl: 5
            cczh: sllz + lgvd
            zczc: 2
            ptdq: humn - dvpt
            dvpt: 3
            lfqf: 4
            humn: 5
            ljgn: 2
            sjmn: drzm * dbpl
            sllz: 4
            pppw: cczh / lfqf
            lgvd: ljgn * ptdq
            drzm: hmdt - zczc
            hmdt: 32
        "});
        let mut solver = Day21::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("152", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            root: pppw + sjmn
            dbpl: 5
            cczh: sllz + lgvd
            zczc: 2
            ptdq: humn - dvpt
            dvpt: 3
            lfqf: 4
            humn: 5
            ljgn: 2
            sjmn: drzm * dbpl
            sllz: 4
            pppw: cczh / lfqf
            lgvd: ljgn * ptdq
            drzm: hmdt - zczc
            hmdt: 32
        "});
        let mut solver = Day21::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("301", result);

        Ok(())
    }
}

/*
{"zczc": ValueNode { name: "zczc", next_branch: Done, yell: 2, left: None, right: None }, "humn": ValueNode { name: "humn", next_branch: Left, yell: 0, left: None, right: None }, "dbpl": ValueNode { name: "dbpl", next_branch: Done, yell: 5, left: None, right: None }, "ljgn": ValueNode { name: "ljgn", next_branch: Done, yell: 2, left: None, right: None }, "pppw": ValueNode { name: "pppw", next_branch: Done, yell: 0, left: None, right: None }, "cczh": ValueNode { name: "cczh", next_branch: Done, yell: 0, left: None, right: None }, "ptdq": ValueNode { name: "ptdq", next_branch: Done, yell: 0, left: None, right: None }, "root": ValueNode { name: "root", next_branch: Done, yell: 0, left: None, right: None }, "lgvd": ValueNode { name: "lgvd", next_branch: Done, yell: 0, left: None, right: None }, "lfqf": ValueNode { name: "lfqf", next_branch: Done, yell: 4, left: None, right: None }, "hmdt": ValueNode { name: "hmdt", next_branch: Done, yell: 32, left: None, right: None }, "sllz": ValueNode { name: "sllz", next_branch: Done, yell: 4, left: None, right: None }, "sjmn": ValueNode { name: "sjmn", next_branch: Done, yell: 0, left: None, right: None }, "dvpt": ValueNode { name: "dvpt", next_branch: Done, yell: 3, left: None, right: None }, "drzm": ValueNode { name: "drzm", next_branch: Done, yell: 0, left: None, right: None }}
*/
