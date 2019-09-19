package main

type SnapshotArray struct {
	data        [][][2]int // index -> slice of version data
	versionId   int
	snapId      int
	snapVersion []int
}

func Constructor(length int) SnapshotArray {
	data := make([][][2]int, length, length)
	for i := range data {
		data[i] = append(data[i], [2]int{0, 0})
	}

	return SnapshotArray{
		data: data,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.versionId++
	this.data[index] = append(this.data[index], [2]int{this.versionId, val})
}

func (this *SnapshotArray) Snap() int {
	this.snapId++
	this.snapVersion = append(this.snapVersion, this.versionId)
	return this.snapId - 1
}

func binarySearch(versionData [][2]int, version int) int {
	i := 0
	j := len(versionData) - 1

	for i <= j {
		mid := (i + j) >> 1
		if versionData[mid][0] == version {
			return versionData[mid][1]
		} else if versionData[mid][0] < version {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return versionData[j][1]
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return binarySearch(this.data[index], this.snapVersion[snap_id])
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
