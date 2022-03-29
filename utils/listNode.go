package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print(",")
		node = node.Next
	}
	fmt.Println()
}
