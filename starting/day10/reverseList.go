package day10

import . "arithmetic/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 * 链表例子：1 -> 2 -> 3 -> 4 -> 5
 */
func ReverseList1(head *ListNode) *ListNode {
	// 递归终止条件，当前节点为空或者下一个节点为空
	if head == nil || head.Next == nil {
		return head
	}

	// 通过递归，将新建的头结点递归到最尾的节点, 此节点将作为新的头结点，然后将next指向前一个节点
	newHead := ReverseList(head.Next)
	// 参考例子，递归到最后面的时候，会返回5作为头结点newHead, 递归回到上一层后，head是4
	// 那么4的next就是5，再next就是空，那么 4.next.next = 4，相当于5的next指回去了4，此时产生了环
	head.Next.Next = head
	// 要把环取消掉，那就是将4.next = nil
	head.Next = nil
	// 以上每次操作都是操作head链表，同时也会一起更新newHead，每次返回head的最后一个节点newHead
	return newHead
}

/**
 * 双指针迭代解法
 */
func ReverseList(head *ListNode) *ListNode {
	// pre指针指向空
	var pre *ListNode
	// 当前指针指向头部
	cur := head
	// 临时存放的指针，用于交换
	var temp *ListNode
	// 循环到链表末尾
	for cur != nil {
		// 存放下一个节点
		temp = cur.Next
		// 当前节点.next 指向pre
		cur.Next = pre
		// pre节点向前移动一次
		pre = cur
		// cur节点向前移动一次，因为temp已经存放了cur的下一个节点
		cur = temp
	}
	return pre
}
