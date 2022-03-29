package day08

import . "arithmetic/utils"

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 节点的值合并
	root1.Val += root2.Val
	// 合并当前节点的左子树
	root1.Left = mergeTrees(root1.Left, root2.Left)
	// 合并当前节点的右子树
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
