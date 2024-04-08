package link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	var (
		fast = head
		slow = head
	)

	if fast == nil {
		return nil
	}

	var first *ListNode
	var hasCycle bool
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return first
		}
		if slow == fast {
			hasCycle = true
			break
		}
	}

	if hasCycle {
		first = head
		for first != slow {
			first = first.Next
			slow = slow.Next
		}
	}

	return first
}
