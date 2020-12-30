package day23

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	raw, err := ioutil.ReadAll(f)
	aoc.NoError(err)

	var cups []byte
	for _, b := range raw {
		if b == '\n' {
			continue
		}
		i, err := strconv.Atoi(string(b))
		aoc.NoError(err)
		cups = append(cups, byte(i))
	}
	if cups == nil {
		panic("cups is nil")
	}

	var cupIndex int

	for n := 0; n < 100; n++ {
		cupValue := cups[cupIndex]
		var pickup []byte

		// pickup 3 cups next to ci
		var ci = cupIndex + 1
		for i := 1; i <= 3; i++ {
			if ci >= len(cups) {
				pickup = append(pickup, cups[0])
				cups = cups[1:]
				continue
			}
			pickup = append(pickup, cups[ci])
			cups = append(cups[:ci], cups[ci+1:]...)
		}

		// find destination
		dest := cupValue - 1
		for {
			if dest == 0 {
				dest = 9
			}
			if pos := bytes.IndexByte(pickup, dest); pos == -1 {
				break // found
			}
			dest--
		}
		ix := bytes.IndexByte(cups, dest)
		if ix == -1 {
			panic("not found")
		}

		// copy
		head := []byte(string(cups[:ix+1]))
		tail := []byte(string(cups[ix+1:]))

		newCups := make([]byte, 0, len(cups))
		newCups = append(head, pickup...)
		newCups = append(newCups, tail...)
		cups = newCups

		cupIndex = bytes.IndexByte(cups, cupValue)
		if cupIndex == len(cups)-1 {
			cupIndex = 0
		} else {
			cupIndex++
		}
	}

	// rotate until 1 is in the front
	for cups[0] != 1 {
		cups = append(cups[1:], cups[0])
	}

	var b bytes.Buffer
	for i, v := range cups {
		if i == 0 {
			continue // skip 1
		}
		b.WriteString(strconv.Itoa(int(v)))
	}

	return b.String(), nil
}

func Part2(f io.Reader) (string, error) {
	return "", nil
}
