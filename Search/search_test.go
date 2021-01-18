package Search

import "testing"

// 1,182s ~ 1.211s - 10.000.000 numbers (reversed 100.000 -> 1)
// Find the position of 87654
func BenchmarkBinarySearch(b *testing.B) {
	var mySlice []int
	for i := 10000000; i > 0; i-- {
		mySlice = append(mySlice, i)
	}
	BinarySearch(mySlice, 87654)
}
