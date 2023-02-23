use std::ptr;

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day20 {
    file_contents: Vec<i64>,
}

impl Day20 {
    pub fn new() -> Self {
        Self {
            file_contents: Vec::new(),
        }
    }
}

impl Solver for Day20 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.file_contents = input
            .trim()
            .lines()
            .map(str::parse::<i64>)
            .map(Result::unwrap)
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut checksum = 0;
        unsafe {
            let mut mixing_buffer = LinkedListMixingBuffer::from_list(&self.file_contents);

            mixing_buffer.mix();
            let (result, zero_index) = mixing_buffer.get_result();

            for index in CHECKSUM_INDICES {
                checksum += result[(zero_index + index) % result.len()];
            }
        }
        Ok(checksum.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut checksum = 0;
        unsafe {
            let keyed_contents = self
                .file_contents
                .iter()
                .map(|x| x * DECRYPTION_KEY)
                .collect::<Vec<_>>();

            let mut mixing_buffer = LinkedListMixingBuffer::from_list(&keyed_contents);

            for _ in 0..FULL_DECRYPTION_ITERATIONS {
                mixing_buffer.mix();
            }

            let (result, zero_index) = mixing_buffer.get_result();

            for index in CHECKSUM_INDICES {
                checksum += result[(zero_index + index) % result.len()];
            }
        }
        Ok(checksum.to_string())
    }
}

const CHECKSUM_INDICES: [usize; 3] = [1000, 2000, 3000];
const DECRYPTION_KEY: i64 = 811_589_153;
const FULL_DECRYPTION_ITERATIONS: usize = 10;

struct LinkedListMixingBuffer {
    original_order: Vec<i64>,
    node_lookup: Vec<*mut LinkedNode>,
    zero_index: usize,
}

impl LinkedListMixingBuffer {
    unsafe fn from_list(list: &Vec<i64>) -> Self {
        let original_order = list.clone();
        let mut node_lookup: Vec<*mut LinkedNode> = vec![ptr::null_mut(); original_order.len()];
        let mut zero_index = 0;

        for (i, &value) in original_order.iter().enumerate() {
            let node: *mut LinkedNode = Box::into_raw(Box::new(LinkedNode::new(value)));
            node_lookup[i] = node;

            if value == 0 {
                zero_index = i;
            }
        }

        for i in 1..=(node_lookup.len() - 1) {
            let prev_index = ((i as i64 - 1) % node_lookup.len() as i64) as usize;
            (*node_lookup[prev_index]).next = node_lookup[i];
            (*node_lookup[i]).prev = node_lookup[prev_index];

            let next_index = ((i as i64 + 1) % node_lookup.len() as i64) as usize;
            (*node_lookup[i]).next = node_lookup[next_index];
            (*node_lookup[next_index]).prev = node_lookup[i];
        }
        (*node_lookup[original_order.len() - 1]).next = node_lookup[0];
        (*node_lookup[0]).prev = node_lookup[original_order.len() - 1];

        Self {
            original_order,
            node_lookup,
            zero_index,
        }
    }

    unsafe fn mix(&mut self) {
        for (i, &target_val) in self.original_order.iter().enumerate() {
            let val = target_val % (self.original_order.len() - 1) as i64;

            if val == 0 {
                continue;
            }

            let mut target_node = self.node_lookup[i];

            (*(*target_node).prev).next = (*target_node).next;
            (*(*target_node).next).prev = (*target_node).prev;

            let mut cursor = target_node;
            if val < 0 {
                for _ in val..=0 {
                    cursor = (*cursor).prev;
                }
            } else {
                for _ in 0..val {
                    cursor = (*cursor).next;
                }
            }

            (*target_node).prev = cursor;
            (*target_node).next = (*cursor).next;
            let temp = (*cursor).next;
            (*cursor).next = target_node;
            (*temp).prev = target_node;
        }
    }

    unsafe fn get_result(&self) -> (Vec<i64>, usize) {
        let mut result: Vec<i64> = Vec::with_capacity(self.original_order.len());
        let mut zero_index = 0;

        let mut cursor = self.node_lookup[self.zero_index];
        for i in 0..self.original_order.len() {
            if (*cursor).val == 0 {
                zero_index = i;
            }
            result.push((*cursor).val);
            cursor = (*cursor).next;
        }

        (result, zero_index)
    }
}

struct LinkedNode {
    val: i64,
    next: *mut LinkedNode,
    prev: *mut LinkedNode,
}

impl LinkedNode {
    fn new(val: i64) -> Self {
        Self {
            val,
            next: ptr::null_mut(),
            prev: ptr::null_mut(),
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
            1
            2
            -3
            3
            -2
            0
            4
        "});
        let mut solver = Day20::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("3", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            1
            2
            -3
            3
            -2
            0
            4
        "});
        let mut solver = Day20::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("1623178306", result);

        Ok(())
    }
}
