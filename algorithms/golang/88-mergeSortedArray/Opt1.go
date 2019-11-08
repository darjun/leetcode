package main

func mergeOpt1(nums1 []int, m int, nums2 []int, n int) {
	i := m
	j := n
	p := m + n
	for i > 0 || j > 0 {
		if i == 0 {
			nums1[p-1] = nums2[j-1]
			j--
		} else if j == 0 {
			nums1[p-1] = nums1[i-1]
			i--
		} else if nums1[i-1] < nums2[j-1] {
			nums1[p-1] = nums2[j-1]
			j--
		} else {
			nums1[p-1] = nums1[i-1]
			i--
		}

		p--
	}
}
