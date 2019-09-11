package main

// 1.创建一对快慢指针，初始指向head，慢指针每次移动一个节点，快指针每次移动两个节点。
// 2.直到两个指针相遇。
// 3.在该点上，慢指针移动了(x+i)，快指针移动了(x+i+y)，其中x为环前面的长度，y是环的长度，i是相遇时慢指针在环中移动的长度。
//   这时(x+i+y)=2*(x+i) -> y=x+i，也就是说慢指针在环中还未移动的长度刚好是x
// 4.新建指针指向head，与慢指针一起移动，相遇点肯定就是环的起点
func detectCycleOpt1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			laterNode := head
			for laterNode != slow {
				laterNode = laterNode.Next
				slow = slow.Next
			}

			return slow
		}
	}

	return nil
}
