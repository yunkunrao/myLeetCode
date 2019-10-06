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

func helper(preorder []int, inorder []int, cur *int) *TreeNode {
	if *cur > len(preorder)-1 || len(inorder) == 0 {
		return nil
	}

	var rootIndex = -1
	for i, n := range inorder {
		if n == preorder[*cur] {
			rootIndex = i
			break
		}
	}

	root := &TreeNode{
		Val: preorder[*cur],
	}
	*cur++

	if rootIndex == -1 {
		fmt.Println(rootIndex)
	}

	root.Left = helper(preorder, inorder[:rootIndex], cur)
	root.Right = helper(preorder, inorder[rootIndex+1:], cur)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var cur = 0
	return helper(preorder, inorder, &cur)
}

func main()  {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	fmt.Println(buildTree(preorder, inorder))
}