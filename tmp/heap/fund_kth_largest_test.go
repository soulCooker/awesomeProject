package heap

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	val := findKthLargest([]int{3, 4, 6, 5}, 2)
	fmt.Println(val)
	//nums := []int{3, 4, 6, 5}
	//s := quickSort(nums, 0, 4, func(val1, val2 int) int {
	//	return val1 - val2
	//})
	//fmt.Println(nums, s)
}
