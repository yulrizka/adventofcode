package day13

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func read(f io.Reader) (earliest int, buses []int) {
	s := bufio.NewScanner(f)

	s.Scan()
	earliest, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}

	s.Scan()
	for _, v := range strings.Split(s.Text(), ",") {
		if v == "" || v == "x" {
			buses = append(buses, 0)
			continue
		}
		id, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		buses = append(buses, id)
	}

	return earliest, buses
}

func Part1(f io.Reader) (string, error) {
	earliest, buses := read(f)
	ts := earliest
	for {
		for _, b := range buses {
			if b == 0 {
				continue // x
			}
			if ts%b == 0 {
				ans := (ts - earliest) * b
				return strconv.Itoa(ans), nil
			}
		}
		ts++
	}
}

func Part2(f io.Reader) (string, error) {
	// first approach (bruteforce) did not yield result.
	// kudos: https://www.reddit.com/r/adventofcode/comments/kc4njx/2020_day_13_solutions/gfncyoc/?context=3
	_, buses := read(f)

	ts, step := 0, 1

	// for each element, find how many steps needed for the condition to be true. At that point,
	// we know that it will collide again after the same amount of 'steps'. So we can do the same for the
	// next element
	// eg part 1: [3,5,6], step=3. we know it will collide at 15 -> step = step*5=15.
	// we then repeat the process to find when 15 & 6x will collide (check 15, 30) -> ans 30
	//
	// part 2 is the same, it's just we add the ts with index of the array to represent i minute after
	for i, b := range buses {
		if b == 0 {
			continue
		}
		for (ts+i)%b != 0 { // + i because it's i minute after ts
			ts += step
		}
		step *= b
	}

	return strconv.Itoa(ts), nil
}
