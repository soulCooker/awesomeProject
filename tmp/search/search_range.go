package search

/*
*
使用两对双指针法，每个指针都是用二分搜索，区别在于找到目标值时，
找最左边界，第二对双指针，总是更新右指针，如果是最左值，则返回，否则更新右指针为目标值下标
找最右边界，第一对双指针，总是更新左指针，如果是最右值则停止，否则左指针=为目标值下标+1

解法二：
两次二分查找，第一次查找target值的最小偏移量或其插入位置，第二次查找target+1的最小偏移量或其插入位置
*
*/

func searchRange2(nums []int, target int) []int {
	start := searchIndex(nums, target)

	//如果找到的是插入位置，返回空
	if len(nums) == start || nums[start] != target {
		return nil
	}

	return []int{start, searchIndex(nums, target+1) - 1}
}

func searchIndex(nums []int, target int) int {
	var (
		l = 0
		r = len(nums) - 1
	)

	//循环不变量 r的右边的数都大于等于target值，l左边的值都小于target值
	//停止条件 l>r，则r+1是大于等于target值中偏移量最小的一个
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r + 1
}

func searchRange(nums []int, target int) []int {
	var (
		lmStart int
		lmEnd   = len(nums)
		rmStart int
		rmEnd   = len(nums)
	)

	res := make([]int, 0)

	for lmStart < lmEnd {
		mid := (lmStart + lmEnd) / 2
		//循环不变量，end开始都是大于等于target的数，start之前，都是小于target的数
		if nums[mid] == target {
			if mid-1 > lmStart && nums[mid-1] == target {
				lmEnd = mid
			} else {
				lmStart = mid
				break
			}
		} else if nums[mid] > target {
			lmEnd = mid
		} else {
			lmStart = mid + 1
		}
	}
	if nums[lmStart] != target {
		return []int{-1, -1}
	}

	res = append(res, lmStart)

	for rmStart < rmEnd {
		mid := (rmStart + rmEnd) / 2
		if nums[mid] == target {
			if mid+1 < rmEnd && nums[mid+1] == target {
				rmStart = mid + 1
			} else {
				rmStart = mid
				break
			}
		} else if nums[mid] > target {
			rmEnd = mid
		} else {
			rmStart = mid + 1
		}
	}

	res = append(res, rmStart)

	return res
}
