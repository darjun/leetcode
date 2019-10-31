package main

func majorityElementOpt1(nums []int) []int {
	var candidate1, count1 int
	var candidate2, count2 int
	var ans []int

	for _, num := range nums {
		if count1 > 0 && candidate1 == num {
			count1++
		} else if count2 > 0 && candidate2 == num {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1++
		} else if count2 == 0 {
			candidate2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		ans = append(ans, candidate1)
	}

	if count2 > len(nums)/3 {
		ans = append(ans, candidate2)
	}

	return ans
}
