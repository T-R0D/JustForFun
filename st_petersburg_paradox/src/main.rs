/*
	Casino
	Casinos have Tables.
	Tables have inconsistent Coins.
	People enter Casinos.
	Players play at Tables.
	Players have a purse.
	Tables have a bankroll.
	Casino policy specifies a starting pot and a buy_in.
	Tables execute games independently of eachother.
	One Player may play at a table at a time.
	Players dweet their results at the end of play.

	If p Players with d dollars in their purse enter the casino to play
	at t Tables,
	how many winners/losers, average/min/max payouts.
*/

extern crate argparse;

/*
mod casino;
mod patrons;

use casino::establishment;
use patrons::person::Person;
*/

extern crate rand;


fn play_round() -> i64 {
	let mut pot = 2;
	'game: loop {
		if rand::random() {
			pot *= 2;
		} else {
			break 'game;
		}
	}

	pot
}


fn main() {
	let mut buy_in = 25;
	let mut n_bets = 1;
	let mut n_trials = 3;

	{
		let mut arg_parser = argparse::ArgumentParser::new();
		arg_parser.refer(&mut buy_in)
			      .add_option(&["-b", "--buyin"], argparse::Store, "Cost to play a round.");
		arg_parser.refer(&mut n_bets)
			      .add_option(&["-p", "--plays"], argparse::Store, "How many rounds to play.");
		arg_parser.refer(&mut n_trials)
			      .add_option(&["-t", "--trials"], argparse::Store, "How many trials to run.");

	    arg_parser.parse_args_or_exit();
	}

	let mut results = Vec::<i64>::new();

	for _ in 0..n_trials {
		let mut total_winnings = 0;
		for _ in 0..n_bets {
			total_winnings += play_round();
		}
		results.push(total_winnings);
	}

	results.sort();
	let mut sum: f64 = 0_f64;
	for result in results.iter() {
		sum += *result as f64;
	}


	println!("---\nbuy-in: {}\tbets: {}\t trials: {}\n---", buy_in, n_bets, n_trials);
	println!("Mean Game:    {}", sum / n_trials as f64);
	println!("Mean Outcome: {}    <===", (sum / n_trials as f64) - (n_bets * buy_in) as f64);
	println!("Best Run:     {}", results.last().unwrap());
	println!("Worst Run:    {}", results.first().unwrap());

	/*
    let mut verbose = false;
    let mut casino_name = "".to_string();
    {
    	let mut arg_parser = argparse::ArgumentParser::new();
    	arg_parser.set_description("A simulator to demonstrate \
    	                            the St. Petersburg Paradox.");

    	arg_parser.refer(&mut casino_name)
    	          .add_argument("casino", argparse::Store, "The casino name.")
    	          .required();

    	arg_parser.refer(&mut verbose)
    	          .add_option(&["-v", "--verbose"], argparse::StoreTrue, 
    	                      "Be verbose.");

    	arg_parser.parse_args_or_exit();
    }

    if verbose {
    }


   	let p = &mut Person::new("Doofus", 100, 100);
   	println!("{}", p);
   	p.dweet();

   	let casino = &mut establishment::Establishment::new();
   	casino.admit(p);
   	*/
}
