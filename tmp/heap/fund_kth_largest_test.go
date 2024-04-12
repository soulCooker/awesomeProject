package heap

import (
	"fmt"
	"testing"
	"time"
)

var arr = make([]int, 0)

var arr2 = []int{1, 1, 1, 1, 1}

func TestFindKthLargest(t *testing.T) {
	start := time.Now()
	//val := findKthLargest(arr, 5000)

	//for i := 0; i < 1000000; i++ {
	//	arr = append(arr, i)
	//}

	val := quickselect(arr2, 0, len(arr2)-1, 4)
	fmt.Println(val, time.Now().Sub(start))
	//nums := []int{3, 4, 6, 5}
	//s := quickSort(nums, 0, 4, func(val1, val2 int) int {
	//	return val1 - val2
	//})
	//fmt.Println(nums, s)
}
