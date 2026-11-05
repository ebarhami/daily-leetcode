func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    answer := 0

    for _, num := range arr1 {
        count := 0
        for _, num2 := range arr2 {
            if dist(num, num2) > d {
                count++
            }
        }
        if count == len(arr2) {
            answer++
        }
    }

    return answer
}

func dist(a, b int) int {
    if a > b {
        return a - b
    }
    return b-a
}