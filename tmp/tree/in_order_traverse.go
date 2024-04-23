package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
深度优先遍历-迭代方式基于栈，除了前序遍历外，m节点会被访问两次
前序：m l r
中序：l m r
后序：l r m

前序
访问栈顶节点
	节点出栈，处理节点，将右节点、左节点进栈

中序
访问栈顶节点
	节点为未访问，标记节点为已访问，左节点进栈
	节点为已访问，节点出栈并处理，右节点进栈

后序
访问栈顶节点
	节点为未访问，标记节点为已访问，右节点、左节点进栈
	节点为已访问，节点出栈并处理
*/

/*
中序
访问栈顶节点

	节点为已访问，节点出栈并处理，右节点进栈
*/
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*VisitedNode, 0)
	res := make([]int, 0)

	if root == nil {
		return nil
	}

	stack = append(stack, &VisitedNode{
		Node:    root,
		Visited: 0,
	})

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		//节点为未访问，标记节点为已访问，左节点进栈
		if cur.Visited == 0 {
			cur.Visited = 1
			if cur.Node.Left != nil {
				stack = append(stack, &VisitedNode{
					Node:    cur.Node.Left,
					Visited: 0,
				})
			}
		} else { //节点为已访问，节点出栈并处理，右节点进栈
			res = append(res, cur.Node.Val)
			if cur.Node.Right != nil {
				stack[len(stack)-1].Visited = 0
				stack[len(stack)-1].Node = cur.Node.Right
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res
}

type VisitedNode struct {
	Node    *TreeNode
	Depth   int
	Visited int
}

func NewVNode(node *TreeNode, depth int) *VisitedNode {
	return &VisitedNode{
		Node:    node,
		Visited: 0,
		Depth:   depth,
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
		if vNode.Visited == 0 {
			if tNode.Left != nil {
				stack = append(stack, NewVNode(tNode.Left, vNode.Depth+1))
			}
			vNode.Visited = 1
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
		if vNode.Visited == 0 {
			if tNode.Left != nil {
				stack = append(stack, NewVNode(tNode.Left, 0))
			}
			res = append(res, tNode.Val)
			vNode.Visited = 1
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
