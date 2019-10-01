package _173_BSTIterator

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	List []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var iter BSTIterator
	var ptr *TreeNode = root

	for ptr != nil {
		iter.List = append(iter.List, ptr)
		ptr = ptr.Left
	}
	return iter
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {

	item := this.List[len(this.List)-1]
	this.List = this.List[:len(this.List)-1]
	ptr := item.Right
	for ptr != nil {
		this.List = append(this.List, ptr)
		ptr = ptr.Left
	}
	return item.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return 0 < len(this.List)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */