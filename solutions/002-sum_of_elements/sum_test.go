package sumofelements

import "testing"

func TestSumofElements(t *testing.T) {
	array := []int{2, 2, 4}
	res := 8

	realRes := sum_of_elements(array)
	if realRes != res {
		t.Error(realRes)
	}

}
