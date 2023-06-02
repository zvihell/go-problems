package issafebridge

import (
	"strings"
)

func is_safe_bridge(bridge string) bool {
	if len(bridge) == 0 {
		return false
	}
	if strings.Contains(bridge, " ") {
		return false
	}
	return true
}
