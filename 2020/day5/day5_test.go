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
	row, col = rowCol("FBFBBFFRLR")
	assert.Equal(t, 44, row)
	assert.Equal(t, 5, col)

	row, col = rowCol("BFFFBBFRRR")
	assert.Equal(t, 70, row)
	assert.Equal(t, 7, col)

	row, col = rowCol("FFFBBBFRRR")
	assert.Equal(t, 14, row)
	assert.Equal(t, 7, col)

	row, col = rowCol("BBFFBBFRLL")
	assert.Equal(t, 102, row)
	assert.Equal(t, 4, col)

	row, col = rowCol("BBFBBBBRRL")
	assert.Equal(t, 111, row)
	assert.Equal(t, 6, col)

}

// getID2 is my first attempt to do do binary search manually. It's not used
//goland:noinspection GoUnusedFunction
func getID2(text string) (uint64, error) {
	row, col := rowCol(text)
	return uint64(row*8 + col), nil
}

func rowCol(text string) (row, col int) {
	// parse row
	l := 0
	r := 127
	for i := 0; i < 6; i++ {
		mid := l + (r-l)/2
		switch text[i] {
		case 'F': // lower
			r = mid - 1
		case 'B': // upper
			l = mid + 1
		default:
			panic("got invalid row input")
		}
	}
	// last bit represents 0 or 1
	row = l
	if text[6] == 'B' {
		row += 1
	}

	l, r = 0, 7
	for i := 7; i < 10; i++ {
		mid := l + (r-l)/2
		switch text[i] {
		case 'L': // lower
			r = mid
		case 'R': // upper
			l = mid + 1
		default:
			panic("got invalid row input")
		}
		if i == 9 {
			if text[i] == 'L' {
				col = l
			} else {
				col = r
			}
		}
	}

	return row, col
}
