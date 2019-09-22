package _654_constructMaximumBinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMaxIndex(array []int) int {
	var index, max int
	for i, v := range array {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := findMaxIndex(nums)

	root := TreeNode {
		Val:nums[maxIndex],
	}

	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	if maxIndex < len(nums) - 1 {
		root.Right = constructMaximumBinaryTree(nums[maxIndex+1:len(nums)])
	} else {
		root.Right = nil
	}

	return &root
}