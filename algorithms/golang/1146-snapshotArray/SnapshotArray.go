package main

import "fmt"

func main() {
	sa := Constructor(3)
	sa.Set(0, 5)
	sa.Snap()
	sa.Set(0, 6)
	fmt.Println(sa.Get(0, 0))
}
