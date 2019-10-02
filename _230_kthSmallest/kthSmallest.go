package _230_kthSmallest

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root *TreeNode, k int, list *[]int) {
	if root == nil {
		return
	}
	if len(*list) >= k {
		return
	}
	helper(root.Left, k, list)
	*list = append(*list, root.Val)
	helper(root.Right, k, list)
}

func kthSmallest(root *TreeNode, k int) int {
	var list []int

	helper(root, k, &list)

	return list[k-1]
}