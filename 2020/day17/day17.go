// the idea is to store all the point in a map (world.m) and initialize it's surrounding neighbour
// for each iteration, range over the map and perform the rule store it in a new map and replace the current map
package day17

import (
	"bufio"
	"io"
	"strconv"
)

const (
	active   byte = '#'
	inactive byte = '.'
)

type point struct {
	x, y, z, a int
}

func (p point) add(p2 point) point {
	return point{
		x: p.x + p2.x,
		y: p.y + p2.y,
		z: p.z + p2.z,
		a: p.a + p2.a,
	}
}

type world struct {
	m          map[point]byte
	neighbours []point
}

func newWorld(f io.Reader, neighbours []point) *world {
	w := world{
		m:          map[point]byte{},
		neighbours: neighbours,
	}
	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		for j, v := range s.Bytes() {
			p := point{i, j, 0, 0}
			w.m[p] = v
		}
		i++
	}

	// add surrounding neighbours for each point
	for p := range w.m {
		for _, n := range neighbours {
			np := p.add(n)
			if _, ok := w.m[np]; !ok {
				w.m[np] = inactive
			}
		}
	}

	return &w
}

func (w *world) simulate() {
	newMap := map[point]byte{}
	for p, v := range w.m {
		var nActive int
		for _, diff := range w.neighbours {
			n := p.add(diff)
			if _, ok := w.m[n]; !ok {
				newMap[n] = inactive
				continue
			} else {
				if w.m[n] == active {
					nActive++
				}
			}
		}
		if v == active {
			if nActive != 2 && nActive != 3 {
				v = inactive
			}
		} else { // inactive
			if nActive == 3 {
				v = active
			}
		}
		newMap[p] = v
	}
	w.m = newMap
}

func (w *world) active() int {
	var sum int
	for _, v := range w.m {
		if v == active {
			sum++
		}
	}
	return sum
}

func Part1(f io.Reader) (string, error) {
	// 3D neighbor
	var neighbours []point
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				neighbours = append(neighbours, point{i, j, k, 0})
			}
		}
	}
	w := newWorld(f, neighbours)

	// simulate
	for i := 0; i < 6; i++ {
		w.simulate()
	}

	return strconv.Itoa(w.active()), nil
}

func Part2(f io.Reader) (string, error) {
	// 4D neighbor
	var neighbours []point
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					neighbours = append(neighbours, point{i, j, k, l})
				}
			}
		}
	}
	w := newWorld(f, neighbours)

	// simulate
	for i := 0; i < 6; i++ {
		w.simulate()
	}

	return strconv.Itoa(w.active()), nil
}
