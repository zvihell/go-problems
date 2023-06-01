package google

import (
	"strings"
)

func google(num int) string {
	count := strings.Repeat("o", num)
	return "G" + count + "gle"
}
