package search

/**
二分法,取mid
将mid与target对比，相等停止

如果存在L，M 两部分有序数组，且 o(M) < o(L)
	将end与target对比，相等停止
	不需要考虑mid==end
	对于 mid > end, l1, l2, m的情况
		对于 target > mid
			取右区间
		对于 target < mid
			target < end
				取右区间
			target > end
				取左区间
	对于 mid < end, l, m1, m2的情况
		对于target < mid 在m1中
			取左区间
		target > mid
			target < end
				取右区间
			target > end
				取左区间

如果只有L’或M‘，那么变为正常的二分搜索

坑点：
1、如果待查找空间为整体有序时，要采用正常二分查找
2、注意比较的是下标所在值，不是下标

**/

func search(nums []int, target int) int {
	var (
		start int
		end   = len(nums) - 1
	)

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[start] < nums[end] {
			if target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target == nums[end] {
				return end
			}
			if nums[mid] > nums[end] {
				// 对于 mid > end, l1 ｜ l2, m的情况
				// 对于 target > mid
				if target > nums[mid] {
					// 	取右区间
					start = mid + 1
				} else {
					if target < nums[end] {
						//在m中，取右区间
						start = mid + 1
					} else {
						// 	target > end
						// 		取左区间
						end = mid - 1
					}
				}
			} else {
				// 对于 mid < end, l, m1 ｜ m2的情况
				if target < nums[mid] {
					// 对于target < mid 在m1中
					// 		取左区间
					end = mid - 1
				} else {
					// 	target < end
					// 		取右区间
					// 	target > end
					// 		取左区间
					if target < nums[end] {
						start = mid + 1
					} else {
						end = mid - 1
					}
				}

			}
		}
	}
	return -1
}
