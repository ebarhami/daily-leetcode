func pivotArray(nums []int, pivot int) []int {
    less, more, equal := make([]int, 0), make([]int, 0), make([]int, 0)

    for _, num := range nums {
        if num > pivot {
            more = append(more, num)
        } else if num < pivot {
            less = append(less, num)
        } else {
            equal = append(equal, num)
        }
    }
    
    less = append(less, equal...)
    less = append(less, more...)

    return less
}