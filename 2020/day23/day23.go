package day23

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	root, index, max := parseNode(f, 9)
	simulate(root, index, max, 100)
	one := index[1]
	c := one.next

	var b bytes.Buffer
	for c != one {
		b.WriteString(strconv.Itoa(c.v))
		c = c.next
	}

	return b.String(), nil
}

func Part2(f io.Reader) (string, error) {
	root, index, max := parseNode(f, 1000_000)
	simulate(root, index, max, 10_000_000)

	one := index[1]
	sum := one.next.v * one.next.next.v
	return strconv.Itoa(sum), nil
}

type node struct {
	v    int
	next *node
}

// parseNode construct a linked list node, index to each node and max number
func parseNode(f io.Reader, addTo int) (*node, map[int]*node, int) {
	root := new(node)

	var max int

	index := map[int]*node{}

	raw, err := ioutil.ReadAll(f)
	aoc.NoError(err)

	c := root
	first := true
	for _, b := range raw {
		if b == '\n' {
			continue
		}
		i, err := strconv.Atoi(string(b))
		aoc.NoError(err)

		if max < i {
			max = i
		}

		if first {
			c.v = i
			index[i] = c
			first = false
		} else {
			c.next = new(node)
			c.next.v = i
			c = c.next
			index[i] = c
		}
	}

	for v := max + 1; v <= addTo; v++ {
		max = v
		c.next = new(node)
		c.next.v = v
		c = c.next
		index[v] = c
	}

	// close the loop
	c.next = root

	return root, index, max
}

func simulate(root *node, index map[int]*node, max int, iteration int) {
	cup := root
	for n := 0; n < iteration; n++ {

		// pickup 3 cups next to the current cup
		var start, end *node
		start = cup.next
		end = start.next.next

		// remove the picked-up cup from the chain
		cup.next = end.next

		// find the destination and repeat util it picks up cup that are not part of the picked-up cup earlier
		pickup := map[int]bool{start.v: true, start.next.v: true, end.v: true}
		destValue := cup.v - 1
		for {
			if destValue <= 0 {
				destValue = max // wrap
			}
			if !pickup[destValue] {
				break // found
			}
			destValue--
		}

		dest := index[destValue]

		// stitch it again
		temp := dest.next
		dest.next = start
		end.next = temp

		cup = cup.next
	}
}
