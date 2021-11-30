package day15

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func TestPart1(t *testing.T) {
	var s string
	var err error
	s, err = Part1(strings.NewReader("0,3,6"))
	require.NoError(t, err)
	assert.EqualValues(t, "436", s)

	s, err = Part1(strings.NewReader("1,3,2"))
	require.NoError(t, err)
	assert.EqualValues(t, "1", s)

	s, err = Part1(strings.NewReader("2,1,3"))
	require.NoError(t, err)
	assert.EqualValues(t, "10", s)

	s, err = Part1(strings.NewReader("1,2,3"))
	require.NoError(t, err)
	assert.EqualValues(t, "27", s)

	s, err = Part1(strings.NewReader("2,3,1"))
	require.NoError(t, err)
	assert.EqualValues(t, "78", s)

	s, err = Part1(strings.NewReader("3,2,1"))
	require.NoError(t, err)
	assert.EqualValues(t, "438", s)

	s, err = Part1(strings.NewReader("3,1,2"))
	require.NoError(t, err)
	assert.EqualValues(t, "1836", s)

	aoc.Test(t, "../../input/day15", "276", Part1)

}

func TestPart2(t *testing.T) {
	aoc.Test(t, "../../input/day15", "31916", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "../../input/day15", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "../../input/day15", Part2)
}
