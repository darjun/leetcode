package main

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]struct{}, len(G))
	for _, v := range G {
		m[v] = struct{}{}
	}

	count := 0
outer:
	for p := head; p != nil; p = p.Next {
		if _, exist := m[p.Val]; exist {
			count++

			p = p.Next
			for {
				if p == nil {
					break outer
				}

				_, exist = m[p.Val]
				if !exist {
					break
				}

				p = p.Next
			}
		}
	}

	return count
}
