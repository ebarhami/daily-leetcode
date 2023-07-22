func findDisappearedNumbers(nums []int) []int {
    curr := 0
    i := 0

    for i < len(nums) {
        x := nums[i] - 1
        if nums[i] != nums[x] {
            nums[x], nums[i] = nums[i], nums[x]
        } else {
            i++
        }
    }

    for i:=0;i<len(nums);i++{
        x := nums[i]
        if x != i+1 {
            nums[curr] = i+1
            curr++
        }
    }
    return nums[:curr]
}