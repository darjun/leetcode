package main

func twoSum(numbers []int, target int) []int {
	lo := 0
	hi := len(numbers) - 1
	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		if sum == target {
			return []int{lo + 1, hi + 1}
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}

	return nil
}
