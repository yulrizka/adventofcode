package day25

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day25", "2679568", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day25", "", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day25", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day25", Part2)
}
