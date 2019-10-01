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
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var nodes []*TreeNode

	if root == nil {
		return nil
	}

	ptr := root

	for ptr != nil || len(nodes) != 0 {

		// left & root
		for ptr != nil {
			nodes = append(nodes, ptr)
			ptr = ptr.Left
		}

		//     *  <--- ptr
		// nil    *
		ptr = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		ret = append(ret, ptr.Val)

		// right
		ptr = ptr.Right
	}
	return ret
}

func main() {
	root := &TreeNode{
		Val:1,
		Right: &TreeNode{
			Val:2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(inorderTraversal(root))
}