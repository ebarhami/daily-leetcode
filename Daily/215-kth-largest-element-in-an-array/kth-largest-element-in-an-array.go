func findKthLargest(nums []int, k int) int {
    l, r := -10000, 10000
    ans := 0
    for l <= r {
        mid := (l+r)/2
        nGreaterThanX := calcNumsGreaterThanX(mid, nums)
        if nGreaterThanX >= k {
            l = mid + 1
            ans = mid
        } else {
            r = mid - 1
        }
    }

    return ans
}

func calcNumsGreaterThanX(x int, nums []int) int {
    total := 0
    for _, num := range nums {
        if num >= x {
            total++
        }
    }
    return total
}

/*


*/