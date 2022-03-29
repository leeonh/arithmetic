package day08

import . "arithmetic/utils"

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	// 每次循环从最左节点开始循环，判断左节点是否为空,每次循环后都移动到下一层的最左节点
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 每次操作连接的节点都是下一层的，所以当前循环的节点都是上一次循环就连接好了的
		// 循环向右连接同一层的节点，每次循环后当前节点切换到下一个连接节点
		for node := leftmost; node != nil; node = node.Next {
			// 将下一层的左节点连接到右节点
			node.Left.Next = node.Right
			// 如果当前节点的下一个连接节点不为空，则将当前节点的右节点连接到下一个连接节点的左节点(头结点没有next，所以只会进行左节点连接右节点)
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
