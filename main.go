package main

import (
	"github.com/mellotonio/Golang-Algorithms/Search"
)

func main() {
	//numbers := []int{3, 2, 1}
	//	max := len(numbers)

	// Sort.SelectionSort(numbers, max)
	// Sort.BubbleSort(numbers, max)
	// Sort.RecursiveBubbleSort(numbers, max)
	// fmt.Println(Sort.MergeSort(numbers))
	// Quick Sort

	Search.BinarySearch([]int{3, 2, 1, 5, 6, 45, 57}, 42)

}
