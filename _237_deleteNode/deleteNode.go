package deleteNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {

	// 题目已经提示给定的节点一定是一个有效的节点，故不需要检查

	var cur *ListNode = node
	var next *ListNode = node.Next
	for ;next.Next != nil; {
		cur.Val = next.Val

		cur = next
		next = cur.Next
	}
	cur.Val = next.Val
	cur.Next = nil
}
