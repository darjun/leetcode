package main

func nextLargerNodesOpt2(head *ListNode) []int {
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
	// position i save the first index whose val is larger than the val at index i
	nextLarger := make([]int, count, count)

	nextLarger[count-1] = -1
	for i := count - 2; i >= 0; i-- {
		cur := nodeVal[i]
		next := nodeVal[i+1]
		if next > cur {
			nextLarger[i] = i + 1
		} else {
			nextLargerIndex := nextLarger[i+1]
			for {
				if nextLargerIndex == -1 {
					nextLarger[i] = -1
					break
				}

				if nodeVal[nextLargerIndex] > cur {
					nextLarger[i] = nextLargerIndex
					break
				}
				nextLargerIndex = nextLarger[nextLargerIndex]
			}
		}
	}

	for i, index := range nextLarger {
		if index == -1 {
			nextLarger[i] = 0
		} else {
			nextLarger[i] = nodeVal[index]
		}
	}
	return nextLarger
}
