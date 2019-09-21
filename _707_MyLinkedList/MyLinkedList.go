package main

import "fmt"

type MyLinkedList struct {
	Size int
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Size: 0,
		Val:  0,
		Next: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	var ptr *MyLinkedList = this
	if index < 0 || index + 1 > this.Size{
		return -1
	}

	for i := 0; i < index; i++ {
		ptr = ptr.Next
	}

	return ptr.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	if this.Size == 0 {
		this.Val = val
		this.Size++
		return
	}

	node := MyLinkedList{
		Val:  this.Val,
		Next: this.Next,
	}

	this.Val = val
	this.Next = &node
	this.Size++
	return
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Size == 0 {
		this.Val = val
		this.Size++
		return
	}

	tail := MyLinkedList{
		Val:  val,
		Next: nil,
	}

	var ptr *MyLinkedList = this
	for ptr.Next != nil {
		ptr = ptr.Next
	}

	ptr.Next = &tail
	this.Size++
	return
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	var ptr, pre *MyLinkedList

	if index == this.Size {
		this.AddAtTail(val)
		return
	} else if index > this.Size {
		return
	} else if index <= 0 {
		this.AddAtHead(val)
		return
	} else {
		// index : (0,length)
		ptr = this
		for i:=0; ptr != nil;{
			if i == index - 1 {
				pre = ptr
				break
			}
			ptr = ptr.Next
			i++
		}

		node := MyLinkedList{
			Val:  val,
			Next: pre.Next,
		}
		pre.Next = &node
		this.Size++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	var ptr *MyLinkedList = this
	var i = 0

	if index < 0 || index >= this.Size {
		return
	}

	if index == 0 {
		if this.Next != nil {
			this.Val = this.Next.Val
			this.Next = this.Next.Next
		} else {
			this.Val = 0
			this.Next = nil
		}

		this.Size--
		return
	}

	for ptr != nil {
		if i == index-1 {
			ptr.Next = ptr.Next.Next
			this.Size--
			return
		}
		ptr = ptr.Next
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//["MyLinkedList","addAtHead(0)","get","addAtTail(2)","addAtIndex(1,4)",
// "addAtHead(4)","addAtIndex(1,4)","addAtTail(5)","addAtTail(2)",
// "addAtIndex(2,0)","get(2)","addAtTail(1)"]
//[,,,,,,,,,,,]

// 4, 4, 0, 0,4, 2, 5, 2

func main() {
	obj := Constructor()
	obj.AddAtHead(0)
	obj.AddAtTail(2)
	obj.AddAtIndex(1, 4)
	obj.AddAtHead(4)
	obj.AddAtIndex(1, 4)
	//obj.AddAtTail(5)
	//obj.AddAtTail(2)
	fmt.Println("###", obj.Size)

	obj.AddAtIndex(2, 0)
	tmp := &obj
	for tmp != nil {
		fmt.Printf("%d ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("")
	//fmt.Println("###", obj.Get(2))
}
