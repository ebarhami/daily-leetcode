func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text1))
    for i:=0;i<len(text1);i++{
        dp[i] = make([]int, len(text2))
        for j:=0;j<len(text2);j++{
            dp[i][j] = -1
        }
    }

    return solve(0,0,dp,text1, text2)
}

func solve(i, j int, dp [][]int, text1, text2 string) int {
    if i >= len(text1) || j >= len(text2) {return 0}

    if dp[i][j] > -1 {return dp[i][j]}
    if text1[i] == text2[j] {
        dp[i][j] = 1 + solve(i+1, j+1, dp, text1, text2)
        return dp[i][j]
    }
    dp[i][j] = max(solve(i+1, j, dp, text1, text2), solve(i, j+1, dp, text1, text2))
    return dp[i][j]
}

func max(a, b int) int {
    if a > b {return a}
    return b
}