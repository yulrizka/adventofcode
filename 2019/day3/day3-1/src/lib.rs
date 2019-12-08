use std::collections::HashMap;

pub fn get_input() -> (String, String) {
    let contents = std::fs::read_to_string("../input").expect("Error opening file!");
    let mut lines = contents.lines();
    let line1 = lines.next().unwrap();
    let line2 = lines.next().unwrap();

    return (line1.to_string(), line2.to_string());
}


fn parse_line(line: &str)-> HashMap<(i32,i32),i32>{
    let steps: Vec<&str> = line.split(",").collect();

    let mut points = HashMap::new();
    let mut start: (i32, i32) = (0,0);

    let mut num_step = 0;
    for step in steps {
        let dir = step.chars().next().unwrap();
        let count: i32 = step[1..].parse().unwrap();
        let movement : (i32,i32) = match dir {
            'L' => (-1, 0),
            'R' => (1, 0),
            'U' => (0,  1),
            'D' => (0, -1),
            _ => unreachable!()
        };

        for _ in 0..count {
            num_step +=1;
            start.0 = start.0 + movement.0;
            start.1 = start.1 + movement.1;

            points.entry(start).or_insert(num_step);

        }

    }
    return points;
}

pub fn compute(input1: &str, input2: &str) -> (i32, i32) {
    let line1 = parse_line(input1);
    let line2 = parse_line(input2);


    // range over hashmap, find intersection find, minimum manhattan distance
    let mut min_distance = 0;
    let mut min_timing = 0;
    for i in line1 {
        let key = i.0;
        if line2.contains_key(&key) {
            // minimal distance (for part1_
            let distance = key.0.abs() + key.1.abs();
                if min_distance == 0 || min_distance > distance{
                min_distance = distance;
            }

            // minimal timing
            let timing1 = i.1;
            let timing2= line2[&key];

            let timing = timing1 + timing2;

            if min_timing == 0 || min_timing > timing{
                min_timing = timing;
            }
        }
    }
    return (min_distance, min_timing);
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_parse_line() {
        parse_line("L5,U5");
    }

    #[test]
    fn test_compute_min() {
        assert_eq!(compute("R8,U5,L5,D3", "U7,R6,D4,L4"), (6, 30));
        assert_eq!(compute("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"), (159, 610));
        assert_eq!(compute("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"), (135, 410));
    }

    #[test]
    fn test_get_input() {
        let input = get_input();
        println!("{:?}", input)
    }
}

