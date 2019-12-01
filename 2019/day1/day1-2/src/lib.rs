use std::fs::File;
use std::io::{BufRead, BufReader};
use std::iter;

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
    // https://doc.rust-lang.org/std/iter/fn.successors.html
    // Creates a new iterator where each successive item is computed based on the preceding one.
    let i  = iter::successors(Some(mass), |fuel| {
        let current = fuel / 3 - 2;
        if current > 0 {
            Some(current)
        } else {
            None
        }
    });

    // for debugging
    // println!("{:?}", i.collect::<Vec<_>>());
    // 0

    // skip the 1st element (original mass)
    i.skip(1).sum()
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

