use aoc2019::*;

fn main() {
    let input = get_input();

    for noun in 0..99 {
        for verb in 0..99 {
            let mut code = input.to_vec();
            code[1] = noun;
            code[2] = verb;

            compute(&mut code);
            let answer = code[0];
            if answer == 19690720 {
                let result = 100 * noun + verb;
                println!("answer:{} noun:{} verb:{}", result, noun, verb);
                return
            }
        }
    }

    panic!("should not be here")
}

