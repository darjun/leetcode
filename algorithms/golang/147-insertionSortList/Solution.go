package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	next := head.Next

	// 排序的过程中反转，方便插入元素
	for cur != nil {
		cur.Next = prev
		val := cur.Val

		q := cur
		for p := prev; p != nil && p.Val > val; p, q = p.Next, p {
			q.Val = p.Val
		}
		q.Val = val

		prev, cur = cur, next
		if next != nil {
			next = next.Next
		}
	}

	// 排序完成后再次反转回来
	return reverse(prev)
}
