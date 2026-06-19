func largestAltitude(gain []int) int {
    max := 0
    curr := 0
    for _, num := range gain {
        curr += num 
        if curr > max {
            max = curr
        }
    }

    return max
}