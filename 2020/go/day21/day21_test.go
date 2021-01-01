package day21

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "input", "2075", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "input", "zfcqk,mdtvbb,ggdbl,frpvd,mgczn,zsfzq,kdqls,kktsjbh", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "input", Part2)
}
