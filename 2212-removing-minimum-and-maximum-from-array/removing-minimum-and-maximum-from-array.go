func minimumDeletions(nums []int) int {
    min, max := 0, 0

    for i, num := range nums {
        if num < nums[min] {
            min = i
        }
        if num > nums[max] {
            max = i
        }
    }
    if min > max {
        temp := min
        min = max
        max = temp
    } 

    a := min+1 + len(nums) - max
    b := max + 1
    c := len(nums) - min

    answer := minn(a,b)
    answer = minn(answer, c)

    return answer
}

func minn(a, b int) int {
    if a < b {
        return a
    }
    return b
}