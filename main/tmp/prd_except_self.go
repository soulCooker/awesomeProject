package main

import "fmt"

/*
*
输入: nums = [1,2,3,4]
输出: [24,12,8,6]
l[0] = 1
r[nums.len-1] = 1
l[i]=l[i-1]*nums[i-1]
r[i-1]=r[i]*nums[i]
mul[i] = l[i]*r[i]
*/

func main() {
	res := productExceptSelf([]int{1, 2, 3, 4})

	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))

	if len(nums) == 1 {
		return nums
	}

	res = append(res, 1)
	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]*nums[i-1])
	}

	var rightMul = 1

	for i := len(nums) - 2; i >= 0; i-- {
		rightMul = rightMul * nums[i+1]
		res[i] = rightMul * res[i]
	}
	return res
}
