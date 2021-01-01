package day19

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	rules, input := parse(f)

	sum := 0
	for _, s := range input {
		if valid(s, rules) {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

func Part2(f io.Reader) (string, error) {
	rules, input := parse(f)

	sum := 0
	for _, s := range input {
		if valid(s, rules) {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

// valid validates string for a given rule. If there is reminder with an empty string, means that
// there is rule where all character are valid
func valid(s string, rules map[string]rule) bool {
	if len(s) == 0 {
		return false
	}
	ok, remainder := validate([]string{s}, "0", rules)
	if ok {
		for _, r := range remainder {
			if r == "" {
				return true
			}
		}
	}
	return false
}

const (
	ruleType = iota
	letterType
)

type rule struct {
	typ    byte
	v      string
	chains [][]string
}

func parse(f io.Reader) (rules map[string]rule, input []string) {
	raw, err := ioutil.ReadAll(f)
	aoc.NoError(err)

	parts := strings.Split(string(raw), "\n\n")

	rules = make(map[string]rule)
	for _, line := range strings.Split(parts[0], "\n") {
		parts := strings.Split(line, ": ")
		num, s := parts[0], parts[1]
		if strings.Contains(s, `"`) { // a or b
			char := strings.ReplaceAll(s, `"`, ``)
			rules[num] = rule{typ: letterType, v: string(char[0])}
		} else {
			r := rule{typ: ruleType}
			for _, list := range strings.Split(s, " | ") {
				r.chains = append(r.chains, strings.Fields(list))
			}
			rules[num] = r
		}
	}

	input = strings.Split(parts[1], "\n")

	return
}

func validate(ss []string, id string, rules map[string]rule) (valid bool, rest []string) {
	r := rules[id]
	// end of chain
	if r.typ == letterType {
		remainders := make([]string, 0)
		for _, s := range ss {
			if s != "" && string(s[0]) == r.v {
				remainders = append(remainders, s[1:])
			}
		}
		return len(remainders) > 0, remainders
	}

	hits := make([]string, 0)
loop:
	for _, rr := range r.chains {
		rem := ss
		for _, id := range rr {
			isValid, rest := validate(rem, id, rules)
			if !isValid {
				continue loop
			}
			rem = rest
		}
		hits = append(hits, rem...)
	}

	return len(hits) > 0, hits
}
