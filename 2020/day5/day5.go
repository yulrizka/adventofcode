package day5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// getID return the set representation of the binary.
// ex BFFFBBFRRR is actually -> 1000110111 which is ID 567
func getID(s string) (uint64, error) {
	b := strings.ReplaceAll(s, "B", "1")
	b = strings.ReplaceAll(b, "F", "0")
	b = strings.ReplaceAll(b, "R", "1")
	b = strings.ReplaceAll(b, "L", "0")
	id, err := strconv.ParseUint(b, 2, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func parse(f io.Reader) (map[uint64]bool, error) {
	s := bufio.NewScanner(f)

	taken := map[uint64]bool{}
	for s.Scan() {
		id, err := getID(s.Text())
		if err != nil {
			return nil, err
		}
		taken[id] = true
	}
	return taken, nil
}

func Part1(f io.Reader) (string, error) {
	seatTaken, err := parse(f)
	if err != nil {
		return "", err
	}

	var max uint64
	for id := range seatTaken {
		if id > max {
			max = id
		}
	}

	return strconv.FormatUint(max, 10), nil
}

func Part2(f io.Reader) (string, error) {
	seatTaken, err := parse(f)
	if err != nil {
		return "", err
	}

	// front row (index 0) is invalid. seat id 0 - 8
	// last seat (row 127) is invalid. staring from (127*8=1016)
	for i := uint64(8); i < 1016; i++ {
		if i > 7 && i < 1016 && !seatTaken[i] {
			return strconv.FormatUint(i, 10), nil
		}
	}
	return "", fmt.Errorf("not found")
}

// getID2 is my first attempt to do do binary search manually. It's not used
func getID2(text string) (row int, col int) {
	// parse row
	l := 0
	r := 127
	for i := 0; i < 6; i++ {
		switch text[i] {
		case 'F': // lower
			r -= ((r - l) / 2) + 1
		case 'B': // upper
			l += ((r - l) / 2) + 1
		default:
			panic("got invalid row input")
		}
		if i == 5 {
			if text[i] == 'F' {
				row = l
			} else {
				row = r
			}
		}
	}

	l, r = 0, 7
	for i := 7; i < 10; i++ {
		switch text[i] {
		case 'L': // lower
			r -= ((r - l) / 2) + 1
		case 'R': // upper
			l += ((r - l) / 2) + 1
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
