use aoc2019::*;

fn main() {
    let input = get_input();

    let (distance, timing) = compute(&input.0, &input.1);
    println!("distance:{:?} timing:{:?}", distance, timing);
}

