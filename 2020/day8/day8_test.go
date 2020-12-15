package day8

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "input", "1487", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "input", "1607", Part2)
}

func BenchmarkScan(b *testing.B) {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = Part1(bytes.NewReader(content))
	}
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "input", Part2)
}
