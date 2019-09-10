package main

func nextLargerNodesOpt1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}

	nodeVal := make([]int, count, count)
	p := head
	for i := 0; i < count; i++ {
		nodeVal[i] = p.Val
		p = p.Next
	}

	nextLarger := make([]int, count, count)

	for i, val1 := range nodeVal {
		for _, val2 := range nodeVal[i+1:] {
			if val2 > val1 {
				nextLarger[i] = val2
				break
			}
		}
	}

	return nextLarger
}
