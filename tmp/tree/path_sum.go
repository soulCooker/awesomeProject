package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type PathSumNode struct {
	*TreeNode
	Visited int
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 前缀路径和列表=以父节点作为终点，任一祖父节点起点的路径上的节点值之和
	suffixPathSum := make([]int, 0)
	res := 0
	stack := make([]*PathSumNode, 0)
	// 初始化节点栈
	stack = append(stack, &PathSumNode{
		TreeNode: root,
	})

	// 节点栈不为空时循环处理
	for len(stack) > 0 {
		// 	取栈顶节点
		curRoot := stack[len(stack)-1]
		// 	节点未访问
		if curRoot.Visited == 0 {
			// 		前缀路径和列表的值为 单前值 + 当前节点值
			for i := range suffixPathSum {
				suffixPathSum[i] += curRoot.Val
			}
			// 		将当前节点值作为路径和添加到前缀路径和列表
			suffixPathSum = append(suffixPathSum, curRoot.Val)
			// 		将节点的左、右子树入栈
			if curRoot.Left != nil {
				stack = append(stack, &PathSumNode{
					TreeNode: curRoot.Left,
				})
			}
			if curRoot.Right != nil {
				stack = append(stack, &PathSumNode{
					TreeNode: curRoot.Right,
				})
			}
			// 		更新节点为已访问
			curRoot.Visited = 1
		} else {
			// 	节点已访问
			// 		判断前缀路径和列表中值为目标值出现次数并累加到结果值中,并将所有路径和减去该节点的值
			for i, v := range suffixPathSum {
				if v == targetSum {
					res++
				}
				suffixPathSum[i] -= curRoot.Val
			}
			// 		将前缀路径和列表并最后一项移除
			suffixPathSum = suffixPathSum[0 : len(suffixPathSum)-1]

			// 		节点出栈
			stack = stack[0 : len(stack)-1]
		}

	}
	// 返回结果值
	return res
}
