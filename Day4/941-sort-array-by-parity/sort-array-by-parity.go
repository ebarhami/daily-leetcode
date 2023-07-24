func sortArrayByParity(nums []int) []int {
    l, r := 0, len(nums)-1
    for l < r {
        for l<len(nums) && nums[l]%2==0{
            l++
        }
        for r >= 0 && nums[r]%2==1{
            r--
        }
        if l < r {
            nums[l], nums[r] = nums[r], nums[l]
        }
    }

    return nums
}