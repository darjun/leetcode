package main

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}

	nextLarger := make([]int, 0, 64)
	nodeVal := make([]int, 0, 64)
	index := 1
	for head != nil {
		for i, val := range nodeVal {
			if nextLarger[i] != 0 {
				continue
			}

			if head.Val > val {
				nextLarger[i] = head.Val
			}
		}

		nodeVal = append(nodeVal, head.Val)
		nextLarger = append(nextLarger, 0)
		head = head.Next
		index++
	}

	return nextLarger
}
