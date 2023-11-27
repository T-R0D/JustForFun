use std::{fs, path::Path, path::PathBuf};

use clap::Parser;
use indoc::formatdoc;

#[derive(Parser)]
#[command(about = "Generate Advent of Code 2023 solution templates", long_about = None)]
struct Args {
    /// Path to the AoC project folder (not the actual `src` directory).
    #[arg(long)]
    pub working_dir: PathBuf,
}

fn main() {
    let args = Args::parse();

    println!("Generating files in project: {}...", args.working_dir.display());

    let src_dir =  args.working_dir.join(format!("src/aoc{}", 2023));

    for day in 0..=25 {
        println!("Generating template for day {day}...");
        match gen_template(src_dir.as_path(), day) {
            Ok(_) => (),
            Err(msg) => {
                println!("Error: {msg}\n");
                return;
            }
        }
    }


}

fn gen_template(working_dir: &Path, day: i32) -> Result<(), String> {
    let day_num_str = &format!("{day:0>2}", day = day);

    let folder_path = working_dir.join(format!("day{day_num_str}"));
    match fs::create_dir_all(folder_path.as_path()) {
        Ok(_) => (),
        Err(_) => {
            return Err(format!(
                "Failed to create directory: {path}",
                path = folder_path.display()
            ))
        }
    }

    gen_mod_file(folder_path.as_path())?;

    gen_solution_file(folder_path.as_path(), day_num_str)?;

    Ok(())
}

fn gen_mod_file(folder_path: &Path) -> Result<(), String> {
    let file_path = folder_path.join("mod.rs");

    match fs::write(file_path, "pub mod solver;\n") {
        Ok(_) => Ok(()),
        Err(err) => Err(format!("Unable to write mod file: {err}")),
    }
}

fn gen_solution_file(folder_path: &Path, day_num: &String) -> Result<(), String> {
    let template_contents = formatdoc! {r##"
        use crate::aoc2023::solver::interface::{{AoCResult, Solver}};

        pub struct Day{day_num} {{}}

        impl Day{day_num} {{
            pub fn new() -> Self {{
                Day{day_num} {{}}
            }}
        }}

        impl Solver for Day{day_num} {{
            fn consume_input(&mut self, _input: &String) -> AoCResult {{
                Ok(String::from(""))
            }}

            fn solve_part_1(&self) -> AoCResult {{
                Err(String::from("Not implemented."))
            }}

            fn solve_part_2(&self) -> AoCResult {{
                Err(String::from("Not implemented."))
            }}
        }}

        #[cfg(test)]
        mod tests {{
            use super::*;
            #[cfg(test)]
            use pretty_assertions::assert_eq;

            #[test]
            fn part_one_solves_correctly() -> Result<(), String> {{
                // Arrange.
                let input = &String::from("");
                let mut solver = Day{day_num}::new();
                solver.consume_input(input)?;

                // Act.
                let result = solver.solve_part_1()?;

                // Assert.
                assert_eq!("", result);

                Ok(())
            }}

            #[test]
            fn part_two_solves_correctly() -> Result<(), String> {{
                // Arrange.
                let input = &String::from("");
                let mut solver = Day{day_num}::new();
                solver.consume_input(input)?;

                // Act.
                let result = solver.solve_part_2()?;

                // Assert.
                assert_eq!("", result);

                Ok(())
            }}
        }}
    "##};

    let file_path = folder_path.join("solver.rs");
    match fs::write(file_path.as_path(), template_contents) {
        Ok(_) => (),
        Err(err) => {
            return Err(format!(
                "Failed to write file contents for {}: {}",
                file_path.display(),
                err
            ));
        }
    }

    Ok(())
}
