func rob(nums []int) int {
    memo := make([]int, len(nums))
    for i:=0;i<len(nums);i++{
        memo[i] = -1
    }

    return dp(memo, 0, nums)
}

func dp(memo []int, idx int, nums []int) int {
    if idx >= len(nums) {
        return 0
    }
    if memo[idx] != -1 {
        return memo[idx]
    }

    memo[idx] = max(dp(memo, idx+1, nums), dp(memo, idx+2, nums) + nums[idx])
    return memo[idx]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}