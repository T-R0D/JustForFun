pub type AoCResult = Result<String, String>;

pub trait Solver {
    fn consume_input(&mut self, input: &String) -> AoCResult;
    fn solve_part_1(&self) -> AoCResult;
    fn solve_part_2(&self) -> AoCResult;
}
