package Sort

import (
	"testing"
)

// 5.460s ~ 6s - 100.000 numbers (reversed 100.000 -> 1)
func BenchmarkSelectionSort(b *testing.B) {
	var mySlice []int
	for i := 100000; i > 0; i-- {
		mySlice = append(mySlice, i)
	}
	SelectionSort(mySlice, len(mySlice))
}

// 0.052s ~ 0.061s - 100.000 numbers (reversed 100.000 -> 1)
func BenchmarkMergeSort(b *testing.B) {
	var mySlice []int
	for i := 100000; i > 0; i-- {
		mySlice = append(mySlice, i)
	}
	MergeSort(mySlice)
}

// 10.318s ~ 12.479s - 100.000 numbers (reversed 100.000 -> 1)
func BenchmarkBubbleSort(b *testing.B) {
	var mySlice []int
	for i := 100000; i > 0; i-- {
		mySlice = append(mySlice, i)
	}
	BubbleSort(mySlice, len(mySlice))
}

// 4.116s ~ 4.205s - 100.000 numbers (reversed 100.000 -> 1)
func BenchmarkQuickSort(b *testing.B) {
	var mySlice []int
	for i := 100000; i > 0; i-- {
		mySlice = append(mySlice, i)
	}
	QuickSort(mySlice, 0, len(mySlice)-1)

}
