use aoc2019::*;

#[macro_use] extern crate log;

fn main() {
    env_logger::init();
    let mut input = get_input();

//    let input5 = &mut [3, 0, 4, 0,  99];
    compute(&mut input);
}

