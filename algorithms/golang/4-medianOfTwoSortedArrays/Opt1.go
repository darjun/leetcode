package main

func findMedianSortedArraysOpt1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	lo := 0
	hi := m
	half := (m + n + 1) >> 1

	for lo < hi {
		i := hi + lo>>1
		j := half - i

		if i == 0 || nums1[i-1] <= nums2[j] {
			if j == 0 || nums2[j-1] <= nums1[i] {
				break
			}

			lo = i + 1
		} else {
			hi = i - 1
		}
	}

	i := hi + lo>>1
	j := half - i
	maxLeft := 0
	if i == 0 {
		maxLeft = nums2[j-1]
	} else if j == 0 {
		maxLeft = nums1[i-1]
	} else {
		if nums1[i-1] > nums2[j-1] {
			maxLeft = nums1[i-1]
		} else {
			maxLeft = nums2[j-1]
		}
	}

	if (m+n)%2 == 1 {
		return float64(maxLeft)
	}

	minRight := 0
	if i == m {
		minRight = nums2[j]
	} else if j == n {
		minRight = nums1[i]
	} else {
		if nums1[i] < nums2[j] {
			minRight = nums1[i]
		} else {
			minRight = nums2[j]
		}
	}

	return float64(maxLeft+minRight) / 2.0
}
