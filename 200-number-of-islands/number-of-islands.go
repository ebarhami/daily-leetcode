var (
    dx = []int{-1, 0, 1, 0}
    dy = []int{0, 1, 0, -1}
)

func numIslands(grid [][]byte) int {
    color := 0
    mep := make([][]int, len(grid))
    for i:=0;i<len(grid);i++{
        mep[i] = make([]int, len(grid[0]))
    }

    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j] == '1' && mep[i][j] == 0 {
                color++
                dfs(grid, i, j, color, mep)
            }
        }
    }

    return color
}

func dfs(grid [][]byte, x, y int, color int, mep [][]int) {
    if mep[x][y] > 0 || grid[x][y] == '0' {
        return
    }
    mep[x][y] = color
    for i:=0;i<4;i++{
        xx := x + dx[i]
        yy := y + dy[i]
        if xx >= 0 && xx < len(grid) && yy >=0 && yy < len(grid[0]) && grid[xx][yy] == '1' {
            dfs(grid, xx, yy, color, mep)
        }
    }
}