package day10

import (
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "1856", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "sample-small", "8", Part2)
	adventofcode.Test(t, "sample", "19208", Part2)
	adventofcode.Test(t, "input", "2314037239808", Part2)
}

func BenchmarkPart1(b *testing.B) {
	adventofcode.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	adventofcode.Bench(b, "input", Part2)
}
