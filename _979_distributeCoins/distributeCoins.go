package _979_distributeCoins

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func helper(root *TreeNode, count *int) int {
	if root == nil {
		return 0
	}

	root.Val += helper(root.Left, count)
	root.Val += helper(root.Right, count)

	gap := root.Val-1
	if gap < 0 {
		gap *= -1
	}
	*count += gap

	return root.Val-1
}

func distributeCoins(root *TreeNode) int {
	var count int = 0
	helper(root, &count)
	return count
}