use std::collections::{HashMap, HashSet};

use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day25 {
    component_graph: HashMap<String, Vec<String>>,
}

impl Day25 {
    pub fn new() -> Self {
        Day25 {
            component_graph: HashMap::new(),
        }
    }
}

impl Solver for Day25 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.component_graph = input
            .trim()
            .lines()
            .map(|line| {
                let parts = line.split(": ").collect::<Vec<_>>();
                let key = String::from(parts[0]);
                let vals = parts[1].split(" ").map(String::from).collect::<Vec<_>>();
                (key, vals)
            })
            .collect::<HashMap<String, Vec<String>>>();

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let mut bisector = GraphBisector::new(&self.component_graph, TARGET_CUTS);
        let (partition_a_size, partition_b_size) = bisector.find_bisection_with_target_number_of_bridges();

        let size_score = partition_a_size * partition_b_size;

        Ok(size_score.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        Ok(String::from("Merry Christmas!"))
    }
}

const TARGET_CUTS: usize = 3;

struct GraphBisector {
    g: HashMap<String, HashSet<String>>,
    s: HashSet<String>,
    target: usize,
}

impl GraphBisector {
    fn new(graph: &HashMap<String, Vec<String>>, target_n_bridges: usize) -> Self {
        let mut g = HashMap::<String, HashSet<String>>::new();
        for (src, dsts) in graph.iter() {
            for dst in dsts.iter() {
                g.entry(src.clone())
                    .or_insert(HashSet::new())
                    .insert(dst.clone());
                g.entry(dst.clone())
                    .or_insert(HashSet::new())
                    .insert(src.clone());
            }
        }

        let s = g.keys().map(|s| s.clone()).collect::<HashSet<String>>();

        Self {
            g,
            s,
            target: target_n_bridges,
        }
    }

    fn find_bisection_with_target_number_of_bridges(&mut self) -> (usize, usize) {
        // Shamefully stolen algorithm:
        // https://www.reddit.com/r/adventofcode/comments/18qbsxs/comment/ketzp94/?utm_source=share&utm_medium=web2x&context=3
        //
        // The idea here is to arrive at 2 sets where there are the target number of "bridge crossings"
        // from a set S to a set T (where T is all the vertices of G not in S). Start with every vertex
        // in the graph in S. Remove the vertices from S one by one until the target number of bridges remain
        // (the number of crossing will start at 0). Remove the vertex that would have the most bridge
        // crossings from S to T (or any if there is a tie).
        //
        // I believe this works because our target number is relatively small (3 of many), and wonder if
        // things would be troublesome if we had to target a relatively large number of crossings?

        let mut total_external_neighbor_count = self
            .s
            .iter()
            .map(|label| self.count_neighbors_in_t_but_not_s(label))
            .sum::<usize>();
        while total_external_neighbor_count != self.target {
            let mut label_of_vertex_with_most_crossings = String::new();
            let mut max_crossings = 0;
            for v in self.s.iter() {
                let crossing_count = self.count_neighbors_in_t_but_not_s(v);
                if crossing_count >= max_crossings {
                    label_of_vertex_with_most_crossings = v.clone();
                    max_crossings = crossing_count;
                }
            }

            self.s.remove(&label_of_vertex_with_most_crossings);

            total_external_neighbor_count = self
                .s
                .iter()
                .map(|label| self.count_neighbors_in_t_but_not_s(label))
                .sum::<usize>();
        }

        (
            self.g
                .keys()
                .map(|label| label.clone())
                .collect::<HashSet<String>>()
                .difference(&self.s)
                .count(),
            self.s.len(),
        )
    }

    fn count_neighbors_in_t_but_not_s(&self, label: &String) -> usize {
        self.g.get(label).unwrap().difference(&self.s).count()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[cfg(test)]
    use indoc::indoc;
    #[cfg(test)]
    use pretty_assertions::assert_eq;

    #[test]
    fn part_one_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            jqt: rhn xhk nvd
            rsh: frs pzl lsr
            xhk: hfx
            cmg: qnr nvd lhk bvb
            rhn: xhk bvb hfx
            bvb: xhk hfx
            pzl: lsr hfx nvd
            qnr: nvd
            ntq: jqt hfx bvb xhk
            nvd: lhk
            lsr: lhk
            rzs: qnr cmg lsr rsh
            frs: qnr lhk lsr
        "});
        let mut solver = Day25::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("54", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from("");
        let mut solver = Day25::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("Merry Christmas!", result);

        Ok(())
    }
}
