func minimumReplacement(nums []int) int64 {
    var answer int64

    next := nums[len(nums)-1]
    for i:=len(nums)-2;i>=0;i--{
        if nums[i] > next {
            x := (ceil(nums[i], next))
            answer += (x-1)
            if nums[i] % next != 0 {
                next = nums[i]/int(x)
            }
        } else {
            next = nums[i]
        }
    }

    return answer
}

func ceil(n, x int) int64 {
    answer := int64(n/x)
    if n%x != 0 {
        answer++
    }
    return answer
}