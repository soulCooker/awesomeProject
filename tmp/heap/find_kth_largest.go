package heap

import (
	"fmt"
	"math/rand"
)

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
		split, isSwap := findSplit(nums, i, m, func(val1, val2 int) int {
			return val1 - val2
		})

		fmt.Println("split=", split)

		if split == k {
			return nums[split]
		} else if split > k {
			if !isSwap {
				return nums[k]
			}
			m = split
		} else {
			i = split
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
	split, isSwap := findSplit(nums, start, end, cmp)
	if !isSwap {
		return
	}
	doQuickSort(nums, start, split, cmp)
	doQuickSort(nums, split+1, end, cmp)
}

/*
*
双指针法
*/
func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	newL := rand.Intn(r - l)
	swap(nums, l, l+newL)
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

func findSplit2(nums []int, start, end int, cmp func(val1, val2 int) int) (int, bool) {
	var (
		i     = start + 1
		k     = end - 1
		split = start
	)

	r := rand.Intn(end - start)
	swap(nums, split, start+r)

	isSwap := false
	for i <= k {
		diff := cmp(nums[split], nums[i])
		if diff == 0 {
			split = i
			i++
		} else if diff < 0 {
			swap(nums, i, split)
			isSwap = true
			split = i
			i++
		} else {
			swap(nums, i, k)
			isSwap = true
			k--
		}
	}
	return split, isSwap
}

/*
*
1 8 7 2
*/
func findSplit(nums []int, start, end int, cmp func(val1, val2 int) int) (int, bool) {
	var (
		i     = start + 1
		k     = end - 1
		split = start
	)

	r := rand.Intn(end - start)
	swap(nums, split, start+r)

	isSwap := false
	for i <= k {
		diff := cmp(nums[split], nums[i])
		if diff == 0 {
			split = i
			i++
		} else if diff < 0 {
			swap(nums, i, split)
			isSwap = true
			split = i
			i++
		} else {
			swap(nums, i, k)
			isSwap = true
			k--
		}
	}
	return split, isSwap
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
