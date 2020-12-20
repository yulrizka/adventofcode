package day20

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	tiles := parse(f)
	mul := 1
	corners, _ := getCorners(tiles)
	for _, t := range corners {
		mul *= t.id
	}

	return strconv.Itoa(mul), nil
}

func Part2(f io.Reader) (string, error) {
	//tiles := parse(f)
	//corners, edgeMatches := getCorners(tiles)

	return "", nil
}

type tile struct {
	id int
	s  []string
}

func (t *tile) edges() []string {
	var top, bottom, left, right []byte
	for i := 0; i < 10; i++ {
		top = append(top, t.s[0][i])
		bottom = append(bottom, t.s[9][i])
		left = append(left, t.s[i][0])
		right = append(right, t.s[i][9])
	}
	return []string{
		string(top),
		string(right),
		string(bottom),
		string(left),
		reverse(string(top)),
		reverse(string(right)),
		reverse(string(bottom)),
		reverse(string(left)),
	}
}

func (t *tile) rotateRight() tile {
	newTile := tile{
		id: t.id,
	}
	for i := 0; i < 10; i++ {
		var b bytes.Buffer
		for j := 9; j >= 0; j-- {
			b.WriteByte(t.s[i][j])
		}
		newTile.s = append(newTile.s, b.String())
	}
	return newTile
}

func (t *tile) flip() tile {
	newTile := tile{
		id: t.id,
	}
	for _, s := range t.s {
		newTile.s = append(newTile.s, reverse(s))
	}
	return newTile
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func getCorners(tiles []*tile) ([]*tile, map[string][]*tile) {
	// map and count matching edges
	edgesMatch := map[string][]*tile{}
	for _, t := range tiles {
		for _, v := range t.edges() {
			edgesMatch[v] = append(edgesMatch[v], t)
		}
	}

	// count per tile, how many edges matched
	tileEdgesMatch := map[*tile]int{}
	for _, t := range tiles {
		for _, s := range t.edges() {
			if i, ok := edgesMatch[s]; ok && len(i) > 1 {
				tileEdgesMatch[t]++
			}
		}
	}

	var c []*tile
	for t, v := range tileEdgesMatch {
		// corner tiles only have 2 edges match with the other
		if v == 4 { // 4 instead of 2 because it matched the flipped one as well
			c = append(c, t)
		}
	}
	return c, edgesMatch
}

func parse(f io.Reader) []*tile {
	input, err := ioutil.ReadAll(f)
	aoc.NoError(err)

	var tiles []*tile

	for _, t := range strings.Split(string(input), "\n\n") {
		if t == "" {
			continue
		}
		var id int
		_, err := fmt.Sscanf(t, "Tile %d:", &id)
		aoc.NoError(err)
		currentTile := tile{
			id: id,
			s:  []string{},
		}
		for i, s := range strings.Split(t, "\n") {
			if i == 0 {
				continue
			}
			currentTile.s = append(currentTile.s, s)
		}
		tiles = append(tiles, &currentTile)
	}
	return tiles
}
