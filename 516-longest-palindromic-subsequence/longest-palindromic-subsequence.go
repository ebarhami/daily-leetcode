func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i:=0;i<n;i++{
        dp[i] = make([]int, n)
        for j:=0;j<n;j++{
            dp[i][j] = -1
        }
    }

    return solve(dp, s, 0, n-1)
}

func solve(dp [][]int, s string, l, r int) int {
    if l > r {
        return 0
    }
    if l == r {
        return 1
    }
    if dp[l][r] != -1 {
        return dp[l][r]
    }
    
    if s[l] == s[r] {
        dp[l][r] = 2 + solve(dp, s, l+1, r-1)
        return dp[l][r]
    }
    goLeft := solve(dp, s, l+1, r)
    goRight := solve(dp, s, l, r-1)

    dp[l][r] = goRight
    if goLeft > goRight {
        dp[l][r] = goLeft
    }

    return dp[l][r]
}

/*
states: left index , right index
*/