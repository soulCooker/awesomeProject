package tmp

import "fmt"

//示例 1：
//
//输入：nums = [1,1,1], k = 2
//输出：2
//示例 2：
//
//输入：nums = [1,2,3], k = 3
//sums[1,3,6]
//
//输出：2
//

func subarraySum3(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func subarraySum(nums []int, k int) int {
	sumArr := make([]int, len(nums))
	bucket := make(map[int]int)

	count := 0
	sum := 0
	for i, num := range nums {
		sum += num
		sumArr[i] = sum
		bucket[sum]++
	}

	// fmt.Println(bucket)

	sum = 0
	for i := 0; i < len(nums); i++ {
		fmt.Println("sum:", sum)
		if i == 0 {
			if bucket[k] > 0 {
				count += bucket[k]
			}
			// fmt.Println(i, ":", k, ":", bucket[k])
		} else {
			bucket[sum]--
			if bucket[k+sumArr[i-1]] > 0 {
				count += bucket[k+sumArr[i-1]]
			}
			// fmt.Println(i, ":", k+sumArr[i-1], ":", bucket[k+sumArr[i-1]])
		}
		sum += nums[i]
	}

	return count
}

func subarraySum2(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}
