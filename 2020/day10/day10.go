package day10

import (
	"io"
	"sort"
	"strconv"

	"github.com/yulrizka/adventofcode"
)

func Part1(f io.Reader) (string, error) {
	ints := adventofcode.MustInts(f)
	ints = append(ints, 0)
	sort.Ints(ints)
	ints = append(ints, ints[len(ints)-1]+3)

	diffs := map[int]int{}
	for i := 1; i < len(ints); i++ {
		d := ints[i] - ints[i-1]
		diffs[d]++
	}

	ans := diffs[1] * diffs[3]

	return strconv.Itoa(ans), nil
}

func Part2(r io.Reader) (string, error) {
	ints := adventofcode.MustInts(r)
	ints = append(ints, 0)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	ints = append([]int{ints[0] + 3}, ints...)

	mem := map[int]int{}
	var f func(int []int) int
	f = func(is []int) int {
		if len(is) <= 1 {
			return 1
		}

		first := is[0]
		if v, ok := mem[first]; ok {
			return v
		}

		sum := 0
		tail := is[1:]
		for i, v := range tail {
			if first-v > 3 {
				break
			}
			ans := f(tail[i:])
			sum += ans
		}
		mem[first] = sum

		return sum
	}

	return strconv.Itoa(f(ints)), nil
}
