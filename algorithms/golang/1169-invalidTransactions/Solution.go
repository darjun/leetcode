package main

import (
	"math"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	times := make([]int64, 0, len(transactions))
	names := make([]string, 0, len(transactions))
	cities := make([]string, 0, len(transactions))
	indices := make([]bool, len(transactions), len(transactions))

	for i, transaction := range transactions {
		parts := strings.Split(transaction, ",")

		name := parts[0]
		time, _ := strconv.ParseInt(parts[1], 10, 64)
		amount, _ := strconv.ParseInt(parts[2], 10, 64)
		city := parts[3]

		if amount > 1000 {
			indices[i] = true
		}
		for j := 0; j < i; j++ {
			if math.Abs(float64(times[j]-time)) <= 60 && names[j] == name && cities[j] != city {
				indices[j] = true
				indices[i] = true
			}
		}

		names = append(names, name)
		times = append(times, time)
		cities = append(cities, city)
	}

	invalids := make([]string, 0, len(transactions))
	for i, invalid := range indices {
		if invalid {
			invalids = append(invalids, transactions[i])
		}
	}
	return invalids
}
