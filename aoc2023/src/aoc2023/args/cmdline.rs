use std::path::PathBuf;

use clap::{Parser, ValueEnum};

#[derive(Parser)]
#[command(about = "Advent of Code 2022 solutions!", long_about = None)]
pub struct Args {
    /// Path to the text file containing the puzzle input.
    #[arg(long)]
    pub input_file: PathBuf,

    /// Which day to solve for.
    #[arg(long, value_parser = clap::value_parser!(u8).range(0..=25))]
    pub day: u8,

    /// Which part to solve for.
    #[arg(long)]
    pub part: Part,
}

#[derive(Copy, Clone, PartialEq, Eq, PartialOrd, Ord, ValueEnum)]
pub enum Part {
    One,
    Two,
}
