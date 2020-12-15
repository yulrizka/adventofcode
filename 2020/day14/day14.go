package day14

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/rxscan"
)

var (
	rxMask = regexp.MustCompile(`mask = (.+)`)
	rxMem  = regexp.MustCompile(`mem\[(\d+)] = (\d+)`)
)

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)

	var (
		mem  = map[uint]uint{}
		mask string
	)
	for s.Scan() {
		var maskStr string
		n, err := rxscan.Scan(rxMask, s.Text(), &maskStr)
		if err != nil {
			panic(err)
		}
		if n > 0 {
			mask = maskStr
			continue
		}

		// parsing memory assignment
		var addr, value uint
		n, err = rxscan.Scan(rxMem, s.Text(), &addr, &value)
		if err != nil || n == 0 {
			log.Fatalf("error:%v n:%d", err, n)
		}
		for i, v := range mask {
			if v == 'X' {
				continue
			}
			pos := len(mask) - i - 1
			if v == '1' {
				value |= 1 << pos
			} else {
				value &^= 1 << pos
			}
		}
		mem[addr] = value
	}
	var sum uint
	for _, v := range mem {
		sum += v
	}

	return strconv.Itoa(int(sum)), nil
}

func Part2(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)

	var (
		mem  = map[uint]uint{}
		mask string
	)
	for s.Scan() {
		var maskStr string
		n, err := rxscan.Scan(rxMask, s.Text(), &maskStr)
		if err != nil {
			panic(err)
		}
		if n > 0 {
			mask = maskStr
			continue
		}

		// parsing memory assignment
		var addr, value uint
		n, err = rxscan.Scan(rxMem, s.Text(), &addr, &value)
		if err != nil || n == 0 {
			log.Fatalf("error:%v n:%d", err, n)
		}

		setMemory(mem, mask, addr, value)
	}
	var sum uint
	for _, v := range mem {
		sum += v
	}

	return strconv.Itoa(int(sum)), nil
}

func setMemory(mem map[uint]uint, maskStr string, addr uint, v uint) {
	i := strings.Index(maskStr, "X")

	if i == -1 {
		// mask is complete
		mask, err := strconv.ParseUint(maskStr, 2, 36)
		if err != nil {
			panic(err)
		}
		addr |= uint(mask)
		mem[addr] = v
		return
	}

	// set address bit at pos to be '0' since the OR on the base will depend on the mas bit
	pos := len(maskStr) - i - 1 // little endian
	addr &^= 1 << pos
	newMask := []byte(maskStr)

	// X = 0
	newMask[i] = '0'
	setMemory(mem, string(newMask), addr, v)

	// X = 1
	newMask[i] = '1'
	setMemory(mem, string(newMask), addr, v)
}
