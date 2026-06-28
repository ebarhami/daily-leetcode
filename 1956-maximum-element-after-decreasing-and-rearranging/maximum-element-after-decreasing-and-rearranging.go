func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    sort.Ints(arr)

    arr[0] = 1
    max := 1

    for i, _ := range arr {
        if i == 0 {continue}
        if arr[i] - arr[i-1] > 1 {
            arr[i] = arr[i-1] + 1
        }
        if arr[i] > max {
            max = arr[i]
        }
    }

    return max   
}