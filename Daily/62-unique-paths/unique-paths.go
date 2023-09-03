func uniquePaths(m int, n int) int {
    return combine(n+m-2, m-1)
}

func combine(n int, k int) int {
    comb := make([][]int,2)
    comb[0] = make([]int, k+1)
    comb[1] = make([]int, k+1)
    prev := 0
    now := 1
    temp := 0
    comb[0][0] = 1
    for i:=0;i<=n;i++{
        for j:=0;j<=i && j<=k;j++{
            if j==0 || j==i || i<=1{
                comb[now][j]=1
            } else if j == 1 {
                comb[now][j]=i
            } else {
                comb[now][j] = comb[prev][j] + comb[prev][j-1]
            }
        }
        temp = prev
        prev = now
        now = temp
    }
    
    return comb[prev][k]
}