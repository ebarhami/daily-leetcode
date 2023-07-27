func nextGreaterElement(nums1 []int, nums2 []int) []int {
    answer := make([]int, len(nums1))
    nextGreater := make(map[int]int)
    stack := make([]int, 0)
    
    for _, num := range nums2 {
        nextGreater[num] = -1

        // monotonic stack
        for len(stack) > 0 && stack[len(stack)-1] < num {
            top := stack[len(stack)-1]
            nextGreater[top] = num
            stack = stack[:len(stack)-1]
        }

        stack = append(stack, num)
    }

    for i, num := range nums1 {
        answer[i] = nextGreater[num]
    }

    return answer
}