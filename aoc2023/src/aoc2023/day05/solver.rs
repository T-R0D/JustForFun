use crate::aoc2023::solver::interface::{AoCResult, Solver};

const MAPPING_SECTION_NAMES: [&str; 7] = [
    "seed-to-soil",
    "soil-to-fertilizer",
    "fertilizer-to-water",
    "water-to-light",
    "light-to-temperature",
    "temperature-to-humidity",
    "humidity-to-location",
];

pub struct Day05 {
    seeds: Vec<isize>,
    resource_maps: Vec<ResourceMap>,
}

impl Day05 {
    pub fn new() -> Self {
        Day05 {
            seeds: vec![],
            resource_maps: vec![],
        }
    }
}

impl Solver for Day05 {
    fn consume_input(&mut self, input: &String) -> AoCResult {
        let sections = input.split("\n\n").collect::<Vec<_>>();

        self.seeds = self.parse_seed_list(sections[0])?;

        self.resource_maps = sections
            .iter()
            .skip(1)
            .zip(MAPPING_SECTION_NAMES.iter())
            .map(|(section, expected_name)| self.parse_map(expected_name, section))
            .collect::<Result<Vec<_>, String>>()?;

        Ok(String::from(""))
    }

    fn solve_part_1(&self) -> AoCResult {
        let min_location_id = self
            .seeds
            .iter()
            .map(|seed_id| self.map_seed_to_location(*seed_id))
            .fold(isize::MAX, isize::min);

        Ok(min_location_id.to_string())
    }

    fn solve_part_2(&self) -> AoCResult {
        let mut seed_ranges = self.seed_list_to_seed_ranges();
        seed_ranges.sort_by(|a, b| b.0.cmp(&a.0));

        let location_ranges = self.map_seed_ranges_to_location_ranges(seed_ranges);

        match location_ranges.first() {
            Some(lowest_range) => Ok(lowest_range.0.to_string()),
            _ => Err(format!("No locations were mapped to"))
        }
    }
}

impl Day05 {
    fn parse_seed_list(&self, seed_list: &str) -> Result<Vec<isize>, String> {
        let parts = seed_list.split("seeds: ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(String::from("Seeds section was invalid"));
        }

        let list = parts[1]
            .split(" ")
            .map(|item| match item.parse::<isize>() {
                Ok(seed_id) => Ok(seed_id),
                Err(_) => Err(format!("Could not parse seed ID: {item}")),
            })
            .collect::<Result<Vec<_>, String>>()?;

        Ok(list)
    }

    fn parse_map(&self, expected_name: &str, section: &str) -> Result<ResourceMap, String> {
        let lines = section.lines().collect::<Vec<_>>();
        if lines.len() < 2 {
            return Err(format!("{expected_name} section is malformed (lines)"));
        }

        if !lines[0].contains(expected_name) {
            return Err(format!("{expected_name} section not found"));
        }

        let mut resource_map = ResourceMap::new();

        for (i, range_line) in lines.iter().skip(1).enumerate() {
            let numbers = range_line
                .split(" ")
                .map(|item| match item.parse::<isize>() {
                    Ok(value) => Ok(value),
                    Err(_) => Err(format!("{item} on line {i} of {expected_name} is invalid")),
                })
                .collect::<Result<Vec<isize>, String>>()?;

            if numbers.len() != 3 {
                return Err(format!("line {i} of {expected_name} was malformed"));
            }

            resource_map.add_range(numbers[0], numbers[1], numbers[2]);
        }

        Ok(resource_map)
    }

    fn seed_list_to_seed_ranges(&self) -> Vec<(isize, isize)> {
        self.seeds
            .iter()
            .step_by(2)
            .zip(self.seeds.iter().skip(1).step_by(2))
            .map(|(start, span)| (*start, *start + *span))
            .collect::<Vec<(isize, isize)>>()
    }

    fn map_seed_to_location(&self, seed_id: isize) -> isize {
        let mut dst_id = seed_id;
        for resource_map in self.resource_maps.iter() {
            dst_id = resource_map.get_resource_id(dst_id);
        }
        dst_id
    }

    fn map_seed_ranges_to_location_ranges(&self, seed_ranges: Vec<(isize, isize)>) -> Vec<(isize, isize)> {
        let mut dst_ranges = seed_ranges;
        for resource_map in self.resource_maps.iter() {
            dst_ranges = resource_map.get_resource_ranges(dst_ranges);
        }
        dst_ranges
    }
}

struct MappingRange {
    dst_start: isize,
    src_start: isize,
    src_end: isize,
    range: isize,
    shift_magnitude: isize,
}

impl MappingRange {
    pub fn new(dst_start: isize, src_start: isize, range: isize) -> Self {
        MappingRange {
            dst_start,
            src_start,
            src_end: src_start + range,
            range,
            shift_magnitude: src_start - dst_start,
        }
    }

