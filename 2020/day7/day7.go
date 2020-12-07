package day7

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	name  string
	edges map[string]int64
}

func newNode(name string) *node {
	return &node{
		name:  name,
		edges: map[string]int64{},
	}
}

type graph struct {
	nodes map[string]*node
}

func (g *graph) addNode(n *node) {
	g.nodes[n.name] = n
}

func (g *graph) addEdge(name1 string, name2 string, weight int64) {
	n1 := g.nodes[name1]
	n2 := g.nodes[name2]
	if n1 == nil {
		g.nodes[name1] = newNode(name1)
		n1 = g.nodes[name1]
	}
	if n2 == nil {
		g.nodes[name2] = newNode(name2)
		n2 = g.nodes[name2]
	}
	n1.edges[name2] = weight
}

func (g graph) hasPath(start string, dst string) bool {
	if start == dst {
		return true
	}

	startNode, ok := g.nodes[start]
	if !ok {
		panic("start node not found")
	}

	for n := range startNode.edges {
		if g.hasPath(n, dst) {
			return true
		}
	}

	return false
}

func (g *graph) countWeight(name string) int64 {
	n, ok := g.nodes[name]
	if !ok {
		panic("node not found")
	}

	if len(n.edges) == 0 {
		return 1
	}

	sum := int64(1)
	for edgeName, weight := range n.edges {
		sum += weight * g.countWeight(edgeName)
	}

	return sum
}

var bagRx = regexp.MustCompile(`(\d) (.+)`)

func parse(f io.Reader) (*graph, error) {
	g := graph{
		nodes: map[string]*node{},
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		// "light red bags contain 1 bright white bag, 2 muted yellow bags."
		// -> "light red  contain 1 bright white , 2 muted yellow "
		text := strings.ReplaceAll(s.Text(), "bags", "")
		text = strings.ReplaceAll(text, "bag", "")
		text = strings.ReplaceAll(text, ".", "")

		// 0 -> "light red  "
		// 1 -> " 1 bright white , 2 muted yellow "
		fields := strings.Split(text, " contain ")
		if len(fields) != 2 {
			return nil, fmt.Errorf("first scan got %d want 2", len(fields))
		}

		// "light red"
		name := strings.TrimSpace(fields[0])
		n := newNode(name)
		g.addNode(n)

		// " 1 bright white , 2 muted yellow"
		fields = strings.Split(fields[1], ",")
		for _, f := range fields {
			// "1 bright white"
			f = strings.TrimSpace(f)
			if strings.Contains(f, "no other") {
				continue
			}

			// 0 -> 1 bright white"
			// 1 -> 1
			// 2 -> bright white
			submatch := bagRx.FindStringSubmatch(f)
			if len(submatch) != 3 {
				return nil, fmt.Errorf("inside got %d want 3", len(submatch))
			}
			count, err := strconv.ParseInt(submatch[1], 10, 64)
			if err != nil {
				return nil, err
			}

			g.addEdge(n.name, submatch[2], count)
		}
	}

	return &g, nil
}

var target = "shiny gold"

func Part1(f io.Reader) (string, error) {
	g, err := parse(f)
	if err != nil {
		return "", err
	}

	var count int64
	for n := range g.nodes {
		if n == target {
			continue // skip it self
		}
		if g.hasPath(n, target) {
			count++
		}
	}

	return strconv.FormatInt(count, 10), nil
}

func Part2(f io.Reader) (string, error) {
	g, err := parse(f)
	if err != nil {
		return "", err
	}

	// -1 because we exclude (sum=1) on the top node
	// Example if the node has no children, it will return 1 instead of 0 which is wrong
	count := g.countWeight(target) - 1

	return strconv.FormatInt(count, 10), nil
}
