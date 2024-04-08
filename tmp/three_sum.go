package tmp

import "sort"

func threeSum(nums []int) [][]int {
	numSet := map[int]int{}
	sort.Ints(nums)
	for i, num := range nums {
		numSet[num] = i
	}
	result := make([][]int, 0)
	var i = 0
	for i < len(nums)-2 {
		var j = i + 1
		if i-1 >= 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		for j < len(nums) {
			//k := numSet[-nums[i]-nums[j]]

			if j-1 > i && nums[j] == nums[j-1] {
				j++
				continue
			}
			k := numSet[nums[i]+nums[j]]

			if k > j {
				if len(result) > 0 {
					last := result[len(result)-1]
					if last[0] != nums[i] ||
						last[1] != nums[j] ||
						last[2] != nums[k] {
						result = append(result, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
			j++
		}
		i++
	}
	return result
}
