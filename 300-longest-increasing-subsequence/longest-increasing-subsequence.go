func lengthOfLIS(nums []int) int {
    n := len(nums)
    LIS := make([]int, n)

    for i := n-1; i >= 0; i-- {
        LIS[i] = 1
        for j := i+1; j<n;j++{
            if nums[i] < nums[j] {
                LIS[i] = max(LIS[i], 1 + LIS[j])
            }
        }
    }

    answer := 0
    for _, val := range LIS {
        answer = max(answer, val)
    }

    return answer
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}