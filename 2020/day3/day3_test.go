package day3

import (
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "171", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "input", "1206576000", Part2)
}

func BenchmarkPart1(b *testing.B) {
	adventofcode.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	adventofcode.Bench(b, "input", Part2)
}
