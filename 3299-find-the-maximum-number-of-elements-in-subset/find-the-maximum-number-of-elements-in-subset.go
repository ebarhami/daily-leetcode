func maximumLength(nums []int) int {
	sort.Ints(nums)

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	ans := 1

	// Special handling for 1
	if cnt, ok := freq[1]; ok {
		if cnt%2 == 1 {
			ans = max(ans, cnt)
		} else {
			ans = max(ans, cnt-1)
		}
		delete(freq, 1) // prevent DP from processing 1
	}

	// dp[x] = number of levels before x
	dp := make(map[int]int)

	for _, num := range nums {
		if num == 1 || freq[num] == 0 {
			continue
		}

		// use num as the center
		cur := dp[num]*2 + 1
		ans = max(ans, cur)

		// propagate only if num can be an intermediate node
		if freq[num] >= 2 {
			next := num * num
			dp[next] = max(dp[next], dp[num]+1)
		}

		// process each distinct number once
		freq[num] = 0
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}