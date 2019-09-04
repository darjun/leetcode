package main

func listLength(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}

	return len
}

// 由于相同后缀的长度必定相同。
// 先计算两个列表的长度，然后较长列表指针向后移动，让其到尾部的距离和较短列表头部到尾部的距离相同。
// 最后每次移动两个列表的指针，比较是否相同，相同则返回当前指针。如果到尾部了，还没有相同的节点，那么返回nil。
func getIntersectionNodeOpt1(headA, headB *ListNode) *ListNode {
	lenA := listLength(headA)
	lenB := listLength(headB)

	if lenA > lenB {
		gap := lenA - lenB
		for gap > 0 {
			headA = headA.Next
			gap--
		}
	} else {
		gap := lenB - lenA
		for gap > 0 {
			headB = headB.Next
			gap--
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
