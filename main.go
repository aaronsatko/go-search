package main

import (
	"fmt"

	"github.com/aaronsatko/go-search/search"
)

func main() {
	// Example array to search in
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Define the value to search for
	target := 5

	// Call the search algorithm
	index := search.LinearSearch(data, target)

	// Display the result
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
