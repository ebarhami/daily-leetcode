
func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0)

	for i:=0;i<k;i++{
		q = insertX(q, i, nums)
	}
	answer := make([]int, 0)
	answer = append(answer, nums[q[0]])

	for i:=k;i<len(nums);i++{
		q = insertX(q, i, nums)
		for q[0] <= i-k {
			q = q[1:]
		}
		answer = append(answer, nums[q[0]])
	}

	return answer
}

func insertX(q []int, i int, nums []int) []int {
	for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
		q = q[:len(q)-1]
	}
	return append(q, i)
}
/*
O(nk). n = k = 10^5


*/ 
