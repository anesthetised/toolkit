package toolkit

import "fmt"

func Must[T any](val1 T, val2 any) T {
	switch val3 := val2.(type) {
	case error:
		if val3 != nil {
			panic(val3)
		}
	case bool:
		if !val3 {
			panic(fmt.Errorf("second argument to Must as a bool must be true"))
		}
	case nil:
	default:
		panic(fmt.Errorf("invalid argument type to Must: %T", val3))
	}

	return val1
}
