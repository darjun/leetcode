package main

func plusOne(digits []int) []int {
	ans := digits
	var maybee int
	var overflow bool
	if len(digits) == 0 || digits[0] == 9 {
		ans = make([]int, len(digits)+1, len(digits)+1)
		maybee = 1
	}

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if sum == 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		ans[i+maybee] = sum
	}

	if carry == 1 {
		ans[0] = 1
		overflow = true
	}

	if maybee == 1 && !overflow {
		return ans[1:]
	}

	return ans
}
