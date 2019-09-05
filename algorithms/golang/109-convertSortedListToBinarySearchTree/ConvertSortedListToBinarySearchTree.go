package main

import "fmt"

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("<nil>")
		return
	}

	level := make([]*TreeNode, 0, 1)

	level = append(level, root)
	for len(level) > 0 {
		nextLevel := level[0:0]
		for _, node := range level {
			fmt.Print(node.Val, " ")

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		level = nextLevel
		fmt.Println()
	}
}

func main() {
	{
		n1 := &ListNode{Val: -10}
		n2 := &ListNode{Val: -3}
		n3 := &ListNode{Val: 0}
		n4 := &ListNode{Val: 5}
		n5 := &ListNode{Val: 9}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		r := sortedListToBSTOpt1(n1)
		printTree(r)
	}
}
