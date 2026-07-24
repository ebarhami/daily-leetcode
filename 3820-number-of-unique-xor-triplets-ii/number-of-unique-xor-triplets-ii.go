func uniqueXorTriplets(nums []int) int {
    pairs := make([]bool, 2048)

    for _, num1 := range nums {
        for _, num2 := range nums {
            val := num1 ^ num2
            pairs[val] = true
        }
    }

    answer := make([]bool, 2048)
    for val, exists := range pairs {
        if !exists {
            continue
        }

        for _, num := range nums {
            triplet := val ^ num
            answer[triplet] = true
        }
    }

    count := 0
    for _, exists := range answer {
        if exists {
            count++
        }
    }

    return count
}

// for range uniquePairs -> 2048
//    for range nums -> 1500

//    4096 x 1500 -> 


// unique pairs = two loops againts nums
// max(len(uniquePairs)) = 2^11 -> 2048