package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	d1 := 0
	for i := start; i != destination; i = (i + 1) % len(distance) {
		d1 += distance[i]
	}

	d2 := 0
	for i := destination; i != start; i = (i + 1) % len(distance) {
		d2 += distance[i]
	}

	if d1 < d2 {
		return d1
	}

	return d2
}
