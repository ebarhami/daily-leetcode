func knightProbability(n int, k int, row int, column int) float64 {
    prob := make([][][]float64, k+1)
    for i:=0;i<k+1;i++{
        prob[i] = make([][]float64, n)
        for j:=0;j<n;j++{
            prob[i][j] = make([]float64, n)
        }
    }
    // probability after 0 moves in {row,colum} vs outside the board is 1, since the knight cannot move outside the board yet
    prob[0][row][column] = 1.0

    for move:=1;move<=k;move++{ 
        for i:=0;i<n;i++{
            for j:=0;j<n;j++{
                for _, prevCoordinate := range getCoordinateBeforeKnightMoves(i,j) {
                    x, y := prevCoordinate[0], prevCoordinate[1]
                    if isValid(x,y,n) {
                        prob[move][i][j] += (prob[move-1][x][y]/8)
                    }
                }
            }
        }
    }

    answer := 0.0
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            answer += prob[k][i][j]
        }
    }
    return answer
}
var (
    xx = []int{-1,-2,-2,-1,1,2,2,1}
    yy = []int{-2,-1,1,2,2,1,-1,-2}
)
func getCoordinateBeforeKnightMoves(x, y int) [][]int {
    coordinate := make([][]int, 8)
    for i:=0;i<8;i++{
        coordinate[i] = []int{x-xx[i], y-yy[i]}
    }

    return coordinate
}

func isValid(x, y, n int) bool {
    return x >= 0 && x < n && y >= 0 && y < n
}