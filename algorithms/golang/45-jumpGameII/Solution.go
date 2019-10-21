package main

func jump(nums []int) int {
	count := 0
	l := len(nums)
	middle := make(map[int]struct{})
	middle[l-1] = struct{}{}
	used := make([]bool, l, l)
	used[l-1] = true

	for _, exist := middle[0]; !exist; {
		middle2 := make(map[int]struct{})
		for p := range middle {
			if nums[0] >= p {
				return count + 1
			}

			for i := 1; i < p; i++ {
				if used[i] {
					continue
				}

				if nums[i] >= p-i {
					middle2[i] = struct{}{}
					used[i] = true
				}
			}
		}

		middle = middle2

		count++
	}

	return 0
}
