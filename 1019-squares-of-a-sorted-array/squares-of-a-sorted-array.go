func sortedSquares(nums []int) []int {
    answer := make([]int, 0)

    for _, num := range nums {
        sq := num*num
        answer = append(answer, sq)
    }

    sort.Ints(answer)
    return answer
}