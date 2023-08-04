func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s))

    for i:=0;i<len(dp);i++{
        for _, word := range wordDict {
            if dp[i] {continue}
            start := i-len(word)+1
            same := start >= 0 && s[start:i+1] == word
            if same{
                dp[i] = start==0 || (start>0 && dp[start-1])
            }
        }
    }
    return dp[len(s)-1]
}