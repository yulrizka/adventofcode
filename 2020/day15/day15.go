package day15

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	var ints []int
	for _, i := range strings.Split(sc.Text(), ",") {
		i, err := strconv.Atoi(i)
		aoc.NoError(err)
		ints = append(ints, i)
	}
	spoken := ints
	turn := len(spoken)
	for {
		if turn >= 30000000 {
			break
		}
		last := spoken[len(spoken)-1]
		spokenBefore := -1
		for i := len(spoken) - 2; i >= 0; i-- {
			if spoken[i] == last {
				spokenBefore = i
				break
			}
		}
		if spokenBefore >= 0 {
			last := turn - 1 - spokenBefore
			spoken = append(spoken, last)
		} else {
			spoken = append(spoken, 0)
		}
		turn++
	}

	return strconv.Itoa(spoken[len(spoken)-1]), nil
}

func Part2(f io.Reader) (string, error) {
	return "", nil
}
