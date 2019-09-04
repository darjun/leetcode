package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q, r := (*ListNode)(nil), head, head.Next
	for q != nil {
		q.Next = p
		p, q = q, r
		if r != nil {
			r = r.Next
		}
	}

	return p
}

// isPalindrome 是否 head 是否是回文列表
// 先通过快慢指针找到中心节点的前一个节点，然后将后半部分反转。最后从头比较反转后的列表。
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head.Next

	count := 1
	for fast != nil {
		fast = fast.Next
		count++
		if fast != nil {
			fast = fast.Next
			count++
		}

		if fast != nil {
			slow = slow.Next
		}
	}

	if count%2 == 0 {
		// 偶数个元素
		slow.Next, slow = nil, slow.Next
	} else {
		// 奇数个元素
		slow.Next, slow = nil, slow.Next.Next
	}

	slow = reverse(slow)

	n1 := head
	n2 := slow
	for ; n1 != nil && n2 != nil; n1, n2 = n1.Next, n2.Next {
		if n1.Val != n2.Val {
			return false
		}
	}

	return true
}

func main() {

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n1.Next = n2
		fmt.Println("1->2 is palindrome?", isPalindromeOpt1(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n1.Next = n2
		n2.Next = n3
		fmt.Println("1->2->3 is palindrome?", isPalindromeOpt1(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 1}
		n1.Next = n2
		n2.Next = n3
		fmt.Println("1->2->1 is palindrome?", isPalindromeOpt1(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		fmt.Println("1->2->3->4 is palindrome?", isPalindromeOpt1(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 2}
		n4 := &ListNode{Val: 1}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		fmt.Println("1->2->2->1 is palindrome?", isPalindromeOpt1(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n5 := &ListNode{Val: 5}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		fmt.Println("1->2->3->4->5 is palindrome?", isPalindromeOpt1(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 2}
		n5 := &ListNode{Val: 1}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		fmt.Println("1->2->3->2->1 is palindrome?", isPalindromeOpt1(n1))
	}
}
