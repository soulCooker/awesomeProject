package link

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var (
		resHead *ListNode
		tail    *ListNode
		next    *ListNode
	)

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			next = list2
			list2 = list2.Next
		} else {
			next = list1
			list1 = list1.Next
		}

		if tail == nil {
			tail = next
			resHead = tail
		} else {
			tail.Next = next
			tail = next
		}
	}

	if list1 != nil {
		next = list1
	} else if list2 != nil {
		next = list2
	}

	if tail == nil {
		resHead = next
	} else {
		tail.Next = next
	}

	return resHead
}
