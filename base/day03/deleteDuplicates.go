package day03

import (
	. "arithmetic/utils"
)

/*
给定一个已排序的链表的头head ，删除原始链表中所有重复数字的节点，只留下不同的数字。返回 已排序的链表。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//指针的next指向head
	ans := &ListNode{0, head}
	cur := ans
	// 每次比较用cur的next和next.next对比
	for cur.Next != nil && cur.Next.Next != nil {
		// 如果值相同
		if cur.Next.Val == cur.Next.Next.Val {
			//cur.next一直迭代去到值不相等为止
			temp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
		} else {
			// 如果不相同
			cur = cur.Next
		}
	}
	return ans.Next
}
