package main

func mergeOpt1(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	p := m + n - 1
	for i > 0 || j > 0 {
		if i == 0 {
			nums1[p] = nums2[j]
			j--
		} else if j == 0 {
			nums1[p] = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[p] = nums2[j]
			j--
		} else {
			nums1[p] = nums1[i]
			i--
		}

		p--
	}
}
