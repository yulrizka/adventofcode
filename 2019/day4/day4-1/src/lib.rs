
pub fn compute(start: usize, end: usize) -> (usize, usize) {
    let mut count1 = 0;
    let mut count2= 0;
    for i in start..end {
        if check(i).0 {
            count1 += 1;
        }

        if check(i).1 {
            count2 += 1;
        }
    }

    return (count1, count2);
}

pub fn check(num: usize) -> (bool, bool) {
    let mut i = num;
    let mut last = 10; // 10 because the first iteration (val > last) will always true

    // answers
    let mut part1 = false;
    let mut part2 = false;

    let mut max_sequence = 0;

    while i > 0 {
        let val = i % 10;
        i /= 10;

        if val > last {
            return (false, false)
        }

        if val == last {
            part1 = true;
            max_sequence += 1;
        } else {
            // sequence is no longer matched
            if max_sequence == 1 {
                part2 = true
            }
            max_sequence = 0;
        }

       last = val;
    }

    // either the last 2 character is the same, or there is 2 character matches
    part2 = max_sequence == 1 || part2;

    return (part1, part2)
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_check() {
        // test part1
        assert_eq!(check(111111).0, true);
        assert_eq!(check(223450).0, false);
        assert_eq!(check(123789).0, false);

        // test part2
        //meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
        assert_eq!(check(112233).1, true);

        // no longer meets the criteria (the repeated 44 is part of a larger group of 444)
        assert_eq!(check(444555).1, true);

        // meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
        assert_eq!(check(111122).1, true);
    }

    #[test]
    fn test_compute() {
        let count = compute(236491, 713787);
        assert_eq!(count.1, 757)
    }
}

