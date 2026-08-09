func stoneGameII(piles []int) int {
	n := len(piles)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	suffixSum := make([]int, n)
	copy(suffixSum, piles)

	for i := n - 2; i >= 0; i-- {
		suffixSum[i] += suffixSum[i+1]
	}

	return maxStones(suffixSum, 1, 0, memo)
}

func maxStones(suffixSum []int, maxTillNow int, currIndex int, memo [][]int) int {
	if currIndex+2*maxTillNow >= len(suffixSum) {
		return suffixSum[currIndex]
	}

	if memo[currIndex][maxTillNow] > 0 {
		return memo[currIndex][maxTillNow]
	}

	res := math.MaxInt

	for i := 1; i <= 2*maxTillNow; i++ {
		opponentStones := maxStones(
			suffixSum,
			max(i, maxTillNow),
			currIndex+i,
			memo,
		)

		res = min(res, opponentStones)
	}

	memo[currIndex][maxTillNow] = suffixSum[currIndex] - res

	return memo[currIndex][maxTillNow]
}