package day20

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "../../input/day20", "27803643063307", Part1)
}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day20", "1644", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day20", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day20", Part2)
}

func TestFlip(t *testing.T) {
	got := `1111111111
2222222222
3333333333
4444444444
5555555555
6666666666
7777777777
8888888888
9999999999
0000000000`

	want := `0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321`

	want1 := `1234567890
1234567890
1234567890
1234567890
1234567890
1234567890
1234567890
1234567890
1234567890
1234567890`

	n := tile{
		s: []string{},
	}
	a := strings.Split(got, "\n")
	for _, s := range a {
		n.s = append(n.s, s)
	}

	n.flip()
	assert.Equal(t, a, n.s)

	n.rotateRight()
	b := strings.Split(want, "\n")
	assert.Equal(t, b, n.s)
	c := strings.Split(want1, "\n")
	n.flip()
	assert.Equal(t, c, n.s)
}

func TestClean(t *testing.T) {
	a := `0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321
0987654321`

	want := `98765432
98765432
98765432
98765432
98765432
98765432
98765432
98765432`

	n := tile{}
	n.load(a)

	n.cleanBorder()

	b := strings.Split(want, "\n")
	assert.Equal(t, b, n.s)

}

func TestCombine(t *testing.T) {
	a := `98765432
98765432
98765432
98765432
98765432
98765432
98765432
98765432`

	b := `9876543298765432
9876543298765432
9876543298765432
9876543298765432
9876543298765432
9876543298765432
9876543298765432
9876543298765432`

	n1 := tile{}
	n1.load(a)

	n2 := tile{}
	n2.load(a)

	c := combine(n1.s, n2.s)
	wantc := strings.Split(b, "\n")

	assert.Equal(t, c, wantc)

}
