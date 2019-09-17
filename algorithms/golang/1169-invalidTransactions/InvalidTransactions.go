package main

import "fmt"

func main() {
	// fmt.Println(invalidTransactions([]string{"alice,20,800,mtv", "alice,50,100,beijing"}))

	// fmt.Println(invalidTransactions([]string{"alice,20,800,mtv", "alice,50,1200,mtv"}))

	// fmt.Println(invalidTransactions([]string{"alice,20,800,mtv", "bob,50,1200,mtv"}))

	fmt.Println(invalidTransactions([]string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"}))
}
