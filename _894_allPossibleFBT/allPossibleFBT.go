package _894_allPossibleFBT

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	var root *TreeNode
	var nodes []*TreeNode

	root = &TreeNode{
		Val: 0,
	}

	if N%2 == 0 {
		return nil
	}

	if N == 1 {
		return []*TreeNode{root}
	}

	for i := 1; i < N-1; i += 2 {
		for _, right := range allPossibleFBT(N - i - 1) {
			for _, left := range allPossibleFBT(i) {
				tmp := &TreeNode{Val: 0}

				tmp.Left = left
				tmp.Right = right
				nodes = append(nodes, tmp)
			}
		}
	}
	return nodes
}
