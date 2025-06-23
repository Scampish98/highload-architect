package strings

import (
	"strconv"
	"strings"

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

func Capitalize(s string) string {
	letters := strings.Split(strings.ToLower(s), "")
	letters[0] = strings.ToUpper(letters[0])

	return strings.Join(letters, "")
}
