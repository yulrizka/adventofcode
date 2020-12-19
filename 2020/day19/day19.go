package day19

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	rules, input := parse(f)
	walk(rules[0], rules)

	valid := map[string]struct{}{}
	for _, s := range rules[0].v {
		valid[s] = struct{}{}
	}

	var sum int
	for _, s := range input {
		if _, ok := valid[s]; ok {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

func Part2(f io.Reader) (string, error) {
	rules, input := parse(f)
	walk(rules[0], rules)

	valid := map[string]struct{}{}
	for _, s := range rules[0].v {
		valid[s] = struct{}{}
	}

	var sum int
	for _, s := range input {
		if _, ok := valid[s]; ok {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

type node struct {
	el [][]int
	v  []string
}

func walk(n *node, rules map[int]*node) []string {
	if n.v != nil {
		// return the cached result
		return n.v
	}
	for _, pair := range n.el {
		var ans []string
		for _, e := range pair {
			combination := walk(rules[e], rules)
			var newAns []string
			if len(ans) == 0 {
				newAns = combination
			} else {
				for _, a := range ans {
					for _, c := range combination {
						newAns = append(newAns, a+c)
					}
				}
			}
			ans = newAns
		}
		n.v = append(n.v, ans...)
	}
	return n.v
}

func parse(f io.Reader) (rules map[int]*node, input []string) {
	rules = make(map[int]*node)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			break // line separator
		}
		text := s.Text()
		i := strings.Index(text, ":")
		idStr, rest := text[:i], text[i+2:]
		id, err := strconv.Atoi(idStr)
		aoc.NoError(err)

		n := node{}
		if strings.Contains(rest, "|") {
			var elements [][]int
			for _, tuple := range strings.Split(rest, "|") {
				var element []int
				for _, f := range strings.Fields(tuple) {
					v, err := strconv.Atoi(f)
					aoc.NoError(err)
					element = append(element, v)
				}
				elements = append(elements, element)
			}
			n.el = elements
			rules[id] = &n
			continue
		}
		if rest == `"a"` {
			rules[id] = &node{v: []string{"a"}}
			continue
		}
		if rest == `"b"` {
			rules[id] = &node{v: []string{"b"}}
			continue
		}
		var element []int
		for _, f := range strings.Fields(rest) {
			v, err := strconv.Atoi(f)
			aoc.NoError(err)
			element = append(element, v)
		}
		rules[id] = &node{el: [][]int{element}}
	}

	// parse text
	for s.Scan() {
		input = append(input, s.Text())
	}
	return
}
