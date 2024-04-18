package heap

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{1}
	quickSort(nums, func(val1, val2 int) int {
		return val1 - val2
	})
	fmt.Println(nums)
}
