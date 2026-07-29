func change(amount int, coins []int) int {
    n := len(coins)

    memo := make([][]int, n)
    for i := 0; i < n; i++ {
        memo[i] = make([]int, amount+1)

        for j := 0; j <= amount; j++ {
            memo[i][j] = -1
        }
    }

    answer := dp(0, amount, coins, memo)

    if answer < 0 {
        return 0
    }

    return answer
}

func dp(idx int, remaining int, coins []int, memo [][]int) int {
    if idx >= len(coins) {
        return 0
    }
    if remaining < 0 {
        return 0
    }
    if remaining == 0 {
        return 1
    }
    if memo[idx][remaining] != -1 {
        return memo[idx][remaining]
    }
    
    skip := dp(idx+1, remaining, coins, memo)
    take := dp(idx, remaining-coins[idx], coins, memo)
    total := skip + take

    memo[idx][remaining] = total
    return memo[idx][remaining]
}