package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[n:], nums1[:m])
	var count int
	var i, j int
	i = n
	for i < m+n || j < n {
		if i == m+n {
			nums1[count] = nums2[j]
			j++
		} else if j == n {
			nums1[count] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			nums1[count] = nums1[i]
			i++
		} else {
			nums1[count] = nums2[j]
			j++
		}

		count++
	}
}
