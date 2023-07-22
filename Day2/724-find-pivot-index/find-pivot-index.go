func pivotIndex(nums []int) int {
    n := len(nums)
    left := make([]int, n)

    if n == 1 {return 0}

    for i:=n-1;i>=0;i--{
        if i == n-1 {
            left[i] = nums[i]
        } else {
            left[i] = left[i+1] + nums[i]
        }
    }
    for i:=1;i<n;i++{
        nums[i] = nums[i] + nums[i-1]
    }

    for i:=0;i<n;i++{
        if i == 0 {
            if n > 1 && left[i+1] == 0 {
                return i
            }
        } else if i == n-1 {
            if nums[i-1] == 0 {
                return i
            }
        } else {
            if nums[i-1] == left[i+1] {
                return i
            }
        }
    }
    return -1
}