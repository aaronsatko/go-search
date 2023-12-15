package main

import (
	"fmt"

	"github.com/aaronsatko/go-search/search"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5

	linear := search.LinearSearch(data, target)
	fmt.Println("Linear Search")
	if linear != -1 {
		fmt.Printf("Element %d found at index %d\n", target, linear)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}

	fmt.Println("")

	binary := search.BinarySearch(data, target)
	fmt.Println("Binary Search")
	if binary != -1 {
		fmt.Printf("Element %d found at index %d\n", target, binary)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}

	fmt.Println("")

	fmt.Println("Jump Search")
	jump := search.JumpSearch(data, target)
	if jump != -1 {
		fmt.Printf("Element %d found at index %d\n", target, jump)
	} else {
		fmt.Println("Element not found in the array")
	}

	fmt.Println("")

	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}

	fmt.Println("Depth First Search")
	search.DepthFirstSearch(graph, "A")

	fmt.Println("")

	fmt.Println("Breadth First Search")
	search.BreadthFirstSearch(graph, "A")

}
