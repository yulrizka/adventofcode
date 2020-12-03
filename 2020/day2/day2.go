package day1

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	rx := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	var i, valid int64
	for s.Scan() {
		i++
		m := rx.FindStringSubmatch(s.Text())
		if len(m) < 2 {
			return "", fmt.Errorf("regex not match on line %d, %q", i, s.Text())
		}

		minStr := m[1]
		maxStr := m[2]
		chr := m[3]
		text := m[4]

		min, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
		    return "", err
		}

		max, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return "", err
		}

		occurance := int64(strings.Count(text, chr))
		if min <= occurance && occurance <= max {
			valid++
		}
	}

	return strconv.FormatInt(valid, 10), nil
}


func Part2(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	rx := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	var i, valid int64
	for s.Scan() {
		i++
		m := rx.FindStringSubmatch(s.Text())
		if len(m) < 2 {
			return "", fmt.Errorf("regex not match on line %d, %q", i, s.Text())
		}

		minStr := m[1]
		maxStr := m[2]
		chr := m[3][0]
		text := []byte(m[4])

		pos1, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return "", err
		}

		pos2, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return "", err
		}
		if pos1 == 0 || pos2 == 0 {
			return "", fmt.Errorf("unexpected 0 position on line %d, %s", i, s.Text())
		}
		// make it 0 based index
		pos1--
		pos2--

		var match = 0
		if text[pos1] == chr {
			match++
		}
		if text[pos2] == chr {
			match++
		}
		if match == 1 {
			valid++
		}
	}

	return strconv.FormatInt(valid, 10), nil
}
