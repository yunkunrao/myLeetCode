package _222_countNodes

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	n1 := countNodes(root.Left)
	n2 := countNodes(root.Right)

	return 1+n1+n2
}
