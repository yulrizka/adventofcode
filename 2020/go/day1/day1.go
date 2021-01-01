package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part1(f io.Reader) (string, error) {
	var numbers []int64
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			return "", err
		}
		numbers = append(numbers, i)
	}

	wantSum := int64(2020)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if a, b := numbers[i], numbers[j]; a+b == wantSum {
				ans := a * b
				return strconv.FormatInt(ans, 10), nil
			}
		}
	}

	return "", fmt.Errorf("not found")
}

func Part2(f io.Reader) (string, error) {
	var numbers []int64
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			return "", err
		}
		numbers = append(numbers, i)
	}

	wantSum := int64(2020)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if i == j || i == k || j == k {
					continue
				}
				if a, b, c := numbers[i], numbers[j], numbers[k]; a+b+c == wantSum {
					//fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
					return strconv.FormatInt(a*b*c, 10), nil
				}
			}
		}
	}

	return "", fmt.Errorf("not found")
}
