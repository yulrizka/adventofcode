package day15

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func solve(ints []int, n int) (string, error) {
	// cache
	lastSpoken := map[int]int{}
	for i := 0; i < len(ints)-1; i++ {
		lastSpoken[ints[i]] = i + 1
	}

	turn := len(ints) - 1
	prev := ints[turn]

	for {
		if len(ints) >= n {
			break
		}

		turn++
		last, ok := lastSpoken[prev]
		if !ok {
			// never spoken before
			lastSpoken[prev] = turn
			prev = 0
			ints = append(ints, 0)
		} else {
			lastSpoken[prev] = turn
			prev = turn - last
			ints = append(ints, prev)
		}
	}

	return strconv.Itoa(ints[n-1]), nil
}
func Part1(f io.Reader) (string, error) {
	return solve(ints(f), 2020)
}

func Part2(f io.Reader) (string, error) {
	return solve(ints(f), 30000000)
}

func ints(f io.Reader) []int {
	sc := bufio.NewScanner(f)
	sc.Scan()
	var ints []int
	for _, i := range strings.Split(sc.Text(), ",") {
		i, err := strconv.Atoi(i)
		aoc.NoError(err)
		ints = append(ints, i)
	}
	return ints
}
