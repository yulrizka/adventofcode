package day9

import (
	"errors"
	"io"
	"sort"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

// check if x is result of adding 2 different element of s
func check(s []int, x int) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			if s[i]+s[j] == x {
				return true
			}
		}
	}
	return false
}

// invalid numbers on lines which did not pass the check
func invalid(lines []int, preamble int) int {
	for i := preamble; i < len(lines); i++ {
		start := i - preamble
		v := lines[i]
		if !check(lines[start:i], v) {
			return v
		}
	}

	return -1
}

func Part1(f io.Reader) (string, error) {
	lines := aoc.MustInts(f)

	v := invalid(lines, 25)
	if v < 0 {
		return "", errors.New("not found")
	}

	return strconv.Itoa(v), nil
}

func Part2(f io.Reader) (string, error) {
	lines := aoc.MustInts(f)

	wantSum := invalid(lines, 25)
	if wantSum < 0 {
		return "", errors.New("not found")
	}

next:
	for i := 0; i < len(lines); i++ {
		a := lines[i]
		sum := a
		for j := i + 1; j < len(lines); j++ {
			b := lines[j]
			sum += b
			if sum == wantSum {
				ranges := lines[i : j+1]
				sort.Ints(ranges)
				min, max := ranges[0], ranges[len(ranges)-1]
				return strconv.Itoa(min + max), nil
			}
			if sum > wantSum {
				continue next
			}
		}
	}

	return "", errors.New("not found")
}
