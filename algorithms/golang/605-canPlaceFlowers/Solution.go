package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var maxNum int
	leftOpen := 1
	rightOpen := 0
	var continuous int
	for _, value := range flowerbed {
		if value == 0 {
			continuous++
		} else {
			maxNum += (continuous + leftOpen + rightOpen - 1) / 2
			continuous = 0
			leftOpen = 0
		}
	}

	if continuous > 0 {
		rightOpen = 1
		maxNum += (continuous + leftOpen + rightOpen - 1) / 2
	}

	return maxNum >= n
}
