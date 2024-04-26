package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type MaxPathSumNode struct {
	*TreeNode
	MaxEdgeSuffixSum int
	Visited          int
	MLeft            *MaxPathSumNode
	MRight           *MaxPathSumNode
}

func maxPathSum(root *TreeNode) int {
	// 初始化栈
	stack := make([]*MaxPathSumNode, 0)
	stack = append(stack, &MaxPathSumNode{
		TreeNode: root,
	})
	// 初始化最大路径和为根节点的值
	maxPathSum := root.Val
	// 栈不为空时循环处理
	for len(stack) > 0 {
		curRoot := stack[len(stack)-1]
		// 	节点未访问
		if curRoot.Visited == 0 {
			// 		左右子树入栈
			if curRoot.Left != nil {
				l := &MaxPathSumNode{
					TreeNode: curRoot.Left,
				}
				curRoot.MLeft = l
				stack = append(stack, l)
			}
			if curRoot.Right != nil {
				r := &MaxPathSumNode{
					TreeNode: curRoot.Right,
				}
				curRoot.MRight = r
				stack = append(stack, r)
			}
			//标记节点为已访问
			curRoot.Visited = 1
		} else { // 	节点已访问
			// 		计算节点的最长枝干前缀和 = max(max(左子树的枝干前缀最大值，右子树的枝干前缀最大值) + 当前节点值, 0)并更新到节点属性上
			lMaxEdgeNullableSuffixSum := 0
			rMaxEdgeNullableSuffixSum := 0
			if curRoot.MLeft != nil {
				lMaxEdgeNullableSuffixSum = max(curRoot.MLeft.MaxEdgeSuffixSum, 0)
			}
			if curRoot.MRight != nil {
				rMaxEdgeNullableSuffixSum = max(curRoot.MRight.MaxEdgeSuffixSum, 0)
			}

			curRoot.MaxEdgeSuffixSum = max(lMaxEdgeNullableSuffixSum, rMaxEdgeNullableSuffixSum) + curRoot.Val
			// 		计算该节点为根的树上的最长路径和为 左子树的枝干前缀最大值+右子树的枝干前缀最大值+当前节点值
			// 		该节点为根的树上的最长路径和是否大于 当前最大路径和，大于则更新最大路径和
			maxPathSum = max(lMaxEdgeNullableSuffixSum+rMaxEdgeNullableSuffixSum+curRoot.Val, maxPathSum)
			// 		节点出栈
			stack = stack[0 : len(stack)-1]
		}
	}

	// 返回结果
	return maxPathSum
}
