func longestSubsequence(nums []int) int {
    zero := 0
    xor := 0

    for _, num := range nums {
        if num == 0 {
            zero++
        }
        xor = xor ^ num
    }
    if xor != 0 {
        return len(nums)
    }
    if zero == len(nums) {
        return 0
    }

    return len(nums)-1
}