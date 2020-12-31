package day24

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(f io.Reader) (string, error) {
	m := buildTiles(f)

	sumBlack := 0
	for _, v := range m {
		if v == black {
			sumBlack++
		}
	}

	return strconv.Itoa(sumBlack), nil
}

func Part2(f io.Reader) (string, error) {
	tileMap := buildTiles(f)

	for day := 1; day <= 100; day++ {
		newM := map[point]byte{}
		for p, tile := range tileMap {
			sumBlack := 0
			for _, nn := range neighbors {
				np := add(p, nn)
				if neighbourColor, ok := tileMap[np]; ok {
					if neighbourColor == black {
						sumBlack++
					}
				} else {
					if _, ok := newM[np]; !ok {
						newM[np] = white
					}
				}
			}

			if tile == black && (sumBlack == 0 || sumBlack > 2) {
				newM[p] = white
			} else if tile == white && sumBlack == 2 {
				newM[p] = black
			} else {
				newM[p] = tile
			}
		}

		tileMap = newM
	}
	sumBlack := 0
	for _, v := range tileMap {
		if v == black {
			sumBlack++
		}
	}

	return strconv.Itoa(sumBlack), nil
}

type direction string

const (
	white = iota
	black

	east      direction = "e"
	southeast direction = "se"
	southwest direction = "sw"
	west      direction = "w"
	northwest direction = "nw"
	northeast direction = "ne"
)

type point struct {
	x, y int
}

var neighbors = []point{{1, -1}, {1, 0}, {0, 1}, {-1, 1}, {-1, 0}, {0, -1}}

func add(p1, p2 point) point {
	return point{x: p1.x + p2.x, y: p1.y + p2.y}
}

// move point in hexagon with Axial Coordinate system (https://youtu.be/z3PaGIQTFSE?t=844)
func move(p *point, d direction) {
	switch d {
	case east:
		p.x++
		p.y--
	case southeast:
		p.x++
	case southwest:
		p.y++
	case west:
		p.x--
		p.y++
	case northwest:
		p.x--
	case northeast:
		p.y--
	default:
		panic("unknown direction " + d)
	}
}

func buildTiles(f io.Reader) map[point]byte {
	s := bufio.NewScanner(f)
	m := map[point]byte{}
	for s.Scan() {
		p := point{}
		instruction := s.Bytes()
		for len(instruction) != 0 {
			ins := direction(instruction[0])
			switch ins {
			case east, west:
				move(&p, ins)
				instruction = instruction[1:]
			default:
				move(&p, direction(instruction[0:2]))
				instruction = instruction[2:]
			}
		}
		if n, ok := m[p]; !ok {
			m[p] = black
		} else {
			if n == white {
				m[p] = black
			} else {
				m[p] = white
			}
		}
	}

	// add neighbors
	newM := map[point]byte{}
	for p, tile := range m {
		for _, nn := range neighbors {
			np := add(p, nn)
			if _, ok := m[np]; !ok {
				newM[np] = white
			}
		}
		newM[p] = tile
	}

	return newM
}
