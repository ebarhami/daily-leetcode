func moveZeroes(nums []int)  {
    l, r := 0, 0
    for l < len(nums) && nums[l] != 0 {
        l++
        r++
    }
    for l <= r && l < len(nums) {
        for r < len(nums) && nums[r] == 0 {
            r++
        }
        for l < len(nums) && nums[l] != 0 {
            l++
        }
        if l >= len(nums) || r >= len(nums) || l > r {
            return
        }
        temp := nums[l]
        nums[l] = nums[r]
        nums[r] = temp

        l++
        r++
    }
    return
}