use std::{cmp::Ordering, convert::TryInto};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day07 {
    starting_hands: Vec<Hand>,
}

impl Day07 {
    pub fn new() -> Self {
        Day07 {
            starting_hands: vec![],
        }
    }
}

impl Solver for Day07 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.starting_hands = input
            .trim()
            .lines()
            .map(|line| Hand::try_from_line(line))
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut hands = self.starting_hands.clone();
        hands.sort_by(Hand::straight_cmp);
        let total_winnings = hands
            .iter()
            .rev()
            .enumerate()
            .map(|(i, hand)| (i + 1) * usize::try_from(hand.bid).unwrap())
            .sum::<usize>();

        Ok(total_winnings.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut hands = self.starting_hands.clone();
        hands.sort_by(Hand::jacks_as_jokers_cmp);
        let total_winnings = hands
            .iter()
            .rev()
            .enumerate()
            .map(|(i, hand)| (i + 1) * usize::try_from(hand.bid).unwrap())
            .sum::<usize>();

        Ok(total_winnings.to_string())
    }
}

const CARDS_IN_HAND: usize = 5;

#[derive(Clone)]
struct Hand {
    cards: [Card; CARDS_IN_HAND],
    bid: u32,
}

impl Hand {
    pub fn try_from_line(line: &str) -> Result<Hand, String> {
        let parts = line.split(" ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("{line} is invalid"));
        }

        let cards_vec = parts[0]
            .chars()
            .map(|c| Card::try_from_char(c))
            .collect::<Result<Vec<_>, String>>()?;
        let cards: [Card; CARDS_IN_HAND] = match cards_vec.try_into() {
            Ok(cards) => Ok(cards),
            Err(_) => Err(format!("{} had the wrong number of cards", parts[0])),
        }?;

        let bid = match parts[1].parse::<u32>() {
            Ok(value) => Ok(value),
            Err(err) => Err(err.to_string()),
        }?;

        Ok(Hand { cards, bid })
    }

    fn straight_cmp(a: &Self, b: &Self) -> Ordering {
        let (a_hand_type, b_hand_type) = (a.straight_hand_type(), b.straight_hand_type());

        if a_hand_type != b_hand_type {
            return a_hand_type.cmp(&b_hand_type);
        }

        for (a_card, b_card) in a.cards.iter().zip(b.cards.iter()) {
            if a_card != b_card {
                return a_card.cmp(b_card);
            }
        }

        Ordering::Equal
    }

    fn jacks_as_jokers_cmp(a: &Self, b: &Self) -> Ordering {
        let (a_hand_type, b_hand_type) =
            (a.jacks_as_jokers_hand_type(), b.jacks_as_jokers_hand_type());

        if a_hand_type != b_hand_type {
            return a_hand_type.cmp(&b_hand_type);
        }

        for (a_card, b_card) in a.cards.iter().zip(b.cards.iter()) {
            if a_card != b_card {
                if *a_card == Card::Jack {
                    return Ordering::Greater;
                }

                if *b_card == Card::Jack {
                    return Ordering::Less;
                }

                return a_card.cmp(b_card);
            }
        }

        Ordering::Equal
    }

    fn straight_hand_type(&self) -> HandType {
        let mut counts = [0; Card::NCardTypes as usize];
        for card in self.cards.iter() {
            counts[card.index()] += 1;
        }

        let mut has_triple = false;
        let mut has_pair = false;
        for &count in counts.iter() {
            if count == 5 {
                return HandType::FiveOfAKind;
            }

            if count == 4 {
                return HandType::FourOfAKind;
            }

            if count == 3 {
                has_triple = true;

                if has_pair {
                    return HandType::FullHouse;
                }
            }

            if count == 2 {
                if has_pair {
                    return HandType::TwoPair;
                } else if has_triple {
                    return HandType::FullHouse;
                }

                has_pair = true;
            }
        }

        if has_triple {
            return HandType::ThreeOfAKind;
        }

        if has_pair {
            return HandType::OnePair;
        }

        HandType::HighCard
    }

    fn jacks_as_jokers_hand_type(&self) -> HandType {
        let mut counts = [0_u8; Card::NCardTypes as usize];
        for card in self.cards.iter() {
            counts[card.index()] += 1;
        }

        let jacks_count = counts[Card::Jack.index()];
        let mut sorted_counts = counts
            .iter()
            .enumerate()
            .map(|(i, count)| (*count, i))
            .collect::<Vec<(u8, usize)>>();
        sorted_counts.sort_by_key(|(count, _)| *count);

        for (count, card) in sorted_counts.iter_mut().rev() {
            if *card != Card::Jack.index() {
                *count += jacks_count;
                break;
            }
        }
        sorted_counts.sort_by_key(|(count, _)| *count);

        let mut has_pair = false;
        let mut has_triple = false;
        for &(count, card) in sorted_counts.iter().rev() {
            if card == Card::Jack.index() {
                continue;
            }

            if count == 5 {
                return HandType::FiveOfAKind;
            }

            if count == 4 {
                return HandType::FourOfAKind;
            }

            if count == 3 {
                has_triple = true;

                if has_pair {
                    return HandType::FullHouse;
                }
            }

            if count == 2 {
                if has_pair {
                    return HandType::TwoPair;
                } else if has_triple {
                    return HandType::FullHouse;
                }

                has_pair = true;
            }
        }

        if has_triple {
            return HandType::ThreeOfAKind;
        }

        if has_pair {
            return HandType::OnePair;
        }

        HandType::HighCard
    }
}

#[derive(Clone, Copy, PartialOrd, Ord, PartialEq, Eq)]
enum HandType {
    FiveOfAKind,
    FourOfAKind,
    FullHouse,
    ThreeOfAKind,
    TwoPair,
    OnePair,
    HighCard,
}

#[derive(Clone, Copy, PartialOrd, Ord, PartialEq, Eq)]
enum Card {
    Ace,
    King,
    Queen,
    Jack,
    Ten,
    Nine,
    Eight,
    Seven,
    Six,
    Five,
    Four,
    Three,
    Two,
    NCardTypes,
}

impl Card {
    fn try_from_char(c: char) -> Result<Card, String> {
        match c {
            'A' => Ok(Card::Ace),
            'K' => Ok(Card::King),
            'Q' => Ok(Card::Queen),
            'J' => Ok(Card::Jack),
            'T' => Ok(Card::Ten),
            '9' => Ok(Card::Nine),
            '8' => Ok(Card::Eight),
            '7' => Ok(Card::Seven),
            '6' => Ok(Card::Six),
            '5' => Ok(Card::Five),
            '4' => Ok(Card::Four),
            '3' => Ok(Card::Three),
            '2' => Ok(Card::Two),
            _ => Err(format!("{c} is not a valid card")),
        }
    }

    pub fn index(&self) -> usize {
        *self as usize
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
            32T3K 765
            T55J5 684
            KK677 28
            KTJJT 220
            QQQJA 483
        "});
        let mut solver = Day07::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("6440", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            32T3K 765
            T55J5 684
            KK677 28
            KTJJT 220
            QQQJA 483
        "});
        let mut solver = Day07::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("5905", result);

        Ok(())
    }
}
