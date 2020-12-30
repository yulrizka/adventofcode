package day23

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	max := 9
	root, index := parseNode(f, max)

	iteration := 100
	simulate(root, index, max, iteration)
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
	max := 1000_000
	root, index := parseNode(f, max)

	iteration := 10_000_000
	simulate(root, index, max, iteration)

	one := index[1]
	sum := one.next.v * one.next.next.v
	return strconv.Itoa(sum), nil
}

type node struct {
	v    int
	next *node
}

// parseNode construct a linked list node, index to each node and max number
func parseNode(f io.Reader, max int) (*node, map[int]*node) {
	root := new(node)

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

	for v := len(index) + 1; v <= max; v++ {
		c.next = new(node)
		c.next.v = v
		c = c.next
		index[v] = c
	}

	// close the loop
	c.next = root

	return root, index
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
		destValue := cup.v
		for {
			destValue--
			if destValue == 0 {
				destValue = max // wrap
			}
			if destValue != start.v && destValue != start.next.v && destValue != end.v {
				break
			}
		}

		dest := index[destValue]

		// stitch it again
		temp := dest.next
		dest.next = start
		end.next = temp

		cup = cup.next
	}
}
