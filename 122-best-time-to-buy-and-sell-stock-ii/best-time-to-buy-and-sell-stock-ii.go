func maxProfit(prices []int) int {
    curr := math.MaxInt32
    profit := 0
    for _, price := range prices {
        if price > curr {
            profit += (price-curr)
        }
        curr = price
    }
    return profit
}

/*
[7,1,5,3,6,4]
curr : 7
price = 1

curr = 1
price 5

*/