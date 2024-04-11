package heap

import "fmt"

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

func findKthLargest(nums []int, k int) int {
	var (
		i = 0
		m = len(nums)
	)
	k--
	for i < m {
		split := findSplit(nums, i, m, func(val1, val2 int) int {
			return val1 - val2
		})
		if split == k {
			return nums[split]
		} else if split < k {
			i = split + 1
		} else {
			m = split
		}
	}

	return nums[i]
}

func quickSort(nums []int, cmp func(val1, val2 int) int) {
	doQuickSort(nums, 0, len(nums), cmp)
}

func doQuickSort(nums []int, start, end int, cmp func(val1, val2 int) int) {
	if end-start <= 1 {
		return
	}
	split := findSplit(nums, start, end, cmp)
	doQuickSort(nums, start, split, cmp)
	doQuickSort(nums, split+1, end, cmp)
}

/*
*
1 8 7 2
*/
func findSplit(nums []int, start, end int, cmp func(val1, val2 int) int) int {
	var (
		i     = start + 1
		k     = end - 1
		split = start
	)
	for i <= k {
		if cmp(nums[split], nums[i]) <= 0 {
			swap(nums, i, split)
			split = i
			i++
		} else {
			swap(nums, i, k)
			k--
		}
		fmt.Println("i=", i, ",k=", k, "nums=", nums)
	}

	return split
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
