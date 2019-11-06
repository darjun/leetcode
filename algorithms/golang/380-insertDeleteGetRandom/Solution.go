package main

import "math/rand"

type RandomizedSet struct {
	data map[int]struct{}
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		data: make(map[int]struct{}),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.data[val]; exist {
		return false
	}

	this.data[val] = struct{}{}
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, exist := this.data[val]; exist {
		delete(this.data, val)
		return true
	}

	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	count := rand.Int31n(int32(len(this.data)))
	for val := range this.data {
		if count == 0 {
			return val
		}

		count--
	}

	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
