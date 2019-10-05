package _24_swapPairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	left := head
	right := head.Next

	tmp := right.Next

	right.Next = left
	left.Next = swapPairs(tmp)

	return right
}