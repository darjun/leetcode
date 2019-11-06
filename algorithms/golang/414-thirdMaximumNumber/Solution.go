package main

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max []int

	for _, num := range nums {
		var hasInsert bool
		for i, m := range max {
			if num == m {
				hasInsert = true
				break
			}

			if num > m {
				if len(max) < 3 {
					max = append(max, 0)
				}
				for j := len(max) - 1; j > i; j-- {
					max[j] = max[j-1]
				}
				max[i] = num
				hasInsert = true
				break
			}
		}

		if !hasInsert && len(max) < 3 {
			max = append(max, num)
		}
	}

	if len(max) < 3 {
		return max[0]
	}

	return max[2]
}
