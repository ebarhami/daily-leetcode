func removeElement(nums []int, val int) int {
    l, r := 0, len(nums)-1
    answer := 0

    for l <= r {
        for r >= 0 && nums[r] == val {
            r--
            answer++
        }
        for l < len(nums) && nums[l] != val {
            l++
        }
        if l < r {
            nums[l], nums[r] = nums[r], nums[l]
            answer++
        }
        l++
        r--
    }

    return len(nums) - answer
}

/*

3
3 2 2 2 3 3 2

*/