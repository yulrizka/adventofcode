use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn get_input() -> Vec<i64> {
    let f = File::open("../input").expect("unable to open file");
    let f = BufReader::new(f);
    f.lines()
        .map(|l| l.expect("could not read line"))
        .map(|l| l.parse::<i64>().expect("could not parse number"))
        .collect()
}

pub fn solve(masses: &[i64]) -> i64 {
    masses.iter()
        .map(|&m| calc_fuel(m))
        .sum()
}

pub fn calc_fuel(mass: i64) -> i64 {
    mass / 3 - 2
}

#[cfg(test)]
mod tests {
    use crate::{calc_fuel};

    #[test]
    fn test_calc_fuel() {
        assert_eq!(calc_fuel(12), 2);
        assert_eq!(calc_fuel(14), 2);
        assert_eq!(calc_fuel(1969), 654);
        assert_eq!(calc_fuel(100756), 33583);
    }
}

