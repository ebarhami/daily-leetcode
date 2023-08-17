func updateMatrix(mat [][]int) [][]int {
    n,m := len(mat), len(mat[0])

    inf := 10005
    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            if mat[i][j] == 1 {
                mat[i][j] = inf
                if i > 0 {
                    mat[i][j] = min(mat[i][j], mat[i-1][j] + 1)
                }
                if j > 0 {
                    mat[i][j] = min(mat[i][j], mat[i][j-1] + 1)
                }
            }
        }
    }

    for i:=n-1; i>=0;i--{
        for j:=m-1;j>=0;j--{
            if i < n-1 {
                mat[i][j] = min(mat[i][j], mat[i+1][j] + 1)
            }
            if j < m-1 {
                mat[i][j] = min(mat[i][j], mat[i][j+1] + 1)
            }
        }
    }

    return mat
}

func min(a, b int) int {
    if a < b {return a}

    return b
}