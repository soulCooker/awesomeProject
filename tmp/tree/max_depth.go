package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
深度优先遍历，先序遍历
节点出栈并更新最大深度=max(当前最大深度,节点深度)，左节点、右节点进栈。子节点深度为节点深度+1
*/

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := 0

	stack := make([]*DepthNode, 0)
	stack = append(stack, &DepthNode{
		TreeNode: root,
		Depth:    1,
	})

	for len(stack) > 0 {
		curRoot := stack[len(stack)-1]
		if curRoot.Depth > maxDepth {
			maxDepth = curRoot.Depth
		}
		stack = stack[:len(stack)-1]
		if curRoot.Left != nil {
			stack = append(stack, &DepthNode{
				Depth:    curRoot.Depth + 1,
				TreeNode: curRoot.Left,
			})
		}
		if curRoot.Right != nil {
			stack = append(stack, &DepthNode{
				Depth:    curRoot.Depth + 1,
				TreeNode: curRoot.Right,
			})
		}
	}

	return maxDepth
}
