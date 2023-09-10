const MOD = 1000000007

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = (result * i) % MOD
	}
	return result
}

func pow(base, exp, mod int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func countOrders(n int) int {
	twoNFact := factorial(2 * n) % MOD
	twoPowN := pow(2, n, MOD)
	return (twoNFact * pow(twoPowN, MOD-2, MOD)) % MOD
}


/*
11
_,_1_,_1_,_ = C(2n+1,2n) + (2n+1)
_X_X_X_X_ = C(n+1,n)  


*/