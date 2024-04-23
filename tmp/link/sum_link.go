package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flipL1 := l1
	flipL2 := l2

	var (
		sumTail  *ListNode
		sumHead  *ListNode
		overflow bool
	)

	for flipL1 != nil && flipL2 != nil {
		node := &ListNode{
			Val: flipL1.Val + flipL2.Val,
		}
		if overflow {
			node.Val += 1
		}
		overflow = false
		if node.Val >= 10 {
			node.Val %= 10
			overflow = true
		}

		if sumTail == nil {
			sumTail = node
			sumHead = node
		} else {
			sumTail.Next = node
			sumTail = node
		}
		flipL1 = flipL1.Next
		flipL2 = flipL2.Next
	}

	var remain *ListNode
	if flipL1 != nil {
		remain = flipL1
	} else if flipL2 != nil {
		remain = flipL2
	}

	for remain != nil {
		node := &ListNode{
			Val: remain.Val,
		}

		if overflow {
			node.Val += 1
		}
		overflow = false
		if node.Val >= 10 {
			node.Val %= 10
			overflow = true
		}

		if sumTail == nil {
			sumTail = node
			sumHead = node
		} else {
			sumTail.Next = node
			sumTail = node
		}
		remain = remain.Next
	}
	if overflow {
		sumTail.Next = &ListNode{
			Val: 1,
		}
	}

	return sumHead
}

// func flipLink(head *ListNode) *ListNode {
// 	var (
// 		prev *ListNode
// 		cur  *ListNode
// 	)
// 	cur = head
// 	for cur != nil {
// 		tmp := cur.Next
// 		cur.Next = prev
// 		prev = cur
// 		cur = tmp
// 	}
// 	return prev
// }
