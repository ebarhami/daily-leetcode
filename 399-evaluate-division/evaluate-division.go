func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    graph := make(map[string]map[string]float64)
    for i:=0;i<len(values);i++{
        a, b := equations[i][0], equations[i][1]
        if _, ok := graph[a]; !ok {
            graph[a] = make(map[string]float64)
        }
        if _, ok := graph[b]; !ok {
            graph[b] = make(map[string]float64)
        }
        graph[a][b] = values[i]
        graph[b][a] = 1/values[i]
    }

    answer := make([]float64, len(queries))
    for i, query := range queries {
        start, target := query[0], query[1]
        visited := make(map[string]bool) // can use struct for more optimize space
        var dfs func(string, float64) float64
        dfs = func(curr string, val float64) float64{
            if _, ok := graph[target]; ok && curr == target {
                return val
            }
            visited[curr] = true
            answer := -1.0

            for adj, nextVal := range graph[curr] {
                if !visited[adj] {
                    res := dfs(adj, val * nextVal)
                    if res > answer {
                        answer = res
                    }
                }
            }
            return answer
        }

        answer[i] = dfs(start, 1.0)
    }

    return answer
}

/*
a/d

a/b 1, b/c 2 , c/d 3 -> a/d 9 
a/c 3

start end

*/