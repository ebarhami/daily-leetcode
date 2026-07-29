func coinChange(coins []int, amount int) int {
    memo := make([]int, amount+1)
    for i:=0;i<=amount;i++{
        memo[i] = -1
    }

    inf := math.MaxInt32
    var dp func(int) int
    dp = func(remaining int) int {
        if remaining == 0 {
            return 0
        }
        if remaining < 0 {
            return inf
        }
        if memo[remaining] != -1 {
            return memo[remaining]
        }
        res := inf
        for _, coin := range coins {
            take := dp(remaining-coin)
            if take != inf {
                if take + 1 < res {
                    res = take + 1
                }
            }
        }
        memo[remaining] = res
        return memo[remaining]
    }

    answer := dp(amount)
    if answer == inf {
        return -1
    }
    return answer
}