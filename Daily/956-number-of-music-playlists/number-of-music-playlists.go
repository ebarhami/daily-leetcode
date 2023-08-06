const MOD int = 1e9 + 7

func power(x, y int) int {
    res := 1
    x %= MOD
    for y > 0 {
        if y & 1 == 1 {
            res = (res * x) % MOD
        }
        y >>= 1
        x = (x * x) % MOD
    }
    return res
}

func numMusicPlaylists(n, goal, k int) int {
    factorial := make([]int, n+1)
    inv_factorial := make([]int, n+1)
    factorial[0] = 1
    inv_factorial[0] = 1
    for i := 1; i <= n; i++ {
        factorial[i] = (factorial[i-1] * i) % MOD
        inv_factorial[i] = power(factorial[i], MOD-2)
    }

    sign := 1
    answer := 0
    for i := n; i >= k; i-- {
        temp := power(i-k, goal-k) * factorial[n] % MOD * inv_factorial[n-i] % MOD
        temp = temp * inv_factorial[i-k] % MOD
        answer = (answer + sign*temp) % MOD
        if answer < 0 {
            answer += MOD
        }
        sign *= -1
    }
    return answer
}