package day8

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	nop = "nop"
	acc = "acc"
	jmp = "jmp"
)

func isTerminate(lines []string) (success bool, accumulator int64) {
	executed := map[int]struct{}{}

	success = true
	for i := 0; i < len(lines)-1; i++ {
		if _, ok := executed[i]; ok {
			success = false
			break
		}

		text := lines[i]
		fields := strings.Fields(text)
		if len(fields) != 2 {
			log.Fatalf("invalid instruction %q line %d", text, i)
		}
		op := fields[0]
		arg, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			panic("failed to parse arg")
		}

		switch op {
		case nop:
			// noop
		case acc:
			accumulator += arg
		case jmp:
			i += int(arg) - 1 // for will inc by +1
		}

		executed[i] = struct{}{}
	}

	return success, accumulator
}

func Part1(f io.Reader) (string, error) {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(content), "\n")

	_, accumulator := isTerminate(lines)

	return strconv.FormatInt(accumulator, 10), nil
}

func Part2(f io.Reader) (string, error) {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(content), "\n")

	check := func(src, dst string) (accumulator string, ok bool) {
		for i := 0; i < len(lines); i++ {
			if !strings.HasPrefix(lines[i], jmp) {
				continue
			}
			mod := make([]string, len(lines))
			copy(mod, lines)
			if len(mod) == 0 {
				panic("empty mod")
			}
			mod[i] = strings.ReplaceAll(mod[i], src, dst)

			terminate, accumulator := isTerminate(mod)
			if terminate {
				fmt.Printf("line %d %q -> %q\n", i, lines[i], mod[i])
				return strconv.FormatInt(accumulator, 10), true
			}
		}
		return "", false
	}

	// change nop -> jmp
	if acc, ok := check(jmp, nop); ok {
		return acc, nil
	}

	// change jmp -> nop
	if acc, ok := check(nop, jmp); ok {
		return acc, nil
	}

	return "", fmt.Errorf("answer not found")
}
