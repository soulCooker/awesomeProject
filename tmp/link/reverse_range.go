package link

func reverseNode(head *ListNode, m, n int) *ListNode {
	start := head
	end := head
	var (
		prev        *ListNode
		remain      *ListNode
		reverseHead *ListNode
	)

	for m > 1 {
		prev = start
		start = start.Next
		m--
	}
	for n > 1 {
		end = end.Next
		n--
	}
	remain = end.Next

	for next := start; ; {
		if reverseHead == nil {
			reverseHead = next
		} else {
			next.Next = reverseHead
			reverseHead = next
		}

		if next == end {
			break
		}
		next = next.Next
	}

	start.Next = remain

	if prev == nil {
		head = reverseHead
	} else {
		prev.Next = reverseHead
	}

	return head
}
