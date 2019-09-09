package main

func count(head *ListNode) int {
	c := 0
	for head != nil {
		c++
		head = head.Next
	}

	return c
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	count1 := count(l1)
	count2 := count(l2)

	p1 := l1
	p2 := l2
	var digits []byte
	if count1 > count2 {
		digits = make([]byte, 0, count1)
		for i := 0; i < count1-count2; i++ {
			digits = append(digits, byte(p1.Val))
			p1 = p1.Next
		}
	} else {
		digits = make([]byte, 0, count2)
		for i := 0; i < count2-count1; i++ {
			digits = append(digits, byte(p2.Val))
			p2 = p2.Next
		}
	}

	for p1 != nil {
		digits = append(digits, byte(p1.Val+p2.Val))
		p1 = p1.Next
		p2 = p2.Next
	}

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] >= 10 {
			digits[i] -= 10
			digits[i-1]++
		}
	}

	overflow := 0
	if len(digits) > 0 && digits[0] >= 10 {
		overflow = 1
		digits[0] -= 10
	}

	var head, tail *ListNode
	if overflow > 0 {
		head = &ListNode{Val: 1}
		tail = head
	}

	for _, digit := range digits {
		if tail == nil {
			head = &ListNode{Val: int(digit)}
			tail = head
		} else {
			tail.Next = &ListNode{Val: int(digit)}
			tail = tail.Next
		}
	}

	return head
}
