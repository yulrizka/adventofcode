package day13

import (
	"strings"
	"testing"

	"github.com/yulrizka/adventofcode/pkg/aoc"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	aoc.Test(t, "input", "296", Part1)
}

func TestPart2(t *testing.T) {
	var ans string
	var err error
	ans, err = Part2(strings.NewReader("0\n17,x,13,19"))
	require.NoError(t, err)
	require.EqualValues(t, "3417", ans)

	ans, err = Part2(strings.NewReader("0\n67,7,59,61"))
	require.NoError(t, err)
	require.EqualValues(t, "754018", ans)

	ans, err = Part2(strings.NewReader("0\n67,x,7,59,61"))
	require.NoError(t, err)
	require.EqualValues(t, "779210", ans)

	ans, err = Part2(strings.NewReader("0\n67,7,x,59,61"))
	require.NoError(t, err)
	require.EqualValues(t, "1261476", ans)

	ans, err = Part2(strings.NewReader("0\n1789,37,47,1889"))
	require.NoError(t, err)
	require.EqualValues(t, "1202161486", ans)

	aoc.Test(t, "input", "535296695251210", Part2)
}

func BenchmarkPart1(b *testing.B) {
	aoc.Bench(b, "input", Part1)
}

func BenchmarkPart2(b *testing.B) {
	aoc.Bench(b, "input", Part2)
}
