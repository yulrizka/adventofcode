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

	// parse other ticket and sum the invalid entries
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
	raw, err := ioutil.ReadAll(f)
	aoc.NoError(err)
	input := strings.Split(string(raw), "\n")

	const (
		numFields         = 20
		endRuleLine       = 20
		myTicketLine      = 22
		startNearbyTicket = 25
	)

	// parse rule
	var rules []rule
	validNum := map[int]struct{}{}
	for i := 0; i < endRuleLine; i++ {
		line := input[i]
		var r rule
		n, err := rxscan.Scan(rxRule, line, &r.name, &r.aMin, &r.aMax, &r.bMin, &r.bMax)
		aoc.NoError(err)
		if n != 5 {
			log.Fatalf("want 5 got %d", n)
		}
		rules = append(rules, r)
		for i := r.aMin; i < r.aMax; i++ {
			validNum[i] = struct{}{}
		}
		for i := r.bMin; i < r.bMax; i++ {
			validNum[i] = struct{}{}
		}

	}

	// parse ticket
	var tickets [][]int
	for i := startNearbyTicket; i < len(input); i++ {
		if ticket := ints(input[i], validNum); ticket != nil {
			tickets = append(tickets, ticket)
		}
	}

	// fields represent columns that contains all the possible rule
	// the idea is eliminate invalid rule for each field by applying the value for each rule
	fields := make([]map[rule]bool, numFields)
	for i := 0; i < len(fields); i++ {
		m := map[rule]bool{}
		for _, r := range rules {
			m[r] = true
		}
		fields[i] = m
	}

	// process tickets
	for col := 0; col < numFields; col++ {
		for row := 0; row < len(tickets); row++ {
			v := tickets[row][col]
			for r, valid := range fields[col] {
				if !valid {
					continue
				}
				if !r.valid(v) {
					fields[col][r] = false
				}
			}
		}
	}

	// at this point there exist a column with only 1 valid rule. we record the position and mark
	// the same rule on other fields as false. In the end we'll end up with one rule per field
	found := 0
	cols := make([]rule, numFields)

loop:
	for {
		if found == numFields {
			break
		}

		// find fields that only has one valid rule
		for i := 0; i < len(fields); i++ {
			var sumValid int
			var matchRule rule

			for r, valid := range fields[i] {
				if valid {
					matchRule = r
					sumValid++
				}
				if sumValid > 1 {
					continue
				}
			}
			if sumValid != 1 {
				continue
			}

			// found field where only one value is true
			cols[i] = matchRule

			// mark this rule as not valid for other fields
			for _, filed := range fields {
				filed[matchRule] = false
			}
			found++
			continue loop
		}
	}

	myTicket := ints(input[myTicketLine], validNum)
	if myTicket == nil {
		panic("nil ticket")
	}
	mul := 1
	for i, col := range cols {
		if strings.HasPrefix(col.name, "departure") {
			mul *= myTicket[i]
		}
	}

	return strconv.Itoa(mul), nil
}

type rule struct {
	name                   string
	aMin, aMax, bMin, bMax int
}

func (r rule) valid(i int) bool {
	valid := (r.aMin <= i && i <= r.aMax) || (r.bMin <= i && i <= r.bMax)
	return valid
}

func ints(s string, valid map[int]struct{}) (ints []int) {
	for _, v := range strings.Split(s, ",") {
		num, err := strconv.Atoi(v)
		if _, ok := valid[num]; !ok {
			return nil
		}
		aoc.NoError(err)
		ints = append(ints, num)
	}
	return
}
