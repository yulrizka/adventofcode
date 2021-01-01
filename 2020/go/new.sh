#!/usr/bin/env bash

if [ "$#" -ne 1 ]; then
    echo "use $0 <day number>"
    exit 1
fi

day="day$1"

mkdir "$day"
cat <<EOF > "$day/$day.go"
package $day

import (
	"bufio"
	"io"
)

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	for s.Scan() {
	}

	return "", nil
}

func Part2(f io.Reader) (string, error) {
	return "", nil
}
EOF

cat <<EOF > "$day/${day}_test.go"
package $day

import (
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "input", "", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "input", "", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "input", Part2)
}
EOF
