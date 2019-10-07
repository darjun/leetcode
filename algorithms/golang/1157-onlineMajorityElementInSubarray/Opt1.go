package main

type Node struct {
	l, r        int
	majority    int
	left, right *Node
}

type MajorityCheckerOpt1 struct {
	arr     []int
	leftCnt map[int][][2]int
	root    *Node
}

func ConstructorOpt1(arr []int) MajorityCheckerOpt1 {
	leftCnt := make(map[int][][2]int)
	for i, elem := range arr {
		if cnts, exist := leftCnt[elem]; !exist {
			leftCnt[elem] = [][2]int{[2]int{i, 1}}
		} else {
			lastEntry := cnts[len(cnts)-1]
			leftCnt[elem] = append(cnts, [2]int{i, lastEntry[1] + 1})
		}
	}

	m := MajorityCheckerOpt1{
		arr:     arr,
		leftCnt: leftCnt,
	}

	m.root = m.buildTree(0, len(arr)-1)
	return m
}

func (this *MajorityCheckerOpt1) buildTree(left, right int) *Node {
	if left == right {
		return &Node{left, right, this.arr[left], nil, nil}
	}

	mid := left + (right-left)>>1
	n := &Node{l: left, r: right}
	n.left = this.buildTree(left, mid)
	n.right = this.buildTree(mid+1, right)
	leftMaj, rightMaj := n.left.majority, n.right.majority
	if leftMaj == rightMaj {
		n.majority = leftMaj
		return n
	}

	leftMajCount := this.count(left, right, leftMaj)
	rightMajCount := this.count(left, right, rightMaj)
	if 2*leftMajCount > right-left+1 {
		n.majority = leftMaj
	} else if 2*rightMajCount > right-left+1 {
		n.majority = rightMaj
	} else {
		n.majority = -1
	}
	return n
}

func (this *MajorityCheckerOpt1) binarySearch(cnts [][2]int, index int) int {
	low := 0
	high := len(cnts) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if cnts[mid][0] == index {
			return cnts[mid][1]
		} else if cnts[mid][0] < index {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if high < 0 {
		return 0
	}

	return cnts[high][1]
}

func (this *MajorityCheckerOpt1) count(left, right, maj int) int {
	indexToCnt := this.leftCnt[maj]
	if indexToCnt == nil {
		return 0
	}

	leftCnt := this.binarySearch(indexToCnt, left-1)
	rightCnt := this.binarySearch(indexToCnt, right)

	return rightCnt - leftCnt
}

func (this *MajorityCheckerOpt1) getMaj(left, right int, n *Node) int {
	if n.l == left && n.r == right {
		return n.majority
	}

	mid := n.l + (n.r-n.l)>>1
	if mid >= right {
		return this.getMaj(left, right, n.left)
	}

	if mid < left {
		return this.getMaj(left, right, n.right)
	}

	leftMaj := this.getMaj(left, mid, n.left)
	rightMaj := this.getMaj(mid+1, right, n.right)

	if leftMaj == rightMaj {
		return leftMaj
	}

	leftMajCount := this.count(left, right, leftMaj)
	if 2*leftMajCount > right-left+1 {
		return leftMaj
	}

	rightMajCount := this.count(left, right, rightMaj)
	if 2*rightMajCount > right-left+1 {
		return rightMaj
	}

	return -1
}

func (this *MajorityCheckerOpt1) Query(left int, right int, threshold int) int {
	maj := this.getMaj(left, right, this.root)
	if maj == -1 {
		return -1
	}

	cnt := this.count(left, right, maj)
	if cnt >= threshold {
		return maj
	}

	return -1
}
