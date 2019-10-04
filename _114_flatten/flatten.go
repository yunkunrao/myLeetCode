package _114_flatten

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	var left, right, ptr *TreeNode
	ptr = root

	if root == nil {
		return
	}

	left = root.Left
	right = root.Right

	root.Left = nil

	flatten(left)
	root.Right = left

	for ptr.Right != nil {
		ptr = ptr.Right
	}

	flatten(right)
	ptr.Right = right
}