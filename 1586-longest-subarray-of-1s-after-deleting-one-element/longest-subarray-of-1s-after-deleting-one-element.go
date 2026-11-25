func longestSubarray(nums []int) int {
    l := 0
    foundZero := false
    answer := 0

    for r:=0;r<len(nums);r++{
        if nums[r] == 0 {
            for foundZero {
                if nums[l] == 0 {
                    foundZero = false
                }
                l++
            }
            foundZero = true
        }

        if r - l + 1 > answer {
            answer = r-l+1
        }
    }

    return answer - 1
}