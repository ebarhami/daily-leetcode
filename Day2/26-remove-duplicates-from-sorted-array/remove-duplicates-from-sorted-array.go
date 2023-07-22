func removeDuplicates(nums []int) int {
    curr := 0
    unique := 0
    for i:=0;i<len(nums);i++{
        if i == 0 {
            unique++
            curr++
        } else {
            if nums[i] != nums[i-1] { // append
                unique++
                nums[curr] = nums[i]
                curr++
            } else { //duplicate

            }
        }
    }
    return unique
}