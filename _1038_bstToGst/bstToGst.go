package main

import "fmt"

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

func updateRootValue(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	updateRootValue(root.Right, nodes)
	if len(*nodes) != 0 {
		root.Val += (*nodes)[len(*nodes)-1].Val
	}
	*nodes = append(*nodes, root)
	updateRootValue(root.Left, nodes)
}

func bstToGst(root *TreeNode) *TreeNode {
	var nodes []*TreeNode
	updateRootValue(root, &nodes)
	return root
}

func main() {
	root := &TreeNode{
		Val:7,
		Right: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
	}
	bstToGst(root)
	fmt.Println(root.Val)
}