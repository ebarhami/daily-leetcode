func canPartition(nums []int) bool {
    total, n := 0, len(nums)
    for _, num := range nums {
        total += num
    }
    target := total/2
    if total % 2 == 1 {
        return false
    }

    dp := make([][]int, n)
    for i:=0;i<n;i++{
        dp[i] = make([]int, total)
        for j:=0;j<total;j++{
            dp[i][j] = -1
        }
    }

    val := knapsack(nums, dp, 0, target)

    if val == target {
        return true
    }
    return false
}

func knapsack(nums []int, dp [][]int, idx int, capacity int) int {
    if idx >= len(nums) || capacity < 0 {
        return 0
    }
    
    if dp[idx][capacity] != -1 {
        return dp[idx][capacity]
    }

    skip := knapsack(nums, dp, idx+1, capacity)

    take := 0
    if nums[idx] <= capacity {
        take += nums[idx]
        take += knapsack(nums, dp, idx+1, capacity-nums[idx])
    }

    dp[idx][capacity] = max(skip, take)
    return dp[idx][capacity]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}