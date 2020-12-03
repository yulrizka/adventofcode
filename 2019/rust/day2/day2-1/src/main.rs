use aoc2019::*;

fn main() {
    let mut input = get_input();

    // replace position 1 with the value of 12
    input[1] = 12;

    // replace position 2 with the value of 2
    input[2] = 2;

    let result = compute(input.as_mut());

    println!("answer: {}", result[0])
}

