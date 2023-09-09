func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    for i:=0;i<=target;i++ {
        dp[i]=-1
    }
    
    return calc(nums, target, dp)
}

func calc(nums []int, target int, dp []int) int {
    if target < 0 {
        return 0
    }
    if target == 0 {
        return 1
    }
    if dp[target] != -1 {
        return dp[target]
    }
    
    total := 0
    for _, num := range nums {
        total += calc(nums, target-num, dp)
    }
    dp[target] = total
    
    return dp[target]
}