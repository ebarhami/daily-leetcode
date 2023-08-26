func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }

    dp := make([][]int, len(s1)+1)
    for i:=0;i<=len(s1);i++{
        dp[i] = make([]int, len(s2)+1)
        for j:=0;j<=len(s2);j++{
            dp[i][j] = -1
        }
    }

    return traverse(dp, 0, 0, 0, s1,s2,s3)
}

func traverse(dp [][]int, i, j, k int, s1,s2,s3 string) bool {
    if k == len(s3) {
        if i == len(s1) && j == len(s2) {
            return true
        }
        return false
    }

    if dp[i][j] != -1 {
        return dp[i][j] == 1
    }

    dp[i][j] = 0
    if i < len(s1) && s1[i] == s3[k] && traverse(dp, i+1, j,k+1, s1,s2,s3) {
        dp[i][j] = 1
    } else if j < len(s2) && s2[j] == s3[k] && traverse(dp, i, j+1,k+1, s1,s2,s3) {
        dp[i][j] = 1
    }

    return dp[i][j] == 1
}
