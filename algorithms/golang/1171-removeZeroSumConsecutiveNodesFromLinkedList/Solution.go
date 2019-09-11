package main

func removeNodes(head *ListNode, i, j int) *ListNode {
	var prev *ListNode
	cur := head
	for i > 0 {
		prev = cur
		cur = cur.Next
		i--
		j--
	}

	for j > 0 {
		cur = cur.Next
		j--
	}

	if prev == nil {
		// i = 0, remove head
		return cur.Next
	}

	prev.Next = cur.Next
	return head
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	// remove prefix zero
	for head != nil && head.Val == 0 {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	// remove node whose value is 0 & count node
	prev := head
	cur := head.Next
	count := 1 // head's val is not zero
	for cur != nil {
		if cur.Val == 0 {
			prev.Next = cur.Next
			cur = prev.Next
		} else {
			prev = cur
			cur = cur.Next
			count++
		}
	}

	sum := make([]int, count, count) // sum[i] represents the sum from node 0 to i.
	i := 0
	for p := head; p != nil; p = p.Next {
		if i == 0 {
			sum[i] = p.Val
		} else {
			sum[i] = sum[i-1] + p.Val
		}
		i++
	}

	for i := 0; i < count; i++ {
		for j := count - 1; j > i; j-- {
			var sumFromItoJ int
			if i == 0 {
				sumFromItoJ = sum[j]
			} else {
				sumFromItoJ = sum[j] - sum[i-1]
			}

			if sumFromItoJ == 0 {
				head = removeNodes(head, i, j)
				count -= j - i + 1
				copy(sum[i:], sum[j+1:])
				break
			}
		}
	}

	return head
}
