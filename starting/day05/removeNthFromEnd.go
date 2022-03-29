package day05

import . "arithmetic/utils"

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := &ListNode{0, head}
	slow := &ListNode{0, head}
	res := slow
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return res.Next
}
