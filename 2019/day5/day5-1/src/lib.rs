use std::{fs, io};
use crate::OpCode::{Add, Multiply, Input, Output, Halt};
use std::io::Write;
use log::{debug};

pub fn get_input() -> Vec<i32> {
    let data = fs::read_to_string("../input").expect("unable to open file");
    data.trim().split(",")
        .map(|l| l.parse::<i32>().expect(&format!("failed to parse input {}", l)))
        .collect()
}

#[derive(PartialEq, Debug)]
enum OpCode {
    Add, Multiply, Input, Output, Halt
}

impl From<usize> for OpCode {
    fn from(item: usize) -> Self {
        match item {
            1 => Add,
            2 => Multiply,
            3 => Input,
            4 => Output,
            99 => Halt,
            _ => unreachable!()
        }
    }
}


fn get_param(code: &mut[i32], sp: usize, mode: usize,  n: usize) -> i32 {
    let p1 = code[sp+n];

    let mut m = mode;
    if n > 1 {
       m = (m/10*(n-1)) % 10
    } else {
        m = m % 10
    }

    // reference mode
    if m == 0 {
        if p1 < 0 {
            panic!("reference is negative");
        }
        let addr = p1 as usize;
            debug!("      GET_PARAM REF-> [{:?}] -> {:?}", addr, code[addr]); // TODO: Debugging
        return code[addr];
    }

    // immediate mode
    debug!("      GET_PARAM IMMEDIATE -> {:?}", p1); // TODO: Debugging
    return p1;
}



pub fn compute(code: &mut [i32]) {
    let mut sp = 0; // instruction pointer

    while sp < code.len() {

        let op = ((code[sp] as usize) % 100).into();
        let mode = (code[sp] / 100) as usize;
        debug!("  code[{}]:{:?}", sp, code[sp]); // TODO: Debugging
        debug!("  op:{:?}", op); // TODO: Debugging
        debug!("  mode:{:?}", mode); // TODO: Debugging
        
        match op {
            OpCode::Add => {
                let p1 = get_param(code, sp, mode, 1);
                let p2 = get_param(code, sp, mode, 2);
                let p3 = code[sp+3] as usize;
                let result = p1 + p2;
                debug!("    ADD p1:{:?}", p1); // TODO: Debugging
                debug!("    ADD p2:{:?}", p2); // TODO: Debugging
                debug!("    RESULT ADD [{:?}] -> {:?}", p3, result); // TODO: Debugging
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
                io::stdin()
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

