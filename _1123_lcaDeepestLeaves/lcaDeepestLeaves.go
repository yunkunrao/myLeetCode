package _1123_lcaDeepestLeaves

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := 1+helper(root.Left)
	rightH := 1+helper(root.Right)
	if leftH >= rightH {
		return leftH
	} else {
		return rightH
	}
}


func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftDepth := helper(root.Left)
	rightDepth := helper(root.Right)

	if leftDepth == rightDepth {
		return root
	} else if leftDepth > rightDepth {
		return lcaDeepestLeaves(root.Left)
	} else {
		return lcaDeepestLeaves(root.Right)
	}
}
