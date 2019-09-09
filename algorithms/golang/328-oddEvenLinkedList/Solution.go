package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead := head
	oddTail := head
	evenHead := head.Next
	evenTail := head.Next

	p := head.Next.Next
	for p != nil {
		oddTail.Next = p
		oddTail = p
		p = p.Next

		if p != nil {
			evenTail.Next = p
			evenTail = p
			p = p.Next
		}
	}

	oddTail.Next = evenHead
	evenTail.Next = nil
	return oddHead
}
