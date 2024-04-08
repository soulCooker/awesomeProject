package tmp

import "fmt"

func test() {
	nums := []int{1, 2}
	// 2,1 1,2 2,1
	// 1,2,3,4

	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	var (
		tmp int
	)

	if len(nums) <= k {
		k = k % len(nums)
	}

	i := 0
	j := len(nums) - 1

	for i < j {
		tmp = nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}

	i = 0
	j = k - 1
	for i < j {
		tmp = nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}

	i = k
	j = len(nums) - 1
	for i < j {
		tmp = nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}
}
