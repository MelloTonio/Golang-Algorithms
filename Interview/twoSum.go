package Interview

import (
	"fmt"
)

/*Given an array of integers nums and an integer target, 
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, 
and you may not use the same element twice.

You can return the answer in any order.
*/

func findTwoSum(targetSum int, numbers []int) []int {
	myMap := map[int]int{}
	for i, v := range numbers {
		complement := targetSum - v

		if _, ok := myMap[complement]; ok {
			return []int{myMap[complement], i}
		}

		myMap[numbers[i]] = i
		fmt.Println(numbers[i], i, complement)
	}
	return nil
}
