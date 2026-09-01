func missingMultiple(nums []int, k int) int {
    multiplier := 1 
    exist := make(map[int]bool)
    for _, num := range nums {
        exist[num] = true
    }

    for true {
        x := k * multiplier
        if !exist[x] {
            return x
        }
        multiplier++
    }
    return 0
}