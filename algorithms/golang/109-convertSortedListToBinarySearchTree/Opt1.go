package main

// 找到中间的节点作为根，然后递归构造左子树和右子树
func sortedListToBSTOpt1(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	preSlow := (*ListNode)(nil)
	slow := head
	fast := head
	// 执行到这里 fast.Next != nil，所以循环至少会被执行一次，preSlow 一定不为nil
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	preSlow.Next = nil

	root := &TreeNode{Val: slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}
