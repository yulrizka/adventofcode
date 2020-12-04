package day1

import (
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "213", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "input", "147", Part2)
}
