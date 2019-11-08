package main

func findDisappearedNumbers(nums []int) []int {
	var ans []int

	var i int
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1],nums[i]
		} else {
			i++
		}
	}

	for i, num := range nums {
		if num != i+1 {
			ans = append(ans, i+1)
		}
	}

    return ans
}