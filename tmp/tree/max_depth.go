package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxDepth(root *TreeNode) int {
	maxDepth := 0
	doInorderTraversal(root, func(node *VisitedNode) interface{} {
		if node.Depth > maxDepth {
			maxDepth = node.Depth
		}
		return maxDepth
	})
	return maxDepth
}
