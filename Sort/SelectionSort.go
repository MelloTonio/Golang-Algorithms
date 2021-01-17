package Sort

import "fmt"

// 1. Find min element between [minimum_index...len(slice)]
// 2. Place this element in the minimum_index
// example -> arr[] = 64 25 12 22 11 (Find the minimum element in arr[0...4] and place it at the beginning)
// 11 25 12 22 64 (minimum_index = 1, so the algorithm must find an element between [1...len(slice)])
// 11 12 25 22 64 (minimum_index = 2 -> [2...len(slice)])
// ...
func SelectionSort(numbers []int, max int) {
	for x := 0; x < max-1; x++ {
		var min_index = x
		for y := x; y < max; y++ {
			if numbers[y] < numbers[min_index] {
				min_index = y
			}
		}
		// numbers[x], numbers[min_index] = numbers[min_index], numbers[x]
		swapSelection(numbers, x, min_index)
	}

	fmt.Printf("SelectionSort -> %v ", numbers)
}

func swapSelection(currentList []int, currentIdx int, minIdx int) {
	temp := currentList[minIdx]
	currentList[minIdx] = currentList[currentIdx]
	currentList[currentIdx] = temp
}
