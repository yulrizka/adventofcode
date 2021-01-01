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

	// front row (index 0) is invalid. seat id 0 - 7 (0*8) + col
	// last seat (row 127) is invalid. staring from (127*8=1016)
	for i := uint64(8); i < 1016; i++ {
		if i > 7 && i < 1016 && !seatTaken[i] {
			return strconv.FormatUint(i, 10), nil
		}
	}
	return "", fmt.Errorf("not found")
}
