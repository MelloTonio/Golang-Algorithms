package Sort

import "fmt"

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
	fmt.Printf("BubbleSort -> %d ", numbers)
}

func RecursiveBubbleSort(numbers []int, max int) {
	if max == 1 {
		fmt.Printf("Recursive BubbleSort -> %d", numbers)
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

func swapBubble(currentList []int, current int) {
	temp := currentList[current+1]
	currentList[current+1] = currentList[current]
	currentList[current] = temp
}
