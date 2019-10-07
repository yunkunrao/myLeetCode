package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	if head == nil {
		return dummyHead.Next
	}

	ptr := dummyHead
	cur := head

	for cur != nil {
		for ptr.Next != nil && cur != nil && ptr.Next.Val < cur.Val {
			ptr = ptr.Next
		}

		curNext := cur.Next
		tmp := ptr.Next
		ptr.Next = cur
		cur.Next = tmp

		cur = curNext

		ptr = dummyHead
	}

	return dummyHead.Next
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	ret := insertionSortList(head)
	fmt.Println(ret)
}
