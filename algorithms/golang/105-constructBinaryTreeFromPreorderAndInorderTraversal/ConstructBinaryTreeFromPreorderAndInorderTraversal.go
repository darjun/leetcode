package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) {
	level := make([]*TreeNode, 0, 1)
	level = append(level, root)
	for len(level) > 0 {
		newLevel := make([]*TreeNode, 0, len(level))
		for _, n := range level {
			if n.Left != nil {
				newLevel = append(newLevel, n.Left)
			}

			if n.Right != nil {
				newLevel = append(newLevel, n.Right)
			}

			fmt.Print(n.Val, " ")
		}
		fmt.Println()

		level = newLevel
	}
}

func main() {
	printTree(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))

	printTree(buildTreeOpt1([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
