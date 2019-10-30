package main

func helper(indice map[int]int, preorder, inorder []int, prestart, preend int, instart, inend int) *TreeNode {
	if prestart > preend {
		return nil
	}

	rval := preorder[prestart]
	root := &TreeNode{Val: rval}
	leftInorderLen := indice[rval] - instart
	root.Left = helper(indice, preorder, inorder, prestart+1, prestart+leftInorderLen, instart, indice[rval]-1)
	root.Right = helper(indice, preorder, inorder, prestart+leftInorderLen+1, preend, indice[rval]+1, inend)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	indice := make(map[int]int, len(inorder))
	for i, num := range inorder {
		indice[num] = i
	}

	return helper(indice, preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}
