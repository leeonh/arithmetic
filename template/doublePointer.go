package template

import . "arithmetic/utils"

/*
双指针，常见的有3种，对撞指针，快慢指针，两个对比对象双指针对比
*/

/*
对撞指针
分别定义两个指针，一个在开头，一个在末尾
然后从两头向中间遍历
对撞指针比较适用于有序数组，例如二分法其实也是双指针的一种
*/
func collisionPointer(nums []int) {
	i, j := 0, len(nums)-1
	for i <= j {
		// 加入一些if判断，然后i++或j--，根据题目变化自行改动
		if nums[i] < nums[j] {
			i++
		} else {
			j--
		}
		//	..........
	}
}

/*
快慢指针
顾名思义就是一个快一个慢
分别定义两个指针，都在开头
每次移动的时候，快指针移动2步，慢指针移动1步
快慢指针尤其适用于链表，例如判断链表是否有环，寻找链表中点
*/
func fastAndLowPointer(head *ListNode) bool {
	// 定义快慢指针
	low, fast := head, head
	// 如果跑得快的指针能遇到空nil，说明不是环形链表
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low == fast {
			return true
		}
	}
	return false
}

/*
两个对比对象双指针对比
这种通常就是两个数组一起遍历，取其中的值出来做对比
例如 比较含退格的字符串，
*/
func twoArrayPointer(A, B []int) {
	// 定义双指针，分别指向两个数组的末尾
	i, j := len(A)-1, len(B)-1
	for i >= 0 || j >= 0 {
		// 分别从两个数组取值对比，然后做处理
		if A[i] < B[j] {
			i++
		} else {
			j++
		}
		//	..........
	}
}
