func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    answer := make([][]int, 0)

    for i:=0;i<len(nums);i++{
        first := nums[i]
        l, r := i+1, len(nums)-1
        if i > 0 && nums[i] == nums[i-1] {continue}
        for l < r {
            total := nums[l] + nums[r]
            if total + first == 0 {
                answer = append(answer, []int{first, nums[l], nums[r]})
                for l+1 < r && nums[l] == nums[l+1] {l++}
                for l < r-1 && nums[r] == nums[r-1] {r--}
                l++
                r--
            } else if total + first > 0 { // decrease
                r--
            } else { // increase
                l++
            }
        }
    }

    return answer
}