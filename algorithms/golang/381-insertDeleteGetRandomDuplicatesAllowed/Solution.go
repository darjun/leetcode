package main

import "math/rand"

type RandomizedCollection struct {
	data map[int]int
	num  int32
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		data: make(map[int]int),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	count := this.data[val]
	this.data[val] = count + 1
	this.num++
	return count == 0
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	count := this.data[val]
	if count == 0 {
		return false
	}

	this.data[val] = count - 1
	this.num--
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	prob := int(rand.Int31n(this.num))
	for val, count := range this.data {
		if prob < count {
			return val
		}
		prob -= count
	}

	return 0
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
