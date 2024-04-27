package backtracking

func subsets(nums []int) [][]int {
	return backtracking(nums, 0)
}

func backtracking(nums []int, start int) [][]int {
	if start == len(nums) {
		return [][]int{{}}
	}
	res := backtracking(nums, start+1)
	newRes := make([][]int, 0)
	for _, v := range res {
		sets := v
		newRes = append(newRes, sets)
		inSets := make([]int, 0)
		inSets = append(inSets, sets...)
		inSets = append(inSets, nums[start])
		newRes = append(newRes, inSets)
	}
	return newRes
}
