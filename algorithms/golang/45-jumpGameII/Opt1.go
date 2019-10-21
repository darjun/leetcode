package main

func jumpOpt1(nums []int) int {
	count := 0
	l := len(nums)
	middle := make(map[int]struct{})
	middle[0] = struct{}{}
	used := make([]bool, l, l)
	used[0] = true

	for _, exist := middle[l-1]; !exist; {
		middle2 := make(map[int]struct{})
		for p := range middle {
			if nums[p] >= l-p-1 {
				return count + 1
			}

			for i := p + 1; i <= p+nums[p]; i++ {
				if used[i] {
					continue
				}

				middle2[i] = struct{}{}
				used[i] = true
			}
		}

		middle = middle2

		count++
	}

	return 0
}
