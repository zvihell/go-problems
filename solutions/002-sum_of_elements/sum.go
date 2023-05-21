package sumofelements

import "fmt"

func main() {
	a := []int{20, 5, 5}
	fmt.Println(sum_of_elements(a))
}

func sum_of_elements(array []int) int {
	sum := 0

	for _, item := range array {
		sum += item
	}
	return sum
}
