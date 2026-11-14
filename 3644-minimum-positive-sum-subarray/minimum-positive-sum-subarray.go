func minimumSumSubarray(nums []int, l int, r int) int {
    inf := 100005
    answer := inf
    prefix := make([]int, len(nums))
    for i := 0;i<len(nums);i++{
        prefix[i] += nums[i]
        if i > 0 {
            prefix[i] += prefix[i-1]
        }
    }

    for size := l; size <= r; size++{
        j := 0
        for i := 0; i < len(nums); i++ {
            for i-j+1 > size {
                j++
            }
            if i-j+1 == size {
                sum := prefix[i]
                if j > 0 {
                    sum -= prefix[j-1]
                }
                if sum > 0 && sum < answer {
                    answer = sum
                }
            }
        }
    }
    
    if answer == inf {
        return -1
    }

    return answer
}

/*
[1, 2, 3, 4, 1, 1, 1], l = 2, r = 4

*/