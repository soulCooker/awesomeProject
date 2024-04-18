package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		cur      *ListNode
		nThNode  *ListNode
		prevNode *ListNode
	)

	cur = head
	for cur != nil && n > 0 {
		cur = cur.Next
		n--
	}

	nThNode = head
	prevNode = nil

	for cur != nil {
		cur = cur.Next
		prevNode = nThNode
		nThNode = nThNode.Next
	}

	if prevNode != nil {
		prevNode.Next = nThNode.Next
	} else {
		return head.Next
	}

	return head
}
