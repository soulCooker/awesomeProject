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
	Depth  int
	Status int
}

func NewVNode(node *TreeNode, depth int) *VisitedNode {
	return &VisitedNode{
		Node:   node,
		Status: 0,
		Depth:  depth,
	}
}

func doInorderTraversal(root *TreeNode, visited func(node *VisitedNode) interface{}) interface{} {
	stack := make([]*VisitedNode, 0)

	if root == nil {
		return nil
	}

	stack = append(stack, NewVNode(root, 1))

	var res interface{}

	for len(stack) > 0 {
		vNode := stack[len(stack)-1]
		tNode := vNode.Node
		if vNode.Status == 0 {
			if tNode.Left != nil {
				stack = append(stack, NewVNode(tNode.Left, vNode.Depth+1))
			}
			vNode.Status = 1
		} else {
			res = visited(vNode)
			if tNode.Right != nil {
				stack[len(stack)-1] = NewVNode(tNode.Right, vNode.Depth+1)
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return res
}

func InorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	doInorderTraversal(root, func(node *VisitedNode) interface{} {
		res = append(res, node.Node.Val)
		return res
	})
	return res
}

func PreorderTraversal(root *TreeNode) []int {
	stack := make([]*VisitedNode, 0)
	res := make([]int, 0)

	if root == nil {
		return nil
	}

	stack = append(stack, NewVNode(root, 0))

	for len(stack) > 0 {
		vNode := stack[len(stack)-1]
		tNode := vNode.Node
		if vNode.Status == 0 {
			if tNode.Left != nil {
				stack = append(stack, NewVNode(tNode.Left, 0))
			}
			res = append(res, tNode.Val)
			vNode.Status = 1
		} else {
			if tNode.Right != nil {
				stack[len(stack)-1] = NewVNode(tNode.Right, 0)
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return res
}
