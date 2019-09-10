package main

func nextLargerNodesOpt1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}

	nodeVal := make([]int, 0, count)
	for p := head; p != nil; p = p.Next {
		nodeVal = append(nodeVal, p.Val)
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
