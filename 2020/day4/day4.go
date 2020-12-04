package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type entry map[string]string

func validatePart1(e entry) bool {
	mustHave := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, s := range mustHave {
		if _, ok := e[s]; !ok {
			return false
		}
	}
	return true
}

func numValid(s string, min, max int64) bool {
	if s == "" {
		return false
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("invalid number %s", s)
		return false
	}

	return i >= min && i <= max
}

var (
	validEyeColor = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	validPid       = regexp.MustCompile(`^\d{9}$`)
	validHairColor = regexp.MustCompile(`^#[0-9a-f]{6}$`)
)

func validatePart2(e entry) bool {
	if !numValid(e["byr"], 1920, 2002) {
		return false
	}
	if !numValid(e["iyr"], 2010, 2020) {
		return false
	}
	if !numValid(e["eyr"], 2020, 2030) {
		return false
	}
	// parse height
	height := e["hgt"]
	if strings.HasSuffix(height, "cm") {
		height = height[:len(height)-2]
		if !numValid(height, 150, 193) {
			return false
		}
	} else if strings.HasSuffix(height, "in") {
		height = height[:len(height)-2]
		if !numValid(height, 59, 76) {
			return false
		}
	} else {
		return false
	}

	// hair color
	if !validHairColor.MatchString(e["hcl"]) {
		return false
	}

	// eye color
	if !validEyeColor[e["ecl"]] {
		return false
	}

	if !validPid.MatchString(e["pid"]) {
		return false
	}

	return true
}

func check(f io.Reader, checker func(entry) bool) (string, error) {
	s := bufio.NewScanner(f)

	e := entry{}
	var line int
	var valid int64
	for s.Scan() {
		line++
		text := s.Text()
		for _, field := range strings.Fields(text) {
			split := strings.Split(field, ":")
			if len(split) != 2 {
				return "", fmt.Errorf("unexpexcted len split %d %q", len(split), field)
			}
			e[split[0]] = split[1]
		}

		if text == "" {
			// empty line evaluate previous
			if checker(e) {
				valid++
			}
			e = entry{}
		}
	}
	// check the last element
	if checker(e) {
		valid++
	}

	return strconv.FormatInt(valid, 10), nil

}

func Part1(f io.Reader) (string, error) {
	return check(f, validatePart1)
}

func Part2(f io.Reader) (string, error) {
	return check(f, validatePart2)
}
