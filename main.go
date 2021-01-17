package main

import (
	"fmt"

	"github.com/mellotonio/Golang-Algorithms/Sort"
)

func main() {
	numbers := []int{3, 2, 1}
	//	max := len(numbers)

	// Sort.SelectionSort(numbers, max)
	// Sort.BubbleSort(numbers, max)
	// Sort.RecursiveBubbleSort(numbers, max)
	fmt.Println(Sort.MergeSort(numbers))
	// Quick Sort
}
