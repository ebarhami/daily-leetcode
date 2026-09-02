func maxArea(height []int) int {
    l, r := 0, len(height)-1

    answer := 0
    for l < r {
        area := min(height[l], height[r]) * (r-l)
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
        answer = max(answer, area)
    }

    return answer
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
} 
