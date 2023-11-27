mod aoc2023;

use std::fs;
use std::time::Instant;

use crate::aoc2023::args::cmdline;
use crate::aoc2023::day00::solver::Day00;
use crate::aoc2023::day01::solver::Day01;
use crate::aoc2023::day02::solver::Day02;
use crate::aoc2023::day03::solver::Day03;
use crate::aoc2023::day04::solver::Day04;
use crate::aoc2023::day05::solver::Day05;
use crate::aoc2023::day06::solver::Day06;
use crate::aoc2023::day07::solver::Day07;
use crate::aoc2023::day08::solver::Day08;
use crate::aoc2023::day09::solver::Day09;
use crate::aoc2023::day10::solver::Day10;
use crate::aoc2023::day11::solver::Day11;
use crate::aoc2023::day12::solver::Day12;
use crate::aoc2023::day13::solver::Day13;
use crate::aoc2023::day14::solver::Day14;
use crate::aoc2023::day15::solver::Day15;
use crate::aoc2023::day16::solver::Day16;
use crate::aoc2023::day17::solver::Day17;
use crate::aoc2023::day18::solver::Day18;
use crate::aoc2023::day19::solver::Day19;
use crate::aoc2023::day20::solver::Day20;
use crate::aoc2023::day21::solver::Day21;
use crate::aoc2023::day22::solver::Day22;
use crate::aoc2023::day23::solver::Day23;
use crate::aoc2023::day24::solver::Day24;
use crate::aoc2023::day25::solver::Day25;
use crate::aoc2023::solver::interface::Solver;
use clap::Parser;

fn main() {
    let args = cmdline::Args::parse();

    let input = match fs::read_to_string(args.input_file) {
        Ok(contents) => contents,
        Err(err) => {
            println!("{err}");
            return;
        }
    };

    let mut solver: Box<dyn Solver> = match args.day {
        0 => Box::new(Day00::new()),
        1 => Box::new(Day01::new()),
        2 => Box::new(Day02::new()),
        3 => Box::new(Day03::new()),
        4 => Box::new(Day04::new()),
        5 => Box::new(Day05::new()),
        6 => Box::new(Day06::new()),
        7 => Box::new(Day07::new()),
        8 => Box::new(Day08::new()),
        9 => Box::new(Day09::new()),
        10 => Box::new(Day10::new()),
        11 => Box::new(Day11::new()),
        12 => Box::new(Day12::new()),
        13 => Box::new(Day13::new()),
        14 => Box::new(Day14::new()),
        15 => Box::new(Day15::new()),
        16 => Box::new(Day16::new()),
        17 => Box::new(Day17::new()),
        18 => Box::new(Day18::new()),
        19 => Box::new(Day19::new()),
        20 => Box::new(Day20::new()),
        21 => Box::new(Day21::new()),
        22 => Box::new(Day22::new()),
        23 => Box::new(Day23::new()),
        24 => Box::new(Day24::new()),
        25 => Box::new(Day25::new()),
        _ => {
            println!("Error: {} is not a valid day!", args.day);
            return;
        }
    };

    let start = Instant::now();

    let result = solver.consume_input(&input);
    match result {
        Ok(_) => (),
        Err(err) => {
            println!("{}", err);
            return;
        }
    }

    let result = match args.part {
        cmdline::Part::One => solver.solve_part_1(),
        cmdline::Part::Two => solver.solve_part_2(),
    };

    let elapsed = start.elapsed();

    match result {
        Ok(solution) => {
            println!("{solution}");
            println!("Computation time: {}ms", elapsed.as_millis());
        },
        Err(err) => {
            println!("{err}");
            return;
        }
    }
}
