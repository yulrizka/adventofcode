package day8

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "1487", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "input", "1607", Part2)
}

func BenchmarkScan(b *testing.B) {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		Part1(bytes.NewReader(content))
	}
}
