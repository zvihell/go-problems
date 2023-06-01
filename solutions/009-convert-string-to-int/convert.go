package convertstring

import (
	"strconv"
)

func convert_string_2_int(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
