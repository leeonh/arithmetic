package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func PrintTreeNode(left *TreeNode, right *TreeNode, depth int) {
	var a = depth
	var alias = "|————"
	for a > 0 {
		alias += "————"
		a--
	}

	if left != nil {
		fmt.Print(alias + " l:")
		fmt.Println(left.Val)
		if left.Left != nil {
			PrintTreeNode(left.Left, nil, depth+1)
		}
		if left.Right != nil {
			PrintTreeNode(nil, left.Right, depth+1)
		}

	}

	if right != nil {
		fmt.Print(alias + " r:")
		fmt.Println(right.Val)
		if right.Left != nil {
			PrintTreeNode(right.Left, nil, depth+1)
		}
		if right.Right != nil {
			PrintTreeNode(nil, right.Right, depth+1)
		}

	}

}
