func validPartition(nums []int) bool {
    dp := make([]bool, len(nums)+1)

    if len(nums) <= 1 {return false}
    dp[0] = true
    for i:=2;i<=len(nums);i++{
        idx := i-1
        dp[i] = (firstCond(nums, idx) && dp[i-2])
        if i > 2 {
            dp[i] = dp[i] || (secondCond(nums, idx) && dp[i-3]) || (thirdCond(nums, idx) && dp[i-3])
        }
    }

    return dp[len(nums)]
}

func firstCond(nums []int, i int) bool {
    if i-1 < 0 {return false}
    return nums[i] == nums[i-1]
}

func secondCond(nums []int, i int) bool {
    if i-2 < 0 {return false}
    return nums[i] == nums[i-1] && nums[i-1] == nums[i-2]
}

func thirdCond(nums []int, i int) bool {
    if i-2 < 0 {return false}
    return nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1
}