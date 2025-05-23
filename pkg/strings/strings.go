package strings

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseInt[T constraints.Signed](s string) (T, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return T(i), err
}

func ParseUint[T constraints.Unsigned](s string) (T, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	return T(i), err
}
