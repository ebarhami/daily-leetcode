type Point struct {
    paths int
    cost int
}

var (
    dx = []int{0, 1, 1}
    dy = []int{1, 1, 0}
)

func pathsWithMaxScore(board []string) []int {
    n, m := len(board), len(board[0])    
    dp := make([][]Point, n)
    for i:=0;i<n;i++{
        dp[i] = make([]Point, m)
        for j:=0;j<m;j++{
            dp[i][j] = Point{cost: -1}
        }
    }
    dp[n-1][m-1] = Point{
        paths: 1,
        cost: 0,
    }

    for i:=n-1; i>=0; i--{
        for j:=m-1; j>=0; j--{
            if i == n-1 && j == m-1 {continue}
            if board[i][j] == 'X' {continue}
            dp[i][j] = getPoint(dp, board, i, j)
        }
    }


    goal := dp[0][0]
    if goal.paths == 0 {
        return []int{0, 0}
    }
    return []int{goal.cost, goal.paths}
}

func getPoint(dp [][]Point, board []string, x, y int) Point {
    n,m := len(board), len(board[0])
    val := getInt(board[x][y])
    mod := 1000000007
    
    curCost := 0
    curPath := 0
    for i:=0;i<3;i++{
        xx := dx[i] + x
        yy := dy[i] + y
        if xx >= n || yy >= m {
            continue
        }
        if dp[xx][yy].cost > curCost {
            curCost = dp[xx][yy].cost 
            curPath = dp[xx][yy].paths % mod
        } else if dp[xx][yy].cost == curCost {
            curPath += dp[xx][yy].paths % mod
        }
    }
    return Point{
        cost: val + curCost,
        paths: curPath % mod,
    }
}

func getInt(c byte) int {
    if c=='E' || c=='S' {
        return 0
    }
    return int(c - '0')
}





/*
"E23",
"2X2",
"12S"

checks 3 direction
get the max value from 3 dir, curr value = curr + max

get the paths 
- if only one of the dir is max, then paths is that dir paths 
- if 2-3 dir is max, then paths is 2-3 dir paths sum


"E12",
"1X1",
"21S"
*/