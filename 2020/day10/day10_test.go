package day10

import (
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "1856", Part1)
}

func TestPart2(t *testing.T) {
	//adventofcode.Test(t, "sample-small", "8", Part2)
	//adventofcode.Test(t, "sample", "19208", Part2)
	adventofcode.Test(t, "input", "", Part2)
}
