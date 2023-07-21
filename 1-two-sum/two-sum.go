func twoSum(nums []int, target int) []int {
    appear := make(map[int]int)

    for i, num := range nums {
        if j, ok := appear[target-num]; ok {
            return []int{i, j}
        }
        appear[num] = i
    }
    return nil
}