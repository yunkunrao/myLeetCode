package _938_rangeSumBST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int = 0

	// 检查参数合法性
	if root == nil || L > R {
		return 0
	}

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	sum += rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)

	return sum
}