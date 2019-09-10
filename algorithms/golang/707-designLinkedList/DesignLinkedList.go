package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head  *ListNode
	tail  *ListNode
	count int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) getNthNode(index int) *ListNode {
	if index < 0 || index >= this.count {
		return nil
	}

	p := this.head
	for index > 0 {
		p = p.Next
		index--
	}

	return p
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	p := this.getNthNode(index)
	if p == nil {
		return -1
	}

	return p.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	n := &ListNode{Val: val}
	n.Next = this.head
	this.head = n
	if this.count == 0 {
		this.tail = n
	}

	this.count++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	n := &ListNode{Val: val}
	if this.count == 0 {
		this.head = n
		this.tail = n
	} else {
		this.tail.Next = n
		this.tail = n
	}

	this.count++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.count {
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.count {
		this.AddAtTail(val)
		return
	}

	prev := this.getNthNode(index - 1)
	n := &ListNode{Val: val}
	n.Next = prev.Next
	prev.Next = n
	this.count++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.count {
		return
	}

	if index == 0 {
		this.head = this.head.Next
	} else {
		p := this.getNthNode(index - 1)
		p.Next = p.Next.Next
		if p.Next == nil {
			this.tail = p
		}
	}

	this.count--
	if this.count == 0 {
		this.tail = nil
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

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" -> ")
		head = head.Next
	}

	fmt.Println("NULL")
}

func main() {
	{
		l := Constructor()
		l.AddAtHead(1)
		printList(l.head)
		l.AddAtTail(3)
		printList(l.head)
		l.AddAtIndex(1, 2)
		printList(l.head)
		fmt.Println(l.Get(1))
		l.DeleteAtIndex(1)
		fmt.Println(l.Get(1))
	}

	{
		l := Constructor()
		l.AddAtIndex(-1, 0)
		printList(l.head)
		fmt.Println(l.Get(0))
		l.DeleteAtIndex(-1)
	}

	{
		l := Constructor()
		l.AddAtIndex(-1, 0)
		printList(l.head)
		fmt.Println(l.Get(0))
		l.DeleteAtIndex(-1)
	}
}
