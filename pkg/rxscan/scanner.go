package rxscan

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

func parse(match string, arg interface{}) (err error) {
	switch v := arg.(type) {
	case *bool:
		*v, err = strconv.ParseBool(match)
	//case *complex64:
	//vv, err := strconv.ParseComplex(match, 64)
	//if err != nil {
	//	return err
	//}
	//*v = complex64(vv)
	//case *complex128:
	//vv, err := strconv.ParseComplex(match, 128)
	//if err != nil {
	//	return err
	//}
	//*v = vv
	case *int:
		*v, err = strconv.Atoi(match)
	case *int8:
		vv, err := strconv.ParseInt(match, 10, 8)
		if err != nil {
			return err
		}
		*v = int8(vv)
	case *int16:
		vv, err := strconv.ParseInt(match, 10, 16)
		if err != nil {
			return err
		}
		*v = int16(vv)
	case *int32:
		vv, err := strconv.ParseInt(match, 10, 32)
		if err != nil {
			return err
		}
		*v = int32(vv)
	case *int64:
		vv, err := strconv.ParseInt(match, 10, 64)
		if err != nil {
			return err
		}
		*v = vv
	case *uint:
		vv, err := strconv.ParseUint(match, 10, 64)
		if err != nil {
			return err
		}
		*v = uint(vv)
	case *uint8:
		vv, err := strconv.ParseUint(match, 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(vv)
	case *uint16:
		vv, err := strconv.ParseUint(match, 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(vv)
	case *uint32:
		vv, err := strconv.ParseUint(match, 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(vv)
	case *uint64:
		vv, err := strconv.ParseUint(match, 10, 64)
		if err != nil {
			return err
		}
		*v = vv
	case *uintptr:
		err = errors.New("uintptr is not supported yet")
	case *float32:
		vv, err := strconv.ParseFloat(match, 32)
		if err != nil {
			return err
		}
		*v = float32(vv)
	case *float64:
		vv, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return err
		}
		*v = vv
	case *string:
		*v = match
	case *[]byte:
		*v = []byte(match)
	default:
		err = errors.New("can't scan type: " + reflect.TypeOf(arg).String())
	}

	return err
}

func Scan(re *regexp.Regexp, s string, args ...interface{}) (n int, err error) {
	matches := re.FindStringSubmatch(s)
	if len(matches) <= 1 {
		return 0, nil
	}

	if len(args) > len(matches)-1 {
		return 0, errors.New("got " + strconv.Itoa(len(args)) + " arguments for " + strconv.Itoa(len(matches)-1) + " matches")
	}

	for i, arg := range args {
		if arg == nil {
			continue
		}
		if err := parse(matches[i+1], arg); err != nil {
			return 0, err
		}
		n++
	}
	return n, err
}

type scanner struct {
	matches   [][]string
	i         int
	numParsed int
	args      []interface{}
	err       error
}

func (s *scanner) Error() error {
	return s.err
}

func (s *scanner) NumParsed() int {
	return s.numParsed
}

func NewScanner(re *regexp.Regexp, s string) *scanner {
	return &scanner{
		matches: re.FindAllStringSubmatch(s, -1),
	}
}

func (s *scanner) More() bool {
	return s.err == nil && s.i < len(s.matches)
}

func (s *scanner) Scan(args ...interface{}) bool {
	m := s.matches[s.i]
	if len(s.args) > len(m)-1 {
		s.err = errors.New("got " + strconv.Itoa(len(s.args)) + " arguments for " + strconv.Itoa(len(m)-1) + " matches")
		return false
	}

	if len(m) > 1 {
		parsed := 0
		for i, arg := range args {
			if arg == nil {
				continue
			}
			if s.err = parse(m[i+1], arg); s.err != nil {
				return false
			}
			parsed++
		}
		s.numParsed = parsed

	}

	s.i++

	return s.i <= len(s.matches)
}
