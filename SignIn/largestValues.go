package SignIn

import (
	. "arithmetic/utils"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := []int{}
	// 思路，广度搜索，一层层遍历
	// 将节点放进队列，循环队列
	for len(queue) > 0 {
		// 每一层节点，都从最小值开始对比
		max := math.MinInt64

		// 为啥要将队列交个一个变量来遍历? 为啥要将queue每次循环都设置为空?
		// 因为这里要一层层遍历循环，所以不能用弹出的方式，而是for循环，这样才能将i的next指向到i+1
		// 将queue置为空，代表着一层结束，下次再放进来队列的节点就是下一层的树节点了，不会保留上一层节点
		temp := queue
		queue = nil
		for _, node := range temp {
			// 如果比最大值大，就更新最大值
			max = getMax(max, node.Val)
			// 如果当前节点有左指针
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 如果当前节点有右指针
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, max)

	}
	return res
}
func getMax(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}
