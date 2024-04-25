package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)

	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		curLevelVal := make([]int, 0)
		//遍历当前层节点
		for _, v := range curLevel {
			//将节点值插入到当前值列表中
			curLevelVal = append(curLevelVal, v.Val)
			//将左右子树放到下一层节点的列表中
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}

		}
		//更新当前层节点列表为下一层节点列表
		curLevel = nextLevel
		//将当前层的值列表插入到遍历列表中
		res = append(res, curLevelVal)
	}
	return res
}
