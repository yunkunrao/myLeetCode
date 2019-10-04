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

func helper(inorder []int, postorder []int, begin, end int, tail *int) *TreeNode {
	var index int = -1

	if begin > end {
		return nil
	}

	root := &TreeNode{
		Val: postorder[*tail],
	}
	*tail--

	for i := begin; i <= end; i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	if index == -1 {
		return nil
	}

	root.Right = helper(inorder, postorder, index+1, end, tail)
	root.Left = helper(inorder, postorder, begin, index-1, tail)

	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	N := len(inorder)
	tail := N-1
	return helper(inorder, postorder, 0, N-1, &tail)

}

func main() {

	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	fmt.Println(buildTree(inorder, postorder))
}
