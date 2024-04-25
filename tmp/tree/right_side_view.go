package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 初始化当前层列表
	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)

	// 当前层节点列表不为空时循环处理
	for len(curLevel) > 0 {
		// 	遍历当前层节点列表
		nextLevel := make([]*TreeNode, 0)
		for _, v := range curLevel {
			// 		将左右子节点添加到下一层列表
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		// 	将当前层列表的最后元素添加到结果数组中
		res = append(res, curLevel[len(curLevel)-1].Val)
		// 	更新当前层节点列表 为下一次列表
		curLevel = nextLevel
	}

	return res
	// 返回结果数组
}
