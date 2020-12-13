package day13

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func read(f io.Reader) (earliest int, buses []int) {
	s := bufio.NewScanner(f)
	s.Scan()
	earliest, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	s.Scan()
	for _, v := range strings.Split(s.Text(), ",") {
		if v == "" || v == "x" {
			buses = append(buses, 0)
			continue
		}
		id, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		buses = append(buses, id)
	}

	return earliest, buses
}

func Part1(f io.Reader) (string, error) {
	earliest, buses := read(f)
	ts := earliest
	for {
		for _, b := range buses {
			if b == 0 {
				continue // x
			}
			if ts%b == 0 {
				ans := (ts - earliest) * b
				return strconv.Itoa(ans), nil
			}
		}
		ts++
	}
}

func Part2(f io.Reader) (string, error) {
	_, buses := read(f)

	busesSort := make([]int, len(buses))
	copy(busesSort, buses)
	sort.Sort(sort.Reverse(sort.IntSlice(busesSort)))

	max := busesSort[0]
	var maxI int

	offsets := map[int]int{}
	for i := 0; i < len(buses); i++ {
		b := buses[i]
		if b == max {
			maxI = i
			// find ts offset from the max bus id to the left
			for l := i - 1; l >= 0; l-- {
				bl := buses[l]
				if bl > 0 {
					offsets[bl] = l - i
				}
			}
			// find ts offset from the max bus id to the right
			for r := i + 1; r < len(buses); r++ {
				bl := buses[r]
				if bl > 0 {
					offsets[bl] = r - i
				}
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	search := func(ctx context.Context, ts int) int {
		var i int
		for {
			select {
			case <-ctx.Done():
				return 0
			default:
			}
			found := true
			for _, b := range busesSort[1:] {
				if b == 0 {
					break
				}

				offset := offsets[b]
				leaveAt := ts + offset
				if leaveAt%b != 0 {
					found = false
					break
				}
			}
			if found {
				ans := ts - maxI
				return ans
			}
			ts += max
			if i%100000 == 0 {
				fmt.Printf("ts = %+v\n", ts) // TODO:for debugging
			}

		}

	}

	var ans int

	for w := 1; w < 16; w++ {
		w := w
		go func() {
			as := search(ctx, w*max)
			if as > 0 {
				ans = as
				fmt.Printf("ans = %+v\n", ans) // TODO:for debugging
			}
			cancel()
		}()
	}
	<-ctx.Done()

	return strconv.Itoa(ans), nil
}
