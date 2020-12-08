package adventofcode

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Scan(re *regexp.Regexp, s string, args ...interface{}) (err error) {
	matches := re.FindStringSubmatch(s)
	if len(matches) <= 1 {
		return nil
	}

	if len(args) > len(matches)-1 {
		return fmt.Errorf("got %d arguments for %d matches", len(args), len(matches)-1)
	}

	for i, arg := range args {
		if arg == nil {
			continue
		}
		match := matches[i+1]
		switch v := arg.(type) {
		case *bool:
			*v, err = strconv.ParseBool(match)
		case *complex64:
			vv, err := strconv.ParseComplex(match, 64)
			if err != nil {
				return err
			}
			*v = complex64(vv)
		case *complex128:
			vv, err := strconv.ParseComplex(match, 128)
			if err != nil {
				return err
			}
			*v = vv
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
			//val := reflect.ValueOf(v)
			//ptr := val
			//if ptr.Kind() != reflect.Ptr {
			//	s.errorString("type not a pointer: " + val.Type().String())
			//	return
			//}
			//switch v := ptr.Elem(); v.Kind() {
			//case reflect.Bool:
			//	v.SetBool(s.scanBool(verb))
			//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			//	v.SetInt(s.scanInt(verb, v.Type().Bits()))
			//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			//	v.SetUint(s.scanUint(verb, v.Type().Bits()))
			//case reflect.String:
			//	v.SetString(s.convertString(verb))
			//case reflect.Slice:
			//	// For now, can only handle (renamed) []byte.
			//	typ := v.Type()
			//	if typ.Elem().Kind() != reflect.Uint8 {
			//		s.errorString("can't scan type: " + val.Type().String())
			//	}
			//	str := s.convertString(verb)
			//	v.Set(reflect.MakeSlice(typ, len(str), len(str)))
			//	for i := 0; i < len(str); i++ {
			//		v.Index(i).SetUint(uint64(str[i]))
			//	}
			//case reflect.Float32, reflect.Float64:
			//	s.SkipSpace()
			//	s.notEOF()
			//	v.SetFloat(s.convertFloat(s.floatToken(), v.Type().Bits()))
			//case reflect.Complex64, reflect.Complex128:
			//	v.SetComplex(s.scanComplex(verb, v.Type().Bits()))
			//default:
			//	s.errorString("can't scan type: " + val.Type().String())
			//}
		}
	}
	return err
}
