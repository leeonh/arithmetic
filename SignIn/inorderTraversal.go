package SignIn

import . "arithmetic/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	// 创建结果数组
	var res []int
	// 递归左子树-根节点-右子树
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	// 如果有左子树，则递归左子树
	if root.Left != nil {
		dfs(root.Left, res)
	}
	// 将根节点放进去结果集合
	*res = append(*res, root.Val)

	// 如果有右子树，则递归右子树
	if root.Right != nil {
		dfs(root.Right, res)
	}
}
