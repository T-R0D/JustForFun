use std::{cell::RefCell, collections::BTreeMap, rc::Rc};

use crate::aoc2022::solver::interface::{AoCResult, Solver};

pub struct Day07 {
    io_lines: Vec<IOLine>,
}

impl Day07 {
    pub fn new() -> Self {
        Day07 { io_lines: vec![] }
    }
}

impl Solver for Day07 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let lines = input.trim().split("\n").map(|s| String::from(s));

        self.io_lines = lines
            .map(|line| {
                if line.starts_with("$") {
                    let parts = line.split(" ").collect::<Vec<_>>();
                    return match parts[1] {
                        "ls" => IOLine::LSCommand,
                        "cd" => IOLine::CDCommand {
                            arg: match parts[2] {
                                "/" => CDArg::Root,
                                ".." => CDArg::UpOne,
                                _ => CDArg::Into(parts[2].to_string()),
                            },
                        },
                        _ => panic!("Unknown command: {}", parts[1]),
                    };
                }

                let parts = line.split(" ").collect::<Vec<_>>();
                let name = String::from(parts[1]);
                if parts[0] == "dir" {
                    IOLine::FSListing(FSObj::Dir { name })
                } else {
                    let size = parts[0].parse::<usize>().expect("should be an integer");
                    IOLine::FSListing(FSObj::File { name, size })
                }
            })
            .collect::<Vec<_>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let root = fs_tree_from_io_lines(&self.io_lines);

        let (_, limited_sum_size) = compute_dir_sizes(&root);

        Ok(limited_sum_size.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let root = fs_tree_from_io_lines(&self.io_lines);

        let (total_size, _) = compute_dir_sizes(&root);

        let free_space = 70_000_000 - total_size;
        let target_delete_size = 30_000_000 - free_space;

        let (_, optimal_delete_size) =
            find_optimal_deleted_directory_size(&root, target_delete_size);

        Ok(optimal_delete_size.to_string())
    }
}

enum IOLine {
    CDCommand { arg: CDArg },
    LSCommand,
    FSListing(FSObj),
}

enum CDArg {
    Root,
    UpOne,
    Into(String),
}

enum FSObj {
    Dir { name: String },
    File { name: String, size: usize },
}

struct Dir {
    #[allow(dead_code)]
    name: String,
    dirs: RefCell<BTreeMap<String, Rc<Dir>>>,
    files: RefCell<Vec<Rc<File>>>,
}

impl Dir {
    fn new(name: &String) -> Self {
        Dir {
            name: name.to_owned(),
            dirs: RefCell::new(BTreeMap::new()),
            files: RefCell::new(vec![]),
        }
    }
}

struct File {
    #[allow(dead_code)]
    name: String,
    size: usize,
}

impl File {
    fn new(name: &String, size: usize) -> Self {
        File {
            name: name.to_owned(),
            size,
        }
    }
}

fn fs_tree_from_io_lines(io_lines: &Vec<IOLine>) -> Rc<Dir> {
    let root = Rc::new(Dir::new(&String::from("/")));

    let mut cwd = root.clone();
    let mut path: Vec<Rc<Dir>> = Vec::new();

    for line in io_lines.iter() {
        match line {
            IOLine::CDCommand { arg } => match arg {
                CDArg::Root => {
                    cwd = root.clone();
                }
                CDArg::UpOne => {
                    cwd = path.pop().expect("path should not be empty");
                }
                CDArg::Into(child_dir_name) => {
                    path.push(cwd.clone());

                    let strong_cwd = cwd;
                    let dirs = strong_cwd.dirs.borrow();
                    let subdir = dirs.get(child_dir_name).expect("");

                    cwd = subdir.clone();
                }
            },
            IOLine::LSCommand => {
                if cwd.dirs.borrow().len() > 0 || cwd.files.borrow().len() > 0 {
                    panic!("Revisiting a dir!")
                }
            }
            IOLine::FSListing(obj) => match obj {
                FSObj::Dir { name } => {
                    let new_dir = Rc::new(Dir::new(name));
                    cwd.dirs.borrow_mut().insert(name.to_owned(), new_dir);
                }
                FSObj::File { name, size } => {
                    let new_file = Rc::new(File::new(name, *size));
                    cwd.files.borrow_mut().push(new_file);
                }
            },
        }
    }
    root
}

fn compute_dir_sizes(current_dir: &Rc<Dir>) -> (usize, usize) {
    let mut current_dir_size = 0;
    let mut limited_sum_size = 0;

    for dir in current_dir.dirs.borrow().values() {
        let (subdir_size, limited_size) = compute_dir_sizes(&dir);
        current_dir_size += subdir_size;
        limited_sum_size += limited_size;
    }

    for file in current_dir.files.borrow().iter() {
        current_dir_size += file.size;
    }
    if current_dir_size <= 100_000 {
        limited_sum_size += current_dir_size
    }

    (current_dir_size, limited_sum_size)
}

fn find_optimal_deleted_directory_size(current_dir: &Rc<Dir>, target: usize) -> (usize, usize) {
    let mut current_dir_size = 0;
    let mut optimal_delete_size = usize::MAX;

    for dir in current_dir.dirs.borrow().values() {
        let (subdir_size, candidate_delete_size) = find_optimal_deleted_directory_size(dir, target);
        current_dir_size += subdir_size;
        if target <= candidate_delete_size && candidate_delete_size < optimal_delete_size {
            optimal_delete_size = candidate_delete_size;
        }
    }

    for file in current_dir.files.borrow().iter() {
        current_dir_size += file.size;
    }
    if target <= current_dir_size && current_dir_size < optimal_delete_size {
        optimal_delete_size = current_dir_size;
    }

    (current_dir_size, optimal_delete_size)
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            $ cd /
            $ ls
            dir a
            14848514 b.txt
            8504156 c.dat
            dir d
            $ cd a
            $ ls
            dir e
            29116 f
            2557 g
            62596 h.lst
            $ cd e
            $ ls
            584 i
            $ cd ..
            $ cd ..
            $ cd d
            $ ls
            4060174 j
            8033020 d.log
            5626152 d.ext
            7214296 k
        "});
        let mut solver = Day07::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("95437", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            $ cd /
            $ ls
            dir a
            14848514 b.txt
            8504156 c.dat
            dir d
            $ cd a
            $ ls
            dir e
            29116 f
            2557 g
            62596 h.lst
            $ cd e
            $ ls
            584 i
            $ cd ..
            $ cd ..
            $ cd d
            $ ls
            4060174 j
            8033020 d.log
            5626152 d.ext
            7214296 k
        "});
        let mut solver = Day07::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("24933642", result);

        Ok(())
    }
}
