package day21

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"

	"github.com/yulrizka/rxscan"
)

var rxIngridiets = regexp.MustCompile(`(.+) \(contains (.+)\)`)

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	imap := map[string][][]string{}
	for s.Scan() {
		var is, as string
		n, err := rxscan.Scan(rxIngridiets, s.Text(), &is, &as)
		aoc.NoError(err)
		if n == 0 {
			panic("not parsed")
		}
		ingridients := strings.Split(is, " ")
		alergens := strings.Split(as, ", ")
		for _, a := range alergens {
			imap[a] = append(imap[a], ingridients)
		}
	}

	freqMap := map[string]map[string]int{}
	for alergen, slices := range imap {
		m := map[string]int{}
		for _, ss := range slices {
			for _, s := range ss {
				m[s]++
			}
		}
		m2 := map[string]int{}
		for k, v := range m {
			if v <= 5 {
				continue
			}
			m2[k] = v
		}
		freqMap[alergen] = m2
	}

	return "", nil
}

func Part2(f io.Reader) (string, error) {
	return "", nil
}
