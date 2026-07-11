type Graph struct {
	node  int
	edges int
}

func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	ans := 0

	var dfs func(int, *Graph)
	dfs = func(u int, curr *Graph) {
		visited[u] = true
		curr.node++
		curr.edges += len(graph[u])

		for _, v := range graph[u] {
			if !visited[v] {
				dfs(v, curr)
			}
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		curr := &Graph{}
		dfs(i, curr)

		// each edge is counted twice
		totalEdges := curr.edges / 2

		if totalEdges == curr.node*(curr.node-1)/2 {
			ans++
		}
	}

	return ans
}