type Item struct {
    val int
    idx int
}

func nextGreaterElements(nums []int) []int {
    newArr := append(nums, nums...)

    nextGreat := make(map[int]int)

    stack := make([]Item, 0)
    for i, num := range newArr {
        // popping
        for len(stack) > 0 && num > stack[len(stack) - 1].val {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]

            nextGreat[top.idx] = num
        }
        
        if len(stack) == 0 || num <= stack[len(stack)-1].val {
            stack = append(stack, Item{
                val: num,
                idx: i,
            })
        }
    }

    answer := make([]int, len(nums))
    for i, _ := range nums {
        if val, exist := nextGreat[i]; exist {
            answer[i] = val
        } else {
            answer[i] = -1
        }
    }

    return answer
}