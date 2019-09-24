package _814_pruneTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func containOnlyZero(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Val == 1 {
		return false
	}

	return containOnlyZero(root.Left) && containOnlyZero(root.Right)
}

func pruneTree(root *TreeNode) *TreeNode {
	if containOnlyZero(root) {
		return nil
	}

	if containOnlyZero(root.Left) {
		root.Left = nil
	} else {
		root.Left = pruneTree(root.Left)
	}

	if containOnlyZero(root.Right) {
		root.Right = nil
	} else {
		root.Right = pruneTree(root.Right)
	}

	return root
}