func gcdSum(nums []int) int64 {
    n:=len(nums)
    prefix := make([]int, n)
    maks := 0
    for i := range n {
        maks = max(maks, nums[i])
        prefix[i] = gcd(maks, nums[i])
    }
    sort.Ints(prefix)
    sum := int64(0)
    for i:=0;i<n/2;i++{
        sum += int64(gcd(prefix[i], prefix[n-i-1]))
    }
    return sum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return  gcd(b, a%b)
}