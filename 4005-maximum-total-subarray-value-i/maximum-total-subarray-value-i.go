func maxTotalValue(nums []int, k int) int64 {
    max, min := -1, 1000000001

    for _, num := range nums {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }

    return int64(max-min) * int64(k)
}