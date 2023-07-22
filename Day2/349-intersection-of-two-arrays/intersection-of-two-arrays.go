func intersection(nums1 []int, nums2 []int) []int {
    answer := make(map[int]struct{})
    check := make(map[int]struct{}) 

    for _, num := range nums1 {
        check[num] = struct{}{}
    }

    for _, num := range nums2 {
        if _,ok:=check[num]; ok {
            answer[num] = struct{}{}
        }
    }

    ans := make([]int, len(answer))
    i := 0
    for num := range answer {
        ans[i] = num
        i++
    }

    return ans
}