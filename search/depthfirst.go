package search

import "fmt"

func DepthFirstSearch(graph map[string][]string, startNode string) {
	visited := make(map[string]bool)

	var dfs func(node string)
	dfs = func(node string) {
		if !visited[node] {
			fmt.Println(node)
			visited[node] = true
			for _, neighbor := range graph[node] {
				dfs(neighbor)
			}
		}
	}
	dfs(startNode)
}
