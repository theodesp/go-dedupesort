package main

import (
	"fmt"
	utils "github.com/theodesp/go-dedupesort"
)

func main() {
	array := []int{7, 2, 4, 2, 1, 2}

	fmt.Println(utils.DedupeSort(array)) // prints [1 2 4 7]
}
