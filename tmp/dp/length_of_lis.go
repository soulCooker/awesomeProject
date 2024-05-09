package dp

/*
*
贪心算法+二分法查找
相同长度的子序列中，只需考虑末尾值最小的子序列
用dp[i]存储当前长度为i的子序列的末尾值
遍历所有数组元素，对比元素与dp的最大值

	大于最大值，写入dp[len+1]
	如果小于最大值，则通过二分法查找第一个比他小的数，位置为k，则更新dp[k+1]为当前值

返回dp的长度
*
*/
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)
	dp = append(dp, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			var (
				l int
				r = len(dp) - 1
			)
			found := false
			for l <= r { //考虑长度为1的dp，l=r的情况也需要处理
				mid := (l + r) / 2
				if dp[mid] == nums[i] {
					found = true
					break
				} else if nums[i] > dp[mid] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			if !found && dp[r+1] > nums[i] {
				dp[r+1] = nums[i]
			}
		}
	}
	return len(dp)
}
