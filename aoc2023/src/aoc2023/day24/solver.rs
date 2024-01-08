use crate::aoc2023::solver::interface::{AoCResult, Solver};

pub struct Day24 {
    snapshots: Vec<HailStoneSnapshot>,
}

impl Day24 {
    pub fn new() -> Self {
        Day24 { snapshots: vec![] }
    }
}

impl Solver for Day24 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        self.snapshots = input
            .trim()
            .lines()
            .map(|line| {
                let parts = line.split(" @ ").collect::<Vec<_>>();
                if parts.len() != 2 {
                    return Err(String::from("line didn't have 2 parts"));
                }

                let position = Position::try_from_line(parts[0])?;
                let velocity = Velocity::try_from_line(parts[1])?;

                Ok(HailStoneSnapshot(position, velocity))
            })
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let crossovers = count_path_crossings(
            &self.snapshots,
            LOWER_TEST_WINDOW_BOUND,
            UPPER_TEST_WINDOW_BOUND,
        );

        Ok(crossovers.to_string())
    }

    // 1004774995964533 too low
    // 1004774995964534 is right, there are rounding errors somewhere. I got the right answer using a different subset of data.
    fn solve_part_2(&self) -> AoCResult {
        let HailStoneSnapshot(start_position, ..) =
            find_start_for_perfect_collisions(&self.snapshots);

        let coordinate_sum = start_position.x + start_position.y + start_position.z;

        Ok(coordinate_sum.to_string())
    }
}

const LOWER_TEST_WINDOW_BOUND: i64 = 200_000_000_000_000;
const UPPER_TEST_WINDOW_BOUND: i64 = 400_000_000_000_000;

fn count_path_crossings(
    snapshots: &Vec<HailStoneSnapshot>,
    lower_test_window_bound: i64,
    upper_test_window_bound: i64,
) -> usize {
    let mut crossovers = 0_usize;
    for i in 0..snapshots.len() {
        let a = &snapshots[i];
        for j in (i + 1)..snapshots.len() {
            let b = &snapshots[j];

            if let Some(((x, y), t, u)) = find_x_y_intersection(a, b) {
                if t >= 0
                    && u >= 0
                    && (lower_test_window_bound <= x && x <= upper_test_window_bound)
                    && (lower_test_window_bound <= y && y <= upper_test_window_bound)
                {
                    crossovers += 1;
                }
            }
        }
    }

    crossovers
}

fn find_x_y_intersection(
    snapshot_0: &HailStoneSnapshot,
    snapshot_1: &HailStoneSnapshot,
) -> Option<((i64, i64), i64, i64)> {
    // https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line_segment
    // This works, but beware rounding errors.

    let HailStoneSnapshot(p_0, v_0) = snapshot_0;
    let HailStoneSnapshot(p_1, v_1) = snapshot_1;

    let epsilon = 1e-5_f64;

    let (x_0, x_1) = (p_0.x as f64, (p_0.x + v_0.x) as f64);
    let (y_0, y_1) = (p_0.y as f64, (p_0.y + v_0.y) as f64);
    let (x_2, x_3) = (p_1.x as f64, (p_1.x + v_1.x) as f64);
    let (y_2, y_3) = (p_1.y as f64, (p_1.y + v_1.y) as f64);

    let t_num = (x_0 - x_2) * (y_2 - y_3) - (y_0 - y_2) * (x_2 - x_3);
    let t_den = (x_0 - x_1) * (y_2 - y_3) - (y_0 - y_1) * (x_2 - x_3);
    if t_den.abs() <= epsilon {
        return None;
    }
    let t = t_num / t_den;

    let u_num = (x_0 - x_2) * (y_0 - y_1) - (y_0 - y_2) * (x_0 - x_1);
    let u_den = (x_0 - x_1) * (y_2 - y_3) - (y_0 - y_1) * (x_2 - x_3);
    if u_den.abs() == epsilon {
        return None;
    }
    let u = u_num / u_den;

    let x = (x_0 as f64 + t * v_0.x as f64) as i64;
    let y = (y_0 as f64 + t * v_0.y as f64) as i64;

    Some(((x as i64, y as i64), t as i64, u as i64))
}

