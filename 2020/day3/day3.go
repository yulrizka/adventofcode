package day1

import (
	"bufio"
	"io"
	"strconv"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type world struct {
	maps [][]byte
	x    int
	y    int
}

func (w *world) look() byte {
	return w.maps[w.y][w.x]
}

func (w *world) move(d direction, amount int) {
	switch d {
	case up:
		if newValue := w.y - amount; newValue >= 0 {
			w.y = newValue
		} else {
			// wrap
			amount = amount - w.y // amount now contains how many step from the bottom
			w.y = len(w.maps) - amount
		}
	case right:
		maxX := len(w.maps[w.y])
		w.x = (w.x + amount) % maxX
	case down:
		w.y = (w.y + amount) % len(w.maps)
	case left:
		maxX := len(w.maps[w.y])
		if newValue := w.x - amount; newValue >= 0 {
			w.x = newValue
		} else {
			// wrap to right
			amount = amount - w.x // amount  to move from most right
			w.x = maxX - amount
		}
	}
}

func ParseWorld(f io.Reader) *world {
	w := world{
		maps: make([][]byte, 0),
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		w.maps = append(w.maps, []byte(s.Text()))
	}
	return &w
}

func checkTree(w *world, r int, d int) int64 {
	w.x = 0
	w.y = 0

	var tree int64
	for w.y != len(w.maps)-1 {
		w.move(right, r)
		w.move(down, d)
		v := w.look()
		vStr := string([]byte{v})
		_ = vStr
		if v == '#' {
			tree++
		}
	}
	return tree
}

func Part1(f io.Reader) (string, error) {
	w := ParseWorld(f)

	tree := checkTree(w, 3, 1)

	return strconv.FormatInt(tree, 10), nil
}

func Part2(f io.Reader) (string, error) {
	w := ParseWorld(f)

	a := checkTree(w, 1, 1)
	b := checkTree(w, 3, 1)
	c := checkTree(w, 5, 1)
	d := checkTree(w, 7, 1)
	e := checkTree(w, 3, 2)

	tree := a * b * c * d * e

	return strconv.FormatInt(tree, 10), nil
}
