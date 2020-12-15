package day12

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/rxscan"
)

type direction byte

const (
	north   direction = 'N'
	south   direction = 'S'
	east    direction = 'E'
	west    direction = 'W'
	left    direction = 'L'
	right   direction = 'R'
	forward direction = 'F'
)

var rx = regexp.MustCompile(`(.)(\d+)`)

var rotation = []direction{north, east, south, west}
var rotationMap = map[direction]int{north: 0, east: 1, south: 2, west: 3}

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)

	var x, y int
	bearing := east
	for s.Scan() {
		var dirs string
		var amount int
		if _, err := rxscan.Scan(rx, s.Text(), &dirs, &amount); err != nil {
			panic(err)
		}
		dir := direction(dirs[0])
	loop:
		switch dir {
		case north:
			y += amount
		case south:
			y -= amount
		case east:
			x += amount
		case west:
			x -= amount
		case left:
			amount = amount / 90
			i := rotationMap[bearing] - amount
			if i < 0 {
				i = len(rotation) + i
			}
			bearing = rotation[i]
		case right:
			amount = amount / 90
			i := rotationMap[bearing] + amount
			bearing = rotation[i%4]
		case forward:
			dir = bearing
			goto loop
		default:
			panic("unknown dir")
		}
	}

	return strconv.Itoa(manhattan(x, y)), nil
}

func manhattan(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func Part2(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)

	var x, y int
	wx, wy := 10, 1

	rotate := func(deg int, dir direction, x, y int) (x1, y1 int) {
	loop:
		switch deg {
		case 90:
			switch dir {
			case left:
				x1 = -y
				y1 = x
			case right:
				x1 = y
				y1 = -x
			default:
				panic("invalid dir")
			}
		case 180:
			y1 = -y
			x1 = -x
		case 270:
			if dir == left {
				dir = right
			} else {
				dir = left
			}
			deg = 90
			goto loop
		default:
			panic("invalid dir")
		}
		return x1, y1
	}

	for s.Scan() {
		var dirs string
		var amount int
		if _, err := rxscan.Scan(rx, s.Text(), &dirs, &amount); err != nil {
			panic(err)
		}
		dir := direction(dirs[0])
		switch dir {
		case north:
			wy += amount
		case south:
			wy -= amount
		case east:
			wx += amount
		case west:
			wx -= amount
		case left:
			wx, wy = rotate(amount, left, wx, wy)
		case right:
			wx, wy = rotate(amount, right, wx, wy)
		case forward:
			x += wx * amount
			y += wy * amount
		default:
			panic("unknown dir")
		}
	}

	return strconv.Itoa(manhattan(x, y)), nil
}
