package main

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

func helper(nodes *[]*TreeNode, index *int) []int {
	var ret []int

	if len(*nodes) == 0 {
		return nil
	}

	for _, node := range (*nodes)[:*index+1] {
		if node.Left != nil {
			*nodes = append(*nodes, node.Left)
		}
		if node.Right != nil {
			*nodes = append(*nodes, node.Right)
		}
	}
	ret = append(ret, (*nodes)[*index].Val)
	*nodes = (*nodes)[*index+1:]
	*index = len(*nodes) - 1

	ret = append(ret, helper(nodes, index)...)

	return ret
}

func rightSideView(root *TreeNode) []int {
	var cur int = 0
	var nodeList []*TreeNode

	if root == nil {
		return nil
	}

	nodeList = append(nodeList, root)
	return helper(&nodeList, &cur)
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	rightSideView(root)
}
