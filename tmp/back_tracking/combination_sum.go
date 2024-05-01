package backtracking

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	var res = make([][]int, 0)
	//遍历候选数集合，对每个候选数a出现次数从0次开始，将target-候选数累加值sum(a)的target'，将a添加到结果集中， 尝试找候选集合-当前候选数可以组成target‘的组合的子问题的解
	backing(target, 0, candidates, []int{}, &res)

	return res
}

func backing(target int, index int, candidates []int, current []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, current)
		return
	}
	if index == len(candidates) {
		return
	}

	for target >= 0 {
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		backing(target, index+1, candidates, newCurrent, res)
		target -= candidates[index]
		current = append(current, candidates[index])
	}
}

// 优化后
func combinationSum2(candidates []int, target int) [][]int {
	//排序候选集
	slices.SortFunc(candidates, func(a, b int) int {
		return a - b
	})

	res := make([][]int, 0)
	path := make([]int, 0)

	//声明递归函数
	var dfs func(sum int, start int)
	dfs = func(sum, start int) {
		//当和为目标值时，找到解
		if sum == target {
			ans := make([]int, len(path))
			copy(ans, path)
			res = append(res, ans)
		}
		//当和大于等于目标值时，不可能存在包含当前的组合的未发现的解，停止迭代
		if sum >= target {
			return
		}

		//遍历候选数集合，对每个候选数a
		for i := start; i < len(candidates); i++ {
			// 累加sum+=a
			sum += candidates[i]
			// 将a添加到结果组合中
			path = append(path, candidates[i])
			// 尝试找候选集合可以组成target‘的组合的解，转化为子问题，这里找的候选集下标不变，用来搜索包含了n个重复的当前候选数的组合解。
			dfs(sum, i)
			// 将a从结果组合中移除
			sum -= candidates[i]
			path = path[0 : len(path)-1]
		}
	}
	dfs(0, 0)

	return res
}
