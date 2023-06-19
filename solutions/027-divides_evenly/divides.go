package dividesevenly

import "errors"

func divides_evenly(a, b int) (bool, error) {
	if b == 0 {
		return false, errors.New("divide by zero is not allowed")
	}
	return a%b == 0, nil
}
