package main

import "fmt"

/**
medium：给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
输入：nums = [100,4,101,5,3,2,102,1]
输出：5
*/

func TestFindLongest() {
	res := findLongest([]int32{100, 4, 101, 5, 3, 2, 102, 1})
	fmt.Println(res)
}

func findLongest(nums []int32) int {
	if len(nums) == 0 {
		return 0
	}

	numsMap := make(map[int32]interface{})
	for _, num := range nums {
		numsMap[num] = 1
	}

	var res = 1

	for _, num := range nums {
		if _, ok := numsMap[num-1]; !ok {
			for i := num + 1; ; i++ {
				if _, ok := numsMap[i]; !ok {
					if int(i-num) > res {
						res = int(i - num)
					}
					break
				}
			}
		}
	}

	return res
}
