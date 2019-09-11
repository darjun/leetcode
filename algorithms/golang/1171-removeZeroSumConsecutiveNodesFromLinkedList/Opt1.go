package main

// 如果sum[0,i] = sum[0, j] (j >= i)，那么sum[i+1, j] = 0
// 完美解法！！！
func removeZeroSumSublistsOpt1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	// sum from node 0 to value
	m := make(map[int]*ListNode)
	sum := 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		m[sum] = p
	}

	sum = 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		p.Next = m[sum].Next
	}

	return dummy.Next
}
