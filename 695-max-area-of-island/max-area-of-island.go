var (
    dx = []int{-1, 0, 1, 0}
    dy = []int{0, 1, 0, -1}
)

func maxAreaOfIsland(grid [][]int) int {
    visited := make([][]bool, len(grid))
    for i:=0;i<len(grid);i++{
        visited[i] = make([]bool, len(grid[0]))
    }
    
    answer := 0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            area := getArea(i, j, visited, grid)
            if area > answer {
                answer = area
            }
        }
    }

    return answer
}

func getArea(x, y int, visited [][]bool, grid [][]int) int {
    if visited[x][y] || grid[x][y] == 0 {
        return 0
    }
    visited[x][y] = true
    area := 1 
    for i:=0;i<4;i++{
        xx := x + dx[i]
        yy := y + dy[i]
        if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[0]) {
            area += getArea(xx,yy, visited, grid)
        }
    }

    return area
}