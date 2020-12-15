package day15

import (
	"strings"
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
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
