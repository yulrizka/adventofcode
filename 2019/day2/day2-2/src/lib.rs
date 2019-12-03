use std::fs;

pub fn get_input() -> Vec<i64> {
    let data = fs::read_to_string("../input").expect("unable to open file");
    data.trim().split(",")
        .map(|l| l.parse::<i64>().expect(&format!("failed to parse input {}", l)))
        .collect()
}

pub fn compute(code: &mut [i64]) -> Vec<i64> {
    let mut ip = 0; // instruction pointer
    while ip < code.len() {
        let opcode = code[ip];
        if opcode == 99 {
            return code.to_vec()
        }

        // get 1st parameter
        let pos_param1 = code[ip +1];
        let param1 = code[pos_param1 as usize];

        // get 2nd parameter
        let pos_param2 = code[ip +2];
        let param2 = code[pos_param2 as usize];

        // get 3rd parameter (output position)
        let param3 = code[ip +3];

        let mut result = 0;
        if opcode == 1 {
            result = param1 + param2;
        } else if opcode == 2 {
            result = param1 * param2;
        }
        code[param3 as usize] = result;

        ip += 4
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

