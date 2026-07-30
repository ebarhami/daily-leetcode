func maxSubArray(nums []int) int {
    total := 0
    answer := -20000
    for _, num := range nums {
        total += num 
        if total < num {
            total = num
        }

        if total > answer {
            answer = total
        }
    }

    return answer
}