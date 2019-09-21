package _617_mergeTrees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var l1, r1, l2, r2 *TreeNode
	t3 := new(TreeNode)

	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 != nil {
		t3.Val = t1.Val
		l1 = t1.Left
		r1 = t1.Right
	}
	if t2 != nil {
		t3.Val += t2.Val
		l2 = t2.Left
		r2 = t2.Right
	}

	t3.Left = mergeTrees(l1, l2)
	t3.Right = mergeTrees(r1, r2)

	return t3
}