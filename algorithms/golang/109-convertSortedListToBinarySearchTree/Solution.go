package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		if fast != nil {
			slow = slow.Next
		}
	}

	secondHead := slow.Next
	slow.Next = nil
	return secondHead
}

// 找到中间的节点作为根，然后递归构造左子树和右子树
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	mid := splitList(head)
	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}
