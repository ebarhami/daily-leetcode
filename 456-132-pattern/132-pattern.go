func find132pattern(nums []int) bool {    
    minSoFar := make([]int, len(nums))
    stack := make([]int, 0)

    min := math.MaxInt32
    for i, num := range nums {
        if num < min {
            min = num
        }
        minSoFar[i] = min        
    }

    for i:=len(nums)-1;i>=0;i--{
        num := nums[i]
        if num <= minSoFar[i] {
            continue
        }

        for len(stack) > 0 && (stack[len(stack)-1]) <= minSoFar[i]{
            stack = stack[:len(stack)-1]
        }

        if len(stack) > 0 {
            top := stack[len(stack)-1]
            if top < num && top > minSoFar[i] {
                return true
            }
        }

        stack = append(stack, num)
    }

    return false
}

/*
132

3,5,10,3,4
3 3 10 3 4 -> min so far
    1 3 4

3 5 4
*/