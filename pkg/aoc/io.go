package aoc

import (
	"bufio"
	"io"
	"strconv"
)

func MustInts(r io.Reader) []int {
	var lines []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err.Error())
		}
		lines = append(lines, i)
	}

	return lines
}

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
