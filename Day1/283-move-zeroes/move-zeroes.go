func moveZeroes(nums []int)  {
    curr := 0
    for i:=0;i<len(nums);i++{
        if nums[i] != 0 {
            nums[curr], nums[i] = nums[i], nums[curr]
            curr++
        }
    }
    return 
}