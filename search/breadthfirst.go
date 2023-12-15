package search

import (
	"fmt"

	"github.com/aaronsatko/go-search/datastructures"
)

func BreadthFirstSearch(graph map[string][]string, startNode string) {
	q := datastructures.NewQueue()
	visited := make(map[string]bool)

	q.Enqueue(startNode)
	visited[startNode] = true

	for !q.IsEmpty() {
		current := q.Dequeue().(string)
		fmt.Println(current)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				q.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}
}
