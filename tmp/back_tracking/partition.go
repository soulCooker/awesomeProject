package backtracking

func partition(s string) [][]string {
	return doPartiton(s, 0, len(s)-1)
}

func doPartiton(s string, start, end int) [][]string {
	// 初始化结果集合
	res := make([][]string, 0)
	// 处理当前位置的每一种解方式：从当前下标开始匹配回文子串
	for j := start; j <= end; j++ {
		if isSymmetric(s, start, j) {
			symmetric := string([]byte(s)[start : j+1])
			//     是否已经最小子问题：如果匹配到末尾，返回结果集合
			if j == end {
				res = append(res, []string{symmetric})
				return res
			}
			//     递归处理子问题，并根据子问题结果计算出父问题结果： 从下一个匹配位置递归切割回文串，返回子结果集
			subPartion := doPartiton(s, j+1, end)
			//     回溯处理：
			//         子结果集非空，将当前回文子串添加到所有子结果集合中,将更新后的子结果集合合并到结果集合中
			for i := range subPartion {
				oneAns := make([]string, 0)
				oneAns = append(oneAns, symmetric)
				oneAns = append(oneAns, subPartion[i]...)
				res = append(res, oneAns)
			}
		}
	}

	return res
}

func isSymmetric(s string, start, end int) bool {
	for i := start; i <= (start+end)/2; i++ {
		if s[i] != s[start+end-i] {
			return false
		}
	}
	return true
}
