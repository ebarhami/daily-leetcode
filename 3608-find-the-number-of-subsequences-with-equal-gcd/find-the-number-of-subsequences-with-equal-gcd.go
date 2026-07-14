const MOD = 1000000007

func subsequencePairCount(nums []int) int {
	m := 0
	for _, num := range nums {
		m = max(m, num)
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1

	for _, num := range nums {
		ndp := make([][]int, m+1)
		for i := range ndp {
			ndp[i] = make([]int, m+1)
		}

		for j := 0; j <= m; j++ {
			divisor1 := gcd(j, num)
			for k := 0; k <= m; k++ {
				val := dp[j][k]
				if val == 0 {
					continue
				}

				divisor2 := gcd(k, num)
				ndp[j][k] = (ndp[j][k] + val) % MOD
				ndp[divisor1][k] = (ndp[divisor1][k] + val) % MOD
				ndp[j][divisor2] = (ndp[j][divisor2] + val) % MOD
			}
		}
		dp = ndp
	}

	ans := 0
	for j := 1; j <= m; j++ {
		ans = (ans + dp[j][j]) % MOD
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}