package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var (
		fast = head
		slow = head
	)

	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		if slow == fast {
			return true
		}
	}
	return false
}
