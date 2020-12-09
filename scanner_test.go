package adventofcode

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	var (
		b bool

		c64  complex64
		c128 complex128

		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64

		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64

		f32 float32
		f64 float64

		s    string
		byts []byte
	)
	err := Scan(regexp.MustCompile(`it is (\w+)$`), "it is true", &b)
	require.NoError(t, err)
	require.True(t, b)

	// boolean
	wantB := true
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is true", []interface{}{&b}, []interface{}{&wantB})
	wantB = false
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is false", []interface{}{&b}, []interface{}{&wantB})

	// complex
	wantC64 := complex64(3 + 5.5i)
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is (3.0+5.5i)", []interface{}{&c64}, []interface{}{&wantC64})
	wantC128 := 3 + 5.5i
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is (3.0+5.5i)", []interface{}{&c128}, []interface{}{&wantC128})

	// int
	wantI := 11
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 11", []interface{}{&i}, []interface{}{&wantI})
	wantI8 := int8(127)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 127", []interface{}{&i8}, []interface{}{&wantI8})
	wantI16 := int16(32767)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 32767", []interface{}{&i16}, []interface{}{&wantI16})
	wantI32 := int32(2147483647)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 2147483647", []interface{}{&i32}, []interface{}{&wantI32})
	wantI64 := int64(9223372036854775807)
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is 9223372036854775807", []interface{}{&i64}, []interface{}{&wantI64})

	// uint
	wantUI := uint(11)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 11", []interface{}{&ui}, []interface{}{&wantUI})
	wantUI8 := uint8(255)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 255", []interface{}{&ui8}, []interface{}{&wantUI8})
	wantUI16 := uint16(65535)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 65535", []interface{}{&ui16}, []interface{}{&wantUI16})
	wantUI32 := uint32(4294967295)
	ok(t, regexp.MustCompile(`it is (\w+)$`), "it is 4294967295", []interface{}{&ui32}, []interface{}{&wantUI32})
	wantUI64 := uint64(18446744073709551615)
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is 18446744073709551615", []interface{}{&ui64}, []interface{}{&wantUI64})

	// float
	wantFloat32 := float32(0.123456)
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is 0.123456", []interface{}{&f32}, []interface{}{&wantFloat32})
	wantFloat64 := 0.123456
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is 0.123456", []interface{}{&f64}, []interface{}{&wantFloat64})

	wantS := "some cool text"
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is some cool text", []interface{}{&s}, []interface{}{&wantS})
	wantByts := []byte("some cool text")
	ok(t, regexp.MustCompile(`it is (.+)$`), "it is some cool text", []interface{}{&byts}, []interface{}{&wantByts})
}

func ok(t *testing.T, re *regexp.Regexp, s string, args []interface{}, want []interface{}) {
	t.Helper()
	err := Scan(re, s, args...)
	require.NoError(t, err)

	for i, v := range want {
		require.EqualValuesf(t, v, args[i], `got "%v" want "%v"`, reflect.ValueOf(args[i]).Elem(), reflect.ValueOf(v).Elem())
	}
}

func TestScanAll(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	lines := strings.Split(input, "\n")

	var got [][]interface{}

	rx := regexp.MustCompile(`(\d+) (\w+ \w+) bag`)
	for _, l := range lines {
		var count int
		var color string
		sc := ScanAll(rx, l)
		for sc.Scan(&count, &color) {
			got = append(got, []interface{}{count, color})
		}
		assert.NoError(t, sc.err)
	}

	want := [][]interface{}{
		{1, "bright white"}, {2, "muted yellow"},
		{3, "bright white"}, {4, "muted yellow"},
		{1, "shiny gold"},
		{2, "shiny gold"}, {9, "faded blue"},
		{1, "dark olive"}, {2, "vibrant plum"},
		{3, "faded blue"}, {4, "dotted black"},
		{5, "faded blue"}, {6, "dotted black"},
	}

	assert.EqualValues(t, want, got)
}

var (
	text = "jmp +32"
	rx   = regexp.MustCompile(`(\w+) (\w+)`)
)

func withScan() {
	var op string
	var arg int64
	_, _ = fmt.Sscanf(text, "%s %d", &op, &arg)
}

func withRegex() {
	var op string
	var arg int64
	_ = Scan(rx, text, &op, &arg)
}

// $ benchstat scan.txt regex.txt
// name    old time/op  new time/op  delta
// Scan-8   108µs ± 1%   107µs ± 1%   ~     (p=0.222 n=5+5)
func BenchmarkScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if os.Getenv("TEST_SCAN") != "" {
			withScan()
		} else {
			withRegex()
		}
	}
}
