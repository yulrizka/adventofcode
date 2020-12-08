package adventofcode

import (
	"reflect"
	"regexp"
	"testing"

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
