package day25

import (
	"io"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	keys := aoc.MustInts(f)
	dorPub := keys[1]
	cardPub := keys[0]

	transform := func(v int, subject int) int {
		return (v * subject) % 20201227
	}

	loop := 0
	v := 1
	for v != dorPub {
		loop++
		v = transform(v, 7)
	}

	v = 1
	for i := 0; i < loop; i++ {
		v = transform(v, cardPub)
	}

	return strconv.Itoa(v), nil
}

func Part2(_ io.Reader) (string, error) {
	return "", nil
}
