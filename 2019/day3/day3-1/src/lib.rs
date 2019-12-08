use std::str::FromStr;

pub fn get_input() -> (String, String) {
    let contents = std::fs::read_to_string("../input").expect("Error opening file!");
    let mut lines = contents.lines();
    let line1 = lines.next().unwrap();
    let line2 = lines.next().unwrap();

    return (line1.to_string(), line2.to_string());
}

#[derive(Debug, Clone, Eq, PartialEq, Hash)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug, Clone)]
struct Instruction {
    dir: Point,
    distance: usize
}

impl FromStr for Instruction {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Instruction{
            dir: match s.chars().nth(0).unwrap() {
                'U' => Point{ x: 0 , y: 1},
                'D' => Point{ x: 0 , y: -1},
                'L' => Point{ x: -1 , y: 0},
                'R' => Point{ x: 1 , y: 0},
                _ => unreachable!()
            },
            distance: s[1..]
                .parse::<usize>().expect("failed to parse")
        })
    }
}

fn parse_line(line: &str) {
//    fn parse_line(line: &str) -> HashMap<Point, i32> {
    let _instructions: Vec<_> = line
        .split(",")
        .map(|token| token.parse::<Instruction>().unwrap())
        .collect();

        println!("{:?}", _instructions.iter()
            .flat_map(|i| std::iter::repeat(i).take((i).distance))
            .scan((Point{x:0 , y:0}, 0), |(prev, step), Instruction{dir, ..}|{
                prev.x += dir.x;
                prev.y += dir.y;
                *step += 1;
                Some((prev, step))
            }));
}

pub fn compute(_code: &mut [usize]) {
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_parse_line() {
        parse_line("L5,U5");
    }

    #[test]
    fn test_get_input() {
        let input = get_input();
        println!("{:?}", input)
    }
}

