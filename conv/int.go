package conv

import "strconv"

//nolint:funlen,gocognit,cyclop,gocyclo
/*
Int converts the given value to int.
Should be used in a known environment, when values are expected to be correct.
Panics in case of failure.

Usage:

	conv.Int("123") // 123
	conv.Int("asd") // panic!
*/
func Int(val any) int {
	switch val := val.(type) {
	case bool:
		if val {
			return 1
		}
		// Return int representation
		return 0
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case *int:
		if val == nil {
			return 0
		}
		// Return int representation
		return *val
	case *int8:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *int16:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *int32:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *int64:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *uint:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *uint8:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *uint16:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *uint32:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *uint64:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)

	case *float32:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case *float64:
		if val == nil {
			return 0
		}
		// Return int representation
		return int(*val)
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		// Return int representation
		return i
	case *string:
		if val == nil {
			return 0
		}
		// Convert string to int
		i, err := strconv.Atoi(*val)
		if err != nil {
			panic(err)
		}
		// Return int representation
		return i
	case nil:
		return 0
	default:
		panic("unknown type for Int()")
	}
}
