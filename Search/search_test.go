package Search

import "testing"

// 0.067s ~ 0.0101s - 100.000 numbers (reversed 100.000 -> 1)
// Find the position of 80000
func BenchmarkBinarySearch(b *testing.B) {
	var mySlice []int
	for i := 100000; i > 0; i-- {
		mySlice = append(mySlice, i)
	}
	BinarySearch(mySlice, 80000)
}
