// The ugliest code for this aoc so far but I did not give up!
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

// Part2 assemble the pieces, clean up border and find the monster
func Part2(f io.Reader) (string, error) {
	tiles := parse(f)
	corners, edgeMatches := getCorners(tiles)

	// pick a corner and
	corner := corners[0]
	m := map[*tile]struct{}{}
	for _, t := range tiles {
		if t == corner {
			continue
		}
		m[t] = struct{}{}
	}
	picture := [][]*tile{
		{corner},
	}

	// align orientation so that top and left has no match
	for len(edgeMatches[corner.top()]) > 1 {
		corner.rotateRight()
	}
	for len(edgeMatches[corner.left()]) > 1 {
		corner.flip()
	}
	picture = assemble(picture, m, 0, 0)

	for _, row := range picture {
		for _, t := range row {
			t.cleanBorder()
		}
	}

	var clean []string
	for _, row := range picture {
		var c []string
		for i, t := range row {
			if i == 0 {
				c = t.s
				continue
			}
			c = combine(c, t.s)
		}
		clean = append(clean, c...)
	}

	monsterStr := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `
	monster := strings.Split(monsterStr, "\n")

	var sum int
	clean = flip(clean)
	for sum == 0 {
		for k := 0; k < 4; k++ {
			clean = rotateRight(clean)
			for i := 0; i < 96; i++ {
				for j := 0; j < 96; j++ {
					if monsterExists(monster, clean, i, j) {
						sum++
					}
				}
			}
		}
	}

	fill := strings.Count(strings.Join(clean, "\n"), "#")
	ans := fill - (sum * strings.Count(monsterStr, "#"))

	return strconv.Itoa(ans), nil
}

func monsterExists(monster []string, clean []string, i int, j int) bool {
	// y does not match
	if len(clean)-i < len(monster) {
		return false
	}

	// x does not match
	if len(clean[0])-j < len(monster[0]) {
		return false
	}

	for row, line := range monster {
		for col, v := range line {
			if v != '#' {
				continue
			}
			if clean[i+row][j+col] != '#' {
				return false
			}
		}
	}

	return true
}

// assemble picture by starting from the top row, find piece that match the right side.
// if there are no pieces matched the right side, start the next row from the left which
// the top should match with the bottom part of the pieces above
func assemble(picture [][]*tile, tiles map[*tile]struct{}, row int, col int) [][]*tile {
	// continue
	for len(tiles) != 0 {
		// find to the right
		ref := picture[row][col]
		found := false

	next:
		for t := range tiles {
			for i := 0; i < 4; i++ {
				t.rotateRight()
				if matchLeftRight(ref, t) {
					picture[row] = append(picture[row], t)
					delete(tiles, t)
					found = true
					break next
				}
			}

			t.flip()
			for i := 0; i < 4; i++ {
				t.rotateRight()
				if matchLeftRight(ref, t) {
					picture[row] = append(picture[row], t)
					delete(tiles, t)
					found = true
					break next
				}
			}
		}
		if found {
			col++
		} else {
			// to pieces match to the right, start the next row which top part matches the pieces above
			ref = picture[row][0]
			row += 1
			col = 0
			picture = append(picture, []*tile{})

		next2:
			for t := range tiles {
				for i := 0; i < 4; i++ {
					t.rotateRight()
					if matchTopBottom(ref, t) {
						picture[row] = append(picture[row], t)
						delete(tiles, t)
						found = true
						break next2
					}
				}

				t.flip()
				for i := 0; i < 4; i++ {
					t.rotateRight()
					if matchTopBottom(ref, t) {
						picture[row] = append(picture[row], t)
						delete(tiles, t)
						found = true
						break next2
					}
				}
			}
			if !found {
				panic("should not be here")
			}
		}
	}
	return picture
}

type tile struct {
	id int
	s  []string
}

func (t *tile) load(s string) {
	t.s = strings.Split(s, "\n")
}

// edges return the pieces edges including when it's flipped
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

func (t *tile) bottom() string {
	return reverse(t.s[9])
}

func (t *tile) top() string {
	return t.s[0]
}

func (t *tile) right() string {
	var b bytes.Buffer
	for i := 0; i < 10; i++ {
		b.WriteByte(t.s[i][9])
	}
	return b.String()
}

func (t *tile) left() string {
	var b bytes.Buffer
	for i := 9; i >= 0; i-- {
		b.WriteByte(t.s[i][0])
	}
	return b.String()
}

func (t *tile) rotateRight() {
	t.s = rotateRight(t.s)
}

func rotateRight(a []string) []string {
	var s []string
	for i := 0; i < len(a[0]); i++ {
		var b bytes.Buffer
		for j := len(a) - 1; j >= 0; j-- {
			b.WriteByte(a[j][i])
		}
		s = append(s, b.String())
	}
	return s
}

func (t *tile) flip() {
	t.s = flip(t.s)
}

func flip(a []string) []string {
	var n []string
	for _, s := range a {
		n = append(n, reverse(s))
	}
	return n
}

func (t *tile) print() {
	for _, v := range t.s {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("\n")
}

// cleanBorder removes outer part of the pieces
func (t *tile) cleanBorder() {
	var n []string
	for i, s := range t.s {
		if i == 0 || i == 9 {
			continue
		}
		n = append(n, s[1:9])
	}
	t.s = n
}

// matchLeftRight returns true if right side of left pieces matches with left side of right piece
func matchLeftRight(l *tile, r *tile) bool {
	for i := 0; i < 10; i++ {
		if l.s[i][9] != r.s[i][0] {
			return false
		}
	}
	return true
}

// matchTopBottom returns true if top part of bottom pieces match the bottom part of above piece
func matchTopBottom(t *tile, b *tile) bool {
	for i := 0; i < 10; i++ {
		if t.s[9][i] != b.s[0][i] {
			return false
		}
	}
	return true
}

// combine martix by appending r to l for each row
func combine(l []string, r []string) []string {
	var n []string
	for i := 0; i < len(l); i++ {
		n = append(n, l[i]+r[i])
	}
	return n
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

// getCorners by mapping and count each edge that match with each other
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
