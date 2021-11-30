package day9

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day9", "393911906", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day9", "59341885", Part2)
}
func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day9", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day9", Part2)
}
