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

func sortList(head *ListNode) *ListNode {
	//var tmp55 []int
	//for tmp := head; tmp != nil; tmp = tmp.Next {
	//	tmp55 = append(tmp55, tmp.Val)
	//}
	//fmt.Println(tmp55)

	var slow, fast = head, head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if fast != nil {
			slow = slow.Next
		}
	}

	tmp := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(tmp)

	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dummyHead

	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			cur = left
			left = left.Next
		} else {
			cur.Next = right
			cur = right
			right = right.Next
		}
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return dummyHead.Next
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	sortList(head)

	fmt.Println("hello")
}
