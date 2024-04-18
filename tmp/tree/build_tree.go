package tree

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return doBuildSubTree(preorder, inorder, 0, len(preorder), 0, len(preorder))
}

func doBuildSubTree(preorder []int, inorder []int, pStart, pEnd, iStart, iEnd int) *TreeNode {
	//fmt.Println("pStart:", pStart, ", pEnd:", pEnd, ",iStart:", iStart, ",iEnd:", iEnd)
	if pStart == pEnd {
		return nil
	}

	split := 0

	for split = iStart; split < iEnd; split++ {
		if inorder[split] == preorder[pStart] {
			break
		}
	}

	subTreeRoot := &TreeNode{
		Val:   preorder[pStart],
		Left:  doBuildSubTree(preorder, inorder, pStart+1, pStart+1+(split-iStart), iStart, split),
		Right: doBuildSubTree(preorder, inorder, pStart+1+(split-iStart), pEnd, split+1, iEnd),
	}
	return subTreeRoot
}
