package tmp

import "fmt"

/**
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

*/

func TestMaxSubArray() {
	res := maxSubArray([]int{-2, -2})
	fmt.Println(res)
}

func maxSubArray(arr []int) int {
	var maxSum = 0
	var res = arr[0]

	for _, val := range arr {
		maxSum += val
		if maxSum > res {
			res = maxSum
		}
		if maxSum < 0 {
			maxSum = 0
		}
	}

	return res
}
