package main

import "fmt"

func main() {
	{
		majorityChecker := Constructor([]int{1, 1, 2, 2, 1, 1})
		fmt.Println(majorityChecker.Query(0, 5, 4)) // returns 1
		fmt.Println(majorityChecker.Query(0, 3, 3)) // returns -1
		fmt.Println(majorityChecker.Query(2, 3, 2)) // returns 2
	}

	{
		// majorityChecker := ConstructorOpt1([]int{1, 1, 2, 2, 1, 1})
		// fmt.Println(majorityChecker.Query(0, 5, 4)) // returns 1
		// fmt.Println(majorityChecker.Query(0, 3, 3)) // returns -1
		// fmt.Println(majorityChecker.Query(2, 3, 2)) // returns 2
	}

	{
		majorityChecker := ConstructorOpt1([]int{2, 2, 1, 2, 1, 2, 2, 1, 1, 2})
		fmt.Println(majorityChecker.Query(5, 9, 5))
	}
}
