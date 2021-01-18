package Sort

func partition(number []int, start, end int) int {
	pivot := number[end]

	for index := start; index < end; index++ {
		// if the current element is smaller than the pivot
		if number[index] < pivot {
			// swap the current number with the number "start"
			number[index], number[start] = number[start], number[index]
			// Advance 1 block
			start++
		}
	}

	number[start], pivot = pivot, number[start]

	// Everything before "start" is smaller than pivot, everything after "start" is greater than pivot
	return start
}

func QuickSort(numbers []int, start, end int) []int {

	if start > end {
		return []int{}
	}

	pivot := partition(numbers, start, end)

	// numbers to the Left <-
	QuickSort(numbers, start, pivot-1)

	// numbers to the Right ->
	QuickSort(numbers, pivot+1, end)

	return numbers
}
