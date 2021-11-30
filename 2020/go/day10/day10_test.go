package day10

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day10", "1856", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "sample-small", "8", Part2)
	aoc.Test(t, "sample", "19208", Part2)
	aoc.Test(t, "../../input/day10", "2314037239808", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day10", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day10", Part2)
}
