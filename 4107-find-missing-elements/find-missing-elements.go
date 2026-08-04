func findMissingElements(nums []int) []int {
    answer := make([]int, 0)
    exist := make(map[int]bool)
    min, max := 101, 0

    for _, num := range nums {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
        exist[num] = true
    }

    for i:=min;i<=max;i++{
        if !exist[i] {
            answer = append(answer, i)
        }
    }

    return answer
}