package backtracking

/**
空间节省的做法。
在源数组中构造排列，将选中的值移动到当前排列的末端，然后求后续剩余元素的全排列，接受后再将选择元素返回到原来位置。
不会在子递归中重复处理元素
**/

func permute(nums []int) [][]int {
	numSet := make(map[int]interface{})
	for _, v := range nums {
		numSet[v] = 1
	}
	return doPermute(numSet, nums)
}

func doPermute(numSet map[int]interface{}, nums []int) [][]int {
	res := make([][]int, 0)
	for _, num := range nums {
		if numSet[num] == 0 {
			continue
		}
		numSet[num] = 0

		subPremutation := doPermute(numSet, nums)
		for _, p := range subPremutation {
			p = append(p, num)
			res = append(res, p)
		}

		numSet[num] = 1
	}
	if len(res) == 0 {
		res = append(res, []int{})
	}
	return res
}
