package Sort

import "fmt"

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}

	mid := int((len(slice)) / 2)
	// fmt.Println(slice)
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(right, left []int) []int {

	size, top_right, top_left := len(right)+len(left), 0, 0
	result := make([]int, size)
	for index := 0; index < size; index++ {
		switch true {
		case top_left == len(left):
			// All numbers in left slice have been used
			// so we can only pick numbers from the right slice
			result[index] = right[top_right]
			fmt.Printf("result 1: %d\n", result)
			top_right += 1
		case top_right == len(right):
			// All numbers in right slice have been used
			// so we can only pick numbers from the left slice
			result[index] = left[top_left]
			fmt.Printf("result 2: %d\n", result)
			top_left += 1
		case right[top_right] > left[top_left]:
			// The "current number from right > current number from left"?
			// So we pick the lowest number (left) and put it on the result slice
			result[index] = left[top_left]
			fmt.Printf("result 3: %d\n", result)
			top_left += 1
		case right[top_right] < left[top_left]:
			//increment the index of a array because its present index
			//element is already appended to the result array
			result[index] = right[top_right]
			fmt.Printf("result 4: %d\n", result)
			top_right += 1
		case right[top_right] == left[top_left]:
			//in case of equality
			result[index] = left[top_left]
			fmt.Printf("result 5: %d\n", result)
			top_left += 1
		}
	}
	return result
}

// MergeSort funciona como orgnaizar cartas de um baralho
// Inicialmente começamos com varias cartas desorganizadas
// Quando dividimos por 2 cada porção desse baralho até chegar a ponta, tendo apenas uma unidade por lado podemos começar o Sort
// O sort funciona comparando os dois lados, vendo qual e maior.
// A divisão então vai refazendo o caminho (pra cima) organizando o array.
