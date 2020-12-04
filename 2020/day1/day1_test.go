package day1

import (
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "744475", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "input", "70276940", Part2)
}
