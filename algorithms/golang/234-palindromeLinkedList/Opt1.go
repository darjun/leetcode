package main

func isPalindromeOpt1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	preSlow := (*ListNode)(nil)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		preSlow = slow
		slow = slow.Next
	}

	preSlow.Next = nil
	if fast == nil {
		// 偶数个元素
	} else {
		// 奇数个元素
		slow = slow.Next
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
