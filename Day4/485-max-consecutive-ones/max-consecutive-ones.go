func findMaxConsecutiveOnes(nums []int) int {
    maks := 0
    answer := 0
    for i:=0;i<len(nums);i++{
        if nums[i] == 1 {
            maks++
        } else {
            maks = 0
        }
        if maks > answer {answer = maks}
    }
    return answer
}