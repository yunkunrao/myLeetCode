package _1008_bstFromPreorder

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode

	if len(preorder) == 0 {
		return nil
	}

	var mid int = len(preorder)
	for i, v := range preorder {
		if i != 0 && v > preorder[0] {
			mid = i
			break
		}
	}

	root = &TreeNode{
		Val: preorder[0],
	}

	root.Left = bstFromPreorder(preorder[1:mid])
	root.Right = bstFromPreorder(preorder[mid:])


	return root
}
