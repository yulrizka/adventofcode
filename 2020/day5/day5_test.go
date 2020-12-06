package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yulrizka/adventofcode"
)

func TestPart1(t *testing.T) {
	adventofcode.Test(t, "input", "908", Part1)
}

func TestPart2(t *testing.T) {
	adventofcode.Test(t, "input", "619", Part2)
}

func TestParse(t *testing.T) {
	var row, col int
	//row, col = getID2("FBFBBFFRLR")
	//assert.Equal(t, 44, row)
	//assert.Equal(t, 5, col)

	row, col = getID2("BFFFBBFRRR")
	assert.Equal(t, 70, row)
	assert.Equal(t, 7, col)

	row, col = getID2("FFFBBBFRRR")
	assert.Equal(t, 14, row)
	assert.Equal(t, 7, col)

	row, col = getID2("BBFFBBFRLL")
	assert.Equal(t, 102, row)
	assert.Equal(t, 4, col)
}
