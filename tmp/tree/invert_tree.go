package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		//栈顶节点出栈
		curRoot := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		//交换左右子树。将左右子树进栈
		curRoot.Left, curRoot.Right = curRoot.Right, curRoot.Left

		if curRoot.Left != nil {
			stack = append(stack, curRoot.Left)
		}
		if curRoot.Right != nil {
			stack = append(stack, curRoot.Right)
		}
	}

	return root
}
