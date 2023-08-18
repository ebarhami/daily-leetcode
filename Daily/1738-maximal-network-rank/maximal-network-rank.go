func maximalNetworkRank(n int, roads [][]int) int {
    graph := make(map[int]map[int]int)
    for i:=0;i<n;i++{
        graph[i] = make(map[int]int)
    }
    for _, road := range roads {
        graph[road[0]][road[1]] = 1
        graph[road[1]][road[0]] = 1
    }

    answer := 0
    for i:=0;i<n-1;i++{
        for j:=i+1;j<n;j++{
            tot := len(graph[i]) + len(graph[j]) - graph[i][j]
            answer = max(answer, tot)
        }
    }
    return answer
}

func max(a, b int) int {
    if a > b {return a}
    return b
}