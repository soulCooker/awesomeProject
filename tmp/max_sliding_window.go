package tmp

import (
	"container/heap"
	"fmt"
)

/*
*
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]

heap

*/

func TestMaxSlidingWindow() {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	res := maxSlidingWindow(input, 3)
	fmt.Println(res)
}

var a []int

type hp struct {
	Index []int
}

func (h hp) Len() int {
	return len(h.Index)
}

func (h hp) Less(i, j int) bool {
	return a[h.Index[i]] > a[h.Index[j]]
}

func (h hp) Swap(i, j int) {
	tmp := h.Index[i]
	h.Index[i] = h.Index[j]
	h.Index[j] = tmp
}

func (h *hp) Push(x interface{}) {
	h.Index = append(h.Index, x.(int))
}

func (h *hp) Pop() interface{} {
	top := h.Index[len(h.Index)-1]
	h.Index = h.Index[0 : len(h.Index)-1]
	return top
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	res := make([]int, 0)

	maxHeap := &hp{Index: make([]int, 0)}
	for i := 0; i < k-1; i++ {
		maxHeap.Push(i)
	}
	heap.Init(maxHeap)
	for i := k - 1; i < len(nums); i++ {
		heap.Push(maxHeap, i)
		for maxHeap.Len() > 0 {
			maxIndex := maxHeap.Index[0]
			if maxIndex > i-k && maxIndex <= i {
				res = append(res, nums[maxIndex])
				break
			} else {
				heap.Pop(maxHeap)
			}
		}
	}

	return res
}
