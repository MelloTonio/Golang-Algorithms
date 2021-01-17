package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 34, 54, 2, 3, 5, 45, 34, 65, 67, 1245, 8, 90}
	// max := len(numbers)

	// SelectionSort(numbers, max)
	// BubbleSort(numbers, max)
	// RecursiveBubbleSort(numbers, max)
	fmt.Println(MergeSort(numbers))
	// Quick Sort
}

// BubbleSort
// Percorre a lista (len(lista) * len(lista)) vezes
func BubbleSort(numbers []int, max int) {
	for outside := 0; outside < max-1; outside++ {
		for inside := 0; inside < max-1; inside++ {
			// Se o numero atual for maior que o proximo
			if numbers[inside] > numbers[inside+1] {
				// Então trocamos eles de lugar
				swapBubble(numbers, inside)
			}
		}
	}
	fmt.Println("BubbleSort -> %d", numbers)
}

func swapBubble(currentList []int, current int) {
	temp := currentList[current+1]
	currentList[current+1] = currentList[current]
	currentList[current] = temp
}

// Selection Sort
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

	fmt.Println("SelectionSort -> %d", numbers)
}

func swapSelection(currentList []int, currentIdx int, minIdx int) {
	temp := currentList[minIdx]
	currentList[minIdx] = currentList[currentIdx]
	currentList[currentIdx] = temp
}

func RecursiveBubbleSort(numbers []int, max int) {
	if max == 1 {
		fmt.Println("Recursive BubbleSort -> %d", numbers)
		return
	}
	for outside := 0; outside < max-1; outside++ {
		if numbers[outside] > numbers[outside+1] {
			// Então trocamos eles de lugar
			swapBubble(numbers, outside)
		}
		// fmt.Println("step %d -> %d", outside, numbers)
	}

	// Max - 1 -> after the first iteration the higher number will be in the end.
	RecursiveBubbleSort(numbers, max-1)

}

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}

	mid := (len(slice)) / 2
	// fmt.Println(slice)
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}
func Merge(left, right []int) []int {

	size, top_left, top_right := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for index := 0; index < size; index++ {
		// Se tivermos esgotado as possibilidades no lado esquerdo e o lado direito estiver disponivel
		// Incrementamos no lado direito
		if top_left > len(left)-1 && top_right <= len(right)-1 {
			slice[index] = right[top_right]
			top_right++
			// Oposto do de cima
		} else if top_right > len(right)-1 && top_left <= len(left)-1 {
			slice[index] = left[top_left]
			top_left++
			// Se o topo do da esquerda for menor do que o topo da esquerda
		} else if left[top_left] < right[top_right] {
			// Colocamos ele na frente
			slice[index] = left[top_left]
			top_left++
		} else {
			// Se não o direito que ficará na frente
			slice[index] = right[top_right]
			top_right++
		}
	}
	// fmt.Println(slice)
	return slice
}
