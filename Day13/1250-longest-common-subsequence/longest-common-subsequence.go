func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text1))
    for i:=0;i<len(text1);i++{
        dp[i] = make([]int, len(text2))
        for j:=0;j<len(text2);j++{
            dp[i][j] = -1
        }
    }

    return solve(text1, text2, 0, 0, dp)
}

func solve(a, b string, i, j int, dp [][]int) int {
    if i >= len(a) || j >= len(b) {
        return 0
    }

    if dp[i][j] != -1 {
        return dp[i][j]
    }

    if a[i] == b[j] {
        dp[i][j] = 1 + solve(a,b,i+1,j+1,dp)
        return dp[i][j]
    }
    moveA := solve(a,b,i+1,j,dp)
    moveB := solve(a,b,i,j+1,dp)

    dp[i][j] = max(moveA, moveB)
    return dp[i][j]
}

func max(a, b int) int {
    if a > b {return a}
    return b
}