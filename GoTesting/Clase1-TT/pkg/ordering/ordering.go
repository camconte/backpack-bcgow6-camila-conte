package ordering

import "sort"

func OrderIntSlice (numbers ...int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		//ascending order
		return numbers[i] < numbers[j]
		//descending order
		//return numbers[i] > numbers[j]
	})
	return numbers
}