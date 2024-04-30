package search

func searchInsert(nums []int, target int) int {
	var (
		start = 0
		end   = len(nums)
	)

	for start < end {
		mid := (end + start) / 2
		if target < nums[mid] {
			end = mid
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}

	return start
}
