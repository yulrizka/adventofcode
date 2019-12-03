use std::fs;

pub fn get_input() -> Vec<i64> {
    let data = fs::read_to_string("../input").expect("unable to open file");
    data.trim().split(",")
        .map(|l| l.parse::<i64>().expect(&format!("failed to parse input {}", l)))
        .collect()
}

pub fn compute(code: &mut [i64]) -> &[i64] {
    let mut i = 0;
    while i < code.len() {
        let opcode = code[i];
        if opcode == 99 {
            return code
        }

        let pos_op1 = code[i+1];
        let op1 = code[pos_op1 as usize];
        let pos_op2 = code[i+2];
        let op2 = code[pos_op2 as usize];
        let out_pos = code[i+3];

        let mut result = 0;
        if opcode == 1 {
            result = op1 + op2;
        } else if opcode == 2 {
            result = op1 * op2;
        }
        code[out_pos as usize] = result;

        i += 4
    }

    panic!("should not come to this");
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_solve() {
        assert_eq!(compute(&mut[1,0,0,0,99]), [2,0,0,0,99]);
        assert_eq!(compute(&mut[2,3,0,3,99]), [2,3,0,6,99]);
        assert_eq!(compute(&mut[2,4,4,5,99,0]), [2,4,4,5,99,9801]);
        assert_eq!(compute(&mut[1,1,1,4,99,5,6,0,99]), [30,1,1,4,2,5,6,0,99]);
    }

    #[test]
    fn test_get_input() {
        let input = get_input();
        println!("{:?}", input)
    }
}