    pub fn find_resource_id(&self, src_id: isize) -> Option<isize> {
        if self.src_start <= src_id && src_id < self.src_start + self.range {
            let offset = src_id - self.src_start;
            return Some(self.dst_start + offset);
        }

        None
    }
}

struct ResourceMap {
    mapping_ranges: Vec<MappingRange>,
}

impl ResourceMap {
    pub fn new() -> Self {
        ResourceMap {
            mapping_ranges: vec![],
        }
    }

    pub fn add_range(&mut self, dst_start: isize, src_start: isize, range: isize) {
        self.mapping_ranges
            .push(MappingRange::new(dst_start, src_start, range));
        self.mapping_ranges.sort_by_key(|range| range.src_start);
    }

    pub fn get_resource_id(&self, src_id: isize) -> isize {
        for mapping_range in self.mapping_ranges.iter() {
            if let Some(dst_id) = mapping_range.find_resource_id(src_id) {
                return dst_id;
            }
        }

        src_id
    }

    fn get_resource_ranges(&self, src_ranges: Vec<(isize, isize)>) -> Vec<(isize, isize)> {
        let mut dst_ranges = Vec::<(isize, isize)>::new();

        'src_processing: for (src_start, src_end) in src_ranges.iter() {
            match (self.mapping_ranges.first(), self.mapping_ranges.last()) {
                (Some(first_range), Some(last_range)) => {
                    if *src_end <= first_range.src_start || last_range.src_end <= *src_start {
                        dst_ranges.push((*src_start, *src_end));
                        continue 'src_processing;
                    }
                }
                _ => (),
            }

            let mut current_start = *src_start;

            'dst_processing: for mapping_range in self.mapping_ranges.iter() {
                if *src_end <= mapping_range.src_start {
                    break 'dst_processing;
                }

                if mapping_range.src_end <= current_start {
                    continue 'dst_processing;
                }

                if current_start < mapping_range.src_start {
                    dst_ranges.push((current_start, mapping_range.src_start));
                    current_start = mapping_range.src_start;
                }

                let end = isize::min(mapping_range.src_end, *src_end);
                dst_ranges.push((
                    current_start - mapping_range.shift_magnitude,
                    end - mapping_range.shift_magnitude,
                ));
                current_start = end;
            }

            if current_start < *src_end {
                dst_ranges.push((current_start, *src_end));
            }
        }

        dst_ranges.sort_by_key(|range| range.0);

        self.merge_ranges(dst_ranges)
    }

    fn merge_ranges(&self, ranges: Vec<(isize, isize)>) -> Vec<(isize, isize)> {
        if ranges.len() <= 1 {
            return ranges;
        }

        let mut merged_ranges = Vec::<(isize, isize)>::new();
        let (mut current_start, mut current_end) = *ranges.first().unwrap();
        for &(next_start, next_end) in ranges.iter().skip(1) {
            if current_end <= next_start {
                merged_ranges.push((current_start, current_end));
                (current_start, current_end) = (next_start, next_end);
                continue;
            }

            current_end = isize::min(current_end, next_end);
        }
        merged_ranges.push((current_start, current_end));

        merged_ranges
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
            seeds: 79 14 55 13

            seed-to-soil map:
            50 98 2
            52 50 48

            soil-to-fertilizer map:
            0 15 37
            37 52 2
            39 0 15

            fertilizer-to-water map:
            49 53 8
            0 11 42
            42 0 7
            57 7 4

            water-to-light map:
            88 18 7
            18 25 70

            light-to-temperature map:
            45 77 23
            81 45 19
            68 64 13

            temperature-to-humidity map:
            0 69 1
            1 0 69

            humidity-to-location map:
            60 56 37
            56 93 4
        "});
        let mut solver = Day05::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_1()?;

        // Assert.
        assert_eq!("35", result);

        Ok(())
    }

    #[test]
    fn part_two_solves_correctly() -> Result<(), String> {
        // Arrange.
        let input = &String::from(indoc! {"
            seeds: 79 14 55 13

            seed-to-soil map:
            50 98 2
            52 50 48

            soil-to-fertilizer map:
            0 15 37
            37 52 2
            39 0 15

            fertilizer-to-water map:
            49 53 8
            0 11 42
            42 0 7
            57 7 4

            water-to-light map:
            88 18 7
            18 25 70

            light-to-temperature map:
            45 77 23
            81 45 19
            68 64 13

            temperature-to-humidity map:
            0 69 1
            1 0 69

            humidity-to-location map:
            60 56 37
            56 93 4
        "});
        let mut solver = Day05::new();
        solver.consume_input(input)?;

        // Act.
        let result = solver.solve_part_2()?;

        // Assert.
        assert_eq!("46", result);

        Ok(())
    }
}
