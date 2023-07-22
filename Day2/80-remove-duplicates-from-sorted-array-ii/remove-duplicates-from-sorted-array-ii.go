func removeDuplicates(nums []int) int {
    curr := 0
    for i, _ := range nums {
        num := nums[i]
        if curr < 2 {
            nums[curr] = num
            curr++
        } else {
            if (num == nums[curr-1]) && (num == nums[curr-2]) { // duplicate > twice

            } else {
                nums[curr] = num
                curr++
            }
        }
    }
    return curr
}