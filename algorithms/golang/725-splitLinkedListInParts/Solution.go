package main

func splitListToParts(root *ListNode, k int) []*ListNode {
	parts := make([]*ListNode, k, k)
	count := 0
	for p := root; p != nil; p = p.Next {
		count++
	}

	var prelen int
	var prenum int
	if count%k == 0 {
		prelen = count / k
		prenum = k
	} else {
		prelen = count/k + 1
		prenum = k - (prelen*k - count)
	}

	i := 0
	for p := root; p != nil; {
		parts[i] = p
		for n := prelen - 1; n > 0; n-- {
			p = p.Next
		}

		if p != nil {
			next := p.Next
			p.Next = nil
			p = next
		}

		i++
		if i == prenum {
			// second part len - 1
			prelen--
		}
	}

	return parts
}
