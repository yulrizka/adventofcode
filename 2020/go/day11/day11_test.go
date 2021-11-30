package day11

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day11", "2166", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day11", "1955", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day11", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day11", Part2)
}
