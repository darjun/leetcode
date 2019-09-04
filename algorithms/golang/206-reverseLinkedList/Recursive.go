package main

// reverseRecursively 递归解法工具函数，返回新的头部和尾部
func reverseRecursively(head *ListNode) (newHead, newTail *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	newHead, newTail = reverseRecursively(head.Next)
	newTail.Next = head
	newTail = head
	newTail.Next = nil
	return
}

func reverse2(head *ListNode) *ListNode {
	newHead, _ := reverseRecursively(head)
	return newHead
}
