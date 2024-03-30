package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type DepthNode struct {
	*TreeNode
	Depth int
}

func minDepth(root *TreeNode) int {
	stack := make([]*DepthNode, 0)

	if root == nil {
		return 0
	}

	stack = append(stack, &DepthNode{
		TreeNode: root,
		Depth:    1,
	})

	min := -1

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if min > 0 && node.Depth >= min {
			continue
		}
		if node.Left == nil && node.Right == nil {
			if min == -1 || min > node.Depth {
				min = node.Depth
			}
		}

		if node.Right != nil {
			stack = append(stack, &DepthNode{
				TreeNode: node.Right,
				Depth:    node.Depth + 1,
			})
		}
		if node.Left != nil {
			stack = append(stack, &DepthNode{
				TreeNode: node.Left,
				Depth:    node.Depth + 1,
			})
		}
	}
	return min
}
