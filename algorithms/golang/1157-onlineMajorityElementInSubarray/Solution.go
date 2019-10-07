package main

type MajorityChecker struct {
	arr []int
}

func Constructor(arr []int) MajorityChecker {
	return MajorityChecker{
		arr: arr,
	}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	candidate := this.arr[left]
	num := 1
	for _, value := range this.arr[left+1 : right+1] {
		if value == candidate {
			num++
		} else {
			num--
			if num == 0 {
				candidate = value
				num = 1
			}
		}
	}

	num = 0
	for _, value := range this.arr[left : right+1] {
		if value == candidate {
			num++
		}
	}

	if num >= threshold {
		return candidate
	}

	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
