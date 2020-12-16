package day16

import (
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/yulrizka/rxscan"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

var rxRule = regexp.MustCompile(`([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)`)

func Part1(f io.Reader) (string, error) {
	raw, err := ioutil.ReadAll(f)
	aoc.NoError(err)

	input := strings.Split(string(raw), "\n")

	// parse rule
	valid := map[int]struct{}{}
	for i := 0; i < 20; i++ {
		line := input[i]
		var name string
		var aMin, aMax, bMin, bMax int
		n, err := rxscan.Scan(rxRule, line, &name, &aMin, &aMax, &bMin, &bMax)
		aoc.NoError(err)
		if n != 5 {
			log.Fatalf("want 5 got %d", n)
		}
		for i := aMin; i < aMax; i++ {
			valid[i] = struct{}{}
		}
		for i := bMin; i < bMax; i++ {
			valid[i] = struct{}{}
		}
	}

	// parse other ticket
	var errRate int
	for i := 25; i < 270; i++ {
		line := input[i]
		for _, v := range strings.Split(line, ",") {
			num, err := strconv.Atoi(v)
			aoc.NoError(err)
			if _, ok := valid[num]; !ok {
				errRate += num
			}
		}
	}

	return strconv.Itoa(errRate), nil
}

func Part2(f io.Reader) (string, error) {
	return "", nil
}
