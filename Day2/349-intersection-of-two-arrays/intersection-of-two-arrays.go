func intersection(nums1 []int, nums2 []int) []int {
    answer := make([]int, 0)
    check := make(map[int]struct{}) 

    for _, num := range nums1 {
        check[num] = struct{}{}
    }

    for _, num := range nums2 {
        if _,ok:=check[num]; ok {
            answer = append(answer, num)
            delete(check, num)
        }
    }

    return answer
}