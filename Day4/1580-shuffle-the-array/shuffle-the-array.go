func shuffle(nums []int, n int) []int {
    answer := make([]int, 2*n)
    for i:=0;i<n;i++{
        answer[i*2] = nums[i]
        answer[i*2+1] = nums[i+n]
    }
    return answer
}

/*

1,2,3,4,10,20,30,40
1  2  3  4
 10 20 30 40


1,4,2,3,3,2,4,1
*/