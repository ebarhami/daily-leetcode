func findCircleNum(isConnected [][]int) int {
    visited := make([]bool, len(isConnected))
    province := 0
    for i:=0;i<len(isConnected);i++ {
        if !visited[i] {
            province++
            dfs(isConnected, visited, i)
        }
    }

    return province
}

func dfs(isConnected [][]int, visited []bool, i int) {
    if visited[i] {
        return
    }
    visited[i] = true

    for neighbor, connect := range isConnected[i] {
        if connect == 1 {
            dfs(isConnected, visited, neighbor)
        }
    }
}