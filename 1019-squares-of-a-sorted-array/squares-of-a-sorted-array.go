func sortedSquares(nums []int) []int {
        n := len(nums)
        res := make([]int, n, n)
        i, j := 0, n-1
        for k := n - 1; k >= 0; k-- {
            
            i2 := nums[i]*nums[i]
            j2 := nums[j]*nums[j]
            
            
                if i2 > j2{
                        res[k] = i2
                        i++
                } else {
                        res[k] = j2
                        j--
                }
        }

        return res
}