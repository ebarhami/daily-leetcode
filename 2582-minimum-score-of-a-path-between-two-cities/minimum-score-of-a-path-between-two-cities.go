func minScore(n int, roads [][]int) int {
	graph := make([][][2]int, n+1)

	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		graph[u] = append(graph[u], [2]int{v, w})
		graph[v] = append(graph[v], [2]int{u, w})
	}

	visited := make([]bool, n+1)
	ans := int(^uint(0) >> 1) // MaxInt

	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true

		for _, edge := range graph[u] {
			v, w := edge[0], edge[1]
			if w < ans {
				ans = w
			}
			if !visited[v] {
				dfs(v)
			}
		}
	}

	dfs(1)
	return ans
}