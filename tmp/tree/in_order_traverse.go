package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type VisitedNode struct {
	Node   *TreeNode
	Status int
}

func NewVNode(node *TreeNode) *VisitedNode {
	return &VisitedNode{
		Node:   node,
		Status: 0,
	}
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*VisitedNode, 0)
	res := make([]int, 0)

	if root == nil {
		return nil
	}

	stack = append(stack, NewVNode(root))

	for len(stack) > 0 {
		vNode := stack[len(stack)-1]
		tNode := vNode.Node
		if vNode.Status == 0 {
			if tNode.Left != nil {
				stack = append(stack, NewVNode(tNode.Left))
			}
			vNode.Status = 1
		} else {
			res = append(res, tNode.Val)
			if tNode.Right != nil {
				stack[len(stack)-1] = NewVNode(tNode.Right)
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return res
}

func PreorderTraversal(root *TreeNode) []int {
	stack := make([]*VisitedNode, 0)
	res := make([]int, 0)

	if root == nil {
		return nil
	}

	stack = append(stack, NewVNode(root))

	for len(stack) > 0 {
		vNode := stack[len(stack)-1]
		tNode := vNode.Node
		if vNode.Status == 0 {
			if tNode.Left != nil {
				stack = append(stack, NewVNode(tNode.Left))
			}
			res = append(res, tNode.Val)
			vNode.Status = 1
		} else {
			if tNode.Right != nil {
				stack[len(stack)-1] = NewVNode(tNode.Right)
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return res
}
