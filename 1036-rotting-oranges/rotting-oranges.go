var (
    dx = []int{-1, 0, 1, 0}
    dy = []int{0, 1, 0, -1}
)

type cell struct {
    idx int
    idy int
}

func orangesRotting(grid [][]int) int {
    fresh := 0
    queue := make([]cell, 0)
    n,m := len(grid), len(grid[0])

    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j] == 1 {
                fresh++
            } else if grid[i][j] == 2 {
                queue = append(queue, cell{
                    idx: i,
                    idy: j,
                })
            }
        }
    }

    minutes := 0
    for len(queue) > 0 {
        size := len(queue)
        for i:=0;i<size;i++{
            x, y := queue[0].idx, queue[0].idy
            queue = queue[1:]

            for j:=0;j<4;j++{
                xx, yy := x + dx[j], y + dy[j]
                if xx >= 0 && xx < n && yy >= 0 && yy < m && grid[xx][yy] == 1 {
                    grid[xx][yy] = 2
                    queue = append(queue, cell{idx: xx, idy: yy})
                    fresh--
                }
            }
        }
        if len(queue) > 0 {
            minutes++
        }
    }

    if fresh > 0 {
        return -1
    }

    return minutes
}

/*
for each cell, if it's rotten, bfs to all fresh
if curr less than already filled, replace

loop again, find maximum number, if there exist fresh with no visited, return -1 

*/