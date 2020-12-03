use aoc2019::*;
use std::io::stdin;

fn main() {
    env_logger::init();
    let mut input = get_input();

    compute(&mut input, &mut stdin().lock());
}

