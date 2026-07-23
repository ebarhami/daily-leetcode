func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	ans := 1
	for ans <= n {
		ans <<= 1
	}
	return ans
}

/*
1 -> 1 
2 -> 2
3 -> 4
4 -> 8    -> 2^n
5 -> 8
6 -> 8
7 -> 8
8 -> 16

001
010
011
100


101



*/