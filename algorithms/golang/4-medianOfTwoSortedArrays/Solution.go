package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	low := 0
	high := m
	half := (m + n + 1) >> 1
	for {
		i := high + low>>1
		j := half - i
		if i < high && nums2[j-1] > nums1[i] {
			low = i + 1
		} else if i > low && nums1[i-1] > nums2[j] {
			high = i - 1
		} else {
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
				if nums1[i] > nums2[j] {
					minRight = nums2[j]
				} else {
					minRight = nums1[i]
				}
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
}
