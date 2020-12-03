use std::{fs, io};
use std::io::Write;

use log::debug;
use std::collections::HashMap;

pub fn get_input() -> Vec<i32> {
    let data = fs::read_to_string("../input").expect("unable to open file");
    data.trim().split(",")
        .map(|l| l.parse::<i32>().expect(&format!("failed to parse input {}", l)))
        .collect()
}

pub fn solve() {
    let content = fs::read_to_string("../input").unwrap();

    let orbits: HashMap<_, _> = content
        .lines()
        .map(|line| line.split(')').collect::<Vec<&str>>())
        .map(|pair| (pair[1], pair[0]))
        .collect();

    println!("orbits:{:?}", orbits); // TODO: Debugging
}


#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_solve() {
        solve()
    }

    #[test]
    fn test_get_input() {
        let input = get_input();
        println!("{:?}", input)
    }
}

