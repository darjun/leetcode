package main

func find(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return -1
}

func helper(indice map[int]int, inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rval := postorder[len(postorder)-1]
	root := &TreeNode{Val: rval}
	rindex := find(inorder, rval)
	leftInorder := inorder[:rindex]
	rightInorder := inorder[rindex+1:]
	root.Left = helper(indice, leftInorder, postorder[:len(leftInorder)])
	root.Right = helper(indice, rightInorder, postorder[len(leftInorder):len(postorder)-1])
	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	indice := make(map[int]int, len(inorder))
	for i, num := range inorder {
		indice[num] = i
	}

	return helper(indice, inorder, postorder)
}
