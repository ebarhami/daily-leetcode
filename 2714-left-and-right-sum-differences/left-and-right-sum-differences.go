func leftRightDifference(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return []int{}
    }

    left, right := make([]int, n), make([]int, n)

    for i := 1;i<n;i++{
        left[i] += left[i-1] + nums[i-1]
    }

    for i := n-2;i>=0;i--{
        right[i] += right[i+1] + nums[i+1]
    }
    
    answer := make([]int, n)
    for i := 0; i < n; i++ {
        answer[i] = abs(left[i], right[i])
    }
    return answer
}

func abs(a,b int) int {
    if a > b {
        return a-b
    }
    return b-a
}