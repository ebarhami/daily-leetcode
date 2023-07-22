func intersect(nums1 []int, nums2 []int) []int {
    check := make(map[int]int)
    for _, num := range nums1 {
        check[num]++
    }

    answer := make([]int, 0)
    for _, num := range nums2 {
        if check[num] > 0 {
            check[num]--
            answer = append(answer, num)
        }
    }

    return answer
}