package main

func moveZeroes(nums []int) {
	var (
		zeroStart = -1
	)

	for i, num := range nums {
		if num != 0 { //提交为不等于0
			if zeroStart >= 0 {
				nums[zeroStart] = num
				nums[i] = 0
				zeroStart++
			}
		} else if zeroStart < 0 {
			zeroStart = i
		}
	}
}
