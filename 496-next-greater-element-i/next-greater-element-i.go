func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nextGreat := make(map[int]int)

    stack := make([]int, 0)
    for _, num := range nums2 {
        // popping
        for len(stack) > 0 && num > stack[len(stack) - 1] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]

            nextGreat[top] = num
        }
        
        if len(stack) == 0 || num < stack[len(stack)-1] {
            stack = append(stack, num)
        }
    }

    answer := make([]int, len(nums1))
    for i, num := range nums1 {
        if val, exist := nextGreat[num]; exist {
            answer[i] = val
        } else {
            answer[i] = -1
        }
    }

    return answer
}