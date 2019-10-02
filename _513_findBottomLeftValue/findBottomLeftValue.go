package _513_findBottomLeftValue

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func helper(root *TreeNode, height int)(val int, h int) {
	var left, right = -1, -1
	var lh, rh = -1, -1

	if root.Left == nil && root.Right == nil {
		return root.Val, height
	}
	if root.Left != nil {
		left, lh = helper(root.Left, height+1)
	}
	if root.Right != nil {
		right, rh = helper(root.Right, height+1)
	}

	if lh >= rh {
		return left, lh + 1
	} else {
		return right, rh + 1
	}
}

func findBottomLeftValue(root *TreeNode) int {

	v, _ := helper(root, 0)
	return v
}
