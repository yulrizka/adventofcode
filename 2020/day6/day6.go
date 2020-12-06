package day6

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)

	var sum int
	ans := map[byte]struct{}{}
	for s.Scan() {
		for _, c := range s.Bytes() {
			ans[c] = struct{}{}
		}

		if s.Text() == "" {
			// end of group count
			sum += len(ans)
			ans = map[byte]struct{}{}
		}
	}
	sum += len(ans)

	return strconv.FormatInt(int64(sum), 10), nil
}

func Part2(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)

	var sum, numPerson int
	ans := map[byte]int{}
	var i int64

	check := func(text string) {
		if text != "" {
			numPerson++
			return
		}

		// end of group
		for _, num := range ans {
			if num == numPerson {
				sum++
			}
		}

		ans = map[byte]int{}
		numPerson = 0
	}

	for s.Scan() {
		i++
		for _, c := range s.Bytes() {
			ans[c] += 1
		}
		check(s.Text())
	}
	// check the final one
	check(s.Text())

	return strconv.FormatInt(int64(sum), 10), nil
}
