func countServers(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    col, row := make([]int, n), make([]int, m)

    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            if grid[i][j] == 1 {
                col[i]++
                row[j]++
            }
        }
    }

    answer := 0
    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            if grid[i][j] == 1 {
                if col[i] > 1 || row[j] > 1 {
                    answer++
                }
            }
        }
    }

    return answer
}
