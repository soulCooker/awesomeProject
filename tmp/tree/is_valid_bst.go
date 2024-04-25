package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTNode struct {
	*TreeNode
	Visited int
}

func isValidBST(root *TreeNode) bool {
	//递推法，中序遍历过程中，如果左右子树是二叉搜索树，且节点大于前序节点，小于后序节点，则节点为根的树也是一个二叉树。
	if root == nil {
		return false
	}

	//前序值=极小值
	preVal := 0
	preValOveride := false
	stack := make([]*BSTNode, 0)
	stack = append(stack, &BSTNode{
		TreeNode: root,
	})

	//取栈顶值
	for len(stack) > 0 {
		curRoot := stack[len(stack)-1]
		//节点未访问
		if curRoot.Visited == 0 {
			//	l进栈 更新为已访问
			if curRoot.Left != nil {
				stack = append(stack, &BSTNode{
					TreeNode: curRoot.Left,
				})
			}
			curRoot.Visited = 1
		} else {
			//节点已访问
			//  节点出栈，如果节点值是否<=前序值，不是BST树，结束
			//  r进栈
			if preValOveride && curRoot.Val <= preVal {
				return false
			}
			preVal = curRoot.Val
			preValOveride = true
			if curRoot.Right != nil {
				stack[len(stack)-1] = &BSTNode{
					TreeNode: curRoot.Right,
				}
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return true
	//是BST树，结束
}
