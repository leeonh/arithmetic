package day10

import . "arithmetic/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 */

func MergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	// 结果链表
	res := &ListNode{0, nil}
	head := res
	// 两个指针向右循环
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next

	}
	if list1 == nil {
		head.Next = list2
	} else {
		head.Next = list1
	}
	return res.Next
}

/**
递归
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果l1已经空了
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// 对比节点大小,如果l1<l2，则l1的下一个节点 = 递归对比l1.Next 和 l2的大小
	if l1.Val <= l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists(l2.Next, l1)
		return l2
	}
}
