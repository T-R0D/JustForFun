extern crate hyper;

// use std::io::Read;

use self::hyper::Client;
// use self::hyper::header::Connection;


pub struct Person {
	name: String,
	purse: i64,
	willing_to_spend: i64,
	winnings: i64,
}

impl Person {
	pub fn new(name: &str, purse: i64, willing_to_spend:i64) -> Person {
		Person {
			name: name.to_string(),
			purse: purse,
			willing_to_spend: willing_to_spend,
			winnings: 0,
		}
	}

	pub fn dweet(&self) {
		let client = Client::new();

		let res = client.get("https://dweet.io/dweet/for/middle-finger?fuck=you")
		                .send().unwrap();
		assert_eq!(res.status, hyper::Ok);
	}
}

use std::fmt;
impl fmt::Display for Person {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		write!(f, "{} has ${}; willing to spend ${}.", self.name,
			   self.purse, self.willing_to_spend)
	}
}
