package day18

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day18", "25190263477788", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day18", "297139939002972", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day18", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day18", Part2)
}
