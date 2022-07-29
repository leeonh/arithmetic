package day07

import . "arithmetic/utils"

/*
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


示例：



输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指针连接），'#' 表示每层的末尾。


*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	// 思路，广度搜索，一层层遍历
	// 将节点放进队列，循环队列，每次将当前i节点的next指向下一个i+1节点
	for len(queue) > 0 {
		// 为啥要将队列交个一个变量来遍历? 为啥要将queue每次循环都设置为空?
		// 因为这里要一层层遍历循环，所以不能用弹出的方式，而是for循环，这样才能将i的next指向到i+1
		// 将queue置为空，代表着一层结束，下次再放进来队列的节点就是下一层的树节点了，不会保留上一层节点
		temp := queue
		queue = nil
		for i, node := range temp {
			// 如果不是队列最后一个节点
			if i+1 < len(temp) {
				node.Next = temp[i+1]
			}
			// 如果当前节点有左指针
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 如果当前节点有右指针
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}
	return root
}
