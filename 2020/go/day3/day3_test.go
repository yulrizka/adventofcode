package day3

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day3", "171", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day3", "1206576000", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day3", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day3", Part2)
}
