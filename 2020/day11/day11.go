package day11

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

const (
	Empty    = 'L'
	Floor    = '.'
	Occupied = '#'
)

func Part1(r io.Reader) (string, error) {
	s := newSeats(r)

	var i int
	//s.png(i)
	tolerance, direct := 4, true
	for s.iterate(tolerance, direct) {
		i++
		//s.png(i)
	}

	return strconv.Itoa(s.countPeople()), nil
}

func Part2(r io.Reader) (string, error) {
	s := newSeats(r)

	var i int
	//s.png(i)
	tolerance, direct := 5, false
	for s.iterate(tolerance, direct) {
		i++
		//s.png(i)
	}
	return strconv.Itoa(s.countPeople()), nil
}

type point struct{ x, y int }
type seats struct {
	m          map[point]byte
	maxX, maxY int
}

func newSeats(r io.Reader) *seats {
	sc := bufio.NewScanner(r)
	s := seats{
		m: map[point]byte{},
	}
	var y int
	for sc.Scan() {
		for x, v := range sc.Bytes() {
			s.m[point{x, y}] = v
		}
		s.maxY++
		if maxX := len(sc.Bytes()); maxX > s.maxX {
			s.maxX = maxX
		}
		y++
	}
	return &s
}

var adjacent = []point{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 1}, {1, 0}, {0, 1}, {0, -1}}

func (s *seats) occupiedAdjacent(p point, direct bool) int {
	var sum int
	for _, dir := range adjacent {
		x, y := p.x, p.y
		for {
			x, y = x+dir.x, y+dir.y
			v, ok := s.m[point{x, y}]
			if !ok {
				break //outside boundary
			}
			if v == Occupied {
				sum++
				break
			}
			if v == Empty || direct {
				break
			}
		}
	}
	return sum
}

func (s *seats) countPeople() int {
	var sum int
	for _, v := range s.m {
		if v == Occupied {
			sum++
		}
	}
	return sum
}

func (s *seats) iterate(tolerance int, direct bool) bool {
	newMap := make(map[point]byte)
	changed := false

	for p, v := range s.m {
		newMap[p] = s.m[p]
		switch v {
		case Empty:
			// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
			if s.occupiedAdjacent(p, direct) == 0 {
				newMap[p] = Occupied
			}
		case Occupied:
			// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
			if s.occupiedAdjacent(p, direct) >= tolerance {
				newMap[p] = Empty
			}
		case Floor:
			// Otherwise, the seat's state does not change.
		}
		if s.m[p] != newMap[p] {
			changed = true
		}
	}

	s.m = newMap
	return changed
}

func (s *seats) png(i int) {
	f, err := os.Create(fmt.Sprintf("%d.png", i))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := image.NewRGBA(image.Rect(0, 0, s.maxX, s.maxY))
	for y := 0; y < s.maxY; y++ {
		for x := 0; x < s.maxX; x++ {
			v := s.m[point{x, y}]
			var c color.Color
			switch v {
			case Floor:
				c = color.RGBA{R: 0, G: 0, B: 0, A: 255}
			case Empty:
				c = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			case Occupied:
				c = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			}
			img.Set(x, y, c)
		}
	}
	img2 := resize.Resize(uint(s.maxX*4), 0, img, resize.NearestNeighbor)

	if err := png.Encode(f, img2); err != nil {
		panic(err)
	}
}
