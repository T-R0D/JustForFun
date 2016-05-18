use std::fmt;
use std::collections::HashMap;
use casino::table::Table;
use patrons::person::Person;

pub struct Establishment {
	name: String,
	guests: HashMap<String, Person>,
	tables: Vec<Table>,
	pot_start: i64,
	buy_in: i64,
}

impl Establishment {
	pub fn new() -> Establishment {
		let bank_roll = 1_000_000;
		let pot_start = 2;
		let buy_in = 25;
		let mut tables = Vec::new();
		tables.push(Table::with_policy(bank_roll, pot_start, buy_in));

		Establishment {
			name: "Buttlo's House of Vice".to_string(),
			guests: HashMap::new(),
			tables: tables,
			pot_start: pot_start,
			buy_in: buy_in,
		}
	}

	pub fn admit(&mut self, person: &mut Person) -> () {
		let mut first_table = &mut self.tables[0];
		first_table.admit(person);
	}
}

impl fmt::Display for Establishment {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		write!(f, "Welcome to {} where the pot starts at {} and the buy-in is just {}", self.name, self.pot_start, self.buy_in)
	}
}
