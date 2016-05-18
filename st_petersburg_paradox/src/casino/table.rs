
use casino::coin::Coin;
use patrons::person;

pub struct Table {
	coin: Coin,
	bankroll: i64,
	pot_start: i64,
	buy_in: i64,
	player: Option<person::Person>,
}


impl Table {
	pub fn with_policy(bankroll: i64, pot_start: i64, buy_in: i64) -> Table {
		Table {
			coin: Coin::new(),
			bankroll: bankroll,
			pot_start: pot_start,
			buy_in: buy_in,
			player: None,
		}
	}

	pub fn admit(&mut self, player: &mut person::Person) -> () {
		match self.player {
			None => {self.player = Some(player)}
			Some(_) => {println!("Too many people at the table!")}
		}
	}
}
