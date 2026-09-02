func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    memo := make([]int, len(nums))
    memo2 := make([]int, len(nums))
    for i:=0;i<len(nums);i++{
        memo[i] = -1
        memo2[i] = -1
    }

    return max(dp(memo, 0, nums[:len(nums)-1]), dp(memo2, 1, nums))
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