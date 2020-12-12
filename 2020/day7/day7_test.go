package day7

import (
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "261", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "input", "3765", Part2)
}

func BenchmarkPart1(b *testing.B) {
	adventofcode.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	adventofcode.Bench(b, "input", Part2)
}
