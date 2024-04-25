package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type KthSmallestNode struct {
	*TreeNode
	Visited int
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*KthSmallestNode, 0)
	stack = append(stack, &KthSmallestNode{
		TreeNode: root,
	})
	for len(stack) > 0 {
		// 取栈顶节点
		curRoot := stack[len(stack)-1]
		if curRoot.Visited == 0 {
			//     节点未访问
			//         节点左节点进栈，节点标记为已读
			if curRoot.Left != nil {
				stack = append(stack, &KthSmallestNode{
					TreeNode: curRoot.Left,
				})
			}
			curRoot.Visited = 1
		} else {
			//     节点已访问
			//       序号递减，如果序号==0，返回节点值
			k--
			if k == 0 {
				return curRoot.Val
			}

			//          节点出栈，节点的右节点进栈
			if curRoot.Right != nil {
				stack[len(stack)-1] = &KthSmallestNode{
					TreeNode: curRoot.Right,
				}
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return 0
	// 结束
}
