func maxIceCream(costs []int, coins int) int {
    max := 0
    answer := 0

    freq := make(map[int]int)

    for _, cost := range costs {
        freq[cost]++
        if cost > max {
            max = cost
        }
    }

    for i:=1; i<=max; i++{
        if _, exist := freq[i]; !exist {
            continue
        }
        if freq[i] * i <= coins {
            coins -= (freq[i]*i)
            answer += freq[i]
        } else {
            answer += (coins / i)
            coins -= (int(coins / i) * i)
        }
    }

    return answer
} 

/*
3 3 3
coins = 8

if 3 * 3 <= coins -> coins -= 3*3
else -> coins/cost -> 2


*/