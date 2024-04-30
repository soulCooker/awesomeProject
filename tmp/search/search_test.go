package search

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	input := []int{1, 1, 2, 2, 2, 4}
	res := searchRange(input, 2)
	fmt.Println("res:", res)

	insert := searchInsert(input, 2)
	fmt.Println(insert)
}

func TestSearch(t *testing.T) {
	res := search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)
	fmt.Println(res)
}

func TestMedianFind(m *testing.T) {
	res := findMedianSortedArrays([]int{}, []int{1})
	fmt.Println(res)
}
