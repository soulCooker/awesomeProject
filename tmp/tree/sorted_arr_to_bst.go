package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return doSortedArrayToBST(nums, 0, len(nums))
}

func doSortedArrayToBST(nums []int, start, end int) *TreeNode {
	// 如果切片长度 <= 0，直接返回nil值
	if end <= start {
		return nil
	}

	// 取当前切片中位数作为根节点，对下标小于和大于中位数的部分分别做递归调用，递归调用结果分别作为根节点的左右子树。
	mid := (start + end) / 2
	curRoot := &TreeNode{
		Val:   nums[mid],
		Left:  doSortedArrayToBST(nums, start, mid),
		Right: doSortedArrayToBST(nums, mid+1, end),
	}
	return curRoot
}
