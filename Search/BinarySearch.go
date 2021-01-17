package Search

import (
	"fmt"

	"github.com/mellotonio/Golang-Algorithms/Sort"
)

func BinarySearch(numbers []int, value int) {
	numbers = Sort.MergeSort(numbers)

	lowerBound := 0
	upperBound := len(numbers)

	var found bool = false
	var location int

	for found != true {
		if upperBound < lowerBound {
			found = true
			return
		}

		// Set mid
		midPoint := lowerBound + (upperBound-lowerBound)/2

		switch true {

		// +Right
		case numbers[midPoint] < value:
			lowerBound = midPoint + 1

		// +Left
		case numbers[midPoint] > value:
			upperBound = midPoint - 1

		// Found
		case numbers[midPoint] == value:
			found = true
			location = midPoint
			fmt.Printf("Location: %d\n", location)

		}

	}

}
