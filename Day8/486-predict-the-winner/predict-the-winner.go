func PredictTheWinner(nums []int) bool {
    n := len(nums)
    dp := make([][]int, n)
    for i:=0;i<n;i++{
        dp[i] = make([]int, n)
        for j:=0;j<n;j++{
            dp[i][j] = -1
        }
    }
    l, r := 0, n-1
    return maxDiff(l, r, nums, dp) >= 0
}

func maxDiff(l, r int, nums []int, dp [][]int) int {
    if l == r {
        return nums[l]
    }
    if dp[l][r] != -1 {
        return dp[l][r]
    }

    takeLeft := nums[l] - maxDiff(l+1, r, nums, dp) // minus the max diff in the perspective of the next player
    takeRight := nums[r] - maxDiff(l, r-1, nums, dp)

    dp[l][r] = max(takeLeft, takeRight)

    return dp[l][r]
}

func max(a, b int) int {
    if a > b {return a}
    return b
}