package link

import (
	"fmt"
	"testing"
)

func TestReverseRange(t *testing.T) {
	head := buildLink([]int{1, 2, 3, 4, 5, 6})
	printLink(head)
	head = reverseNode(head, 2, 3)
	printLink(head)
}

func printLink(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}

func buildLink(vals []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, val := range vals {
		node := &ListNode{
			Val: val,
		}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}
