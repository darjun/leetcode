package main

import "fmt"

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastIntervalOpt1([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
