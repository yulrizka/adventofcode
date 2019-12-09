use std::{fs, io};
use crate::OpCode::{Add, Multiply, Input, Output, JumpIfTrue, JumpIfFalse, Halt, LessThan, Equals};
use std::io::{Write, BufRead};

pub fn get_input() -> Vec<i32> {
    let data = fs::read_to_string("../input").expect("unable to open file");
    data.trim().split(",")
        .map(|l| l.parse::<i32>().expect(&format!("failed to parse input {}", l)))
        .collect()
}

#[derive(PartialEq, Debug)]
enum OpCode {
    Add, Multiply, Input, Output, JumpIfTrue, JumpIfFalse, LessThan, Equals, Halt
}

impl From<usize> for OpCode {
    fn from(item: usize) -> Self {
        match item {
            1 => Add,
            2 => Multiply,
            3 => Input,
            4 => Output,
            5 => JumpIfTrue,
            6 => JumpIfFalse,
            7 => LessThan,
            8 => Equals,
            99 => Halt,
            _ => unreachable!()
        }
    }
}


// get_param get parameter (1 based index) from the code on the stack pointer
// considering if the mode is immediate (return the value) or a pointer
fn get_param(code: &mut[i32], sp: usize, mode: usize,  n: usize) -> i32 {
    let p1 = code[sp+n];

    // figure out the mode by shifting the integer
    let mut m = mode;
    if n > 1 {
       m = m/(10*(n-1))
    }
    m = m % 10;

    if m == 0 {
        // reference mode
        if p1 < 0 {
            panic!("reference is negative");
        }
        let addr = p1 as usize;
        return code[addr];
    }

    // immediate mode
    return p1;
}



pub fn compute(code: &mut [i32], input: &mut dyn BufRead) {
    let mut sp = 0; // instruction pointer

    while sp < code.len() {

        let op = ((code[sp] as usize) % 100).into();
        let mode = (code[sp] / 100) as usize;

        match op {
            OpCode::Add => {
                let p1 = get_param(code, sp, mode, 1);
                let p2 = get_param(code, sp, mode, 2);
                let p3 = code[sp+3] as usize;
                code[p3] = p1 + p2;
                sp += 4
            },
            OpCode::Multiply => {
                let p1 = get_param(code, sp, mode, 1);
                let p2 = get_param(code, sp, mode, 2);
                let p3 = code[sp+3] as usize;
                code[p3] = p1 * p2;
                sp += 4;
            }
            OpCode::Input => {
                let p1 = code[sp+1];

                print!("Input  [{}]: ", p1);
                io::stdout().flush().expect("flush failed");

                let mut input_text = String::new();
                input
                    .read_line(&mut input_text)
                    .expect("failed to read from stdin");
                let val =  input_text.trim().parse::<i32>().expect("failed parsing value");

                code[p1 as usize] = val;
                sp += 2;
            },
            OpCode::Output => {
                let p1 = get_param(code, sp, mode, 1) as usize; // address
                println!("Output [{}]: {}", p1, p1);
                sp += 2;
            },
            OpCode::JumpIfTrue => {
                let p1 = get_param(code, sp, mode, 1) as usize; // address
                let p2 = get_param(code, sp, mode, 2) as usize; // address
                if p1 != 0 {
                    sp = p2;
                } else {
                    sp += 3;
                }
            },
            OpCode::JumpIfFalse => {
                let p1 = get_param(code, sp, mode, 1) as usize; // address
                let p2 = get_param(code, sp, mode, 2) as usize; // address
                if p1 == 0 {
                    sp = p2;
                } else {
                    sp += 3;
                }
            },
            OpCode::LessThan => {
                let p1 = get_param(code, sp, mode, 1) as usize; // address
                let p2 = get_param(code, sp, mode, 2) as usize; // address
                let p3 = code[sp+3] as usize;
                if p1 < p2 {
                    code[p3] = 1;
                } else {
                    code[p3] = 0;
                }
                sp += 4;
            },
            OpCode::Equals => {
                let p1 = get_param(code, sp, mode, 1) as usize; // address
                let p2 = get_param(code, sp, mode, 2) as usize; // address
                let p3 = code[sp+3] as usize;
                if p1 == p2 {
                    code[p3] = 1;
                } else {
                    code[p3] = 0;
                }
                sp += 4;
            },
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
//    use std::io::BufReader;


    #[test]
    fn test_solve() {
        env_logger::init();

        let input = &mut [1, 0, 0, 0, 99];
        compute(input, &mut io::Cursor::new(b""));
        assert_eq!(input.to_vec(), [2,0,0,0,99]);

        let input2 = &mut [2, 3, 0, 3, 99];
        compute(input2, &mut io::Cursor::new(b""));
        assert_eq!(input2.to_vec(), [2,3,0,6,99]);

        let input3 = &mut [2, 4, 4, 5, 99, 0];
        compute(input3, &mut io::Cursor::new(b""));
        assert_eq!(input3.to_vec(), [2,4,4,5,99,9801]);

        let input4 = &mut [1, 1, 1, 4, 99, 5, 6, 0, 99];
        compute(input4, &mut io::Cursor::new(b""));
        assert_eq!(input4.to_vec(), [30,1,1,4,2,5,6,0,99]);

        // testing input and output
        let input5 = &mut [3, 0, 4, 0, 99];
        compute(input5, &mut io::Cursor::new(b"8\n"));
        assert_eq!(input5.to_vec(), [8, 0, 4, 0, 99]);

        // test jump if true
        // got input 1 output (last element) should not change
        let input6 = &mut [3,3,1105,-1,9,1101,0,0,12,4,12,99,1];
        compute(input6, &mut io::Cursor::new(b"1\n"));
        assert_eq!(input6.to_vec(), [3,3,1105,1,9,1101,0,0,12,4,12,99,1]);

        // got input 0 output (last element) should be 0
        let input7 = &mut [3,3,1105,-1,9,1101,0,0,12,4,12,99,1];
        compute(input7, &mut io::Cursor::new(b"0\n"));
        assert_eq!(input7.to_vec(), [3,3,1105,0,9,1101,0,0,12,4,12,99,0]);

        // test jump with reference
        // got input 1
        let input8 = &mut [3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9];
        compute(input8, &mut io::Cursor::new(b"1\n"));
        assert_eq!(input8.to_vec(), [3,12,6,12,15,1,13,14,13,4,13,99,1,1,1,9]);

        // got input 0
        let input8 = &mut [3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9];
        compute(input8, &mut io::Cursor::new(b"0\n"));
        assert_eq!(input8.to_vec(), [3,12,6,12,15,1,13,14,13,4,13,99,0,0,1,9]);

    }

    #[test]
    fn test_get_input() {
        let input = get_input();
        println!("{:?}", input)
    }
}

