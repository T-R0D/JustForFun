extern crate rand;

#[derive(PartialEq)]
pub enum FlipOutcome {
	Heads,
	Tails
}

pub struct Coin {
	heads_probability: f64,
}

impl Coin {
	pub fn new() -> Coin {
		let mut head_prob: f64 = 0.5;
		let bias: f64 = rand::random::<f64>() / 1_000_f64;

		if rand::random() {
			head_prob += bias;
		} else {
			head_prob -= bias;
		}

		Coin {
			heads_probability: head_prob,
		}
	}

	pub fn flip(&self) -> FlipOutcome {
		let result = rand::random::<f64>();

		if result < self.heads_probability {
			return FlipOutcome::Heads;
		} else {
			return FlipOutcome::Tails;
		}
	}
}

use std::fmt;
impl fmt::Display for Coin {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		write!(f, "P(heads): {}", self.heads_probability)
	}
}
