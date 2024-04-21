package heap

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	input := []int{1, 2}
	res := topKFrequent(input, 2)
	fmt.Println("res:", res)
}
