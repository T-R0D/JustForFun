use std::{cell::RefCell, vec};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day13 {
    packet_pairs: Vec<PacketPair>,
}

impl Day13 {
    pub fn new() -> Self {
        Day13 {
            packet_pairs: vec![],
        }
    }
}

impl Solver for Day13 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.packet_pairs = input
            .trim()
            .split("\n\n")
            .map(|pair_text| {
                let mut packets = pair_text
                    .lines()
                    .map(|line| Packet::from_line(line))
                    .collect::<Vec<_>>();
                let p1 = packets.pop().unwrap();
                let p0 = packets.pop().unwrap();
                PacketPair(p0, p1)
            })
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut in_order_pair_indices: Vec<usize> = Vec::new();

        for (i, pair) in self.packet_pairs.iter().enumerate() {
            if pair.is_in_correct_order() {
                in_order_pair_indices.push(i + 1);
            }
        }

        Ok(in_order_pair_indices.iter().sum::<usize>().to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let packet_pairs = self.packet_pairs.clone();
        let mut packets = packet_pairs
            .iter()
            .map(|pair| Vec::from([pair.0.clone(), pair.1.clone()]))
            .flatten()
            .collect::<Vec<_>>();

        let first_divider_packet = &Packet::List(vec![Packet::List(vec![Packet::Int(2)])]);
        let second_divider_packet = &Packet::List(vec![Packet::List(vec![Packet::Int(6)])]);
        packets.push(first_divider_packet.clone());
        packets.push(second_divider_packet.clone());

        packets.sort_by(|a, b| {
            let pair = PacketPair(a.clone(), b.clone());
            if pair.is_in_correct_order() {
                return std::cmp::Ordering::Less;
            }
            std::cmp::Ordering::Greater
        });

        let mut product = 1;
        for divider_packet in [first_divider_packet, second_divider_packet] {
            product *= packets.iter().position(|x| *x == *divider_packet).unwrap() + 1;
        }

        Ok(product.to_string())
    }
}

#[derive(Clone)]
struct PacketPair(Packet, Packet);

impl PacketPair {
    fn is_in_correct_order(&self) -> bool {
        match (&self.0, &self.1) {
            (Packet::List(contents_a), Packet::List(contents_b)) => {
                match self.list_contents_correctly_ordered(contents_a, contents_b) {
                    Ordering::Correct | Ordering::Same => true,
                    Ordering::Incorrect => false,
                }
            }
            _ => panic!("must start with 2 lists"),
        }
    }

    fn list_contents_correctly_ordered(
        &self,
        contents_a: &Vec<Packet>,
        contents_b: &Vec<Packet>,
    ) -> Ordering {
        for (a, b) in std::iter::zip(contents_a.iter(), contents_b.iter()) {
            let element_comparison = match (&a, &b) {
                (Packet::List(new_a), Packet::List(new_b)) => {
                    self.list_contents_correctly_ordered(new_a, new_b)
                }
                (Packet::List(new_a), Packet::Int(y)) => {
                    let new_b = &vec![Packet::Int(*y)];
                    self.list_contents_correctly_ordered(new_a, new_b)
                }
                (Packet::Int(x), Packet::List(new_b)) => {
                    let new_a = &vec![Packet::Int(*x)];
                    self.list_contents_correctly_ordered(new_a, new_b)
                }
                (Packet::Int(x), Packet::Int(y)) => {
                    if x < y {
                        Ordering::Correct
                    } else if x == y {
                        Ordering::Same
                    } else {
                        Ordering::Incorrect
                    }
                }
            };
            match element_comparison {
                Ordering::Correct => {
                    return Ordering::Correct;
                }
                Ordering::Same => {
                    continue;
                }
                Ordering::Incorrect => return Ordering::Incorrect,
            }
        }

        if contents_a.len() < contents_b.len() {
            return Ordering::Correct;
        } else if contents_a.len() == contents_b.len() {
            return Ordering::Same;
        } else {
            return Ordering::Incorrect;
        }
    }
}

#[derive(Clone, Eq, PartialEq)]
enum Packet {
    List(Vec<Packet>),
    Int(u8),
}

enum Ordering {
    Correct,
    Same,
    Incorrect,
}

impl Packet {
    fn from_line(line: &str) -> Packet {
        let mut component_stack: Vec<RefCell<Vec<Packet>>> = Vec::new();
        let mut int_builder = String::from("");
        let mut list_contents: Vec<Packet> = Vec::new();

        let mut chars = line.bytes().skip(1);
        chars.next_back();

        for c in chars {
            if c == b'[' {
                component_stack.push(RefCell::new(list_contents));
                list_contents = Vec::new();
            } else if c == b']' {
                if int_builder.len() > 0 {
                    let n = int_builder.parse::<u8>().unwrap();
                    int_builder.clear();
                    list_contents.push(Packet::Int(n));
                }

                let l = Packet::List(list_contents);

                list_contents = component_stack.pop().unwrap().take();
                list_contents.push(l);
            } else if c == b',' {
                if int_builder.len() > 0 {
                    let n = int_builder.parse::<u8>().unwrap();
                    int_builder.clear();
                    list_contents.push(Packet::Int(n));
                }
            } else {
                int_builder.push(c as char);
            }
        }

        if int_builder.len() > 0 {
            let n = int_builder.parse::<u8>().unwrap();
            int_builder.clear();
            list_contents.push(Packet::Int(n));
        }

        Packet::List(list_contents)
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
            [1,1,3,1,1]
            [1,1,5,1,1]

            [[1],[2,3,4]]
            [[1],4]

            [9]
            [[8,7,6]]

            [[4,4],4,4]
            [[4,4],4,4,4]

            [7,7,7,7]
            [7,7,7]

            []
            [3]

            [[[]]]
            [[]]

            [1,[2,[3,[4,[5,6,7]]]],8,9]
            [1,[2,[3,[4,[5,6,0]]]],8,9]
        "});
        let mut solver = Day13::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("13", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            [1,1,3,1,1]
            [1,1,5,1,1]

            [[1],[2,3,4]]
            [[1],4]

            [9]
            [[8,7,6]]

            [[4,4],4,4]
            [[4,4],4,4,4]

            [7,7,7,7]
            [7,7,7]

            []
            [3]

            [[[]]]
            [[]]

            [1,[2,[3,[4,[5,6,7]]]],8,9]
            [1,[2,[3,[4,[5,6,0]]]],8,9]
        "});
        let mut solver = Day13::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("140", result);

        Ok(())
    }
}
