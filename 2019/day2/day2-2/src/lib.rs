use std::fs;
use crate::OpCode::{Add, Multiply, Halt};

pub fn get_input() -> Vec<usize> {
    let data = fs::read_to_string("../input").expect("unable to open file");
    data.trim().split(",")
        .map(|l| l.parse::<usize>().expect(&format!("failed to parse input {}", l)))
        .collect()
}

#[derive(PartialEq)]
enum OpCode {
    Add, Multiply, Halt
}

impl From<usize> for OpCode {
    fn from(item: usize) -> Self {
        match item {
            1 => Add,
            2 => Multiply,
            99 => Halt,
            _ => unreachable!()
        }
    }
}

pub fn compute(code: &mut [usize]) {
    let mut sp = 0; // instruction pointer

    while sp < code.len() {
        let opcode = code[sp].into();

        match opcode {
            OpCode::Add => {
                let p1 = code[sp+1];
                let p2 = code[sp+2];
                let p3 = code[sp+3];
                code[p3] = code[p1] + code[p2];
                sp += 4
            },
            OpCode::Multiply => {
                let p1 = code[sp+1];
                let p2 = code[sp+2];
                let p3 = code[sp+3];
                code[p3] = code[p1] * code[p2];
                sp += 4;
            }
            OpCode::Halt => {
                return
            }
        }
    }

    unreachable!()
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_solve() {
        let input = &mut [1, 0, 0, 0, 99];
        compute(input);
        assert_eq!(input.to_vec(), [2,0,0,0,99]);

        let input2 = &mut [2, 3, 0, 3, 99];
        compute(input2);
        assert_eq!(input2.to_vec(), [2,3,0,6,99]);

        let input3 = &mut [2, 4, 4, 5, 99, 0];
        compute(input3);
        assert_eq!(input3.to_vec(), [2,4,4,5,99,9801]);

        let input4 = &mut [1, 1, 1, 4, 99, 5, 6, 0, 99];
        compute(input4);
        assert_eq!(input4.to_vec(), [30,1,1,4,2,5,6,0,99]);
    }

    #[test]
    fn test_get_input() {
        let input = get_input();
        println!("{:?}", input)
    }
}

