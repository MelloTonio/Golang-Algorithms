package main

import (
	"fmt"

	"github.com/mellotonio/Golang-Algorithms/Sort"
)

func main() {
	numbers := []int{10, 654, 34, 4, 23, 56, 4, 346, 5, 23, 43, 54}
	//	max := len(numbers)

	// Sort.SelectionSort(numbers, max)
	// Sort.BubbleSort(numbers, max)
	// Sort.RecursiveBubbleSort(numbers, max)
	// fmt.Println(Sort.MergeSort(numbers))
	// Quick Sort

	x := Sort.QuickSort(numbers, 0, len(numbers)-1)

	fmt.Println(x)

}
