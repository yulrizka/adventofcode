package day14

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day14", "14862056079561", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day14", "3296185383161", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day14", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day14", Part2)
}
