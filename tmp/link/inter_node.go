package link

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		i = headA
		j = headB
	)

	for i != nil && j != nil {
		i = i.Next
		j = j.Next
	}

	var (
		s      *ListNode
		l      *ListNode
		remain *ListNode
	)

	if i == nil {
		s = headA
		l = headB
		remain = j
	} else if j == nil {
		s = headB
		l = headA
		remain = i
	}

	for ; remain != nil; remain = remain.Next {
		l = l.Next
	}

	for s != nil && s != l {
		l = l.Next
		s = s.Next
	}

	return s
}
