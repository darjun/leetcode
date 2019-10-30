package main

func find(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return -1
}

func helperOpt1(indice map[int]int, preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rval := preorder[0]
	root := &TreeNode{Val: rval}
	rindex := find(inorder, rval)
	leftInorder := inorder[:rindex]
	rightInorder := inorder[rindex+1:]
	root.Left = helperOpt1(indice, preorder[1:1+len(leftInorder)], leftInorder)
	root.Right = helperOpt1(indice, preorder[1+len(leftInorder):], rightInorder)
	return root
}

func buildTreeOpt1(preorder []int, inorder []int) *TreeNode {
	indice := make(map[int]int, len(inorder))
	for i, num := range inorder {
		indice[num] = i
	}

	return helperOpt1(indice, preorder, inorder)
}