fn find_start_for_perfect_collisions(snapshots: &Vec<HailStoneSnapshot>) -> HailStoneSnapshot {
    // Shamefully ripped off algorithm from
    // https://github.com/ash42/adventofcode/blob/main/adventofcode2023/src/nl/michielgraat/adventofcode2023/day24/Day24.java

    let mut coefficients = vec![vec![0_f64; 4]; 4];
    let mut rhs = vec![0_f64; 4];
    for i in 0..4 {
        // These silly offsets of `+1` seem to trump the rounding errors I got previously. ¯\_(ツ)_/¯
        let (HailStoneSnapshot(p_0, v_0), HailStoneSnapshot(p_1, v_1)) =
            (&snapshots[i + 1], &snapshots[i + 1 + 1]);

        coefficients[i][0] = (v_1.y - v_0.y) as f64;
        coefficients[i][1] = (v_0.x - v_1.x) as f64;
        coefficients[i][2] = (p_0.y - p_1.y) as f64;
        coefficients[i][3] = (p_1.x - p_0.x) as f64;

        rhs[i] = (-p_0.x * v_0.y + p_0.y * v_0.x + p_1.x * v_1.y - p_1.y * v_1.x) as f64;
    }

    gaussian_elimination(&mut coefficients, &mut rhs);

    let x = rhs[0].round() as i64;
    let y = rhs[1].round() as i64;
    let vx = rhs[2].round() as i64;
    let vy = rhs[3].round() as i64;

    let mut coefficients_2 = vec![vec![0_f64; 2]; 2];
    let mut rhs_2 = vec![0_f64; 2];
    for i in 0..2 {
        let (HailStoneSnapshot(p_0, v_0), HailStoneSnapshot(p_1, v_1)) =
            (&snapshots[i], &snapshots[i + 1]);

        coefficients_2[i][0] = (v_0.x - v_1.x) as f64;
        coefficients_2[i][1] = (p_1.x - p_0.x) as f64;

        rhs_2[i] = (-p_0.x * v_0.z + p_0.z * v_0.x + p_1.x * v_1.z
            - p_1.z * v_1.x
            - ((v_1.z - v_0.z) * x)
            - ((p_0.z - p_1.z) * vx)) as f64;
    }

    gaussian_elimination(&mut coefficients_2, &mut rhs_2);

    let z = rhs_2[0].round() as i64;
    let vz = rhs_2[1].round() as i64;

    HailStoneSnapshot(
        Position { x, y, z },
        Velocity {
            x: vx,
            y: vy,
            z: vz,
        },
    )
}

fn gaussian_elimination(coefficients: &mut Vec<Vec<f64>>, rhs: &mut Vec<f64>) -> () {
    let m = coefficients.len();

    for i in 0..m {
        let pivot = coefficients[i][i];

        for j in 0..m {
            coefficients[i][j] /= pivot;
        }
        rhs[i] /= pivot;

        for k in 0..m {
            if k == i {
                continue;
            }

            let factor = coefficients[k][i];
            for j in 0..m {
                coefficients[k][j] -= factor * coefficients[i][j];
            }
            rhs[k] -= factor * rhs[i];
        }
    }
}

struct HailStoneSnapshot(Position, Velocity);

struct Position {
    x: i64,
    y: i64,
    z: i64,
}

impl Position {
    fn try_from_line(line: &str) -> Result<Self, String> {
        let parts = line
            .split(", ")
            .map(|x| match x.parse::<i64>() {
                Ok(val) => Ok(val),
                Err(err) => Err(err.to_string()),
            })
            .collect::<Result<Vec<_>, String>>()?;

        if parts.len() != 3 {
            return Err(format!("Got {} parts, not 3", parts.len()));
        }

        Ok(Self {
            x: parts[0],
            y: parts[1],
            z: parts[2],
        })
    }
}

struct Velocity {
    x: i64,
    y: i64,
    z: i64,
}

impl Velocity {
    fn try_from_line(line: &str) -> Result<Self, String> {
        let parts = line
            .split(", ")
            .map(|x| match x.parse::<i64>() {
                Ok(val) => Ok(val),
                Err(err) => Err(err.to_string()),
            })
            .collect::<Result<Vec<_>, String>>()?;

        if parts.len() != 3 {
            return Err(format!("Got {} parts, not 3", parts.len()));
        }

        Ok(Self {
            x: parts[0],
            y: parts[1],
            z: parts[2],
        })
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
            19, 13, 30 @ -2, 1, -2
            18, 19, 22 @ -1, -1, -2
            20, 25, 34 @ -2, -2, -4
            12, 31, 28 @ -1, -2, -1
            20, 19, 15 @ 1, -5, -3
        "});
        let mut solver = Day24::new();
        solver.consume_input(input)?;

        // Act.
        let result = count_path_crossings(&solver.snapshots, 7, 27).to_string();

        // Assert.
        assert_eq!("2", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            19, 13, 30 @ -2, 1, -2
            18, 19, 22 @ -1, -1, -2
            20, 25, 34 @ -2, -2, -4
            12, 31, 28 @ -1, -2, -1
            20, 19, 15 @ 1, -5, -3
        "});
        let mut solver = Day24::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("47", result);

        Ok(())
    }
}
