package go_dedupesort

import (
	"sort"
)

// UniqueSort removes duplicates and sorts a slice of integers in increasing order.
func DedupeSort(array []int) []int {
	array = removeDuplicates(array)
	sort.Ints(array)

	return array
}

func removeDuplicates(array []int) []int {
	seen := map[int]bool{}
	result := []int{}

	for _, value := range array {
		if seen[value] == false {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}
