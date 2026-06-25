func countMajoritySubarrays(nums []int, target int) int {
    answer := 0
    n := len(nums)
    count := make([]int, n)

    for i, num := range nums {
        sum := 0
        if num == target {
            sum = 1
        }

        count[i] += sum
        if i > 0 {
            count[i] += count[i-1]
        }
    }

    for i:=0;i<n;i++{
        for j:=i;j<n;j++{
            length := j-i+1
            total :=  count[j] - count[i]
            if nums[i] == target {
                total++
            }
            if total > (length)/2 {
                answer++
            }
            
        }
    }

    return answer
}

