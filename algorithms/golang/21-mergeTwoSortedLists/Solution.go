package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		var newNode *ListNode
		if l1 == nil || (l2 != nil && l1.Val >= l2.Val) {
			newNode = l2
			l2 = l2.Next
		} else {
			newNode = l1
			l1 = l1.Next
		}

		newNode.Next = nil
		if head == nil {
			head = newNode
		} else {
			tail.Next = newNode
		}
		tail = newNode
	}

	return head
}
