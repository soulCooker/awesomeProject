package search

/*
*
更新当前最小值为当前值与mid索引值的最小值
左区间 [start, mid-1] 右区间[mid+1, end]
对于L ， M1｜ M2，比mid小的数在M1中，取左区间
对于 L1 | L2, M，对mid小的数在M中，取右区间

*
*/
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		start = 0
		end   = len(nums) - 1
		res   = nums[0]
	)

	for start <= end {
		//更新当前最小值为当前值与mid索引值的最小值
		mid := (start + end) / 2
		if res > nums[mid] {
			res = nums[mid]
		}
		//左区间 [start, mid-1] 右区间[mid+1, end]

		if nums[mid] < nums[end] { //对于L ， M1｜ M2，比mid小的数在M1中，取左区间
			end = mid - 1
		} else { //对于 L1 | L2, M，对mid小的数在M中，取右区间
			start = mid + 1
		}
	}
	return res
}
