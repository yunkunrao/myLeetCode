package _109_sortedListToBST

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedList2BST(list []int)(*TreeNode, int) {
	var length = len(list)
	var root *TreeNode

	if length == 0 {
		return nil, 0
	}

	mid := length/2
	left, lh := sortedList2BST(list[:mid])
	right, rh := sortedList2BST(list[mid+1:])
	gap := lh - rh
	hMax := lh
	if gap < 0 {
		gap *= -1
		hMax = rh
	}
	if gap <= 1 {
		root = &TreeNode{
			Val: list[mid],
			Left: left,
			Right: right,
		}
		return root, hMax+1
	}
	return nil, 0
}

func sortedListToBST(head *ListNode) *TreeNode {

	var list []int

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	root, _ := sortedList2BST(list)

	return root
}

