func strangePrinter(s string) int {
    dp := make([][]int, len(s))
    for i := range dp {
        dp[i] = make([]int, len(s))
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dpCalc func(int, int) int
    dpCalc = func(l, r int) int {
        if dp[l][r] != -1 { return dp[l][r] }
                
        dp[l][r] = len(s)
        j := -1
        for i := l; i < r; i++ {
            if s[i] != s[r] && j == -1 { j = i }

            if j != -1 {
                dp[l][r] = min(
                    dp[l][r], 
                    1 + dpCalc(j, i) + dpCalc(i+1, r),
                )
            }
        }

        if j == -1 { dp[l][r] = 0 }
        return dp[l][r]
    }
    return dpCalc(0, len(s)-1) + 1
}

func min(a, b int) int {
    if a < b { return a }
    return b
}